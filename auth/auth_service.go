package auth

import (
	"github.com/minisource/apiclients/auth/models"
	"github.com/minisource/common_go/http/helper"
)

var authService *AuthService

type AuthService struct {
	client *helper.APIClient
}

func NewAuthService(client *helper.APIClient) *AuthService {
	return &AuthService{
		client: client,
	}
}

func GetAuthService() *AuthService {
	return authService
}

func (s *AuthService) ValidateToken(req models.ValidateTokenReq) (*models.ValidateTokenRes, error) {
	data, err := s.client.MakeRequest("POST", ValidateToken, req)
	if err != nil {
		return nil, err
	}

	// Prepare a variable to hold the deserialized result (User in this case)
	var token *models.ValidateTokenRes

	// Use the GenericDeserializer to deserialize the response
	err = helper.DeserializeResponse(data, &token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
