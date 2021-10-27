package collection

import "strconv"

type simplePropertySetter struct {
	props Properties
}

func (inst *simplePropertySetter) _Impl() PropertySetter {
	return inst
}

func (inst *simplePropertySetter) SetString(name string, value string) {
	inst.props.SetProperty(name, value)
}

func (inst *simplePropertySetter) SetBool(name string, value bool) {
	str := strconv.FormatBool(value)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetInt(name string, value int) {
	str := strconv.FormatInt(int64(value), 10)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetInt8(name string, value int8) {
	str := strconv.FormatInt(int64(value), 10)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetInt16(name string, value int16) {
	str := strconv.FormatInt(int64(value), 10)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetInt32(name string, value int32) {
	str := strconv.FormatInt(int64(value), 10)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetInt64(name string, value int64) {
	str := strconv.FormatInt(value, 10)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetUint(name string, value uint) {
	str := strconv.FormatUint(uint64(value), 10)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetUint8(name string, value uint8) {
	str := strconv.FormatUint(uint64(value), 10)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetUint16(name string, value uint16) {
	str := strconv.FormatUint(uint64(value), 10)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetUint32(name string, value uint32) {
	str := strconv.FormatUint(uint64(value), 10)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetUint64(name string, value uint64) {
	str := strconv.FormatUint(value, 10)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetFloat32(name string, value float32) {
	str := strconv.FormatFloat(float64(value), 'f', -1, 32)
	inst.SetString(name, str)
}

func (inst *simplePropertySetter) SetFloat64(name string, value float64) {
	str := strconv.FormatFloat(value, 'f', -1, 64)
	inst.SetString(name, str)
}
