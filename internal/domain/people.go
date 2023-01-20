package domain

type ListPeople []People

type People struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	PhoneNumber int    `json:"phone_number"`
	BirthYear   int    `json:"birth_year"`
	Nationality string `json:"nationality"`
}
