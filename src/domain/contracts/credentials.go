package contracts

type Credentials struct {
	PrivateKeyId string `json:"private_key_id"`
	ProjectId    string `json:"project_id"`
	Crypt		 string `json:"crypt"`
}
