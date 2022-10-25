package model

type TagIcon string

func NewTagIcon(icon string) (TagIcon, error) {
	return TagIcon(icon), nil
}
