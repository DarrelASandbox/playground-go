package topics

import (
	"encoding/json"
	"fmt"
	"log"
)

// person struct already exist in E7
type someone struct {
	First   string
	Last    string
	Sayings []string
}

func E11() {
	so1 := someone{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	fmt.Println("\n\n##################################################")
	errorHandlingEx1(so1)
	errorHandlingEx2(so1)
	errorHandlingEx3()
	errorHandlingEx4()
}

func errorHandlingEx1(so1 someone) {
	fmt.Println("\n\nerrorHandlingEx1:")

	bs, err := json.Marshal(so1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}

func errorHandlingEx2(so1 someone) {
	fmt.Println("\n\nerrorHandlingEx2:")

	// json.Marshal return already a []byte, error:
	// https://pkg.go.dev/encoding/json#Marshal
	bs, err := json.Marshal(so1)
	if err != nil {
		err = fmt.Errorf("error in toJSON: %v", err)
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}

/* #################################################################################################### */

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("error: %v", ce.info)
}

func errorHandlingEx3() {
	fmt.Println("\n\nerrorHandlingEx3:")

	foo := func(e error) {
		fmt.Println("foo ran = ", e)
	}

	// Assertion: e.(customErr).info
	// Explicitly states that it is customErr
	// foo := func(e error) {
	// 	fmt.Println("foo ran = ", e.(customErr).info)
	// }

	ce1 := customErr{
		info: "some error msg",
	}
	foo(ce1)
}

/* #################################################################################################### */

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func errorHandlingEx4() {
	fmt.Println("\n\nerrorHandlingEx4:")

	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// e := errors.New("more coffee needed")
		e := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}

/* #################################################################################################### */
