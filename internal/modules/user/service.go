package user

type UserService interface {
	CreateUser(dto *CreateUserRequestDTO) error
	GetUser(id string) (*GetUserResponseDTO, error)
	GetUsersPaginated(page int, limit int) ([]*GetUserResponseDTO, error)
}

type userService struct {
	repository UserRepository
}

func (s *userService) CreateUser(dto *CreateUserRequestDTO) error {
	return s.repository.CreateUser(dto)
}

func (s *userService) GetUser(id string) (*GetUserResponseDTO, error) {
	return s.repository.GetUser(id)
}

func (s *userService) GetUsersPaginated(page int, limit int) ([]*GetUserResponseDTO, error) {
	return s.repository.GetUsersPaginated(page, limit)
}
