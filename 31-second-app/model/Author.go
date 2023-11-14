package model

type Author struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Bio       string `json:"bio"`
	Username  string `json:"username"`
	Highlight string `json:"highlights"`
}
