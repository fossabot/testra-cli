package main

import "github.com/testra-tech/testra-cli/cmd"

func main() {

	/*// create the transport
	transport := httptransport.New("10.210.200.82:8080", apiclient.DefaultBasePath, []string{"http"})

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	// to override the host for the default client
	// apiclient.Default.SetTransport(transport)

	// make the request to get all items
	//params := project.GetProjectParams{ ID: "5b3b558e52faff0006df1597" }
	//resp, err := client.Project.GetProject(project.NewGetProjectParams().WithID("5b3b558e52faff0006df1597"))
	resp, err := client.Project.GetProjects(nil)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%#v\n", *resp.Payload)
	//fmt.Printf("%#v\n", *resp.Payload)
	for _, element := range resp.Payload {
		fmt.Printf("%s, %s\n", *element.ID, *element.Name)
	}*/

	cmd.Execute()
}
