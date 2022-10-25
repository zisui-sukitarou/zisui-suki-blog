package userapp

type Status struct{}

func (s *Status) OK() uint {
	return 0
}

func (s *Status) UserNotExists() uint {
	return 10
}

func (s *Status) InvalidPassword() uint {
	return 11
}

func (s *Status) UserAlreadyExists() uint {
	return 20
}