package routes

import (
	"errors"
	"forum/internal/helper/constraints"
	model "forum/internal/models"
	"net/http"
)

func (route *Routes) SignUpPage(w http.ResponseWriter, r *http.Request) {
	forum := model.Forum{}

	GetUserBySession(route, &forum, w, r)

	if forum.UserIn {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodGet {
		route.ExecuteErrorTemplate(forum, http.StatusMethodNotAllowed, w)
		return
	}

	route.ExecuteTemplate(forum, "signup", w)
}

func (route *Routes) SignUp(w http.ResponseWriter, r *http.Request) {
	forum := model.Forum{}

	GetUserBySession(route, &forum, w, r)

	if forum.UserIn {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user := model.User{
		Nickname:  r.FormValue("nickname"),
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}
	forum.User = &user

	if r.Method != http.MethodPost {
		route.ExecuteErrorTemplate(forum, http.StatusMethodNotAllowed, w)
		return
	}

	err := route.service.User.Register(forum.User)
	var validateError *constraints.ValidateError
	var existsError *constraints.ExistsError
	if err == nil {
	} else if errors.As(err, &validateError) {
		forum.ErrStatus = http.StatusBadRequest
		forum.ErrMsg = validateError.Error()
		route.ExecuteTemplate(forum, "signup", w)
		return
	} else if errors.As(err, &validateError) {
		forum.ErrStatus = http.StatusBadRequest
		forum.ErrMsg = validateError.Error()
		route.ExecuteTemplate(forum, "signup", w)
		return
	} else if errors.As(err, &validateError) {
		forum.ErrStatus = http.StatusBadRequest
		forum.ErrMsg = validateError.Error()
		route.ExecuteTemplate(forum, "signup", w)
		return
	} else if errors.As(err, &existsError) {
		forum.ErrStatus = http.StatusBadRequest
		forum.ErrMsg = existsError.Error()
		route.ExecuteTemplate(forum, "signin", w)
		return
	} else {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	session, err := route.service.Session.GenerateSession(forum.User.Id)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	SetSessionCookie(session.Uuid, w)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (route *Routes) SignInPage(w http.ResponseWriter, r *http.Request) {
	forum := model.Forum{}

	GetUserBySession(route, &forum, w, r)

	if forum.UserIn {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodGet {
		route.ExecuteErrorTemplate(forum, http.StatusMethodNotAllowed, w)
		return
	}

	route.ExecuteTemplate(forum, "login", w)
}

func (route *Routes) SignIn(w http.ResponseWriter, r *http.Request) {
	forum := model.Forum{}

	GetUserBySession(route, &forum, w, r)

	if forum.UserIn {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user := model.User{
		Nickname: r.FormValue("nickname"),
		Password: r.FormValue("password"),
	}
	forum.User = &user

	if r.Method != http.MethodPost {
		route.ExecuteErrorTemplate(forum, http.StatusMethodNotAllowed, w)
		return
	}

	err := route.service.User.Login(forum.User)
	var passErr *constraints.ValidateError
	var existsErr *constraints.ExistsError
	if errors.As(err, &passErr) {
		forum.ErrStatus = http.StatusBadRequest
		forum.ErrMsg = passErr.Error()
		route.ExecuteTemplate(forum, "login", w)
		return
	} else if errors.As(err, &existsErr) {
		forum.ErrStatus = http.StatusBadRequest
		forum.ErrMsg = existsErr.Error()
		route.ExecuteTemplate(forum, "signup", w)
		return
	} else if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	session, err := route.service.Session.GetByUserId(forum.User.Id)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	if session != nil {
		err = route.service.Session.DeleteSession(session.Id)
		if err != nil {
			route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
			return
		}
	}

	session, err = route.service.Session.GenerateSession(forum.User.Id)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	SetSessionCookie(session.Uuid, w)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (route *Routes) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err == nil {
		session, err := route.service.Session.GetByUuid(cookie.Value)
		if err != nil || session == nil {
			DeleteSessionCookie(w, r)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			DeleteSessionCookie(w, r)
			route.service.Session.DeleteSession(session.Id)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
