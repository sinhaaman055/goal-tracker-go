package auth

import "net/http"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) SignupHandler(w http.ResponseWriter,r *http.Request){

//abhii code likha jyega guys

}
func (h *Handler) LoginHandler(w http.ResponseWriter,r *http.Request){

	//abhi code likha jyega Guyz

}
func (h *Handler) ForgetPassword(w http.ResponseWriter,r *http.Request){

   //likha jyega code 

}
func (h *Handler) ChangePasswordVerify(w http.ResponseWriter,r *http.Request){

	//abhi code likha jyega guys

}
func (h *Handler) ChangePassword(w http.ResponseWriter,r *http.Request){

//abhi code likha jyega guys

}