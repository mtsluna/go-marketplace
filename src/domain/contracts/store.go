package contracts

type Store struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	ImageUrl string `json:"image_url"`
	UserId   string `json:"user_id"`
}