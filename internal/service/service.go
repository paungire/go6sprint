package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(text string) string {
	if strings.ContainsAny(text, ".-") {
		return morse.DefaultConverter.ToText(text)
	}
	return morse.DefaultConverter.ToMorse(text)
}
