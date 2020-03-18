package users

/*
curl -X POST -d '{"first_name": "JC", "last_name": "Gomez", "email": "jcm@gmail.com"}' http://localhost:8080/users

{
	"first_name": "JC",
	"last_name": "Gomez",
	"email": "jcm@gmail.com"
}
*/

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
}
