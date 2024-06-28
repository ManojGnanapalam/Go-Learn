package main

type Users struct {
	Name     string `json:"Username"`
	Email    string
	Password string   `json:"-"`
	Role     []string `json:"Role,omitempty"`
}

var userData []Users

func (c *Users) IsEmpty() bool {
	return c.Name == ""
}

func main() {

}
