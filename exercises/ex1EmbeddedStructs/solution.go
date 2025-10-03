package ex1EmbeddedStructs

import (
	"fmt"
	"time"
)

type Human struct {
	Name      string
	BirthDate time.Time
	Gender    bool
	Stamina   int
}

func (h *Human) SayMyName() string {
	return "My name is " + h.Name
}

func (h *Human) Run(dist int) {
	if dist < h.Stamina {
		h.Stamina -= dist
	} else {
		h.Stamina = 0
	}

}

func (h *Human) ChangeGender(gender bool) error {
	return fmt.Errorf("you can't change gender")
}

func (h *Human) CelebrateHappyBirthday() (string, error) {
	if h.BirthDate.Day() == time.Now().Day() {
		return "Happy Birthday " + h.Name + "!", nil
	}

	return "Today is not your birthday", fmt.Errorf("sad, not your birthday today")
}

type Action struct {
	Human
}
