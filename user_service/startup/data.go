package startup

import "github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"

var users = []*domain.User{
	{
		ID:               1,
		Name:             "Pera",
		Username:         "Pera",
		Password:         "pera",
		Email:            "pera@gmail.com",
		Address:          "kodpere10",
		Gender:           "male",
		IsItConfirmed:    true,
		IsPrivateProfile: true,
		UserId:           1,
	},
	{
		ID:               2,
		Name:             "Joka",
		Username:         "Joka",
		Password:         "joka",
		Email:            "joka@gmail.com",
		Address:          "kodjoke10",
		Gender:           "female",
		IsItConfirmed:    true,
		IsPrivateProfile: false,
		UserId:           2,
	},
	{
		ID:               3,
		Name:             "Marko",
		Username:         "Marko",
		Password:         "marko",
		Email:            "marko@gmail.com",
		Address:          "bb10",
		Gender:           "male",
		IsItConfirmed:    true,
		IsPrivateProfile: false,
		UserId:           3,
	},
	{
		ID:               4,
		Name:             "Zeksa",
		Username:         "Zeksa",
		Password:         "zeksa",
		Email:            "zeksa@gmail.com",
		Address:          "bb10",
		Gender:           "female",
		IsItConfirmed:    true,
		IsPrivateProfile: false,
		UserId:           4,
	},
	{
		ID:               5,
		Name:             "Sanja",
		Username:         "Sanja",
		Password:         "sanja",
		Email:            "sanja@gmail.com",
		Address:          "bb10",
		Gender:           "female",
		IsItConfirmed:    true,
		IsPrivateProfile: true,
		UserId:           5,
	}, {
		ID:               6,
		Name:             "Tanja",
		Username:         "Tanja",
		Password:         "tanja",
		Email:            "tanja@gmail.com",
		Address:          "bb10",
		Gender:           "female",
		IsItConfirmed:    true,
		IsPrivateProfile: true,
		UserId:           6,
	},
	{
		ID:               7,
		Name:             "Lale",
		Username:         "Lale",
		Password:         "lale",
		Email:            "lale@gmail.com",
		Address:          "bb10",
		Gender:           "male",
		IsItConfirmed:    true,
		IsPrivateProfile: true,
		UserId:           7,
	},
	{
		ID:               8,
		Name:             "Nena",
		Username:         "Nena",
		Password:         "nena",
		Email:            "nena@gmail.com",
		Address:          "bb10",
		Gender:           "female",
		IsItConfirmed:    true,
		IsPrivateProfile: false,
		UserId:           8,
	},
}
