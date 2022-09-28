package blogapp

type BlogFindByIdRequest struct {
	BlogId int `json:"blog_id"`
}

type BlogFindByTagIdRequest struct {
	TagId string `json:"tag_id"`
}

type BlogFindByUserIdRequest struct {
	UserId string `json:"user_id"`
	UpTo   int    `json:"up_to"`
}

type BlogDeleteRequest struct {
	BlogId int    `json:"blog_id"`
	UserId string `json:"user_id"`
}

type BlogUpdateRequest struct {
	BlogId   int    `json:"blog_id"`
	Content  string `json:"content"`
	Abstract string `json:"abstract"`
	Title    string `json:"title"`
}

type BlogRegisterRequest struct {
	UserId   string `json:"user_id"`
	Content  string `json:"content"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
}
