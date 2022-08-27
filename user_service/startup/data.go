package startup

import (
	"github.com/XWS-Dislinkt-Developers/Dislinkt-backend/user_service/domain"
	"time"
)

// birthdays
var pera_user01_date, _ = time.Parse("2006-01-02", "1988-08-03")
var joka_user02_date, _ = time.Parse("2006-01-02", "1988-08-05")
var marko_user03_date, _ = time.Parse("2006-01-02", "1950-10-10")
var zeksa_user04_date, _ = time.Parse("2006-01-02", "1999-11-11")
var sanja_user05_date, _ = time.Parse("2006-01-02", "2000-08-08")
var tanja_user06_date, _ = time.Parse("2006-01-02", "1988-12-12")
var lale_user07_date, _ = time.Parse("2006-01-02", "1999-06-16")
var nane_user08_date, _ = time.Parse("2006-01-02", "1999-08-03")

// [ 1 -> 8 ] is users
//Password for everyone: Pera12345*
var users = []*domain.User{
	// USERS
	{
		ID:               1,
		Name:             "Pera",
		Username:         "Pera",
		Password:         "pera",
		Email:            "pera@gmail.com",
		Address:          "kodpere10",
		Gender:           "male",
		Biography:        "Pera is a C++ back-end developer who is obsessed with Gangnam style.",
		DateOfBirth:      pera_user01_date,
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
		Biography:        "Joka is a Vue.js front-end developer who loves ABBA.",
		DateOfBirth:      joka_user02_date,
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
		Biography:        "Marko is a Data scientist who watches a lot of anime.",
		DateOfBirth:      marko_user03_date,
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
		Biography:        "Zeksa is a Python developer and she only loves programming.",
		DateOfBirth:      zeksa_user04_date,
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
		Biography:        "Sanja is a student and she also works in a Coffee shop.",
		DateOfBirth:      sanja_user05_date,
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
		Biography:        "Tanja is a student and she wants to be a professor.",
		DateOfBirth:      tanja_user06_date,
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
		Biography:        "Lale is a student and he loves to make short movies.",
		DateOfBirth:      lale_user07_date,
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
		Biography:        "Nena is a Spring Boot developer and she loves Montenegro.",
		DateOfBirth:      nane_user08_date,
		IsItConfirmed:    true,
		IsPrivateProfile: false,
		UserId:           8,
	},
	{
		ID:            9,
		Name:          "Admin",
		Username:      "Admin",
		IsItConfirmed: true,
		UserId:        9,
	},
}
