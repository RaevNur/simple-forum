package models

type Thread struct {
	Id       int64
	Title    string
	Post     *Post
	Author   *User
	Tags     []*Tag
	Comments []*Comment
}

type IThreadRepo interface {
	Create(thread *Thread) error
	GetById(id int64) (*Thread, error)
	GetByUserId(userId int64) ([]*Thread, error)
	GetByLiked(userId int64) ([]*Thread, error)
	GetByTag(tagId int64) ([]*Thread, error)
	GetRecentQuestions() ([]*Thread, error)
	// for future pagination
	// GetRecentQuestions(page int) ([]*Thread, error)
	// GetByUserId(userId int64, page int) ([]*Thread, error)
	// GetByLiked(userId int64, page int) ([]*Thread, error)
	// GetByTag(tagId int64, page int) ([]*Thread, error)
}

type IThreadService interface {
	Create(thread *Thread) error
	GetById(id int64) (*Thread, error)
	GetUserCreatedThreads(userId int64) ([]*Thread, error)
	GetUserLikedThreads(userId int64) ([]*Thread, error)
	GetByTag(tagId int64) ([]*Thread, error)
	GetRecentQuestions() ([]*Thread, error)
	// for future pagination
	// GetRecentQuestions(page int) ([]*Thread, error)
	// GetUserCreatedThreads(userId int64, page int) ([]*Thread, error)
	// GetUserLikedThreads(userId int64, page int) ([]*Thread, error)
	// GetByTag(tagId int64, page int) ([]*Thread, error)
}
