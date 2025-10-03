package ex1EmbeddedStructs

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreationOfAction(t *testing.T) {
	action := &Action{
		Human{
			Name:      "Chelik",
			BirthDate: time.Date(2000, 10, 5, 0, 0, 0, 0, time.UTC),
			Gender:    true,
			Stamina:   100,
		},
	}

	expectedName := "Chelik"
	expectedBirthDate := time.Date(2000, 10, 5, 0, 0, 0, 0, time.UTC)
	expectedGender := true
	expectedStamina := 100

	assert.Equal(t, expectedName, action.Name)
	assert.Equal(t, expectedBirthDate, action.BirthDate)
	assert.Equal(t, expectedGender, action.Gender)
	assert.Equal(t, expectedStamina, action.Stamina)
}

func TestMethodSayMyName(t *testing.T) {
	action := &Action{
		Human{
			Name:      "Chelik",
			BirthDate: time.Date(2000, 10, 5, 0, 0, 0, 0, time.UTC),
			Gender:    true,
			Stamina:   100,
		},
	}

	expectedString := "My name is Chelik"
	actualString := action.SayMyName()
	assert.Equal(t, expectedString, actualString)
}

func TestMethodRun(t *testing.T) {
	action := &Action{
		Human{
			Name:      "Chelik",
			BirthDate: time.Date(2000, 10, 5, 0, 0, 0, 0, time.UTC),
			Gender:    true,
			Stamina:   100,
		},
	}

	dist := 50
	expectedStamina := 100 - dist
	action.Run(dist)
	actualStamina := action.Stamina

	assert.Equal(t, expectedStamina, actualStamina)
}

func TestMethodRunWithDistMoreThenStamina(t *testing.T) {
	action := &Action{
		Human{
			Name:      "Chelik",
			BirthDate: time.Date(2000, 10, 5, 0, 0, 0, 0, time.UTC),
			Gender:    true,
			Stamina:   100,
		},
	}

	dist := 200
	expectedStamina := 0
	action.Run(dist)
	actualStamina := action.Stamina

	assert.Equal(t, expectedStamina, actualStamina)
}

func TestMethodChangeGender(t *testing.T) {
	action := &Action{
		Human{
			Name:      "Chelik",
			BirthDate: time.Date(2000, 10, 5, 0, 0, 0, 0, time.UTC),
			Gender:    true,
			Stamina:   100,
		},
	}

	expectedResult := errors.New("you can't change gender")
	actualResult := action.ChangeGender(true)

	assert.Equal(t, expectedResult, actualResult)
}

func TestMethodCelebrateHappyBirthdayInIncorrectDay(t *testing.T) {
	action := &Action{
		Human{
			Name:      "Chelik",
			BirthDate: time.Date(2000, 10, 5, 0, 0, 0, 0, time.UTC),
			Gender:    true,
			Stamina:   100,
		},
	}

	expectedResult, expectedErr := "Today is not your birthday", errors.New("sad, not your birthday today")

	actualResult, actualErr := action.CelebrateHappyBirthday()

	assert.Equal(t, expectedResult, actualResult)
	assert.Equal(t, expectedErr, actualErr)
}

func TestMethodCelebrateHappyBirthday(t *testing.T) {
	action := &Action{
		Human{
			Name:      "Chelik",
			BirthDate: time.Now(),
			Gender:    true,
			Stamina:   100,
		},
	}

	expectedResult := "Happy Birthday Chelik!"
	actualResult, err := action.CelebrateHappyBirthday()

	assert.Equal(t, expectedResult, actualResult)
	assert.Nil(t, err)
}
