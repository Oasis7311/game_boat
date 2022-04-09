package common_model

type Image struct {
	Uri    string `json:"web_uri,omitempty"`
	Width  int32  `json:"width,omitempty"`
	Height int32  `json:"height,omitempty"`
	Type   string `json:"type,omitempty"`
}
