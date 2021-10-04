package main

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	if len(words) == 0 {
		return nil
	}
	var result []string

	indexs := splitToLines(words, maxWidth)
	if len(indexs) == 1 {
		content := strings.Join(words, " ")
		spaceWidth := maxWidth - len(content)
		content += strings.Repeat("", spaceWidth)
		result = append(result, content)
		return result
	}
	for i := 0; i < len(indexs)-1; i++ {
		result = append(result, expand(words, indexs[i], maxWidth))
	}
	var lastLineItem []string
	for i := indexs[len(indexs)-1][0]; i <= indexs[len(indexs)-1][1]; i++ {
		lastLineItem = append(lastLineItem, words[i])
	}
	content := strings.Join(lastLineItem, " ")
	spaceWidth := maxWidth - len(content)
	content += strings.Repeat(" ", spaceWidth)
	result = append(result, content)
	return result
}

// split into multiple lines, each line has begin and end
func splitToLines(words []string, maxWidth int) [][]int {
	var (
		begin, end = 0, -1
		charCount  = 0
		itemCount  = 0
		result     [][]int
		space      = 0
	)
	for i, word := range words {
		switch {
		case itemCount > 0:
			space = 1
		default:
			space = 0
		}
		if i > 0 {
			if charCount+space+len(word) > maxWidth {
				result = append(result, []int{begin, end})
				begin = end + 1
				end = begin - 1
				charCount = 0
				itemCount = 0
				space = 0
			}
		}

		charCount += space + len(word)
		itemCount += 1
		end += 1
	}

	result = append(result, []int{begin, len(words) - 1})
	return result

}

func expand(words []string, index []int, maxWidth int) string {
	var (
		result    = ""
		begin     = index[0]
		end       = index[1]
		charCount = 0
		itemCount = end - begin + 1
	)

	if begin == end {
		return words[begin] + strings.Repeat(" ", maxWidth-len(words[begin]))
	}
	for i := begin; i <= end; i++ {
		charCount += len(words[i])
	}
	emptyCount := maxWidth - charCount
	base := int(emptyCount / (itemCount - 1))
	extra := int(emptyCount % (itemCount - 1))
	order := 1

	for i := begin; i < end; i++ {
		result += words[i]
		if order <= extra {
			result += strings.Repeat(" ", base+1)
		} else {
			result += strings.Repeat(" ", base)
		}
		order += 1
	}

	result += words[end]
	return result
}
