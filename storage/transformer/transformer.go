package transformer

import (
	"strings"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/model"
	storageModel "github.com/wonderstream/twitch/storage/model"
)

// Transformer manager
type Transformer struct {
}

// TransformCoreChannelToStorageChannel transform model from Core to Storage
func TransformCoreChannelToStorageChannel(c model.Channel) storageModel.Channel {
	return storageModel.Channel{
		Mature:               c.Mature,
		Status:               c.Status,
		BroadcasterLanguage:  c.BroadcasterLanguage,
		DisplayName:          c.DisplayName,
		Game:                 c.Game,
		Language:             c.Language,
		ChannelID:            c.IDTwitch,
		Name:                 c.Name,
		CreatedAt:            c.CreatedAt,
		UpdatedAt:            c.UpdatedAt,
		Partner:              c.Partner,
		VideoBanner:          c.VideoBanner,
		ProfileBanner:        c.ProfileBanner,
		ProfileBannerBGColor: c.ProfileBannerBGColor,
		URL:                  c.URL,
		Views:                c.Views,
		Followers:            c.Followers,
		BroadcasterType:      c.BroadcasterType,
		StreamKey:            c.StreamKey,
		Email:                c.Email,
	}
}

// TransformCoreVideoToStorageVideo transform model from Core to Storage
func TransformCoreVideoToStorageVideo(v model.Video) storageModel.ChannelVideo {
	return storageModel.ChannelVideo{
		Title:           v.Title,
		Description:     v.Description,
		DescriptionHTML: v.DescriptionHTML,
		BroadcastID:     v.BroadcastID,
		BroadcastType:   v.BroadcastType,
		Status:          v.Status,
		TagList:         v.TagList,
		Views:           v.Views,
		URL:             v.URL,
		Language:        v.Language,
		CreatedAt:       v.CreatedAt,
		Viewable:        v.Viewable,
		ViewableAt:      v.ViewableAt,
		PublishedAt:     v.PublishedAt,
		VideoID:         v.VideoID,
		RecordedAt:      v.RecordedAt,
		Game:            v.Game,
		Length:          v.Length,
	}
}

// TransformCoreUserToStorageUser transform model from Core to Storage
func TransformCoreUserToStorageUser(u model.User) storageModel.User {
	return storageModel.User{
		DisplayName: u.DisplayName,
		UserID:      u.UserID,
		Name:        u.Name,
		Type:        u.Type,
		Bio:         u.Bio,
		Logo:        u.Logo,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
	}
}

// TransformCoreTokenResponseToStorageCredential transform TokenResponse to Storage Token
func TransformCoreTokenResponseToStorageCredential(t core.TokenResponse) storageModel.Credential {
	return storageModel.Credential{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		ExpiresIn:    t.ExpiresIn,
		Scopes:       strings.Join(t.Scopes, " "),
	}
}

// TransformStorageCredentialToCoreTokenResponse transform TokenResponse to Storage Token
func TransformStorageCredentialToCoreTokenResponse(t storageModel.Credential) core.TokenResponse {
	return core.TokenResponse{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		ExpiresIn:    t.ExpiresIn,
		Scopes:       strings.Split(t.Scopes, " "),
	}
}
