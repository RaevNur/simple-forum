package routes

import (
	"forum/internal/helper/constraints"
	model "forum/internal/models"
	"net/http"
	"strconv"
)

func (route *Routes) Like(w http.ResponseWriter, r *http.Request) {
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

	postReq := r.FormValue("postId")
	postId, err := strconv.ParseInt(postReq, 10, 64)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusBadRequest, w)
		return
	}

	post, err := route.service.Post.GetById(postId)
	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	if post == nil {
		route.ExecuteErrorTemplate(forum, http.StatusBadRequest, w)
		return
	}

	like := model.Like{
		UserId: forum.User.Id,
		PostId: postId,
	}

	switch r.PostFormValue("selector") {
	case "like":
		like.Liked = constraints.LikeValue
		err = route.service.Like.Like(&like)
	case "clike":
		err = route.service.Like.Unlike(&like)
	case "dislike":
		like.Liked = constraints.DislikeValue
		err = route.service.Like.Dislike(&like)
	case "cdislike":
		err = route.service.Like.Unlike(&like)
	}

	if err != nil {
		route.ExecuteErrorTemplate(forum, http.StatusInternalServerError, w)
		return
	}

	http.Redirect(w, r, r.Referer(), http.StatusSeeOther)
}
