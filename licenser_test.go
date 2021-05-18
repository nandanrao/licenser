package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeContributorListMakesGoodList(t *testing.T) {
	var s string

	s = MakeCopyright([]string{}, "baz", "qux", "2020")
	exp := "Copyright (c) 2020 qux contributors\n(https://github.com/baz/qux/graphs/contributors)"
	assert.Equal(t, exp, s)

	s = MakeCopyright([]string{"Foo Dee"}, "baz", "qux", "2020")
	exp = "Copyright (c) 2020 Foo Dee and qux contributors\n(https://github.com/baz/qux/graphs/contributors)"
	assert.Equal(t, exp, s)

	s = MakeCopyright([]string{"Foo Dee", "The Bar"}, "baz", "qux", "2020")
	exp = "Copyright (c) 2020 Foo Dee, The Bar and qux contributors\n(https://github.com/baz/qux/graphs/contributors)"
	assert.Equal(t, exp, s)
}

func TestParseProjectOnGoodString(t *testing.T) {
	org, proj, e := ParseProject("foo/bar")
	assert.Nil(t, e)
	assert.Equal(t, "foo", org)
	assert.Equal(t, "bar", proj)
}

func TestParseProjectErrorsOnBadString(t *testing.T) {
	_, _, e := ParseProject("foo")
	assert.Error(t, e)
}
