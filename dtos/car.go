package dtos

type CarDTO struct {
	ID             uint   `json:"id"`
	Owner          string `json:"owner"`
	Color          string `json:"color"`
	NumberOfWheels int    `json:"number_of_wheels"`
}
