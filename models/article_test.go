package models

import "testing"

func TestGetArticles(t *testing.T) {
	aList := GetArticles()
	if len(aList) != 2 {
		t.Fail()
	}
}
