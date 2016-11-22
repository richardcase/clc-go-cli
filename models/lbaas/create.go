package lbaas

type CreateReq struct {
	DataCenter  string `json:"-" valid:"required" URIParam:"yes"`
	Name        string `json:"name" valid:"required"`
	Description string `json:"description" valid:"required"`
}
