package bunnyrequestservice

import (
	"encoding/json"
	"fmt"

	jsierralibs "github.com/jSierraB3991/jsierra-libs"
	bunnyrequestlibs "github.com/jSierraB3991/simple-bunny-api-req/domain/bunny_request_libs"
	bunnyrequestrequest "github.com/jSierraB3991/simple-bunny-api-req/infrastructure/bunny_request_request"
	bunnyrequestrest "github.com/jSierraB3991/simple-bunny-api-req/infrastructure/bunny_request_rest"
)

func (s *BunnyRequestService) CreateCollection(libraryId uint, videoLibraryApiKey, collectionName string) (*bunnyrequestrest.CreateCollectionResponse, error) {
	var result bunnyrequestrest.CreateCollectionResponse

	headers := []jsierralibs.HeaderRequest{s.GetHeaderApiKey()}

	body := bunnyrequestrequest.CreateCollectionRequest{
		Name: collectionName,
	}
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	err = jsierralibs.Post(s.videoCdnUrl, fmt.Sprintf(bunnyrequestlibs.CREATE_COLLECTION_URL, libraryId), jsonData, &result, headers)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
