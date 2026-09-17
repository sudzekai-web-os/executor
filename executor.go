package executor

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/types"
)

type Executor struct {
	logger abstractions.ILogger
}

func NewExecutor(factory abstractions.ILoggerFactory) abstractions.IExecutor {
	return &Executor{
		logger: factory.NewLogger("executor"),
	}
}

func (ex *Executor) Execute(command string, args ...string) types.CommandResult {
	cmd := exec.Command(command, args...)

	return ex.executeCommand(cmd)
}

func (ex *Executor) ExecuteInDirectory(directory string, command string, args ...string) types.CommandResult {
	cmd := exec.Command(command, args...)

	cmd.Dir = directory

	return ex.executeCommand(cmd)
}

func (ex *Executor) ExecuteWithInput(input, command string, args ...string) types.CommandResult {
	cmd := exec.Command(command, args...)

	cmd.Stdin = strings.NewReader(input)

	return ex.executeCommand(cmd)
}

func (ex *Executor) executeCommand(cmd *exec.Cmd) (result types.CommandResult) {
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer

	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	ex.logger.LogDebug("выполнение команды: %s %s", cmd.Path, strings.Join(cmd.Args, " "))

	if result.Error = cmd.Start(); result.Error != nil {
		return
	}

	result.Error = cmd.Wait()

	result.Stdout = stdoutBuf.String()
	result.Stderr = stderrBuf.String()

	ex.logger.LogDebug("вывод команды:\nstdout: %s\nstderr: %s", result.Stdout, result.Stderr)

	return
}
