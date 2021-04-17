package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrentMatches(t *testing.T) {
	client, err := fifa.NewClient(&fifa.Options{})
	if !assert.Nil(t, err, "expected no error with NewClient, got: %s", err) {
		t.FailNow()
	}
	resp, err := client.GetCurrentMatches()
	if !assert.Nil(t, err, "expected no error with GetCurrentMatches, got: %s", err) {
		t.FailNow()
	}
	t.Logf("%+v", resp)
}

func TestGetTodaysMatches(t *testing.T) {
	client, err := fifa.NewClient(&fifa.Options{})
	if !assert.Nil(t, err, "expected no error with NewClient, got: %s", err) {
		t.FailNow()
	}
	_, err = client.GetTodaysMatches()
	if !assert.Nil(t, err, "expected no error with GetTodaysMatches, got: %s", err) {
		t.FailNow()
	}
}
