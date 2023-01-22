package request

type Humans struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Address string `json:"address"`
}
type Req struct {
	Id      string `form:"id"`
	Name    string `form:"name"`
	Address string `form:"address"`
}
