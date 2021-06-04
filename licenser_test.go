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

func TestParseProjectOnSimpleString(t *testing.T) {
	org, proj, e := ParseProject("git@github.com:nandanrao/licenser.git")
	assert.Nil(t, e)
	assert.Equal(t, "nandanrao", org)
	assert.Equal(t, "licenser", proj)
}

func TestParseProjectOnHyphenString(t *testing.T) {
	org, proj, e := ParseProject("git@github.com:foo-bar/baz_qux.git")
	assert.Nil(t, e)
	assert.Equal(t, "foo-bar", org)
	assert.Equal(t, "baz_qux", proj)
}

func TestParseProjectErrorsOnBadString(t *testing.T) {
	_, _, e := ParseProject("foo")
	assert.Error(t, e)
}
