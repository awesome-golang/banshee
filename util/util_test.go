// Copyright 2015 Eleme Inc. All rights reserved.

package util

import (
	"github.com/eleme/banshee/util/assert"
	"os"
	"testing"
)

func TestMatch(t *testing.T) {
	assert.Ok(t, Match("a*cd*fg", "abcdefg"))
	assert.Ok(t, !Match("a*cd*fg", "cbcdefg"))
	assert.Ok(t, !Match("a*cd*fg", "abcdef"))
	assert.Ok(t, !Match("a*cd*fg", "abxdef"))
	assert.Ok(t, Match("abxdef", "abxdef"))
}

func TestIsFileExist(t *testing.T) {
	assert.Ok(t, IsFileExist(os.Args[0]))
	assert.Ok(t, !IsFileExist("file-not-exist"))
}
