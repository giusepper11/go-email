package campaign

import (
	"errors"
	"go-email/internal/contract"
	internalerrors "go-email/internal/internal_errors"
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

func (r *repositoryMock) Get() ([]Campaign, error) {
	return nil, nil
}

var service = Service{}

func generateCampaignFixture() contract.NewCampaign {
	return contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body of campaign",
		Emails:  []string{"test1@abc.com", "test2@def.com"},
	}

}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repository_mock := new(repositoryMock)
	repository_mock.On("Save", mock.Anything).Return(nil)
	service.Repository = repository_mock
	id, err := service.Create(generateCampaignFixture())

	assert.NotEmpty(id)
	assert.Nil(err)
}

func Test_Create_Campaign_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign := generateCampaignFixture()
	newCampaign.Name = ""
	_, err := service.Create(newCampaign)

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	newCampaign := generateCampaignFixture()
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

	newCampaign := generateCampaignFixture()
	repository_mock := new(repositoryMock)
	repository_mock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repository_mock
	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternal, err))

}
