package tests

import (
	"certSys/controllers"
	"regexp"
	"testing"
)

func TestCreateAddress(t *testing.T) {

	s := controllers.RequestAddress()
	match, _ := regexp.MatchString("[a-zA-Z1-9]{27,35}$", s)

	if !match {
		t.Errorf("The address generated is not valid: %s", s)
	}

}
