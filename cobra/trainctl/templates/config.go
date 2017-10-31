package templates

type Conf struct {
	Twitter   string
	Author    string
	Email     string
	ModuleDir string
	CourseDir string
}

const Config = `twitter: {{.Twitter}}
author: {{.Author}}
email: {{.Email}}
moduledir: {{.ModuleDir}}
coursedir: {{.CourseDir}}
`
