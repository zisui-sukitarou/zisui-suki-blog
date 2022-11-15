package blogapp

type BlogFindByIdRequest struct {
	BlogId string `json:"blogId"`
}

type BlogFindByTagRequest struct {
	TagName string `json:"tagName"`
	Begin   uint   `json:"begin"`
	End     uint   `json:"end"`
}

type BlogFindByUserIdRequest struct {
	UserId string `query:"userId"`
	Begin  uint   `query:"begin"`
	End    uint   `query:"end"`
}

type BlogFindByUserNameRequest struct {
	UserName string `query:"userName"`
	Begin  uint   `query:"begin"`
	End    uint   `query:"end"`
}

type BlogFindByUserIdAndTagRequest struct {
	TagName string `json:"tagName"`
	UserId  string `json:"userId"`
	Begin   uint   `json:"begin"`
	End     uint   `json:"end"`
}

type BlogFindByUserNameAndTagRequest struct {
	TagName string `json:"tagName"`
	UserName  string `json:"userName"`
	Begin   uint   `json:"begin"`
	End     uint   `json:"end"`
}

type BlogDeleteRequest struct {
	BlogId string `json:"blogId"`
	UserId string `json:"userId"`
}

type BlogUpdateRequest struct {
	BlogId     string   `json:"blogId"`
	Content    string   `json:"content"`
	Abstract   string   `json:"abstract"`
	Title      string   `json:"title"`
	Evaluation uint     `json:"evaluation"`
	Tags       []string `json:"tags"`
}

type BlogRegisterRequest struct {
	UserId     string   `json:"userId"`
	Content    string   `json:"content"`
	Title      string   `json:"title"`
	Abstract   string   `json:"abstract"`
	Evaluation uint     `json:"evaluation"`
	Status     uint     `json:"status"`
	Tags       []string `json:"tags"`
}
