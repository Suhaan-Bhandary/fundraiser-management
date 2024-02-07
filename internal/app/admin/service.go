package admin

// struct containing all the dependencies need for the methods
type service struct{}

type Service interface{}

func NewService() Service {
	return &service{}
}
