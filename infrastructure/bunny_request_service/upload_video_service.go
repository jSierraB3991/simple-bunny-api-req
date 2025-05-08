package bunnyrequestservice

import (
	"fmt"
	"mime/multipart"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	bunnyrequestlibs "github.com/jSierraB3991/simple-bunny-api-req/domain/bunny_request_libs"
	bunnyrequestrest "github.com/jSierraB3991/simple-bunny-api-req/infrastructure/bunny_request_rest"
)

func (s *BunnyRequestService) UploadVideoService(file *multipart.FileHeader, libraryId, videoCount uint, librartName, videoGuid, videoLibraryApiKey string) (*bunnyrequestrest.UploadVideoResponse, error) {
	var result bunnyrequestrest.UploadVideoResponse
	file.Filename = string(librartName) + "_" + jsierralibs.GetStringUNumberFor(videoCount)

	headers := []jsierralibs.HeaderRequest{s.GetHeaderApiKey(videoLibraryApiKey)}

	err := jsierralibs.PostFile(s.videoCdnUrl, fmt.Sprintf(bunnyrequestlibs.UPLOAD_VIDEO_URL, libraryId, videoGuid), file, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
