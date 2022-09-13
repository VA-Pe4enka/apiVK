package userModel

type ResponseList struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Closed    bool   `json:"is_closed"`
}
type ProfileData struct {
	Response []ResponseList `json:"response"`
}

var UserID int
