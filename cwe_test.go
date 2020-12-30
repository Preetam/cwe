package cwe

import "testing"

func TestList(t *testing.T) {
	cwe918 := CWEs["CWE-918"]

	checkField := func(val, expected string) {
		if val != expected {
			t.Errorf("expected %s, got %s", expected, val)
		}
	}
	checkField(cwe918.Name, "Server-Side Request Forgery (SSRF)")
	checkField(cwe918.Description, "The web server receives a URL or similar request from an upstream component and retrieves the contents of this URL, but it does not sufficiently ensure that the request is being sent to the expected destination.")
}
