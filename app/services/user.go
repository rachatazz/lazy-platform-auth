package service

import (
	"fmt"
	"strings"
	"time"

	"lazy-platform-auth/app/adapter"
	entity "lazy-platform-auth/app/entities"
	model "lazy-platform-auth/app/models"
	repository "lazy-platform-auth/app/repositories"
	"lazy-platform-auth/config"
	"lazy-platform-auth/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(appId string, user model.CreateUserRequest) (*model.UserResponse, *fiber.Error)
	UpdateUser(id string, user model.UpdateUserRequest) (*model.UserResponse, *fiber.Error)
	ChangePassword(id string, user model.ChangePasswordRequest) *fiber.Error
	ForgotPassword(appId string, email string) *fiber.Error
	ResetPassword(resetPassReq model.ResetPasswordRequest) *fiber.Error
	GetUser(query model.UserQuery) ([]model.UserResponse, *fiber.Error)
	GetUserById(id string) (*model.UserResponse, *fiber.Error)
	Token(appId string, tokenReq model.TokenRequest) (*model.TokenResponse, *fiber.Error)
	RefreshToken(refeashTokenReq model.RefreshTokenRequest) (*model.TokenResponse, *fiber.Error)
	DeleteToken(id string) *fiber.Error
	GetTemplate() (*model.TemplateResponse, *fiber.Error)
}

type userService struct {
	config             config.Config
	appRepository      repository.AppRepository
	accountRepository  repository.UserRepository
	sessionRepository  repository.SessionRepository
	elasticMailAdapter adapter.ElasticEmailAdapter
}

func NewAccountService(
	config config.Config,
	appRepository repository.AppRepository,
	accountRepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
	elasticMailAdapter adapter.ElasticEmailAdapter,
) UserService {
	return userService{
		config:             config,
		appRepository:      appRepository,
		accountRepository:  accountRepository,
		sessionRepository:  sessionRepository,
		elasticMailAdapter: elasticMailAdapter,
	}
}

func (s userService) CreateUser(
	appId string,
	userReq model.CreateUserRequest,
) (*model.UserResponse, *fiber.Error) {
	app, err := s.appRepository.FindById(appId)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrUnprocessableEntity.Code, "app not found")
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), 10)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	user, err := s.accountRepository.Create(entity.UserEntity{
		AppID:       app.ID,
		DisplayName: userReq.DisplayName,
		FirstName:   userReq.FirstName,
		LastName:    userReq.LastName,
		PhoneNumber: userReq.PhoneNumber,
		Email:       userReq.Email,
		Password:    string(hasedPassword),
		VerifyFlag:  false,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, fiber.NewError(
				fiber.ErrUnprocessableEntity.Code,
				"email or phone number is already exists",
			)
		}
		return nil, fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	res := model.UserResponse{
		ID:                         user.ID.String(),
		AppID:                      user.AppID.String(),
		DisplayName:                user.DisplayName,
		FirstName:                  user.FirstName,
		LastName:                   user.LastName,
		Email:                      user.Email,
		VerifyFlag:                 user.VerifyFlag,
		RequiredChangePasswordFlag: user.RequiredChangePasswordFlag,
		CreatedAt:                  user.CreatedAt.String(),
		UpdatedAt:                  user.UpdatedAt.String(),
	}

	return &res, nil
}

func (s userService) UpdateUser(
	id string,
	userReq model.UpdateUserRequest,
) (*model.UserResponse, *fiber.Error) {
	preUpdateUser, err := s.accountRepository.FindById(id)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	if userReq.DisplayName != "" {
		preUpdateUser.DisplayName = userReq.DisplayName
	}
	if userReq.FirstName != "" {
		preUpdateUser.FirstName = userReq.FirstName
	}
	if userReq.LastName != "" {
		preUpdateUser.LastName = userReq.LastName
	}

	user, err := s.accountRepository.Update(*preUpdateUser)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	res := model.UserResponse{
		ID:                         user.ID.String(),
		AppID:                      user.AppID.String(),
		DisplayName:                user.DisplayName,
		FirstName:                  user.FirstName,
		LastName:                   user.LastName,
		Email:                      user.Email,
		VerifyFlag:                 user.VerifyFlag,
		RequiredChangePasswordFlag: user.RequiredChangePasswordFlag,
		CreatedAt:                  user.CreatedAt.String(),
		UpdatedAt:                  user.UpdatedAt.String(),
	}

	return &res, nil
}

func (s userService) ChangePassword(id string, userReq model.ChangePasswordRequest) *fiber.Error {
	user, err := s.accountRepository.FindById(id)
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.OldPassword))

	if err != nil {
		return fiber.NewError(
			fiber.ErrUnauthorized.Code,
			"email or password is incorrect",
		)
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.NewPassword), 10)
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	user.Password = string(hasedPassword)
	_, err = s.accountRepository.Update(*user)

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	return nil
}

func (s userService) ForgotPassword(appId string, email string) *fiber.Error {
	user, err := s.accountRepository.FindByEmail(appId, email)
	if err != nil {
		return fiber.NewError(
			fiber.ErrNotFound.Code,
			"email not found",
		)
	}

	tokenExp, _ := time.ParseDuration("24h")
	expiredInTime := time.Now().Add(tokenExp)
	user.TicketExpired = expiredInTime
	user.Ticket = "forgot_" + uuid.NewString()
	_, err = s.accountRepository.Update(*user)

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	// TODO: implement with send email

	return nil
}

func (s userService) ResetPassword(resetPassReq model.ResetPasswordRequest) *fiber.Error {
	user, err := s.accountRepository.FindByTicket(resetPassReq.Ticket)
	if err != nil {
		return fiber.NewError(
			fiber.ErrNotFound.Code,
			"email not found",
		)
	}

	today := time.Now()
	isExpired := today.After(user.TicketExpired)

	if isExpired {
		return fiber.NewError(
			fiber.ErrUnprocessableEntity.Code,
			"expired ticket",
		)
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(resetPassReq.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	user.Password = string(hasedPassword)
	user.Ticket = ""
	_, err = s.accountRepository.Update(*user)

	if err != nil {
		return fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	return nil
}

func (s userService) GetUserById(
	id string,
) (*model.UserResponse, *fiber.Error) {
	user, err := s.accountRepository.FindById(id)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	res := model.UserResponse{
		ID:                         user.ID.String(),
		AppID:                      user.AppID.String(),
		DisplayName:                user.DisplayName,
		FirstName:                  user.FirstName,
		LastName:                   user.LastName,
		Email:                      user.Email,
		VerifyFlag:                 user.VerifyFlag,
		RequiredChangePasswordFlag: user.RequiredChangePasswordFlag,
		CreatedAt:                  user.CreatedAt.String(),
		UpdatedAt:                  user.UpdatedAt.String(),
	}

	return &res, nil
}

func (s userService) GetUser(
	query model.UserQuery,
) ([]model.UserResponse, *fiber.Error) {
	users, err := s.accountRepository.Find(query)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	res := []model.UserResponse{}
	for _, user := range users {
		res = append(res, model.UserResponse{
			ID:                         user.ID.String(),
			AppID:                      user.AppID.String(),
			DisplayName:                user.DisplayName,
			FirstName:                  user.FirstName,
			LastName:                   user.LastName,
			Email:                      user.Email,
			VerifyFlag:                 user.VerifyFlag,
			RequiredChangePasswordFlag: user.RequiredChangePasswordFlag,
			CreatedAt:                  user.CreatedAt.String(),
			UpdatedAt:                  user.UpdatedAt.String(),
		})
	}

	return res, nil
}

func (s userService) Token(
	appId string,
	tokenReq model.TokenRequest,
) (*model.TokenResponse, *fiber.Error) {
	user, err := s.accountRepository.FindByEmailOrPhoneNumber(appId, tokenReq.Email, tokenReq.Email)
	if err != nil {
		return nil, fiber.NewError(
			fiber.ErrUnauthorized.Code,
			"email or password is incorrect",
		)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tokenReq.Password))

	if err != nil {
		return nil, fiber.NewError(
			fiber.ErrUnauthorized.Code,
			"email or password is incorrect",
		)
	}

	return s.buildToken(user.ID)
}

func (s userService) RefreshToken(
	refreshTokenReq model.RefreshTokenRequest,
) (*model.TokenResponse, *fiber.Error) {
	claims, ok := utils.ValidateToken(
		refreshTokenReq.RefreshToken,
		s.config.JwtRefreshTokenSecret,
	)
	if !ok {
		return nil, fiber.NewError(fiber.ErrUnauthorized.Code)
	}

	id := fmt.Sprint(claims["sub"])
	_, err := s.sessionRepository.FindByAccountId(id)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrUnauthorized.Code, "refresh token not found")
	}

	userId, err := uuid.Parse(id)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrInternalServerError.Code, "UUID is incorrect")
	}

	return s.buildToken(userId)
}

func (s userService) DeleteToken(id string) *fiber.Error {
	_, err := s.sessionRepository.Delete(id)
	if err != nil {
		return fiber.NewError(fiber.ErrUnauthorized.Code, "refresh token not found")
	}

	return nil
}

func (s userService) buildToken(id uuid.UUID) (*model.TokenResponse, *fiber.Error) {
	tokenExp, _ := time.ParseDuration(s.config.JwtTokenExpired)
	expiredInTime := time.Now().Add(tokenExp).Unix()
	token, err := utils.GenerateToken(
		id.String(),
		expiredInTime,
		s.config.JwtTokenSecret,
	)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	refreshTokenExp, _ := time.ParseDuration(
		s.config.JwtRefreshTokenExpired,
	)
	refreshTokenExpiredInTime := time.Now().Add(refreshTokenExp).Unix()
	refreshToken, err := utils.GenerateToken(
		id.String(),
		refreshTokenExpiredInTime,
		s.config.JwtRefreshTokenSecret,
	)
	if err != nil {
		return nil, fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	session, err := s.sessionRepository.Upsert(entity.RefreshTokenEntity{
		UserID:       id,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Unix(refreshTokenExpiredInTime, 0),
	})
	if err != nil {
		return nil, fiber.NewError(fiber.ErrInternalServerError.Code)
	}

	tokenResponse := model.TokenResponse{
		UserID:                id,
		AccessToken:           token,
		AccessTokenExpiresAt:  time.Unix(expiredInTime, 0),
		RefreshToken:          session.RefreshToken,
		RefreshTokenExpiresAt: time.Unix(refreshTokenExpiredInTime, 0),
	}

	return &tokenResponse, nil
}

func (s userService) GetTemplate() (*model.TemplateResponse, *fiber.Error) {
	res, err := s.elasticMailAdapter.GetEmailTemplate("verify_email")
	if err != nil {
		return nil, fiber.NewError(fiber.ErrUnprocessableEntity.Code)
	}
	return res, nil
}
