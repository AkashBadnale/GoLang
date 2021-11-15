# GOLANG POSTGRES + RESTAPI 

## TO RUN
* Set your configurations in main.go constants.
* `go build`
* `./postgres-api`


## ENDPOINTS
* `http://localhost:8080/create` : Method : POST
   body : </br>
           {
             "name" : "A",
             "age" : 25
           }
           
* `/read/{id}` Method:GET
