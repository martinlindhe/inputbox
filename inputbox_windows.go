package inputbox

import (
	"strings"

	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

// InputBox displays a dialog box, returning the entered value and a bool for success
func InputBox(title, message, defaultAnswer string) (string, bool) {
	shell, err := ps.New(&backend.Local{})
	if err != nil {
		panic(err)
	}
	defer shell.Exit()

	out, _, err := shell.Execute(`
		[void][Reflection.Assembly]::LoadWithPartialName('Microsoft.VisualBasic')
		$title = '` + title + `'
		$msg = '` + message + `'
		$default = '` + defaultAnswer + `'
		$answer = [Microsoft.VisualBasic.Interaction]::InputBox($msg, $title, $default)
		Write-Output $answer
		`)
	// FIXME: if cancel button is pressed in dialog, we should return false
	if err != nil {
		return "", false
	}
	return strings.TrimSpace(string(out)), true
}
