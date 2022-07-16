package routes

import (
	"forum/internal/helper"
	"forum/internal/helper/constraints"
	model "forum/internal/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (route *Routes) CreateThreadPage(w http.ResponseWriter, r *http.Request) {
	forum := model.Forum{}

	GetUserBySession(route, &forum, w, r)

	if !forum.UserIn {
		http.Redirect(w, r, "/signin", http.StatusSeeOther)
		return
	}

	if r.Method != http.MethodGet {
		route.ExecuteErrorTemplate(forum, http.StatusMethodNotAllowed, w)
		return
	}

	route.ExecuteTemplate(forum, "createthread", w)
}

func (route *Routes) CreateThread(w http.ResponseWriter, r *http.Request) {
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

	title := r.FormValue("title")
	content := r.FormValue("text")
	if !helper.ValidateNonEmpty(title, content) {
		forum.ErrStatus = http.StatusBadRequest
		err := &constraints.ValidateError{
			Field:       "title",
			Description: "title cant be empty",
		}
		forum.ErrMsg = err.Error()
		route.ExecuteTemplate(forum, "thread", w)
		return
	}

	tags := r.MultipartForm.Value["addmore[]"]
	if !helper.ValidateTags(tags...) {
		forum.ErrStatus = http.StatusBadRequest
		err := &constraints.ValidateError{
			Field:       "tags",
			Description: "tags cant be empty, have spaces and longer than 40 symbols",
		}
		forum.ErrMsg = err.Error()
		route.ExecuteTemplate(forum, "thread", w)
		return
	}

	post := model.Post{
		Content:   content,
		UserId:    forum.User.Id,
		CreatedAt: time.Now(),
	}

	// TODO transaction|rollback
	err := route.service.Post.Create(&post)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	thread := model.Thread{
		Title: title,
		Post:  &post,
	}

	err = route.service.Thread.Create(&thread)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	err = route.service.Tag.CreateTags(thread.Id, tags)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (route *Routes) ThreadPage(w http.ResponseWriter, r *http.Request) {
	forum := model.Forum{}

	GetUserBySession(route, &forum, w, r)

	if r.Method != http.MethodGet {
		route.ExecuteErrorTemplate(forum, http.StatusMethodNotAllowed, w)
		return
	}

	uri := r.URL.RequestURI()
	if uri == "/thread/" {
		route.ExecuteErrorTemplate(forum, http.StatusNotFound, w)
		return
	}

	reqThread := strings.Split(uri, "/")[2]
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
		route.ExecuteErrorTemplate(forum, http.StatusNotFound, w)
		return
	}

	thread.Author, err = route.service.User.GetByID(thread.Post.UserId)
	if err != nil || thread.Author == nil {
		thread.Author = &model.User{
			Id:       0,
			Nickname: "Unknown",
		}
		return
	}

	thread.Tags, err = route.service.Tag.GetTagsByThread(thread.Id)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	err = route.service.Post.GetLikesCount(thread.Post)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	if forum.UserIn {
		err = route.service.Post.IsLikedByUser(forum.User.Id, thread.Post)
		if err != nil {
			route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
			return
		}
	}

	thread.Comments, err = route.service.Comment.GetThreadComments(thread.Id)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	for i := range thread.Comments {
		thread.Comments[i].Author, err = route.service.User.GetByID(thread.Comments[i].Post.UserId)
		if err != nil || thread.Comments[i].Author == nil {
			thread.Comments[i].Author = &model.User{
				Id:       0,
				Nickname: "Unknown",
			}
			return
		}

		err = route.service.Post.GetLikesCount(thread.Comments[i].Post)
		if err != nil {
			route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
			return
		}

		if forum.UserIn {
			err = route.service.Post.IsLikedByUser(forum.User.Id, thread.Comments[i].Post)
			if err != nil {
				route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
				return
			}
		}
	}

	forum.Thread = thread

	route.ExecuteTemplate(forum, "thread", w)
}
