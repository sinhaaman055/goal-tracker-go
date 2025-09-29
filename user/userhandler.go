package user

import "net/http"

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}
func (h *Handler) GetProfile(w http.ResponseWriter, r *http.Request){

//will write the code

}

func (h *Handler) UpdateProfile(w http.ResponseWriter, r *http.Request){

//will write the code

}

