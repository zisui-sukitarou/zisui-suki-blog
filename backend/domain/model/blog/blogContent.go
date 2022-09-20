package model

import "errors"

type BlogContent string

func NewBlogContent(content string) (BlogContent, error){
	if len(content) <= 0 {
		return BlogContent(""), errors.New("blog content is empty")
	}

	return BlogContent(content), nil
}