package loader

import (
	"strings"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/vlog"
)

type bannerDisplayer struct {
	resources  collection.Resources
	properties collection.Properties
}

func (inst *bannerDisplayer) display() error {
	text, err := inst.loadText()
	if err != nil {
		return nil
	}
	const prefix = "\n"
	text = inst.injectValues(prefix + text)
	vlog.Info(text)
	return nil
}

func (inst *bannerDisplayer) loadText() (string, error) {
	const key = "banner.location"
	path := inst.properties.GetProperty(key, "res:///banner.txt")
	return inst.resources.GetText(path)
}

func (inst *bannerDisplayer) injectValues(text string) string {

	const token1 = "${"
	const token2 = "}"
	buffer := text
	builder := strings.Builder{}

	for {
		p1, p2, p3, ok := inst.parseNextVar(buffer, token1, token2)
		if !ok {
			builder.WriteString(buffer)
			break
		}
		builder.WriteString(p1)
		builder.WriteString(inst.getProperty(p2))
		buffer = p3
	}

	return builder.String()
}

func (inst *bannerDisplayer) parseNextVar(text string, token1 string, token2 string) (string, string, string, bool) {

	index1 := strings.Index(text, token1)
	index2 := strings.Index(text, token2)

	if index1 < 0 || index2 < 0 {
		return "", "", "", false
	}

	if index1 >= index2 {
		return "", "", "", false
	}

	len1 := len(token1)
	len2 := len(token2)

	p1 := text[0:index1]
	p2 := text[index1+len1 : index2]
	p3 := text[index2+len2:]
	return p1, p2, p3, true
}

func (inst *bannerDisplayer) getProperty(key string) string {
	key = strings.TrimSpace(key)
	value, err := inst.properties.GetPropertyRequired(key)
	if err != nil {
		return "{{" + key + "}}"
	}
	return value
}
