package model

import "errors"

type BlogTitle string

func NewBlogTitle(title string) (BlogTitle, error){
	if len(title) <= 0 {
		return BlogTitle(""), errors.New("title is empty")
	}

	return BlogTitle(title), nil
}