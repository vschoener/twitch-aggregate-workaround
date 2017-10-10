package transformer

import (
	"github.com/wonderstream/twitch/core/model"
	storageModel "github.com/wonderstream/twitch/storage/model"
)

// Transformer manager
type Transformer struct {
}

// TransformCoreChannelToStorageChannel transform model from Core to Storage
func TransformCoreChannelToStorageChannel(cc model.Channel) storageModel.Channel {
	return storageModel.Channel{
		Channel: cc,
	}
}

// TransformCoreVideoToStorageVideo transform model from Core to Storage
func TransformCoreVideoToStorageVideo(v model.Video) storageModel.ChannelVideo {
	return storageModel.ChannelVideo{
		Video: v,
	}
}

// TransformCoreUserToStorageUser transform model from Core to Storage
func TransformCoreUserToStorageUser(u model.User) storageModel.User {
	return storageModel.User{
		User: u,
	}
}
