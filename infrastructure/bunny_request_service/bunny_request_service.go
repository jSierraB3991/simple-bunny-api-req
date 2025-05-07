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
	CREATE VIDEO
		libraryId == RESPONSE OF VIDEO LIBRARY ID
		POST /library/{libraryId}/videos
		AUTHORIZATION == RESPOONSE OF VIDEO LIBRARY API KEY
		BODY => json {
			"title": "My Video",
		}
		RESPONSE => json {
			"videLibraryId": 123456,
			"guid": "12345678-1234-1234-1234-123456789012",
			"title": "My Video",
			"dateUploaded": "2023-10-01T12:00:00Z",
			"views": 0,
			"isPublic": false,
			"length": 0,
			"status": "Processing",
			"framerate": 0,
			"width": 0,
			"height": 0,
			"availableResolutions": null,
			"thumbnailCount": 0,
			"encodedProgress": 0,
			"storageSize": 0,
			"captions": [],
			"hasMP4Fallback": false,
			"CollectionId": 0,
		}




	https://video.bunnycdn.com
	GET VIDEOS BY LIBRARY ID
		libraryId == RESPONSE OF VIDEO LIBRARY ID
		GET /library/{libraryId}/videos
		AUTHORIZATION == RESPOONSE OF VIDEO LIBRARY API KEY
		RESPONSE => json {
			"totalItems": 1,
			"currentPage": 1,
			"itemsPerPage": 100,
			"items": [
				RESPONSE OF VIDEO
			]
		}
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
	UPLOAD VIDEO
		libraryId == RESPONSE OF VIDEO LIBRARY ID
		PUT /library/{libraryId}/videos/{videoGuid}
		AUTHORIZATION == RESPOONSE OF VIDEO LIBRARY API KEY
		BINARY BY BODY => file



	https://vz-580df404-944.b-cdn.net/13daffd7-e78d-4936-8521-77740e2cc734/playlist.m3u8

	<div style="position:relative;padding-top:56.25%;">
		<iframe
			src="https://iframe.mediadelivery.net/embed/421874/13daffd7-e78d-4936-8521-77740e2cc734?autoplay=true&loop=false&muted=false&preload=true&responsive=true"
				loading="lazy"
				style="border:0;position:absolute;top:0;height:100%;width:100%;"
				allow="accelerometer;gyroscope;autoplay;encrypted-media;picture-in-picture;"
				allowfullscreen="true">
		</iframe>
	</div>


*/
