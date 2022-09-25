package blogapp

type BlogFindByIdRequest struct {
	Id int `json:"id"`
}

type BlogFindByTagIdRequest struct {
	TagId string `json:"tag_id"`
}

type BlogFindByUserIdRequest struct {
	UserId int `json:"user_id"`
	UpTo   int `json:"up_to"`
}

type BlogDeleteRequest struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
}

type BlogUpdateRequest struct {
	Id       int    `json:"id"`
	Content  string `json:"content"`
	Abstract string `json:"abstract"`
	Title    string `json:"title"`
}

type BlogRegisterRequest struct {
	UserId   int    `json:"user_id"`
	Content  string `json:"content"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
}
