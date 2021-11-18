// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type BuyImageInput struct {
	ImageID string  `json:"imageId"`
	Price   float64 `json:"price"`
}

type DeleteImageInput struct {
	ID string `json:"id"`
}

type ImageFilterInput struct {
	ID        *string  `json:"id"`
	UserID    *string  `json:"userId"`
	Title     *string  `json:"title"`
	Labels    []string `json:"labels"`
	AllLabels *bool    `json:"allLabels"`
	Private   *bool    `json:"private"`
	ForSale   *bool    `json:"forSale"`
	Price     *float64 `json:"price"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewImageInput struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Labels      []string       `json:"labels"`
	File        graphql.Upload `json:"file"`
	Private     bool           `json:"private"`
	ForSale     bool           `json:"forSale"`
	Price       float64        `json:"price"`
}

type NewUserInput struct {
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Bio      string         `json:"bio"`
	Avatar   graphql.Upload `json:"avatar"`
}

type UpdateImageInput struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Labels      []string `json:"labels"`
	Private     bool     `json:"private"`
	ForSale     bool     `json:"forSale"`
	Price       float64  `json:"price"`
}

type UpdateUserInput struct {
	Username string          `json:"username"`
	Email    string          `json:"email"`
	Bio      string          `json:"bio"`
	Avatar   *graphql.Upload `json:"avatar"`
}

type UserFilterInput struct {
	ID       *string `json:"id"`
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

type Role string

const (
	RoleAdmin     Role = "ADMIN"
	RoleUser      Role = "USER"
	RoleModerator Role = "MODERATOR"
)

var AllRole = []Role{
	RoleAdmin,
	RoleUser,
	RoleModerator,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleUser, RoleModerator:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
