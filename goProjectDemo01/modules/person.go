package modules

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age,string"`
	Address string `json:"address,omitempty"`
}

type Teacher struct {
	Person
	school string
}
