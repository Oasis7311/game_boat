package common_model

type Image struct {
	Id     int64  `json:"id,omitempty"`
	Uri    string `json:"uri,omitempty"`
	Width  int32  `json:"width,omitempty"`
	Height int32  `json:"height,omitempty"`
	Format string `json:"format,omitempty"`
}
