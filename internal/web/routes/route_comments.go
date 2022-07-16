package routes

import (
	"forum/internal/helper"
	model "forum/internal/models"
	"net/http"
	"strconv"
	"time"
)

func (route *Routes) CreateComment(w http.ResponseWriter, r *http.Request) {
	forum := model.Forum{}

	GetUserBySession(route, &forum, w, r)

	if !forum.UserIn {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodPost {
		route.ExecuteErrorTemplate(forum, http.StatusMethodNotAllowed, w)
		return
	}

	comm := r.FormValue("comment")
	if !helper.ValidateNonEmpty(comm) {
		route.ExecuteErrorTemplate(forum, http.StatusBadRequest, w)
		return
	}

	reqThread := r.FormValue("threadId")
	threadId, err := strconv.ParseInt(reqThread, 10, 64)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusBadRequest, w)
		return
	}

	thread, err := route.service.Thread.GetById(threadId)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	if thread == nil {
		route.ExecuteErrorTemplate(forum, http.StatusBadRequest, w)
		return
	}

	post := model.Post{
		Content:   comm,
		UserId:    forum.User.Id,
		CreatedAt: time.Now(),
	}

	err = route.service.Post.Create(&post)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	comment := model.Comment{
		Post:     &post,
		ThreadId: threadId,
	}

	err = route.service.Comment.Comment(&comment)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}
