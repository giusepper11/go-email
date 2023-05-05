package campaign

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body of the campaign"
	contacts = []string{"abc@abc.com", "cde@cde.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	// arrange
	assert := assert.New(t)

	// act
	campaign, _ := NewCampaign(name, content, contacts)
	fmt.Println(campaign)

	//assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	for ix, email := range contacts {
		assert.Equal(campaign.Contacts[ix].Email, email)
	}

}

func Test_NewCampaignIDNotNill(t *testing.T) {
	// arrange
	assert := assert.New(t)

	// act
	campaign, _ := NewCampaign(name, content, contacts)
	fmt.Println(campaign)

	// assert
	assert.NotNil(campaign.ID)
}

func Test_NewCampaignCreatedAtNotNillAndMustBeNow(t *testing.T) {
	// arrange
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute * 1)

	// act
	campaign, _ := NewCampaign(name, content, contacts)
	fmt.Println(campaign)

	// assert
	assert.NotNil(campaign.CreatedAt)
	assert.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaignMustValidateName(t *testing.T) {
	// arrange
	assert := assert.New(t)

	// act
	_, err := NewCampaign("", content, contacts)

	// assert
	assert.Error(err)
}

func Test_NewCampaignMustValidateContent(t *testing.T) {
	// arrange
	assert := assert.New(t)

	// act
	_, err := NewCampaign(name, "", contacts)

	// assert
	assert.Error(err)
}

func Test_NewCampaignMustValidateContacts(t *testing.T) {
	// arrange
	assert := assert.New(t)

	// act
	_, err := NewCampaign(name, content, []string{})

	// assert
	assert.Error(err)
}
