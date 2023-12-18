# Middleware-M1

https://edu.forestier.re/#/middleware-m1

To launch everything :
`./start_all_APIs.sh`

This will launch :

- Songs API (Go) on port 8080
- Users API (Go) on port 8081
- Ratings API (Go) on port 8082
- Flask on port 8888

POST /register 			#body {username, password}
POST /login				#body {username, password}
POST /logout      (connected only)
DELETE /delete    (connected only)
GET /introspect   (connected only)	
GET /users/       (connected only)
GET /users/<id>   (connected only)
PUT /users/<id>   (connected only)	#body {username, password}


GET /songs/       (connected only)
GET /songs/<id>   (connected only)
POST /songs       (#connected only)   #body {songname, songauthor, songgenre}  (working / can be improved)
PUT /songs/<id>   (#connected only)   #body {songname, songauthor, songgenre}  (working / can be improved)
DELETE /songs/<id>(#connected only) 						  (working / can be improved)

GET /songs/<id>/ratings
GET /songs/<id_song>/ratings/<id_rating>
POST /songs/<id_song>/ratings/  
PUT  /songs/<id_song>/ratings/<id_rating> 
DELETE /songs/<id_song>/ratings/<id_rating> 
## Debug request (delete at the end)
DELETE /users/<id>
