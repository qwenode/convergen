package data

type UserA struct {
	Username string `json:"username"`
}
type UserB struct {
	Username string  `json:"username"`
	Age      int     `json:"age"`
	User     UserA   `json:"user"`
	Slice    []UserA `json:"slice"`
}
