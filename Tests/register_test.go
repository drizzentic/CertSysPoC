package tests

import (
	"certSys/controllers"
	"regexp"
	"testing"
)


func TestRegister(t *testing.T) {
	//Create a bitcoin address and validate it was created and exist
	var student = controllers.Student{"derrick",
		"cs/m/07/123",
		"Computing Science",
		"Science, Engineering and Technology",
		"Bachelor of Computer Science",
		"",
	}
	student.Name = "derrick"
	s := controllers.Register(&student)
	match, _ := regexp.MatchString("[a-zA-Z1-9]{27,35}$", s)

	if !match {
		t.Errorf("The address generated is not valid: %s", s)
	}

}
