package utils

import (
	"PPYun/server/config"
	_ "go.uber.org/zap"
	"io/ioutil"
	"os"
	"strings"
)

//PathExists 判断所给路径文件/文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//UpDateUserFilesPath 更新用户的文件树
func UpDateUserFilesPath(username string) error {
	var files config.Folder
	basePath := config.File_path + "\\" + username
	dfsFiles(basePath, "", &files)
	jsonAns, err := MarshalFiles(&files)
	if err != nil {
		return err
	}
	err = WriteData(username, jsonAns)
	return err
}

//dfsFiles 深度优先搜索文件
func dfsFiles(basepath string, path string, data *config.Folder) {

	files, err := ioutil.ReadDir(basepath + path)
	if err != nil {
		Zap().Info(err.Error())
	}

	(*data).FilePath = path      //该目录的文件路径
	(*data).FileNum = len(files) //该目录文件下的文件数量
	fileSize := 0

	for _, file := range files {
		if file.IsDir() {
			(*data).FolderSons = append((*data).FolderSons, config.Folder{FilePath: path + "\\" + file.Name()})
			dfsFiles(basepath, path+"\\"+file.Name(), &data.FolderSons[len(data.FolderSons)-1])
		} else {

			md5 := GetMD5ByPath(path + "\\" + file.Name())

			(*data).Files = append((*data).Files, config.File{
				FileName:          file.Name(),
				FilePath:          path + "\\" + file.Name(),
				FileSiz:           int(file.Size()),
				FileType:          GetFileType(strings.Split(file.Name(), ".")[len(strings.Split(file.Name(), "."))-1]),
				AllCanDownload:    false,
				OnlyDownloadOwner: true,
				CanDownloadOnLink: true,
				MD5:               md5,
			}) //默认自己可以下载文件，默认可以通过链接分享下载
			fileSize += GetFileSizeByMD5(md5)
		}
	}

	(*data).FileSize = fileSize
}

//GetFileType 根据文件后缀返回文件类型
func GetFileType(t string) config.FileKind {
	switch t {
	case "exe":
		return config.EXE
	case "txt":
		return config.TXT
	case "png":
		return config.PNG
	case "bat":
		return config.BAT
	case "lnk":
		return config.LNK
	case "doc":
		return config.DOC
	case "ppt":
		return config.PPT
	case "pdf":
		return config.PDF
	default:
		return config.OTHER
	}
	return config.OTHER
}

//GetSuffix 获取文件格式
func GetSuffix(name string) string {
	return strings.Split(name, ".")[len(strings.Split(name, "."))-1]
}
