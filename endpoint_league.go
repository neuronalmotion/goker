package goker

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"time"
)

// ---------------------------------------------------
// Cup endpoint
// ---------------------------------------------------

func (u *Cup) GetAll(w rest.ResponseWriter, r *rest.Request) {
	cups := []Cup{}
	GokerCtx.DB.Find(&cups)
	for i := 0; i < len(cups); i++ {
		FillCupData(&cups[i])
	}
	w.WriteJson(&cups)
}

func (u *Cup) Get(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	cup := Cup{}
	if GokerCtx.DB.First(&cup, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	FillCupData(&cup)
	w.WriteJson(&cup)
}

func (u *Cup) Post(w rest.ResponseWriter, r *rest.Request) {
	cup := Cup{}
	err := r.DecodeJsonPayload(&cup)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := GokerCtx.DB.Save(&cup).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&cup)
}

func (u *Cup) Put(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	cup := Cup{}
	if GokerCtx.DB.First(&cup, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := Cup{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cup.Name = updated.Name
	cup.Users = updated.Users
	cup.Games = updated.Games
	cup.UpdatedAt = time.Now()

	if err := GokerCtx.DB.Save(&cup).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&cup)
}

func (u *Cup) Delete(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	cup := Cup{}
	if GokerCtx.DB.First(&cup, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	if err := GokerCtx.DB.Delete(&cup).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
