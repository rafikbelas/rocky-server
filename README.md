# rocky-server
Golang server that monitors servers using WebSockets.

A server app to periodically fetch the status of services and serve it using WebSockets.

## Getting Started

* Download the project and cd to the project directory.
* Run it using the command 
```go run main.go```
* Run `live-server` on the same directory to start `index.html` which is a client that consumes the server and open the console to see the responses 
OR set up a client application that listens to `ws://localhost:8080/reports` (<a href="https://github.com/rafikbelas/rocky-ui" target="_blank">See Rocky-UI project</a>).
