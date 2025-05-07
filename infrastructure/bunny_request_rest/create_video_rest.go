package bunnyrequestrest

import "time"

type CreateVideoResponse struct {
	VideoLibraryID       int       `json:"videoLibraryId"`
	GUID                 string    `json:"guid"`
	Title                string    `json:"title"`
	Description          *string   `json:"description"` // Puede ser nulo
	DateUploaded         time.Time `json:"dateUploaded"`
	Views                int       `json:"views"`
	IsPublic             bool      `json:"isPublic"`
	Length               int       `json:"length"`
	Status               int       `json:"status"`
	Framerate            int       `json:"framerate"`
	Rotation             *int      `json:"rotation"`
	Width                int       `json:"width"`
	Height               int       `json:"height"`
	AvailableResolutions *string   `json:"availableResolutions"`
	OutputCodecs         *string   `json:"outputCodecs"`
	ThumbnailCount       int       `json:"thumbnailCount"`
	EncodeProgress       int       `json:"encodeProgress"`
	StorageSize          int64     `json:"storageSize"`
	Captions             []string  `json:"captions"`
	HasMP4Fallback       bool      `json:"hasMP4Fallback"`
	CollectionID         string    `json:"collectionId"`
	ThumbnailFileName    string    `json:"thumbnailFileName"`
	AverageWatchTime     int       `json:"averageWatchTime"`
	TotalWatchTime       int       `json:"totalWatchTime"`
	Category             string    `json:"category"`
	Chapters             []string  `json:"chapters"`
	Moments              []string  `json:"moments"`
	MetaTags             []string  `json:"metaTags"`
	TranscodingMessages  []string  `json:"transcodingMessages"`
	JitEncodingEnabled   *bool     `json:"jitEncodingEnabled"`
}
