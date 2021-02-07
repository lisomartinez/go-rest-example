package models

import (
	"math/rand"
	"testing"
)

func NewReview(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   stars,
		Comment: comment,
	}
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func Test_withCorrectParams(t *testing.T) {
	review := NewReview(4, "The iphone X looks good")
	err := review.validate()
	if err != nil {
		t.Error("the validation did not pass")
		t.Fail()
	}
}

func Test_shouldFailWithWrongMaxNumbersOfStars(t *testing.T) {
	review := NewReview(8, "The iphone X looks good")
	err := review.validate()
	if err == nil {
		t.Error("should fail with 5 stars")
		t.Fail()
	}
}

func Test_shouldFailWithWrongMinNumbersOfStars(t *testing.T) {
	review := NewReview(0, "The iphone X looks good")
	err := review.validate()
	if err == nil {
		t.Error("should fail with 5 stars")
		t.Fail()
	}
}

func Test_shouldFailWithWrongCommentLength(t *testing.T) {
	review := NewReview(8, RandomString(500))
	err := review.validate()
	if err == nil {
		t.Error("should fail with 400 max chars")
		t.Fail()
	}
}
