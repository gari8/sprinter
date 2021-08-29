package sprinter

type Response struct {
	Code        int         `json:"code"`
	ContentType string      `json:"content_type,omitempty"`
	Content     []byte      `json:"content,omitempty"`
	Text        string      `json:"text,omitempty"`
	Object      interface{} `json:"object,omitempty"`
	Path        string      `json:"path,omitempty"`
	Err         error       `json:"err,omitempty"`
}
