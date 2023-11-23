package service

import (
	"EspumaBlancaBackend/config"
	"EspumaBlancaBackend/data/request"
	"EspumaBlancaBackend/data/response"
	"EspumaBlancaBackend/models"
	"EspumaBlancaBackend/repository"
	"errors"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	Validate       *validator.Validate
}

func (u UserServiceImpl) SignIn(loginRequest request.LoginRequest) (response.UserResponse, error) {
	var userResponse response.UserResponse
	err := u.Validate.Struct(loginRequest)
	if err != nil {
		return userResponse, errors.New("error to sign in")
	}
	user, err := u.UserRepository.FindByEmail(loginRequest.Email)
	if err != nil {
		return userResponse, errors.New("invalid email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return userResponse, errors.New("invalid password")
	}
	roles := make([]response.RoleResponse, len(user.Roles))
	for i, role := range user.Roles {
		roles[i] = response.RoleResponse{
			Id:   int(role.ID),
			Name: role.Name,
		}
	}
	userResponse = response.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
		Roles:    roles,
	}
	return userResponse, nil
}

func (u UserServiceImpl) Create(user request.UserRequest) {
	err := u.Validate.Struct(user)
	if err != nil {
		panic("error to create the user")
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		panic("error to encrypt the password")
	}
	roleRepository := repository.NewRoleRepositoryImpl(config.DatabaseConnection())
	var roles []models.Role
	for _, role := range user.Roles {
		role, err := roleRepository.FindById(role)
		if err != nil {
			panic("error to find the role")
		}
		roles = append(roles, role)
	}

	userModel := models.User{
		Email:    user.Email,
		FullName: user.FullName,
		Password: string(password),
		Roles:    roles,
	}
	u.UserRepository.Save(userModel)
}

func (u UserServiceImpl) Update(user request.UserRequest) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) Delete(id uint) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) FindById(id uint) response.UserResponse {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) FindAll() []response.UserResponse {
	//TODO implement me
	panic("implement me")
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		Validate:       validate,
	}
}
