package models

type Forum struct {
	User   *User
	UserIn bool

	Thread *Thread

	Threads []*Thread
	Tags    []*Tag

	ErrStatus int
	ErrMsg    string
}
