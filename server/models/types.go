package models

// if a JSON key is missing, the corresponding Go field will be set to nil instead of being omitted.
type Image struct {
	Height   *string `json:"height,omitempty"`
	Width    *string `json:"width,omitempty"`
	Size     *string `json:"size,omitempty"`
	URL      *string `json:"url,omitempty"`
	MP4Size  *string `json:"mp4_size,omitempty"`
	MP4      *string `json:"mp4,omitempty"`
	WebPSize *string `json:"webp_size,omitempty"`
	WebP     *string `json:"webp,omitempty"`
	Frames   *string `json:"frames,omitempty"`
	Hash     *string `json:"hash,omitempty"`
}

type DataItem struct {
	Type             string `json:"type"`
	ID               string `json:"id"`
	URL              string `json:"url"`
	Slug             string `json:"slug"`
	BitlyGifURL      string `json:"bitly_gif_url"`
	BitlyURL         string `json:"bitly_url"`
	EmbedURL         string `json:"embed_url"`
	Username         string `json:"username"`
	Source           string `json:"source"`
	Title            string `json:"title"`
	Rating           string `json:"rating"`
	ContentURL       string `json:"content_url"`
	SourceTLD        string `json:"source_tld"`
	SourcePostURL    string `json:"source_post_url"`
	IsSticker        int    `json:"is_sticker"`
	ImportDatetime   string `json:"import_datetime"`
	TrendingDatetime string `json:"trending_datetime"`
	Images           struct {
		Original               Image `json:"original"`
		FixedHeight            Image `json:"fixed_height"`
		FixedHeightDownsampled Image `json:"fixed_height_downsampled"`
		FixedWidth             Image `json:"fixed_width"`
		FixedWidthDownsampled  Image `json:"fixed_width_downsampled"`
	} `json:"images"`
}

type Pagination struct {
	TotalCount int `json:"total_count"`
	Count      int `json:"count"`
	Offset     int `json:"offset"`
}

type Meta struct {
	Status     int    `json:"status"`
	Message    string `json:"msg"`
	ResponseID string `json:"response_id"`
}

type ResponseData struct {
	Data       []DataItem `json:"data"`
	Pagination Pagination `json:"pagination"`
	Meta       Meta       `json:"meta"`
}
