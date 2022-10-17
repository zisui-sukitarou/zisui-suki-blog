package gateway_test

import (
	"log"
	"testing"
	"time"
	"zisui-suki-blog/adapter/gateway"
	"zisui-suki-blog/domain/model"
	"zisui-suki-blog/domain/repository"
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

/* find by user_id */
func findByUserId(db *db.DB, userId model.UserId, begin uint, end uint) ([]*repository.BlogOverviewData, error) {
	return gateway.NewBlogGateway(db.Client, &db.Ctx).FindByUserId(userId, begin, end)
}

type findByUserIdTestData struct {
	userId  model.UserId
	begin   uint
	end     uint
}

func TestFindByUserId(t *testing.T) {
	data := []findByUserIdTestData{
		{
			userId: model.UserId("zisui-sukitarou"),
			begin: 0,
			end: 10,
		},
	}

	db, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, d := range data {
		blogs, err := findByUserId(db, d.userId, d.begin, d.end)
		if err != nil {
			t.Errorf(err.Error())
		}
		for i, blog := range blogs {
			log.Println("blog", i, ":")
			log.Println("- blog_id:", blog.BlogId)
			log.Println("- abstract:", blog.Abstract)
			log.Println("- writer:", blog.Writer.UserId)
			log.Println("- tags:")
			for it, tag := range blog.Tags {
				log.Println(" - name", it, ":", tag.TagName)
			}
		}
	}
}

/* find by user_id & tag_name */
func findByTagName(db *db.DB, tagName model.TagName, begin uint, end uint) ([]*repository.BlogOverviewData, error) {
	return gateway.NewBlogGateway(db.Client, &db.Ctx).FindByTagName(tagName, begin, end)
}

type findByTagNameTestData struct {
	tagName model.TagName
	begin   uint
	end     uint
}

func TestFindByTagName(t *testing.T) {
	data := []findByTagNameTestData{
		{
			tagName: model.TagName("pasta"),
			begin: 0,
			end: 10,
		},
		{
			tagName: model.TagName("ramen"),
			begin: 0,
			end: 10,
		},
	}

	db, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, d := range data {
		blogs, err := findByTagName(db, d.tagName, d.begin, d.end)
		if err != nil {
			t.Errorf(err.Error())
		}
		for i, blog := range blogs {
			log.Println("blog", i, ":")
			log.Println("- blog_id:", blog.BlogId)
			log.Println("- abstract:", blog.Abstract)
			log.Println("- writer:", blog.Writer.UserId)
			log.Println("- tags:")
			for it, tag := range blog.Tags {
				log.Println(" - name", it, ":", tag.TagName)
			}
		}
	}
}

/* find by user_id & tag_name */
func findByUserIdAndTagName(db *db.DB, userId model.UserId, tagName model.TagName, begin uint, end uint) ([]*repository.BlogOverviewData, error) {
	return gateway.NewBlogGateway(db.Client, &db.Ctx).FindByUserIdAndTagName(userId, tagName, begin, end)
}

type findByUserIdAndTagNameTestData struct {
	userId  model.UserId
	tagName model.TagName
	begin   uint
	end     uint
}

func TestFindByUserIdAndTagName(t *testing.T) {
	data := []findByUserIdAndTagNameTestData{
		{
			userId: model.UserId("zisui-sukitarou"),
			tagName: model.TagName("pasta"),
			begin: 0,
			end: 10,
		},
		{
			userId: model.UserId("zisui-sukitarou"),
			tagName: model.TagName("ramen"),
			begin: 0,
			end: 10,
		},
	}

	db, err := db.Init()
	if err != nil {
		t.Errorf(err.Error())
	}

	for _, d := range data {
		blogs, err := findByUserIdAndTagName(db, d.userId, d.tagName, d.begin, d.end)
		if err != nil {
			t.Errorf(err.Error())
		}
		for i, blog := range blogs {
			log.Println("blog", i, ":")
			log.Println("- blog_id:", blog.BlogId)
			log.Println("- abstract:", blog.Abstract)
			log.Println("- writer:", blog.Writer.UserId)
			log.Println("- tags:")
			for it, tag := range blog.Tags {
				log.Println(" - name", it, ":", tag.TagName)
			}
		}
	}
}

/* service constructor */
func newBlogService() *service.Blog {
	return &service.Blog{}
}
