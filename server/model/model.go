package model

type Directories struct {
	DirList []Dir
}

type File struct {
	Name       string
	ThumbPath  string
	ActualPath string
	IsDir      bool
}

type FilesResponse struct {
	Files []File
	Next  string
}

type Dir struct {
	Id   string
	Path string
}

type VarsConfig struct {
	Dirs []string `json:"dirs"`
}
