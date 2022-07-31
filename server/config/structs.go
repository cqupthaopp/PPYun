package config

type User struct {
	Username string
	Password string
}

type FileKind uint

const (
	EXE   FileKind = 1
	TXT   FileKind = 2
	PNG   FileKind = 3
	BAT   FileKind = 4
	LNK   FileKind = 5
	DOC   FileKind = 7
	PPT   FileKind = 8
	PDF   FileKind = 9
	OTHER FileKind = 10000
)

type File struct {
	FileName          string   //文件名
	FilePath          string   //文件相对路径
	MD5               string   //MD5值
	FileType          FileKind //文件类型
	FileSiz           int      //文件大小
	AllCanDownload    bool     //是否所有人都可以下载
	CanDownloadOnLink bool     //是否仅可以通过链接下载
	OnlyDownloadOwner bool     //仅主人可见
}

type Folder struct {
	FolderSons []Folder //子文件夹
	Files      []File   //子文件
	FilePath   string   //文件相对路径
	FileSize   int      //该目录下的文件大小
	FileNum    int      //该目录下的文件数量
}

type MD5File struct {
	MD5   string //MD5值
	Size  int    //文件大小
	Count int    //引用量
}

type PathMD5 struct {
	MD5  string //MD5值
	Path string //路径
}
