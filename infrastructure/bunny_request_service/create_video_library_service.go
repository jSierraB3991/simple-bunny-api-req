package bunnyrequestservice

import (
	"encoding/json"
	"fmt"

	bunnyrequestlibs "github.com/jSierraB3991/simple-bunny-api-req/domain/bunny_request_libs"
	bunnyrequestrequest "github.com/jSierraB3991/simple-bunny-api-req/infrastructure/bunny_request_request"
	bunnyrequestrest "github.com/jSierraB3991/simple-bunny-api-req/infrastructure/bunny_request_rest"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
)

func (s *BunnyRequestService) CreateVideoLibrary(nameVideoLIbrary string) (*bunnyrequestrest.CreateVideoLibraryResponse, error) {
	var result bunnyrequestrest.CreateVideoLibraryResponse
	videoLibraryRequest := bunnyrequestrequest.CreateVideoLibraryRequest{
		Name: nameVideoLIbrary,
	}
	jsonData, err := json.Marshal(videoLibraryRequest)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return nil, err
	}
	headers := []jsierralibs.HeaderRequest{{Key: "AccessKey", Value: s.bunnyApiKeyAuth}}

	err = jsierralibs.Post(s.videoLibraryUrl, bunnyrequestlibs.CREATE_VIDEO_LIBREY_URI, jsonData, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
