package model

type Tag struct {
	*Model
	Name string `json:"name"`
	State uint8 `json:"state"`
}

func (t Tag) Create() string {
	return "blog_tag"
}