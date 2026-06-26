package api // Matches package namespace seamlessly

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed banners/*.txt
var bannerFS embed.FS

func readEmbeddedBanner(bannerName string) ([]string, error) {
	filepath := "banners/" + bannerName + ".txt"

	data, err := bannerFS.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not find banner style: %s", bannerName)
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")
	return lines, nil
}

func getCharacterrow(ch rune, row int, bannerLines []string) (string, error) {
	if ch < 32 || ch > 126 {
		ch = 32
	}
	index := int(ch) - 32
	startLine := 1 + index*(8+1)
	lineIndex := startLine + row

	if lineIndex >= 0 && lineIndex < len(bannerLines) {
		return bannerLines[lineIndex], nil
	}
	return "", fmt.Errorf("line index %v out of range", lineIndex)
}

func PrintAscii(banner string, input string, color string, subMatch string) (string, error) {
	bannerLines, err := readEmbeddedBanner(banner)
	if err != nil {
		return "", err
	}

	if input == "" {
		return "", nil
	}

	inputLines := strings.Split(strings.ReplaceAll(input, "\\n", "\n"), "\n")
	var result strings.Builder

	for _, line := range inputLines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		matchIdx := -1
		matchLen := 0
		if subMatch != "" && color != "" {
			matchIdx = strings.Index(line, subMatch)
			matchLen = len(subMatch)
		}

		for row := 0; row < 8; row++ {
			for chIdx, ch := range line {
				asciiLines, err := getCharacterrow(ch, row, bannerLines)
				if err != nil {
					return "", err
				}

				if matchIdx != -1 && chIdx >= matchIdx && chIdx < matchIdx+matchLen {
					result.WriteString(fmt.Sprintf("<span style=\"color:%s;\">%s</span>", color, asciiLines))
				} else {
					result.WriteString(asciiLines)
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}
