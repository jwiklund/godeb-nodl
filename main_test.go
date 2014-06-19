package main

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if "1.3" != version("go1.3.linux-amd64.tar.gz") {
		t.Error("Got version", version("go1.3.linux-amd64.tar.gz"))
	}
}