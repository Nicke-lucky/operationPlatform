package service

import "operationPlatform/utils"

//文件上传
var DirPath string

func FileUpload(file []byte, fname string) {
	//log.Println("filebyte:", string(file[:50]))
	//http.HandleFunc("/upload", upload)

	//文件在服务器上的位置
	fpath := DirPath
	//"/root/test/info.json"
	utils.UploadFile(file, fpath+fname, fname)
}

//
//func upload(w http.ResponseWriter, r *http.Request) {
//	r.ParseMultipartForm(32 << 20)
//	file, handler, err := r.FormFile("uploadfile")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()
//	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer f.Close()
//	io.Copy(f, file)
//	fmt.Fprintln(w, "upload ok!")
//}
