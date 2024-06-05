package usecase

// Buat struct Login
type Login struct{}

// Buat interface LoginInterface dengan 1 method autentikasi dengan 2 param (username, password)
type LoginInterface interface {
	Autentikasi(username, password string) bool
}

// Buat function NewLogin dengan definisi return type LoginInterface
// dan pastikan return &(reference) dari struct Login
func NewLogin() LoginInterface {
	return &Login{}
}

// Buat method yang merupakan member dari struct Login
func (m *Login) Autentikasi(username, password string) bool {
	// logic if username = admin & pass = admin123 return true
	if username == "admin" && password == "admin123" {
		return true
	}
	return false
}
