package model

type Blog struct {
	id       BlogId
	userId   BlogUserId
	content  BlogContent
	title    BlogTitle
	abstract BlogAbstract
	view     BlogView
}

/* constructor */
func NewBlog(command *commandNewBlog) (*Blog, error) {
	return &Blog{
		userId:   command.userId,
		content:  command.content,
		title:    command.title,
		abstract: command.abstract,
		view:     command.view,
	}, nil
}

/* change content */
func (b *Blog) Update(command *commandUpdateBlog) (*Blog, error) {
	b.content = command.content
	b.title = command.title
	b.abstract = command.abstract
	return b, nil
}

/* increment view */
func (b *Blog) IncView() *Blog {
	b.view++
	return b
}
