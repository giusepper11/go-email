package campaign

import (
	"errors"
	"go-email/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"test1@abc.com", "test2@def.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repository_mock := new(repositoryMock)
	repository_mock.On("Save", mock.Anything).Return(nil)
	service.Repository = repository_mock
	id, err := service.Create(newCampaign)

	assert.NotEmpty(id)
	assert.Nil(err)
}

func Test_Create_Campaign_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""
	_, err := service.Create(newCampaign)

	assert.NotNil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	repository_mock := new(repositoryMock)
	repository_mock.On("Save", mock.MatchedBy(
		func(campaign *Campaign) bool {
			if campaign.Name != newCampaign.Name {
				return false
			}
			if campaign.Content != newCampaign.Content {
				return false
			}
			if len(campaign.Contacts) != len(newCampaign.Emails) {
				return false
			}
			return true
		})).Return(nil)

	service.Repository = repository_mock

	service.Create(newCampaign)

	repository_mock.AssertExpectations(t)
}

func Test_Create_Campaign_Validate_Repository_Save(t *testing.T) {
	assert := assert.New(t)
	repository_mock := new(repositoryMock)
	repository_mock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repository_mock
	_, err := service.Create(newCampaign)
	assert.Equal("error to save on database", err.Error())

}
