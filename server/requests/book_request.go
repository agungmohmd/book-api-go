package request

type BookRequest struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Stock     int64  `json:"stock"`
	UpdatedAt string `json:"updated_at"`
}
