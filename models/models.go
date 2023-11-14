package model

type UserData struct {
	Name      string `json:"name"`
	Mobile_no string `json:"mobile_no"`
	Email     string `json:"email"`
	Files     []File
}

type FileStorage struct {
	Dir map[string]string
}

type File struct {
	Id       string
	Filename string
}
