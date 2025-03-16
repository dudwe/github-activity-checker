package activity

import (
	"net/http"
	"testing"
)

func TestGetActivtyRequest(t *testing.T) {
	// Test with a valid URL
	validURL := "https://www.example.com"
	res, err := GetActivityRequest(validURL)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got: %d", res.StatusCode)
	}
	res.Body.Close()

	// Test with an invalid URL
	invalidURL := "ht@tp://invalid-url"
	_, err = GetActivityRequest(invalidURL)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	// Test with an unreachable URL (mocking or use an unlikely address)
	unreachableURL := "http://0.0.0.0:9999"
	_, err = GetActivityRequest(unreachableURL)
	if err == nil {
		t.Fatal("expected error for unreachable URL, got nil")
	}
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.

func TestUserEmptyGetActivity(t *testing.T) {
	expected := "request reponse abnormal: 404"
	msg, err := GetActivity(expected)
	if err.Error() != expected {
		t.Errorf(`GetActivity("") = %q, got "%v"`, expected, msg)
	}
}
func TestUserGetActivity(t *testing.T) {
	name := "kamranahmedse"
	msg, _ := GetActivity(name)
	if msg == "" {
		t.Errorf(`Got empty payload  "%v"`, msg)
	}
}
