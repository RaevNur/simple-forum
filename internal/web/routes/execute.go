package routes

import (
	"fmt"
	"forum/internal/models"
	"net/http"
)

func (route *Routes) ExecuteErrorTemplate(forum models.Forum, errStatus int, w http.ResponseWriter) {
	forum.ErrStatus = errStatus
	forum.ErrMsg = http.StatusText(errStatus)
	w.WriteHeader(errStatus)

	tmpl, err := route.view.ErrorTemplate()
	if err != nil {
		fmt.Fprint(w, forum.ErrMsg)
		return
	}

	tmpl.ExecuteTemplate(w, "bootstrap", forum)
}

func (route *Routes) ExecuteTemplate(forum models.Forum, tmplName string, w http.ResponseWriter) {
	tmpl, err := route.view.NewTemplate(tmplName)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	tmpl.ExecuteTemplate(w, "bootstrap", forum)
}
