package catalogue

import (
	"encoding/json"
	"net/http"

	"github.com/OPYS/internal/pkg/types"
)

type AuthControllers struct {
	srv *AuthService
}

func InitControllers(asrv *AuthService) *AuthControllers {
	return &AuthControllers{
		srv: asrv,
	}
}

func (c *AuthControllers) SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var us types.InfoForReg
		json.NewDecoder(r.Body).Decode(&us)

		err, newUser := c.srv.CreateUser(&us)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		json.NewEncoder(w).Encode(newUser)
		http.Redirect(w, r, "localhost/userspage", 301)

	}
}

func (c *AuthControllers) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var us types.UserLogin
		json.NewDecoder(r.Body).Decode(&us)
		err, usClaims := c.srv.TestuserForSignIn(&us)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		// w.Write([]byte("OK!"))
		json.NewEncoder(w).Encode(usClaims)
		http.Redirect(w, r, "localhost/userpage", 301)

	}
}
