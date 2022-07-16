package routes

import (
	model "forum/internal/models"
	"log"
	"net/http"
)

func (route *Routes) Home(w http.ResponseWriter, r *http.Request) {
	forum := model.Forum{}

	GetUserBySession(route, &forum, w, r)

	if r.URL.Path != "/" {
		route.ExecuteErrorTemplate(forum, http.StatusNotFound, w)
		return
	}

	if r.Method != http.MethodGet {
		route.ExecuteErrorTemplate(forum, http.StatusMethodNotAllowed, w)
		return
	}

	threads, err := route.service.Thread.GetRecentQuestions()
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
