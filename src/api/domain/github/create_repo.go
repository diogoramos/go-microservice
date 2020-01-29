package github

/*
{
 	"name": "Hello-world",
	"description":"This is my first  repository",
	"homepage":"https://github.com",
	"private":false,
	"has_issues": true,
	"has_projects": true,
	"has_wiki": true
}
 */

type CreateRepoRequest struct {
	Name        string	`json:"name"`
	Description string	`json:"description"`
	Homepage    string	`json:"homepage"`
	Private     bool	`json:"private"`
	HasIssues   bool	`json:"has_issues"`
	HasProjects bool	`json:"has_projects"`
	HasWiki     bool	`json:"has_wiki"`
}




/*
//Don't do it please
func () CreateRepo(){
	request := map[string]interface{}{
		"name": "Hello-World",
		"private": false,
	}

	private := request["private"].(bool)
}
*/