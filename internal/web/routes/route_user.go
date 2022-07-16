package routes

import (
	model "forum/internal/models"
	"log"
	"net/http"
)

func (route *Routes) UserLikedThreads(w http.ResponseWriter, r *http.Request) {
	forum := model.Forum{}

	GetUserBySession(route, &forum, w, r)

	if !forum.UserIn {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	threads, err := route.service.Thread.GetUserLikedThreads(forum.User.Id)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	forum.Threads = threads
	for i, thread := range forum.Threads {
		user, err := route.service.User.GetByID(thread.Post.UserId)
		if err != nil {
			thread.Author = &model.User{
				Nickname: "Unknown",
			}
		} else {
			thread.Author = user
		}

		tags, err := route.service.Tag.GetTagsByThread(thread.Id)
		if err == nil {
			thread.Tags = tags
		}

		err = route.service.Post.GetLikesCount(forum.Threads[i].Post)
		if err != nil {
			log.Println(err)
		}

		if forum.UserIn {
			err = route.service.Post.IsLikedByUser(forum.User.Id, forum.Threads[i].Post)
			if err != nil {
				log.Println(err)
			}
		}
	}
	route.ExecuteTemplate(forum, "home", w)
}

func (route *Routes) UserThreads(w http.ResponseWriter, r *http.Request) {
	forum := model.Forum{}

	GetUserBySession(route, &forum, w, r)

	if !forum.UserIn {
		http.Redirect(w, r, "/", http.StatusNotFound)
		return
	}

	threads, err := route.service.Thread.GetUserCreatedThreads(forum.User.Id)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	forum.Threads = threads
	for i, thread := range forum.Threads {
		user, err := route.service.User.GetByID(thread.Post.UserId)
		if err != nil {
			thread.Author = &model.User{
				Nickname: "Unknown",
			}
		} else {
			thread.Author = user
		}

		tags, err := route.service.Tag.GetTagsByThread(thread.Id)
		if err == nil {
			thread.Tags = tags
		}

		err = route.service.Post.GetLikesCount(forum.Threads[i].Post)
		if err != nil {
			log.Println(err)
		}

		if forum.UserIn {
			err = route.service.Post.IsLikedByUser(forum.User.Id, forum.Threads[i].Post)
			if err != nil {
				log.Println(err)
			}
		}
	}

	route.ExecuteTemplate(forum, "home", w)
}
