package models

type Tag struct {
	Id    int64
	Name  string
	Count int
}

type ITagRepo interface {
	Create(tag *Tag) error
	CreateRelation(threadId int64, tag *Tag) error
	// for future pagination
	// GetTags(page int) ([]*Tag, error)
	GetTags() ([]*Tag, error)
	GetTagByName(name string) (*Tag, error)
	GetTagsByThread(threadId int64) ([]*Tag, error)
}

type ITagService interface {
	Create(threadId int64, tag string) error
	CreateTags(threadId int64, tags []string) error
	// for future pagination
	// GetTags(page int) ([]*Tag, error)
	GetTags() ([]*Tag, error)
	GetTagByName(name string) (*Tag, error)
	GetTagsByThread(threadId int64) ([]*Tag, error)
}
