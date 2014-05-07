package gocker

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
	GockerCtx.DB.Find(&users)
	w.WriteJson(&users)
}

func (u *User) Get(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	user := User{}
	if GockerCtx.DB.First(&user, id).Error != nil {
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
	if err := GockerCtx.DB.Save(&user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&user)
}

func (u *User) Put(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	user := User{}
	if GockerCtx.DB.First(&user, id).Error != nil {
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

	if err := GockerCtx.DB.Save(&user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&user)
}

func (u *User) Delete(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	user := User{}
	if GockerCtx.DB.First(&user, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	if err := GockerCtx.DB.Delete(&user).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u *User) GetLeagues(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
	user := User{}
	if GockerCtx.DB.First(&user, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

    user.Leagues = DBGetLeaguesForUser(user.Id)
    w.WriteJson(&user.Leagues)
}
