package another

var PublicVariable string = "another package can access"
var privateVariable string = "another package can not access"

type Linklist []string
type desk []string

type Employee struct {
	name string
}

func (e Employee) getType() string {
	return "someting"
}

func (e Employee) GetLength() int {
	return len(e.name)
}

type account struct {
	name string
}

func RandomNumber() int {
	return 11
}

func randonText() string {
	return "someone"
}
