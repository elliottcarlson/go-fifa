package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetCompetitions(t *testing.T) {
	client, err := fifa.NewClient(nil)
	if !assert.Nil(t, err, "expected no error with NewClient, got: %s", err) {
		t.FailNow()
	}
	_, err = client.GetCompetitions()
	if !assert.Nil(t, err, "expected no error with GetCompetitions, got: %s", err) {
		t.FailNow()
	}
}
