package loader2

import (
	"strings"

	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/vlog"
)

type bannerLoader struct {
	// app application.Context
}

func (inst *bannerLoader) load(loading *contextLoading) error {
	res := loading.context.GetResources()
	text, err := res.GetText("/banner.txt")
	if err != nil {
		return nil
	}
	props := loading.context.GetProperties()
	text = inst.injectProperties(props, text)
	vlog.Info("\n" + text)
	return nil
}

func (inst *bannerLoader) injectProperties(props collection.Properties, text string) string {
	for timeout := 99; timeout > 0; timeout-- {
		text2, ok := inst.tryInjectNextProperty(props, text)
		if ok {
			text = text2
		} else {
			break
		}
	}
	return text
}

func (inst *bannerLoader) tryInjectNextProperty(props collection.Properties, text string) (string, bool) {

	const begin = "${"
	const end = "}"
	index1 := strings.Index(text, begin)

	if index1 < 0 {
		return text, false
	}

	part1 := text[0:index1]
	part23 := text[index1+len(begin):]
	index2 := strings.Index(part23, end)

	if index2 < 0 {
		return text, false
	}

	part2 := part23[0:index2]
	part3 := part23[index2+len(end):]
	key := strings.TrimSpace(part2)
	val := props.GetProperty(key, "<[["+key+"]]>")
	part2 = val

	return part1 + part2 + part3, true
}
