package std

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/vlog"
)

// FileWriter 是 vlog 的文件写入器
type FileWriter struct {
	mDefaultFormatter vlog.Formatter

	// public

	Enable  bool
	Path1   string
	Path2   string
	Context application.Context

	// private

	starting bool
	started  bool
	stopping bool
	stopped  bool

	out myFileWriterBuffer
	err myFileWriterBuffer
}

func (inst *FileWriter) _Impl() vlog.Writer {
	return inst
}

// Open 打开 FileWriter
func (inst *FileWriter) Open() error {

	if !inst.Enable {
		return nil
	}

	path1 := fs.Default().GetPath(inst.Path1)
	path2 := inst.Path2
	if !path1.IsDir() {
		inst.Enable = false
		vlog.Warn("由于输出文件夹[" + path1.Path() + "]不存在，vlog 的 FileWriter 将被禁用。")
		return nil
	}

	inst.starting = true
	inst.out.init(inst.Context, path1, path2+"-out.log")
	inst.err.init(inst.Context, path1, path2+"-error.log")
	go inst.loop()
	return nil
}

// Close 关闭 FileWriter
func (inst *FileWriter) Close() error {
	inst.stopping = true
	for {
		if !inst.starting {
			break
		}
		if inst.stopped {
			break
		}
		time.Sleep(time.Millisecond * 20)
	}
	return nil
}

func (inst *FileWriter) getDefaultFormatter() vlog.Formatter {
	f := inst.mDefaultFormatter
	if f == nil {
		f = &DefaultFormatter{}
		inst.mDefaultFormatter = f
	}
	return f
}

func (inst *FileWriter) Write(rec *vlog.Record) {

	if !inst.Enable {
		return
	}

	if rec == nil {
		return
	}

	const nl = "\n"
	text := rec.Message

	if text == "" {
		text = inst.getDefaultFormatter().Format(rec)
	}

	if rec.Level > vlog.WARN {
		inst.err.writeString(text + nl)
	} else {
		inst.out.writeString(text + nl)
	}
}

func (inst *FileWriter) loop() error {
	defer func() {
		inst.stopped = true
	}()
	inst.started = true
	for {
		if inst.stopping {
			break
		}
		time.Sleep(time.Second)
		inst.flush()
	}
	inst.flush()
	return nil
}

func (inst *FileWriter) flush() {
	inst.out.flush()
	inst.err.flush()
}

////////////////////////////////////////////////////////////////////////////////

type myFileWriterBuffer struct {
	context application.Context

	path1              fs.Path
	path2raw           string
	path2final         string
	currentFile        fs.Path
	currentFileCreated time.Time
	buffer             strings.Builder
}

func (inst *myFileWriterBuffer) init(ctx application.Context, path1 fs.Path, path2 string) error {
	inst.path1 = path1
	inst.path2raw = path2
	inst.context = ctx
	return nil
}

func (inst *myFileWriterBuffer) flush() error {

	text := inst.buffer.String()
	inst.buffer.Reset()
	if text == "" {
		return nil
	}

	file := inst.currentFile
	now := time.Now()
	age := now.Sub(inst.currentFileCreated)
	if age > time.Minute {
		file = nil
	}
	if file == nil {
		file = inst.computeFinalFile(now)
		inst.currentFile = file
		inst.currentFileCreated = now
	}

	return inst.appendStringToFile(file, text)
}

func (inst *myFileWriterBuffer) computeFinalFile(now time.Time) fs.Path {

	path1 := inst.path1
	path2 := inst.path2raw

	const key = "path2"
	props := collection.CreateProperties()
	props.SetProperty(key, path2)
	inst.prepareTimeForProps(props, now)
	inst.prepareAppNameForProps(props)
	collection.ResolvePropertiesVar(props)
	path2 = props.GetProperty(key, path2)

	inst.path2final = path2
	return path1.GetChild(path2)
}

func (inst *myFileWriterBuffer) prepareAppNameForProps(props collection.Properties) {
	const key = "application.simple-name"
	appName := inst.context.GetProperties().GetProperty(key, key)
	props.SetProperty(key, appName)
}

func (inst *myFileWriterBuffer) prepareTimeForProps(props collection.Properties, t time.Time) {

	year, month, day := t.Date()
	hour := t.Hour()
	min := t.Minute()
	sec := t.Second()

	YYYY := inst.stringifyInt(year, 1)
	MM := inst.stringifyInt(int(month), 2)
	DD := inst.stringifyInt(day, 2)
	hh := inst.stringifyInt(hour, 2)
	mm := inst.stringifyInt(min, 2)
	ss := inst.stringifyInt(sec, 2)

	const prefix = "time.now."
	props.SetProperty(prefix+"YYYY", YYYY)
	props.SetProperty(prefix+"MM", MM)
	props.SetProperty(prefix+"DD", DD)
	props.SetProperty(prefix+"hh", hh)
	props.SetProperty(prefix+"mm", mm)
	props.SetProperty(prefix+"ss", ss)

}

func (inst *myFileWriterBuffer) stringifyInt(n, width int) string {
	str := strconv.Itoa(n)
	for len(str) < width {
		str = "0" + str
	}
	return str
}

func (inst *myFileWriterBuffer) appendStringToFile(file fs.Path, text string) error {

	dir := file.Parent()
	if !dir.Exists() {
		dir.Mkdirs()
	}

	flag := os.O_APPEND | os.O_CREATE
	mode := inst.path1.GetMeta().Mode()
	name := file.Path()

	f, err := os.OpenFile(name, flag, mode)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	if err != nil {
		return err
	}
	return nil
}

func (inst *myFileWriterBuffer) writeString(s string) {

	// todo : use chan ->

	inst.buffer.WriteString(s)
}
