package utils_test

import (
	"github.com/JannoTjarks/vidya-intarweb-playlist-cli/internal/utils"
	"testing"
)

func TestGenerateVersionSignature(t *testing.T) {
	want := "vidya-intarweb-playlist-cli dev, commit none, built at unknown, build by unknown"

	msg := utils.GenerateVersionSignature()

	if msg != want {
		t.Errorf(`utils.GenerateVersionSignature() = %v, want match for %#v, nil`, msg, want)
	}
}

func TestGenerateVersionJson(t *testing.T) {
	want := `{"version":"dev","commit":"none","date":"unknown","buildby":"unknown"}`

	msg := utils.GenerateVersionJson()

	if msg != want {
		t.Errorf(`utils.GenerateVersionSignature() = %v, want match for %#v, nil`, msg, want)
	}
}
