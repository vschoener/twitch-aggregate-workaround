package service

import (
	"fmt"

	"github.com/wonderstream/twitch/core"
	"github.com/wonderstream/twitch/core/model"
)

// VideoService handles processes for the user channel
type VideoService struct {
	*core.Request
}

// NewVideoService constructor
func NewVideoService() VideoService {
	return VideoService{}
}

// GetVideosFromIDResult contains result from request of GetVideosFromID
type GetVideosFromIDResult struct {
	Total  int16 `json:"_total"`
	Videos []model.Video
}

// GetVideosFromID returns the videos list of the channel ID
func (s VideoService) GetVideosFromID(channelID int64, r *core.Request, total int) []model.Video {
	result := GetVideosFromIDResult{}
	r.SendRequest(fmt.Sprintf("%s/%d/videos?limit=%d", core.ChannelsURI, channelID, total), &result)

	return result.Videos
}
