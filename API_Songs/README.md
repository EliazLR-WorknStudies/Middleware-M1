# TP middleware example

## Run

Tidy / download modules :
```
go mod tidy
```
Build : 
```
go build -o middleware_collections cmd/main.go
```
Run : 
```
./middleware_collections
```

## Documentation

Documentation is visible in **api** directory ([here](api/swagger.json)).


POST /register 			#body {username, password}
POST /login				#body {username, password}
POST /logout     (connected only)
DELETE /delete   (connected only)
GET /introspect  (connected only)	
GET /users/      (connected only)
GET /users/<id>  (connected only)
GET /songs/      (connected only)
GET /songs/<id>  (connected only)
PUT /users/<id>  (connected only)	#body {username, password}
## Debug request (delete at the end)
DELETE /users/<id>
