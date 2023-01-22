package entities

import "golang.org/x/crypto/bcrypt"

type Humans struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}
func (user *Humans) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Email = string(bytes)
	return nil
}
func (user *Humans) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Email), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}