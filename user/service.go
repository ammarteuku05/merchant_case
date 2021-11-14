package user

import (
	"errors"
	"fmt"
	"merchant-service/entity"
	"merchant-service/utils/formatter"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUser() ([]formatter.UserFormat, error)
	SaveNewUser(user entity.UserInput) (formatter.UserFormat, error)
	GetUserByID(userID string) (formatter.UserFormat, error)
	DeleteUserByID(userID string) (interface{}, error)
	UpdateUserByID(userID string, dataInput entity.UpdateUserInput) (formatter.UserFormat, error)
	LoginUser(input entity.LoginUserInput) (entity.User, error)
	CreateOutletUser(outlet entity.OutletInput, userID string) (entity.Outlet, error)
	ShowAllOutletUser() ([]formatter.OutletFormat, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) LoginUser(input entity.LoginUserInput) (entity.User, error) {
	user, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if user.Id == "" {
		newError := "user is not found"
		return user, errors.New(newError)
	}

	// pengecekan password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return user, errors.New("password invalid")
	}

	return user, nil
}

func (s *service) GetAllUser() ([]formatter.UserFormat, error) {
	users, err := s.repository.FindAll()
	var formatUsers []formatter.UserFormat

	for _, user := range users {
		formatUser := formatter.FormatUser(user)
		formatUsers = append(formatUsers, formatUser)
	}

	if err != nil {
		return formatUsers, err
	}

	return formatUsers, nil
}

func (s *service) SaveNewUser(user entity.UserInput) (formatter.UserFormat, error) {
	genPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)

	useruuid, err := uuid.NewV4()

	if err != nil {
		return formatter.UserFormat{}, err
	}

	var newUser = entity.User{
		Id:        useruuid.String(),
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  string(genPassword),
		Role:      "User",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createUser, err := s.repository.Create(newUser)
	formatUser := formatter.FormatUser(createUser)

	if err != nil {
		return formatUser, err
	}

	return formatUser, nil
}

func (s *service) GetUserByID(userID string) (formatter.UserFormat, error) {
	user, err := s.repository.FindByID(userID)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	var userData = entity.User{
		Id:        user.Id,
		FullName:  user.FullName,
		Email:     user.Email,
		Role:      user.Role,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if user.Id == "" {
		newError := fmt.Sprintf("user id %s not found", userID)
		return formatter.UserFormat{}, errors.New(newError)
	}

	formatUser := formatter.FormatUser(userData)

	return formatUser, nil
}

func (s *service) DeleteUserByID(userID string) (interface{}, error) {
	user, err := s.repository.FindByID(userID)

	if err != nil {
		return nil, err
	}
	if user.Id == "" {
		newError := fmt.Sprintf("user id %s not found", userID)
		return nil, errors.New(newError)
	}

	status, err := s.repository.DeleteByID(userID)

	if status == "error" {
		return nil, errors.New("error delete in internal server")
	}

	msg := fmt.Sprintf("success delete user ID : %s", userID)

	formatDelete := formatter.FormatDeleteUser(msg)

	return formatDelete, nil
}

func (s *service) UpdateUserByID(userID string, dataInput entity.UpdateUserInput) (formatter.UserFormat, error) {
	var dataUpdate = map[string]interface{}{}

	user, err := s.repository.FindByID(userID)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	if user.Id == "" {
		newError := fmt.Sprintf("user id %s not found", userID)
		return formatter.UserFormat{}, errors.New(newError)
	}

	if dataInput.FullName != "" || len(dataInput.FullName) != 0 {
		dataUpdate["full_name"] = dataInput.FullName
	}
	if dataInput.Email != "" || len(dataInput.Email) != 0 {
		dataUpdate["email"] = dataInput.Email
	}

	dataUpdate["updated_at"] = time.Now()

	userUpdated, err := s.repository.UpdateByID(userID, dataUpdate)

	if err != nil {
		return formatter.UserFormat{}, err
	}

	formatUser := formatter.FormatUser(userUpdated)

	return formatUser, nil
}

func (s *service) CreateOutletUser(outlet entity.OutletInput, userID string) (entity.Outlet, error) {

	checkStatus, err := s.repository.FindOutletUserByID(userID)

	if err != nil {
		return checkStatus, err
	}

	if checkStatus.UserId == userID {
		errorStatus := fmt.Sprintf("Outlet for user id : %s has been created", userID)
		return checkStatus, errors.New(errorStatus)
	}

	outletuuid, err := uuid.NewV4()

	if err != nil {
		return entity.Outlet{}, err
	}

	var newOutlet = entity.Outlet{
		Id:         outletuuid.String(),
		OutletName: outlet.OutletName,
		Picture:    outlet.Picture,
		UserId:     userID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	createOutlet, err := s.repository.CreateOutletUser(newOutlet)

	if err != nil {
		return createOutlet, err
	}

	return createOutlet, nil
}

func (s *service) ShowAllOutletUser() ([]formatter.OutletFormat, error) {
	outlet, err := s.repository.ShowAllOutletUser()
	var formatOutlet []formatter.OutletFormat

	for _, outlets := range outlet {
		formatOutlets := formatter.FormatOutlet(outlets)
		formatOutlet = append(formatOutlet, formatOutlets)

	}

	if err != nil {
		return formatOutlet, err
	}

	return formatOutlet, nil
}
