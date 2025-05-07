package bunnyrequestrequest

type CreateVideoRequest struct {
	Title         string `json:"title"`
	CollectionID  string `json:"collectionId"`
	ThumbnailTime *int   `json:"thumbnailTime"`
}
