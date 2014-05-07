package gocker

import (
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"time"
)

// ---------------------------------------------------
// League endpoint
// ---------------------------------------------------

func (u *League) GetAll(w rest.ResponseWriter, r *rest.Request) {
	leagues := []League{}
	GockerCtx.DB.Find(&leagues)
	for i := 0; i < len(leagues); i++ {
		FillLeagueData(&leagues[i])
	}
	w.WriteJson(&leagues)
}

func (u *League) Get(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	league := League{}
	if GockerCtx.DB.First(&league, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	FillLeagueData(&league)
	w.WriteJson(&league)
}

func (u *League) Post(w rest.ResponseWriter, r *rest.Request) {
	league := League{}
	err := r.DecodeJsonPayload(&league)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := GockerCtx.DB.Save(&league).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&league)
}

func (u *League) Put(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	league := League{}
	if GockerCtx.DB.First(&league, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	updated := League{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	league.Name = updated.Name
	league.Users = updated.Users
	league.Games = updated.Games
	league.UpdatedAt = time.Now()

	if err := GockerCtx.DB.Save(&league).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&league)
}

func (u *League) Delete(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	league := League{}
	if GockerCtx.DB.First(&league, id).Error != nil {
		rest.NotFound(w, r)
		return
	}

	if err := GockerCtx.DB.Delete(&league).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
