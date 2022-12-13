package util_test

import (
	"log"
	"testing"

	"github.com/djeniusinvfest/inventora/auth/util"
)

func TestHashAndVerify(t *testing.T) {
	str := "Hallo kau"
	digest, _ := util.HashString(str)
	match := util.VerifyDigest(str, digest)

	if !match {
		log.Fatalf("HashUtil verifier returns false, wants true")
	}
}
