package messagepost

type MessagePostRecord struct {
	ID     int    `json:"id"`
	upvote int    `json:"upvote"`
	body   string `json:"body"`
}
