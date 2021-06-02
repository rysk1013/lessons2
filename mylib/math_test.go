package mylib

import "testing"

var Debug bool = true
func TestAverage(t *testing.T) {
	if Debug {
		t.Skip("Skip Reason")
	}
	v := Average([]int{1, 2, 3, 4, 5})
	if v != 3 {
		t.Error("Expected 3, get", v)
	}
}

// テストのためのフレームワークがある
// "RSpec-rails"に似ているのは"Ginkgo"というフレームワーク