package routes

import (
	"forum/internal/helper/constraints"
	"net/http"
	"time"

	model "forum/internal/models"
)

func GetUserBySession(route *Routes, forum *model.Forum, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == nil {
		session, err := route.service.Session.GetByUuid(cookie.Value)
		if err != nil || session == nil {
			DeleteSessionCookie(w, r)
		} else if session.CreatedAt.Add(constraints.CookieExpireTime).Before(time.Now()) {
			DeleteSessionCookie(w, r)
			route.service.Session.DeleteSession(session.Id)
		} else {
			user, err := route.service.User.GetByID(session.UserId)
			if err != nil || user == nil {
				DeleteSessionCookie(w, r)
				route.service.Session.DeleteSession(session.Id)
			} else {
				forum.User = user
				forum.UserIn = true
			}
		}
	}
}

// sets the session cookie
func SetSessionCookie(uuid string, w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    uuid,
		MaxAge:   int(constraints.CookieExpireTime.Seconds()),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, cookie)
}

// deletes the session cookie
func DeleteSessionCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return
	}

	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}
