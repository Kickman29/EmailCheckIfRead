package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Handler struct {
	l *log.Logger
}

func NewHandler(l *log.Logger) Handler {
	return Handler{l}
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	http.ServeFile(rw, r, "resources/1x1.png")

	h.l.Print("\n\nConnection! Logging to file")

	output := fmt.Sprint(
		"\n\nTimestamp:", time.Now(),
		"\nHeader: ", r.Header.Values("User-Agent"),
		"\nRequested URL: ", r.RequestURI,
		"\nIP Address: ", r.RemoteAddr,
	)

	ci := newConnectionInfo(r.Header.Values("User-Agent"), r.RequestURI, r.RemoteAddr)
	ci.formHTMLPage()

	h.l.Print(output)
	h.writeToFile(output)
}

func (h *Handler) writeToFile(data string) {

	f, err := os.OpenFile("output.md",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		h.l.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(data); err != nil {
		log.Println(err)
	}

	h.l.Print("\nData has been written to file")

}

type connectionInfo struct {
	Timestamp  time.Time
	Header     []string
	RequestURI string
	RemoteAddr string
}

func newConnectionInfo(Header []string, RequestURI string, RemoteAddr string) connectionInfo {
	return connectionInfo{time.Now(), Header, RequestURI, RemoteAddr}
}

// Big piece of crap, adds data to html page as []byte to be served
func (c *connectionInfo) formHTMLPage() {

	// formatting text to make it pretty
	date := fmt.Sprint(c.Timestamp.Day(), "/", c.Timestamp.Month(), "/", c.Timestamp.Year())
	time := fmt.Sprint(c.Timestamp.Hour(), ":", c.Timestamp.Minute(), ":", c.Timestamp.Second())
	datetime := fmt.Sprint(date, ", ", time)

	formattedStruct := []byte(fmt.Sprintf(`
	<tr>
		<td>%v</td>
		<td>%v</td>
		<td>%v</td>
		<td>%v</td>
	</tr>
	`,
		datetime,
		c.Header,
		c.RequestURI,
		c.RemoteAddr,
	))

	fp := "./resources/HTMLLog.html"
	htmlTemplate, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal("Cannot reading html template file", err)
	}

	separator := `<div id="start"></div>`

	body := strings.Split(string(htmlTemplate), separator)

	html := fmt.Sprint(body[0] + separator + string(formattedStruct) + body[1])

	os.WriteFile(fp, []byte(html), 0644)

}

type Dashboard struct {
	l *log.Logger
}

func NewDashboard(l *log.Logger) Dashboard {
	return Dashboard{l}
}

func (d *Dashboard) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	http.ServeFile(rw, r, "./resources/HTMLLog.html")
	// d.l.Println("Dashboard loaded")

}

func ClearLog(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	log.Println("Clearing Log")

	fp := "./resources/HTMLLog.html"
	htmlTemplate, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal("Cannot reading html template file", err)
	}

	startsep := `<div id="start"></div>`
	endsep := `<div id="end"></div>`

	start := strings.Split(string(htmlTemplate), startsep)
	end := strings.Split(string(start[1]), endsep)
	html := fmt.Sprint(start[0] + startsep + endsep + end[1])

	os.WriteFile(fp, []byte(html), 0644)
	http.Redirect(rw, r, "/", http.StatusSeeOther)
}

func CSSStyling(rw http.ResponseWriter, r *http.Request) {

	http.ServeFile(rw, r, `./resources/style.css`)
	log.Println("CSS Served")

}
