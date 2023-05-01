package apptesting

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

func MockUser() User {
	return User{
		ID:        "1",
		Name:      "Test User",
		CreatedAt: "2020-01-01T00:00:00Z",
	}
}

func MockUsers() []User {
	return []User{
		{
			ID:        "1",
			Name:      "Test User",
			CreatedAt: "2020-01-01T00:00:00Z",
		},
	}
}
