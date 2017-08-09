package dispatcher

import (
	"appie/compose"
	"fmt"
	"os/exec"
)

func Dispatch() {

	var composeObject compose.ComposeFile
	composeObject = compose.GetAppCompose()

	travarse(composeObject)

}

func travarse(composer compose.ComposeFile)  {
	for name, app:= range composer.Apps {
		fmt.Println("Start processing", name)
		execute(app)
	}
}

func execute(app compose.App)  {
	if app.Run != "" {
		cmd :=exec.Command("sh", "-c", app.Run)


		output, _ := cmd.StdoutPipe()

		_ = cmd.Start()

		cmd.Wait()

		fmt.Println(output)

	}
}
