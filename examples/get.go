package examples

import "fmt"

type EndPoints struct {
	CurrentUserUrl    string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetEndPoints() (*EndPoints, error) {
	response, err := httpClient.Get("https://api.github.com/", nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(fmt.Printf("Status code %d", response.StatusCode()))
	fmt.Println(fmt.Printf("Status %s ", response.Status()))
	fmt.Println(fmt.Printf("Body %s\n", response.String()))

	var ep EndPoints
	if err := response.UnmarshalJson(&ep); err != nil {
		return nil, err
	}
	fmt.Println(fmt.Printf("Repository URL %s ", ep.RepositoryUrl))
	return &ep, nil
}
