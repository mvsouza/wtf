package bitbucket

type pullRequestPage struct {
	pagelen int
	size int
	page int
	values []pullRequest
}

type author struct {
	nickname string
	display_name string
}

type pullRequest struct {
	title string
	values []string
	source reference
	target reference
	author author
}

type reference struct{
	commit map[string]interface{}
	repository map[string]interface{}
	branch string
}
