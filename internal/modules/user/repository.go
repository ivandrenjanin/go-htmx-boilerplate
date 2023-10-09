package user

import (
	"database/sql"
	"reflect"
)

type UserRepository interface {
	CreateUser(dto *CreateUserRequestDTO) error
	GetUser(id string) (*GetUserResponseDTO, error)
	GetUsersPaginated(page int, limit int) ([]*GetUserResponseDTO, error)
}

type userRepository struct {
	db *sql.DB
}

func (r *userRepository) CreateUser(dto *CreateUserRequestDTO) error {
	statement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`

	_, err := r.db.Exec(statement, dto.Name, dto.Email, dto.Password)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUser(id string) (*GetUserResponseDTO, error) {
	statement := `
	SELECT 
		id, 
		name, 
		email, 
		created_at, 
		updated_at
	FROM 
		users 
	WHERE 
		id = $1 
		AND 
		deleted_at IS NULL;`

	var user GetUserResponseDTO

	rows, err := r.db.Query(statement, id)

	for rows.Next() {
		err := scanData(rows, &user)
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUsersPaginated(page int, limit int) ([]*GetUserResponseDTO, error) {
	statement := `
	SELECT 
		id, 
		name, 
		email, 
		created_at, 
		updated_at
	FROM 
		users 
	WHERE 
		deleted_at IS NULL
	LIMIT $1 OFFSET $2;`

	var users []*GetUserResponseDTO

	rows, err := r.db.Query(statement, limit, (page-1)*limit)

	for rows.Next() {
		var user GetUserResponseDTO
		err := scanData(rows, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err != nil {
		return nil, err
	}

	return users, nil
}

func scanData(rows *sql.Rows, dest interface{}) error {
	s := reflect.ValueOf(dest).Elem()
	numCols := s.NumField()
	columns := make([]interface{}, numCols)

	for i := 0; i < numCols; i++ {
		columns[i] = s.Field(i).Addr().Interface()
	}

	err := rows.Scan(columns...)
	if err != nil {
		return err
	}

	return nil
}
