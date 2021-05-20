package contracts

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Brand    string `json:"brand"`
	Password string `json:"password,omitempty"`
	Username string `json:"username"`
}
