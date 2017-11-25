package credential

import (
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/storage"
)

type FakeLoader struct {
	err        error
	AppSetting AppSetting
}

func (f *FakeLoader) Load(path string, definition interface{}) error {
	f.AppSetting = AppSetting{
		TwitchSettings: core.TwitchSettings{
			ClientID:     "cliendID",
			ClientSecret: "clientSecret",
			TwitchRequestSettings: core.TwitchRequestSettings{
				URL: "http://api.twitch.com",
				Headers: map[string]string{
					"Header": "Foo",
				},
			},
			RedirectURL:        "http://redirect.somewhere",
			ErrorRedirectURL:   "http://redirect.error",
			SuccessRedirectURL: "http://redirect.success",
		},
		Databases: map[DBName]storage.DatabaseSettings{
			DBAggregation: storage.DatabaseSettings{
				User:     "john",
				Password: "doe42",
				URL:      "host",
				Port:     "42",
				Name:     "database",
			},
			DBActivity: storage.DatabaseSettings{
				User:     "johnA",
				Password: "doe42A",
				URL:      "hostA",
				Port:     "42A",
				Name:     "databaseA",
			},
		},
	}

	return f.err
}

// TestLoadingSetting only test if all field are present
func TestLoadingSetting(t *testing.T) {
	fakeLoader := &FakeLoader{}
	log.Println("Load Setting")
	credential := NewCredential(fakeLoader, "")
	credential.LoadSetting()
	credential.AppSetting = fakeLoader.AppSetting

	assert.NotEmpty(t, credential.TwitchSettings, "Should not be tempty")
}

func TestGetTwitch(t *testing.T) {
	fakeLoader := &FakeLoader{}
	credential := NewCredential(fakeLoader, "")
	credential.LoadSetting()
	credential.AppSetting = fakeLoader.AppSetting
	twitchSetting := credential.GetTwitch()

	assert.Equal(t, "cliendID", twitchSetting.ClientID, "They should be equal")
	assert.Equal(t, "clientSecret", twitchSetting.ClientSecret, "They should be equal")
	assert.Equal(t, "http://api.twitch.com", twitchSetting.TwitchRequestSettings.URL, "They should be equal")
	assert.Equal(t, "Foo", twitchSetting.TwitchRequestSettings.Headers["Header"], "They should be equal")
	assert.Equal(t, "http://redirect.somewhere", twitchSetting.RedirectURL, "They should be equal")
	assert.Equal(t, "http://redirect.error", twitchSetting.ErrorRedirectURL, "They should be equal")
	assert.Equal(t, "http://redirect.success", twitchSetting.SuccessRedirectURL, "They should be equal")
}

func TestGetDB(t *testing.T) {
	fakeLoader := &FakeLoader{}
	credential := NewCredential(fakeLoader, "")
	credential.LoadSetting()
	credential.AppSetting = fakeLoader.AppSetting
	DBSetting := credential.GetDB(DBAggregation)

	assert.Equal(t, "john", DBSetting.User, "They should be equal")
	assert.Equal(t, "doe42", DBSetting.Password, "They should be equal")
	assert.Equal(t, "host", DBSetting.URL, "They should be equal")
	assert.Equal(t, "database", DBSetting.Name, "They should be equal")

	DBSetting = credential.GetDB(DBActivity)

	assert.Equal(t, "johnA", DBSetting.User, "They should be equal")
	assert.Equal(t, "doe42A", DBSetting.Password, "They should be equal")
	assert.Equal(t, "hostA", DBSetting.URL, "They should be equal")
	assert.Equal(t, "databaseA", DBSetting.Name, "They should be equal")
}

func (f FakeLoader) isAppSettingEmpty(s core.TwitchSettings) bool {
	return reflect.DeepEqual(s, core.TwitchSettings{})
}

func TestSettingFromDistParametersFile(t *testing.T) {
	f := FakeLoader{}
	credential := NewCredential(YAMLLoader{}, "../parameters.yml.dist")
	credential.LoadSetting()
	// Small test but maybe  stronger one is required
	assert.Equal(t, false, f.isAppSettingEmpty(credential.TwitchSettings), "Should not be empty")
}
