package entity

import (
	"regexp"
	"time"

	"github.com/google/uuid"
	"github.com/nordew/go-errx"
)

var (
	RegexEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	RegexPhone = regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	RegexPassword = regexp.MustCompile(`^.{8,}$`)
)

type User struct {
	ID        string
	ImageURL  string
	ProfileName string
	FullName    string
	Email       string
	Phone       string
	Password    string
	Blocked     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUser(
	id,
	profileName,
	fullName,
	email,
	phone,
	password string,
) (*User, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, errx.NewInternal().WithDescription("invalid id")
	}

	if profileName == "" {
		return nil, errx.NewInternal().WithDescription("profile name is required")
	}

	if fullName == "" {
		return nil, errx.NewInternal().WithDescription("full name is required")
	}

	if !RegexEmail.MatchString(email) {
		return nil, errx.NewInternal().WithDescription("invalid email")
	}

	if !RegexPhone.MatchString(phone) {
		return nil, errx.NewInternal().WithDescription("invalid phone")
	}

	if password == "" {
		return nil, errx.NewInternal().WithDescription("password is required")
	}

	if !RegexPassword.MatchString(password) {
		return nil, errx.NewInternal().WithDescription("password must be at least 8 characters long")
	}

	now := time.Now()

	return &User{
		ID:          id,
		ProfileName: profileName,
		FullName:    fullName,
		Email:       email,
		Phone:       phone,
		Password:    password,
		Blocked:     false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

