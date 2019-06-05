package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

// Admin model
type Admin struct {
	ID       int    `db:"id" json:"id,omitempty"`
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password,omitempty"`
}

// String - print admin as string.
func (a *Admin) String() string {
	return fmt.Sprintf("%d - %s, %s",
		a.ID,
		a.Name,
		a.Email)
}

// CreateAdmin - create a new admin
func (a *Admin) CreateAdmin(db *sqlx.DB) (*Admin, bool, error) {
	var id int
	var err error
	var admin []Admin
	err = db.Get(&admin, `SELECT * FROM admin WHERE email = $1`, a.Email)
	if err != nil {
		return nil, false, err
	}
	if len(admin) != 0 {
		return nil, false, errors.New("Another user with the same email address already exists")
	}

	cost := bcrypt.DefaultCost
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(a.Password), cost)
	if err != nil {
		return nil, false, err
	}
	pwd := string(hash)

	err = db.QueryRow(
		`INSERT INTO admin (name, email, password)
		 VALUES ($1, $2, $3) RETURNING id`,
		a.Name, a.Email, pwd).Scan(&id)

	if err != nil {
		return nil, false, err
	}

	a.ID = id
	return a, true, nil
}

// ListAdmins - returns list of admins
func ListAdmins(db *sqlx.DB) (*[]Admin, error) {
	var admins []Admin

	err := db.Select(&admins, `SELECT a.id, a.name, a.email FROM admin a`)

	if err != nil {
		return nil, err
	}

	if len(admins) == 0 {
		return nil, sql.ErrNoRows
	}

	return &admins, nil
}

// GetAdmin - returns admin based on id
func GetAdmin(db *sqlx.DB, id int) (*Admin, error) {
	admin := Admin{}

	err := db.Get(&admin, "SELECT a.id, a.name, a.email FROM admin a WHERE a.id = $1", id)

	if err != nil {
		return nil, err
	}

	return &admin, nil
}

// GetAdminByEmail - returns admin based on email
func GetAdminByEmail(db *sqlx.DB, email string) (*Admin, error) {
	admin := Admin{}

	err := db.Get(&admin, "SELECT a.id, a.name, a.email FROM admin a WHERE a.email = $1",
		email)

	if err != nil {
		return nil, err
	}

	return &admin, nil
}
