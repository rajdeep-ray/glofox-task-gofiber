package requests

type BookingRequest struct {
	Name      string `json:"name"`
	Date      string `json:"date"`
	ClassName string `json:"class"`
}
