- This study project involves the application of the Graceful Shutdown technique. Essentially, it's the process of orderly shutting down a service/server, ensuring no data loss for the client.
- This example consists of three parts:
  - Creating a local server using the "net/http" library.
  - Goroutine - a thread that controls/listens to HTTP requests.
  - Handling possible system/server interruptions using "syscall".
- How it works:
  - The server is started and listens on port 3000. Within a time interval, in this case, 4 seconds, our client sends a request. If during this period the server detects an issue, the graceful shutdown process is initiated. Our client receives the response and upon confirmation, the server shuts down.
- Testing:
  - To execute this example, run the command "go run main.go" in your terminal or in the embedded code editor terminal. Then, using an API client like Postman or Curl, send a request to "http://localhost:3000". In the middle of sending the request within the 4-second interval, press "CTRL+C" in your terminal where the server is running, resulting in an abrupt stop.
  - At this point, the server recognizes the failure and needs to finish processing the client's request before starting the graceful shutdown process.