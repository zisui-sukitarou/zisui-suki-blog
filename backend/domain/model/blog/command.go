package model

import "errors"

/* create new blog */
type commandNewBlog struct {
	userId   BlogUserId
	content  BlogContent
	title    BlogTitle
	abstract BlogAbstract
	view     BlogView
}

func CreateCommandNewBlog(
	userId int,
	content string,
	title string,
	abstract string,
	view int,
) (
	*commandNewBlog,
	error,
) {
	blogUserId, err := NewBlogUserId(userId)
	if err != nil {
		return &commandNewBlog{}, errors.New("user_id is illegal")
	}

	blogContent, err := NewBlogContent(content)
	if err != nil {
		return &commandNewBlog{}, errors.New("content is illegal")
	}

	blogTitle, err := NewBlogTitle(title)
	if err != nil {
		return &commandNewBlog{}, errors.New("title is illegal")
	}

	blogAbstract, err := NewBlogAbstract(abstract)
	if err != nil {
		return &commandNewBlog{}, errors.New("abstract is illegal")
	}

	blogView, err := NewBlogView(view)
	if err != nil {
		return &commandNewBlog{}, errors.New("view is illegal")
	}

	return &commandNewBlog{
		userId:   blogUserId,
		content:  blogContent,
		title:    blogTitle,
		abstract: blogAbstract,
		view:     blogView,
	}, nil
}

/* create new blog */
type commandUpdateBlog struct {
	content  BlogContent
	title    BlogTitle
	abstract BlogAbstract
}

func CreateCommandUpdateBlog(
	content string,
	title string,
	abstract string,
) (
	*commandUpdateBlog,
	error,
) {
	blogContent, err := NewBlogContent(content)
	if err != nil {
		return &commandUpdateBlog{}, errors.New("content is illegal")
	}

	blogTitle, err := NewBlogTitle(title)
	if err != nil {
		return &commandUpdateBlog{}, errors.New("title is illegal")
	}

	blogAbstract, err := NewBlogAbstract(abstract)
	if err != nil {
		return &commandUpdateBlog{}, errors.New("abstract is illegal")
	}

	return &commandUpdateBlog{
		content:  blogContent,
		title:    blogTitle,
		abstract: blogAbstract,
	}, nil
}