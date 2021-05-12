package configen2

import (
	"go/ast"
	"go/token"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

type commentContentBuilder struct {
	buffer strings.Builder
}

func (inst *commentContentBuilder) append(group *ast.CommentGroup) {
	if group == nil {
		return
	}
	text := group.Text()
	inst.buffer.WriteString(text)
}

func (inst *commentContentBuilder) String() string {
	return inst.buffer.String()
}

////////////////////////////////////////////////////////////////////////////////

type commentSet struct {
	allComments []*ast.CommentGroup
}

func (inst *commentSet) init(comments []*ast.CommentGroup) {
	inst.allComments = comments
}

func (inst *commentSet) findComment(from token.Pos, to token.Pos) (string, error) {
	builder := &commentContentBuilder{}
	all := inst.allComments
	for index := range all {
		group := all[index]
		if group == nil {
			continue
		}
		pos := group.Pos()
		end := group.End()
		if from < pos && end < to {
			builder.append(group)
		}
	}
	return builder.String(), nil
}

func (inst *commentSet) loadInjectionProperties(from token.Pos, to token.Pos) map[string]string {
	text, err := inst.findComment(from, to)
	if err != nil {
		return map[string]string{}
	}
	return map[string]string{"text": text}
}
