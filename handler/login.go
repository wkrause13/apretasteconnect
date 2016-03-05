package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Login handles all user logins
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Login")
}
