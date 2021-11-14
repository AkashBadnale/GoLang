# TO understand this example in depth follow this link : https://golang.org/doc/tutorial/web-service-gin

<h2> This is example with static data, without any database </h2>

In Short :
1. Create a folder web-service-gin
2. Add your code, here example.go
3. Inside the web-service gin run this command : go mod init example/web-service-gin
4. Also run : go get: added github.com/gin-gonic/gin v1.7.2
5. Then to run the project : go run .
6. GET : http://localhost:8080/albums
7. GET/{id} : http://localhost:8080/albums/2
8. POST : http://localhost:8080/albums
          header : "Content-Type: application/json"
          body : '{
                   "id": "4",
                   "title": "The Modern Sound of Betty Carter",
                   "artist": "Betty Carter",
                   "price": 49.99
                   }'
