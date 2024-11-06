package assignments

import (
	"bytes"
	"os"
)

func CaptureOutputOf(function func()) string {
	var buffer bytes.Buffer
	originalOutputSetting := os.Stdout
	readEnd, writeEnd, _ := os.Pipe()
	os.Stdout = writeEnd //temporarily set the os standard out to be the write end of our pipe

	function()

	writeEnd.Close()
	os.Stdout = originalOutputSetting // reset it

	buffer.ReadFrom(readEnd)
	return buffer.String()
}
