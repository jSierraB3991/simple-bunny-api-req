package bunnyrequestservice

import jsierralibs "github.com/jSierraB3991/jsierra-libs"

type BunnyRequestService struct {
	bunnyApiKeyAuth string

	videoLibraryUrl string
	videoCdnUrl     string
}

func NewBunnyRequestService(bunnyApiKeyAuth string,
	videoLibraryUrl string,
	videoCdnUrl string) *BunnyRequestService {

	return &BunnyRequestService{
		bunnyApiKeyAuth: bunnyApiKeyAuth,
		videoLibraryUrl: videoLibraryUrl,
		videoCdnUrl:     videoCdnUrl,
	}
}

func (s *BunnyRequestService) GetHeaderApiKey() jsierralibs.HeaderRequest {
	return jsierralibs.HeaderRequest{
		Key:   "AccessKey",
		Value: s.bunnyApiKeyAuth,
	}
}

/*

	ADD THUMBNAIL
		libraryId == RESPONSE OF VIDEO LIBRARY ID
		POST /library/{libraryId}/videos/{videoGuid}/thumbnails?thumbnailUrl={thumbnail_url}
		AUTHORIZATION == RESPOONSE OF VIDEO LIBRARY API KEY
		RESPONSE => json {
			"success": true,
			"message": "Thumbnail added successfully",
			"statusCode": 200,
		}
	ADD CAPTION
		libraryId == RESPONSE OF VIDEO LIBRARY ID
		lang == "EN"
		langDesc == "English"
		captionFile == "/path/to/caption/file.vtt"
		POST /library/{libraryId}/videos/{videoGuid}/captions/{lang}
		AUTHORIZATION == RESPOONSE OF VIDEO LIBRARY API KEY
		BODY => json {
			"srclang": {lang},
			"label": {langDesc},
			captionFile: {base64(captionFile)},
		}
		RESPONSE => json {
			"success": true,
			"message": "Thumbnail added successfully",
			"statusCode": 200,
		}

*/
