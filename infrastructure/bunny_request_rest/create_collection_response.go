package bunnyrequestrest

type CreateCollectionResponse struct {
	VideoLibraryId   int      `json:"videoLibraryId"`
	Guid             string   `json:"guid"`
	Name             string   `json:"name"`
	VideoCount       int      `json:"videoCount"`
	TotalSize        int      `json:"totalSize"`
	PreviewVideoIds  []string `json:"previewVideoIds"` // or *[]string if you want to allow null
	PreviewImageUrls []string `json:"previewImageUrls"`
}
