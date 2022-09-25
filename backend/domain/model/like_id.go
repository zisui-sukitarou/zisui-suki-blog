package model

type LikeId int

func NewLikeId(likeId int) (LikeId, error){
	return LikeId(likeId), nil
}