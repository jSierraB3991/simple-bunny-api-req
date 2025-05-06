package bunnyrequestservice

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
