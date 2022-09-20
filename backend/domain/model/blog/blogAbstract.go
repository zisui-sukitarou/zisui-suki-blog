package model

import "errors"

type BlogAbstract string

func NewBlogAbstract(abstract string) (BlogAbstract, error){
	if len(abstract) <= 0 {
		return BlogAbstract(""), errors.New("abstract is empty")
	}

	return BlogAbstract(abstract), nil
}