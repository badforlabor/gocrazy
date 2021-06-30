/**
 * Auth :   liubo
 * Date :   2020/1/31 22:32
 * Comment: Cmder
 */

package crazyapp

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"sync"
)

const (
	// cmd输出：stdout, stderr
	CmdStdOut = 0
	CmdStdErr = 1
)

type Cmder struct {
	cmd *exec.Cmd

	wg sync.WaitGroup
	pErr io.ReadCloser
	pOut io.ReadCloser
	pIn io.WriteCloser

	errBuffer bytes.Buffer
	outBuffer bytes.Buffer

	LogOnLine bool
}
func (self *Cmder) GetErrorOutput() string {
	return string(self.errBuffer.Bytes())
}
func (self *Cmder) GetOutput() string {
	return string(self.outBuffer.Bytes())
}
func (self *Cmder) WriteString(data string) {
	self.Write([]byte(data))
}
func (self *Cmder) Write(data []byte) {
	if len(data) == 0 || data[len(data)-1] != '\n' {
		data = append(data, '\n')
	}
	var _ ,err = self.pIn.Write(data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
type CmderCallback func(self *Cmder, level int, words string)
var CmderConsoleOutput CmderCallback
func (self *Cmder) Run(callback CmderCallback) error {

	self.wg.Add(2)

	if CmderConsoleOutput != nil {
		CmderConsoleOutput(self, CmdStdOut, fmt.Sprintln("cmd", self.cmd.String()))
	} else {
		fmt.Println("cmd", self.cmd.String())
	}

	var err = self.cmd.Start()
	if err != nil {
		fmt.Println(err.Error())
	}

	self.errBuffer.Reset()
	self.outBuffer.Reset()

	go func() {
		self.copyAndCapture(CmdStdErr, self.pErr, &self.errBuffer, callback)
		self.wg.Done()
	}()
	go func() {
		self.copyAndCapture(CmdStdOut, self.pOut, &self.outBuffer, callback)
		self.wg.Done()
	}()

	err = self.cmd.Wait()
	self.wg.Wait()

	if err != nil {
		fmt.Println(err.Error())
	}

	return err
}

func (self *Cmder) copyAndCapture(level int, r io.ReadCloser, input *bytes.Buffer, callback CmderCallback) error {

	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {

			d := buf[:n]
			input.Write(d)
			if CmderConsoleOutput != nil {
				CmderConsoleOutput(self, level, string(d))
			} else {
				fmt.Print(string(d))
			}

			if callback != nil {
				callback(self, level, string(d))
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return err
		}
	}
}

func NewCmder(exe string, workDir string, arg ...string) *Cmder {
	var ret = &Cmder{}
	var cmd = exec.Command(exe, arg...)
	ret.cmd	= cmd
	if len(workDir) > 0 {
		ret.cmd.Dir = workDir
	}
	ret.pErr, _ =  cmd.StderrPipe()
	ret.pOut, _ = cmd.StdoutPipe()
	ret.pIn, _ = cmd.StdinPipe()

	return ret
}

func RunCmder(exec string, workDir string, callback CmderCallback, args ...string) (*Cmder, error) {
	var cmd = NewCmder(exec, workDir, args...)
	var err = cmd.Run(callback)
	return cmd, err
}
func RunCmder1(exec string) (*Cmder, error) {
	return RunCmder(exec, "", nil)
}
func RunCmder2(exec string, args ...string) (*Cmder, error) {
	return RunCmder(exec, "", nil, args...)
}
func RunCmder3(exec string, workDir string, args ...string) (*Cmder, error) {
	return RunCmder(exec, workDir, nil, args...)
}