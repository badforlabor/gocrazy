/**
 * Auth :   liubo
 * Date :   2021/5/17 15:59
 * Comment:
 */

package crazylog

import (
	"fmt"
	"github.com/badforlabor/gocrazy/crazy3rd/glog"
	"io"
)

type CategoryLog struct {
	category string
	cachedText string
	writer io.Writer
	autoNewLine bool
	parent *CategoryLog
}

// 如果writer是nil，那么会用glog输出内容
func NewCategoryLogDefault(category string, writer io.Writer) *CategoryLog {
	return NewCategoryLog(category, writer, true, nil)
}

func NewCategoryLog2(category string, parent *CategoryLog) *CategoryLog {
	return NewCategoryLog(category, nil, true, parent)
}

func NewCategoryLog(category string, writer io.Writer, autoNewLine bool, parent *CategoryLog) *CategoryLog {

	var s = ""
	for parent != nil {
		s = fmt.Sprintf("[%s]", parent.category) + s
		parent = parent.parent
	}

	return &CategoryLog{autoNewLine:autoNewLine, writer:writer, category:category, cachedText:fmt.Sprintf("%s[%s]", s, category), parent:parent}
}

func (self *CategoryLog) output(s string) {
	if self.writer != nil {
		self.writer.Write([]byte(s))
	}
}

func getLevelDesc(level int) string {
	if level == 2 {
		return "error"
	} else if level == 1 {
		return "warn"
	} else {
		return "info"
	}
}

func (self *CategoryLog) logf(level int, format string, args... interface{}) {
	var s = fmt.Sprintf(format, args...)
	if self.writer != nil {
		self.output(fmt.Sprintf("[%s]%s %s\n", getLevelDesc(level), self.cachedText, s))
	} else {
		s = fmt.Sprintf("%s %s", self.cachedText, s)
		if level == 2 {
			glog.Error(s)
		} else if level == 1 {
			glog.Warning(s)
		} else {
			glog.Info(s)
		}
	}
}
func (self *CategoryLog) logln(level int, args... interface{}) {
	self.logf(level, fmt.Sprintln(args...))
}

func (self *CategoryLog) Infof(format string, args... interface{}) {
	self.logf(0, format, args...)
}
func (self *CategoryLog) Infoln(args... interface{}) {
	self.logln(0, args...)
}
func (self *CategoryLog) Warnf(format string, args... interface{}) {
	self.logf(1, format, args...)
}
func (self *CategoryLog) Warnln(args... interface{}) {
	self.logln(1, args...)
}
func (self *CategoryLog) Errorf(format string, args... interface{}) {
	self.logf(2, format, args...)
}
func (self *CategoryLog) Errorln(args... interface{}) {
	self.logln(2, args...)
}