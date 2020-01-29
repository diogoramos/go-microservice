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

type CreateRepoResponse struct {
	Id        	int64		`json:"id"`
	NodeId      string		`json:"node_id"`
	Name        string		`json:"name"`
	FullName    string		`json:"full_name"`
	Private     bool		`json:"private"`
	Owner		RepoOwner 	`json:"owner"`
	Permissions	RepoPermissions `json:"permissions"`
}

type RepoPermissions struct {
	IsAdmin	bool	`json:"admin"`
	HasPush	bool	`json:"push"`
	HasPull	bool	`json:"pull"`
}

type RepoOwner struct {
	Id 			int64 	`json:"id"`
	Login       string	`json:"login"`
	Url        	string	`json:"url"`
	HtmlUrl     string	`json:"html_url"`
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