package bunnyrequestserviceinterface

import (
	"mime/multipart"

	bunnyrequestrest "github.com/jSierraB3991/simple-bunny-api-req/infrastructure/bunny_request_rest"
)

type BunnyRequestServiceInterface interface {
	CreateVideoLibrary(nameVideoLIbrary string) (*bunnyrequestrest.CreateVideoLibraryResponse, error)
	CreateCollection(libraryId uint, videoLibraryApiKey, collectionName string) (*bunnyrequestrest.CreateCollectionResponse, error)
	CreateVideoService(titleVideo, collectionId string, thumbnailInMs *int, libreryId uint) (*bunnyrequestrest.CreateVideoResponse, error)
	UploadVideoService(file *multipart.FileHeader, libraryId, videoCount uint, librartName, videoGuid string) (*bunnyrequestrest.UploadVideoResponse, error)
}
