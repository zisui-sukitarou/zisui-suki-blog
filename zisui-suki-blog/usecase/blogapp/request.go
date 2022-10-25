package blogapp

type BlogFindByIdRequest struct {
	BlogId string `json:"blog_id"`
}

type BlogFindByTagRequest struct {
	TagName string `json:"tag_name"`
	Begin   uint   `json:"begin"`
	End     uint   `json:"end"`
}

type BlogFindByUserIdRequest struct {
	UserId string `json:"user_id"`
	Begin  uint   `json:"begin"`
	End    uint   `json:"end"`
}

type BlogFindByUserIdAndTagRequest struct {
	TagName string `json:"tag_name"`
	UserId  string `json:"user_id"`
	Begin   uint   `json:"begin"`
	End     uint   `json:"end"`
}

type BlogDeleteRequest struct {
	BlogId string `json:"blog_id"`
	UserId string `json:"user_id"`
}

type BlogUpdateRequest struct {
	BlogId     string `json:"blog_id"`
	Content    string `json:"content"`
	Abstract   string `json:"abstract"`
	Title      string `json:"title"`
	Evaluation uint   `json:"evaluation"`
}

type BlogRegisterRequest struct {
	UserId     string   `json:"user_id"`
	Content    string   `json:"content"`
	Title      string   `json:"title"`
	Abstract   string   `json:"abstract"`
	Evaluation uint     `json:"evaluation"`
	Status     uint     `json:"status"`
	Tags       []string `json:"tags"`
}
