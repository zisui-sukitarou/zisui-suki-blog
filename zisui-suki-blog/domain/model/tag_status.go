package model

type TagStatus uint

func NewTagStatus(status uint) (TagStatus, error) {
	return TagStatus(status), nil
}