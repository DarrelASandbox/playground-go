package files

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func E8() {
	fmt.Println("\n\n##################################################")
	fmt.Println("E8:")

	u1 := user{First: "James", Last: "Bomb", Age: 81}
	u2 := user{First: "Larry", Last: "Fairy", Age: 18}
	u3 := user{First: "Mong", Last: "Kong", Age: 54, Sayings: []string{
		"James, it is soo good to see you",
		"Would you like me to take care of that for you, James?",
		"I would really prefer to be a secret agent myself.",
	}}

	u2.Sayings = append(u2.Sayings, "Shaken, not stirred")

	users := []user{u1, u2, u3}

	fmt.Println(users)
	marshalFunc(users...)
	unmarshalFunc()
	encoderFunc(users...)
}

func marshalFunc(users ...user) {
	fmt.Println("\n\nmarshalFunc:")

	bs, err := json.Marshal(users) // byte slice
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}

func unmarshalFunc() {
	fmt.Println("\n\nunmarshalFunc:")
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	var users []user
	err := json.Unmarshal([]byte(s), &users)
	if err != nil {
		fmt.Println(err)
	}

	for i, p := range users {
		fmt.Printf("#%#v: %v %v, age: %v\n", i, p.First, p.Last, p.Age)
		for _, saying := range p.Sayings {
			fmt.Println("\t", saying)
		}
		fmt.Println()
	}
}

func encoderFunc(users ...user) {
	fmt.Println("\n\nencoderFunc:")
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}
