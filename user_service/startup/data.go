package startup

import "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"

var users = []*domain.User{
	{
		ID:       1,
		Name:     "Pera",
		Username: "Pera",
		Password: "pera",
		Email:    "pera@gmail.com",
		Address:  "kodpere10",
		Gender:   "male",
	},
	{
		ID:       2,
		Name:     "Joka",
		Username: "Joka",
		Password: "joka",
		Email:    "joka@gmail.com",
		Address:  "kodjoke10",
		Gender:   "female",
	},
	{
		ID:       3,
		Name:     "Marko",
		Username: "Marko",
		Password: "marko",
		Email:    "marko@gmail.com",
		Address:  "bb10",
		Gender:   "male",
	},
	{
		ID:       4,
		Name:     "Sanja",
		Username: "Sanja",
		Password: "sanja",
		Email:    "sanja@gmail.com",
		Address:  "bb10",
		Gender:   "female",
	},
}
