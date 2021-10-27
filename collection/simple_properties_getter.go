package collection

import "strconv"

type simplePropertyGetter struct {
	props Properties
	err   error
}

func (inst *simplePropertyGetter) _Impl() PropertyGetter {
	return inst
}

func (inst *simplePropertyGetter) innerGetStr(name string) (string, bool) {
	str, err := inst.props.GetPropertyRequired(name)
	if err != nil {
		inst.err = err
		return str, false
	}
	return str, true
}

func (inst *simplePropertyGetter) innerGetInt(name string, bits int, defaultValue int64) int64 {
	str, ok := inst.innerGetStr(name)
	if !ok {
		return defaultValue
	}
	value, err := strconv.ParseInt(str, 10, bits)
	if err != nil {
		inst.err = err
		return defaultValue
	}
	return value
}

func (inst *simplePropertyGetter) innerGetUint(name string, bits int, defaultValue uint64) uint64 {
	str, ok := inst.innerGetStr(name)
	if !ok {
		return defaultValue
	}
	value, err := strconv.ParseUint(str, 10, bits)
	if err != nil {
		inst.err = err
		return defaultValue
	}
	return value
}

func (inst *simplePropertyGetter) Error() error {
	return inst.err
}

func (inst *simplePropertyGetter) SetError(err error) {
	if err == nil {
		return
	}
	inst.err = err
}

func (inst *simplePropertyGetter) CleanError() {
	inst.err = nil
}

func (inst *simplePropertyGetter) GetString(name string, defaultValue string) string {
	str, ok := inst.innerGetStr(name)
	if !ok {
		return defaultValue
	}
	return str
}

func (inst *simplePropertyGetter) GetBool(name string, defaultValue bool) bool {
	str, ok := inst.innerGetStr(name)
	if !ok {
		return defaultValue
	}
	b, err := strconv.ParseBool(str)
	if err != nil {
		inst.err = err
		return defaultValue
	}
	return b
}

func (inst *simplePropertyGetter) GetInt(name string, defaultValue int) int {
	value := inst.innerGetInt(name, 0, int64(defaultValue))
	return int(value)
}

func (inst *simplePropertyGetter) GetInt8(name string, defaultValue int8) int8 {
	value := inst.innerGetInt(name, 8, int64(defaultValue))
	return int8(value)
}

func (inst *simplePropertyGetter) GetInt16(name string, defaultValue int16) int16 {
	value := inst.innerGetInt(name, 16, int64(defaultValue))
	return int16(value)
}

func (inst *simplePropertyGetter) GetInt32(name string, defaultValue int32) int32 {
	value := inst.innerGetInt(name, 32, int64(defaultValue))
	return int32(value)
}

func (inst *simplePropertyGetter) GetInt64(name string, defaultValue int64) int64 {
	return inst.innerGetInt(name, 64, defaultValue)
}

func (inst *simplePropertyGetter) GetUint(name string, defaultValue uint) uint {
	value := inst.innerGetUint(name, 0, uint64(defaultValue))
	return uint(value)
}

func (inst *simplePropertyGetter) GetUint8(name string, defaultValue uint8) uint8 {
	value := inst.innerGetUint(name, 8, uint64(defaultValue))
	return uint8(value)
}

func (inst *simplePropertyGetter) GetUint16(name string, defaultValue uint16) uint16 {
	value := inst.innerGetUint(name, 16, uint64(defaultValue))
	return uint16(value)
}

func (inst *simplePropertyGetter) GetUint32(name string, defaultValue uint32) uint32 {
	value := inst.innerGetUint(name, 32, uint64(defaultValue))
	return uint32(value)
}

func (inst *simplePropertyGetter) GetUint64(name string, defaultValue uint64) uint64 {
	return inst.innerGetUint(name, 64, defaultValue)
}

func (inst *simplePropertyGetter) GetFloat32(name string, defaultValue float32) float32 {
	str, ok := inst.innerGetStr(name)
	if !ok {
		return defaultValue
	}
	f, err := strconv.ParseFloat(str, 32)
	if err != nil {
		inst.err = err
		return defaultValue
	}
	return float32(f)
}

func (inst *simplePropertyGetter) GetFloat64(name string, defaultValue float64) float64 {
	str, ok := inst.innerGetStr(name)
	if !ok {
		return defaultValue
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		inst.err = err
		return defaultValue
	}
	return f
}
