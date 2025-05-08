package bunnyrequestservice

import (
	"encoding/json"
	"fmt"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	bunnyrequestlibs "github.com/jSierraB3991/simple-bunny-api-req/domain/bunny_request_libs"
	bunnyrequestrequest "github.com/jSierraB3991/simple-bunny-api-req/infrastructure/bunny_request_request"
	bunnyrequestrest "github.com/jSierraB3991/simple-bunny-api-req/infrastructure/bunny_request_rest"
)

func (s *BunnyRequestService) CreateVideoService(titleVideo, collectionId, videoLibraryApiKey string, thumbnailInMs *int, libraryId uint) (*bunnyrequestrest.CreateVideoResponse, error) {
	var result bunnyrequestrest.CreateVideoResponse

	req := bunnyrequestrequest.CreateVideoRequest{
		Title:         titleVideo,
		CollectionID:  collectionId,
		ThumbnailTime: thumbnailInMs,
	}
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	headers := []jsierralibs.HeaderRequest{s.GetHeaderApiKey(videoLibraryApiKey)}

	err = jsierralibs.Post(s.videoCdnUrl, fmt.Sprintf(bunnyrequestlibs.CREATE_VIDEO_URL, libraryId), jsonData, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
