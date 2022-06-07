# Mail read checker

Small app that is serving 1x1px image file via http. Upon loading file it logs the http request. Port used is _80_
Log of all connections is visible in terminal and at dashboard at _9090_ port.

### Usecase
Main purpose of application is to check if sent email was open by reciever. 
Since Google allows to use dynamic signature you can use it to track if email has been seen.

### How to use
Compile & start binary file.
To check if it is working go to _localhost:9090_ in your web browser.
Dashboard can be working but you still won't be able to log data if you're sitting behind your provider's NAT. So you need to open port 80 for this to work or use a VPN/tunnel/DDNS.
