package service

import (
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
)

type Blog struct{}

func (b *Blog) Exists(blogId model.BlogId, repository repository.BlogRepository) (bool, error) {
	exists, _, err := repository.FindById(blogId)
	if err != nil {
		return false, err
	}
	return exists, nil
}
