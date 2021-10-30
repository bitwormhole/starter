package collection

import (
	"errors"
	"strings"
)

// ResolvePropertiesVar 解析 Properties 中的变量
func ResolvePropertiesVar(props Properties) error {
	return ResolvePropertiesVarWithTokens("${", "}", props)
}

// ResolvePropertiesVarWithTokens 解析 Properties 中的变量，token1：变量开始；token2：变量结束
func ResolvePropertiesVarWithTokens(token1, token2 string, props Properties) error {
	task := myRopertyVariableResolver{
		tokenBegin: token1,
		tokenEnd:   token2,
	}
	return task.Resolve(props)
}

////////////////////////////////////////////////////////////////////////////////

type myRopertyVariableResolver struct {
	table      map[string]string
	keys       map[string]bool
	tokenBegin string
	tokenEnd   string
}

func (inst *myRopertyVariableResolver) Resolve(props Properties) error {
	table := props.Export(nil)
	inst.init(table)
	done := false
	for timeout := 9; timeout > 0; timeout-- {
		cnt, err := inst.scanOnce()
		if err != nil {
			return err
		}
		if cnt <= 0 {
			done = true
			break
		}
	}
	props.Import(inst.table)
	if !done {
		return errors.New("还有一些变量没有被解析")
	}
	return nil
}

func (inst *myRopertyVariableResolver) init(table map[string]string) {
	keys := make(map[string]bool)
	for key := range table {
		keys[key] = true
	}
	inst.table = table
	inst.keys = keys
}

func (inst *myRopertyVariableResolver) scanOnce() (int, error) {
	cnt := 0
	for k, v := range inst.table {
		c2, err := inst.scanItem(k, v)
		if err != nil {
			return cnt, err
		}
		cnt += c2
	}
	return cnt, nil
}

func (inst *myRopertyVariableResolver) scanItem(name, value string) (int, error) {
	parts := strings.Split(value, inst.tokenBegin)
	if len(parts) < 2 {
		return 0, nil // 没有发现变量
	}
	cnt := 0
	builder := strings.Builder{}
	for _, part := range parts {
		const n = 2
		keyAndText := strings.SplitN(part, inst.tokenEnd, n)
		if len(keyAndText) < n {
			// no token-end
			builder.WriteString(part)
		} else {
			key := keyAndText[0]
			val, err := inst.resolveVar(key)
			if err != nil {
				return cnt, err
			}
			builder.WriteString(val)
			builder.WriteString(keyAndText[1])
			cnt++
		}
	}
	inst.table[name] = builder.String()
	return cnt, nil
}

func (inst *myRopertyVariableResolver) resolveVar(name string) (string, error) {
	name = strings.TrimSpace(name)
	value := inst.table[name]
	exists := inst.keys[name]
	if exists {
		return value, nil
	}
	return "", errors.New("no value for name:" + name)
}
