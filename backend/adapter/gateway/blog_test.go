package gateway_test

import (
	"log"
	"testing"
	"time"
	"zisui-suki-blog/adapter/gateway"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/service"
	"zisui-suki-blog/infrastructure/db"
)

func TestBlogRegister(t *testing.T) {
	blogId, err := newBlogService().GenULID()
	if err != nil {
		t.Errorf(err.Error())
	}

	blog := model.NewBlog(
		model.BlogId(blogId.String()),
		model.UserId("zisui-sukitarou"),
		model.BlogContent("<h1>World</h1>"),
		model.BlogTitle("Pasta"),
		model.BlogAbstract("Very Delicious"),
		model.BlogEvaluation(8),
		time.Now(),
		time.Now(),
	)

	infra, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}

	err = gateway.NewBlogGateway(infra.Client, &infra.Ctx).Register(blog)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestFindBlogById(t *testing.T) {
	infra, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}

	_, blog, err := gateway.NewBlogGateway(infra.Client, &infra.Ctx).FindById(
		model.BlogId("01GFDYE3JGSZ8EG6RB9AP55GCF"),
	)
	if err != nil {
		t.Errorf(err.Error())
	}
	log.Println("blog:", blog)
}

func TestRegisterTags(t *testing.T) {
	infra, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}

	err = gateway.NewBlogGateway(infra.Client, &infra.Ctx).RegisterTags(
		model.BlogId("01GFFVN4Y993A1PA5M7TN6DKVW"),
		[]model.TagName{
			model.TagName("pasta"),
		},
	)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func newBlogService() *service.Blog {
	return &service.Blog{}
}
