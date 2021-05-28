/**
 * Auth :   liubo
 * Date :   2021/5/17 15:59
 * Comment:
 */

package crazylog

import (
	"fmt"
	"io"
)

type CategoryLog struct {
	category string
	cachedText string
	writer io.Writer
	autoNewLine bool
	parent *CategoryLog
}

func NewCategoryLogDefault(category string, writer io.Writer) *CategoryLog {
	return NewCategoryLog(category, writer, true, nil)
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

func (self *CategoryLog) logf(level, format string, args... interface{}) {
	var s = fmt.Sprintf(format, args...)
	self.output(fmt.Sprintf("[%s]%s %s\n", level, self.cachedText, s))
}
func (self *CategoryLog) Infof(format string, args... interface{}) {
	self.logf("info", format, args...)
}
func (self *CategoryLog) Warnf(format string, args... interface{}) {
	self.logf("warn", format, args...)
}
func (self *CategoryLog) Errorf(format string, args... interface{}) {
	self.logf("error", format, args...)
}