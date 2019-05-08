package dto

// PostAccountRequest DTO
type PostAccountRequest struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	RoleID   string `json:"RoleId"`
}
