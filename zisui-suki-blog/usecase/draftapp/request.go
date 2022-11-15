package draftapp

type DraftFindByIdRequest struct {
	DraftId string `json:"draftId"`
}

type DraftFindByUserIdRequest struct {
	UserId string `query:"userId"`
	Begin  uint   `query:"begin"`
	End    uint   `query:"end"`
}

type DraftFindByUserNameRequest struct {
	UserName string `query:"userName"`
	Begin    uint   `query:"begin"`
	End      uint   `query:"end"`
}

type DraftDeleteRequest struct {
	DraftId string `json:"draftId"`
	UserId  string `json:"userId"`
}

type DraftUpdateRequest struct {
	DraftId    string   `json:"draftId"`
	Content    string   `json:"content"`
	Abstract   string   `json:"abstract"`
	Title      string   `json:"title"`
	Evaluation uint     `json:"evaluation"`
	Tags       []string `json:"tags"`
}

type DraftRegisterRequest struct {
	UserId     string   `json:"userId"`
	Content    string   `json:"content"`
	Title      string   `json:"title"`
	Abstract   string   `json:"abstract"`
	Evaluation uint     `json:"evaluation"`
	Status     uint     `json:"status"`
	Tags       []string `json:"tags"`
}
