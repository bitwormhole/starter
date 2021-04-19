package config

import (
	"embed"
	"io"
	"strings"
)

type simpleEmbedResFS struct {
	fs     *embed.FS
	prefix string
}

func (inst *simpleEmbedResFS) computeResPath(path string) string {

	tmp := inst.prefix + "/" + path
	builder := &strings.Builder{}
	tmp = strings.ReplaceAll(tmp, "\\", "/")
	array := strings.Split(tmp, "/")

	for index := range array {
		part := array[index]
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		if builder.Len() > 0 {
			builder.WriteString("/")
		}
		builder.WriteString(part)
	}

	return builder.String()
}

func (inst *simpleEmbedResFS) GetText(path string) (string, error) {
	path = inst.computeResPath(path)
	data, err := inst.fs.ReadFile(path)
	if err != nil {
		return "", err
	}
	text := string(data)
	return text, nil
}

func (inst *simpleEmbedResFS) GetBinary(path string) ([]byte, error) {
	path = inst.computeResPath(path)
	data, err := inst.fs.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (inst *simpleEmbedResFS) GetReader(path string) (io.ReadCloser, error) {
	return nil, nil
}
