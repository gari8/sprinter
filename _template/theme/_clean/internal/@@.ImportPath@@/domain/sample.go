package domain

type Sample struct {
	ID   uint32 `json:"id"`
	Text string `json:"text,omitempty"`
}
