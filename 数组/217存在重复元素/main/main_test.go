package main

import "testing"

func TestContainsDuplicate2(t *testing.T) {
	a:=containsDuplicate2([]int{1,2,3,1})
	t.Log(a)
}
func TestContainsDuplicate1(t *testing.T) {
	a:=containsDuplicate1([]int{1,2,3,1})
	t.Log(a)
}