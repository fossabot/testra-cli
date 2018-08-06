package tests

import (
	"fmt"
	"testing"

	"github.com/testra-tech/testra-cli/api/models"

	"gopkg.in/jarcoal/httpmock.v1"

	"github.com/spf13/cobra"
	"bytes"
					"github.com/testra-tech/testra-cli/internal/project"
	"github.com/testra-tech/testra-cli/cmd"
)

func TestFetchArticles(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	s := "567467467467"
	projectPayload := &models.Project{
		CreationDate: 12345678,
		Description:  "Description",
		ID:           &s,
		Name:         &s,
	}

	bytes, _ := projectPayload.MarshalBinary()

	httpmock.RegisterResponder("GET", "/api/v1/projects/123",
		httpmock.NewBytesResponder(200, bytes))

	c := cmd.ShowProjectCmd

	c.SetArgs([]string { "projects", "show" })
	c.Flags().Set("projectId", "123")

	project.HandleGet(c)

	//executeCommand(cmd.ShowProjectCmd, "projects", "show")

	/*out := capturer.CaptureStdout(func() {
		output, err := executeCommand(cmd.RootCmd, "projects", "show", "--projectId=123")

		if err != nil {
			fmt.Println(">" + output)
		} else {
			fmt.Println(err)
		}
	})*/

	//fmt.Println(out)
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	var buf bytes.Buffer
	root.SetOutput(&buf)
	root.SetArgs(args)

	root.Flags().Set("projectId", "123")

	project.HandleGet(root)


	/*root.InitDefaultHelpCmd()
	command, flags, err := root.Find(args)

	reflectedCmd := reflect.ValueOf(command).Elem()

	//reflect.ValueOf(command).Elem().Set(reflect.ValueOf(command).Elem())

	fieldByName := reflectedCmd.FieldByName("flags")

	actual := fieldByName.FieldByName("actual")

	fmt.Println(actual.CanSet())

	fmt.Println(flags)
	fmt.Println("flags")
	fmt.Println(command.Flags().GetString("projectId"))*/


	c, err = root.ExecuteC()

	fmt.Println(c.Flags().GetString("projectId"))

	return c, buf.String(), err
}
