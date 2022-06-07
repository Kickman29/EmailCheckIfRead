# Mail read checker

Small app that is serving 1x1px image file via http. Upon loading file it logs the http request. Port used is ```80```
Log of all connections is visible in terminal and at dashboard at ```9090``` port.

### Use case
Main purpose of application is to check if sent email was open by reciever. 
Since Google allows to use dynamic signature you can use it to track if email has been seen.

### How to use
- Compile & start binary file.
- To check if it is working go to ```localhost:9090``` in your web browser. Then use ```curl your_public_ip_address/anyString```, console should log data and it will show up in dashboard.


Dashboard can be working but you still won't be able to log data if ports are closed by your router or internet provider. So you need to open port 80 for this to work or use a VPN/tunnel/DDNS.

In case of using Google Mail to track email you need to add image to your signature in this format: ```http://your_public_ip_address/stringThatWillBeShownInLogger```
Also you can add multiple signatures with different strings, so you can differentiate receiver of email
