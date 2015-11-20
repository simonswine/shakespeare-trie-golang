package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func assertList(t *testing.T, exp []string, result []string) {
	sort.Strings(result)
	sort.Strings(exp)
	assert.Equal(t, exp, result)
}

func TestShakespeareTrie(t *testing.T) {
	t1 := NewShakespeareTrie()
	t1.AddString("")
	r := t1.GetMatches("")
	assert.Equal(t, []string{""}, r)
	r = t1.GetMatches("a")
	assert.Equal(t, []string{}, r)

	t2 := NewShakespeareTrie()
	t2.AddString("foo")
	t2.AddString("fuu")
	t2.AddString("buu")
	assertList(
		t,
		[]string{"foo", "fuu", "buu"},
		t2.GetMatches(""),
	)
	assertList(
		t,
		[]string{"foo", "fuu"},
		t2.GetMatches("f"),
	)
	assertList(
		t,
		[]string{},
		t2.GetMatches("fa"),
	)

	t3 := NewShakespeareTrie()
	t3.AddString("äöüß")
	t3.AddString("abc")
	assertList(
		t,
		[]string{"äöüß", "abc"},
		t3.GetMatches(""),
	)
	assertList(
		t,
		[]string{"äöüß"},
		t3.GetMatches("ä"),
	)
	assertList(
		t,
		[]string{},
		t3.GetMatches("fa"),
	)
}
