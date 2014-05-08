package goker

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"time"
)

// ---------------------------------------------------
// User endpoint
// ---------------------------------------------------

func (u *User) GetAll(w rest.ResponseWriter, r *rest.Request) {
	users := []User{}
	GokerCtx.DB.Find(&users)
	w.WriteJson(&users)
}

func (u *User) Get(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	user := User{}
	if GokerCtx.DB.First(&user, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&user)
}

func (u *User) Post(w rest.ResponseWriter, r *rest.Request) {
	user := User{}
	err := r.DecodeJsonPayload(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := GokerCtx.DB.Save(&user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&user)
}

func (u *User) Put(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	user := User{}
	if GokerCtx.DB.First(&user, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := User{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Login = updated.Login
	user.Email = updated.Email
	user.Name = updated.Name
	user.UpdatedAt = time.Now()

	if err := GokerCtx.DB.Save(&user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&user)
}

func (u *User) Delete(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	user := User{}
	if GokerCtx.DB.First(&user, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	if err := GokerCtx.DB.Delete(&user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u *User) GetCups(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	user := User{}
	if GokerCtx.DB.First(&user, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	user.Cups = DBGetCupsForUser(user.Id)
	w.WriteJson(&user.Cups)
}
