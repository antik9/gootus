package max

import (
	"testing"
)

type User struct {
	name string
	age  uint
}

func TestMaxElementByAge(t *testing.T) {
	got := []User{
		User{"Mr.Orange", 45},
		User{"Mr.Black", 36},
		User{"Ms.White", 38},
	}

	reference := got[0]
	user, _ := FindMax(got, func(i, j int) bool { return got[i].age < got[j].age })

	if user != reference {
		t.Errorf("get %T %v; want %T %v", user, user, reference, reference)
	}
}

func TestMaxElementByName(t *testing.T) {
	got := []User{
		User{"Mr.Orange", 45},
		User{"Mr.Black", 36},
		User{"Ms.White", 38},
	}

	reference := got[2]
	user, _ := FindMax(got, func(i, j int) bool { return got[i].name < got[j].name })

	if user != reference {
		t.Errorf("get %T %v; want %T %v", user, user, reference, reference)
	}
}

func TestMaxElementPassNil(t *testing.T) {
	_, err := FindMax(nil, func(i, j int) bool { return i < j })

	if err == nil || err.Error() != "nil is given" {
		t.Errorf("get %s; want nil is given", err)
	}
}
