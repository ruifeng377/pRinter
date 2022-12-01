package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// type uploadFileRequest struct {
// }

// type uploadFileResponse struct {
// }

func (server *Server) saveFormFile(ctx *gin.Context) (string, error) {
	file, err := ctx.FormFile("file")
	if err != nil {
		return "", err
	}

	username := ctx.Request.Form["name"]
	if len(username) == 0 {
		username = []string{ctx.ClientIP()}
	}

	// Upload the file to specific dst.
	dir := filepath.Join(server.store.GetDir(), username[0])
	os.Mkdir(dir, os.ModePerm)
	fileId := uuid.New().String()
	fileType := filepath.Ext(file.Filename)
	filePath := filepath.Join(dir, fmt.Sprintf("%s.%s", fileId, fileType))

	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		return "", nil
	}

	return filePath, nil
}

func (server *Server) uploadFile(ctx *gin.Context) {
	filePath, err := server.saveFormFile(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	ctx.String(http.StatusOK, fmt.Sprintf("The file has been uploaded and is saved as '%s'.", filePath))
}

func (server *Server) printFile(ctx *gin.Context) {
	filePath, err := server.saveFormFile(ctx)
	if err != nil {
		ctx.String(http.StatusBadRequest, "print file err: %s", err.Error())
		return
	}

	out, errout, err := server.printer.Print(filePath)
	if err != nil {
		ctx.String(http.StatusBadRequest, "failed, out: %s, errout: %s, err: %v", out, errout, err)
	}

	ctx.String(http.StatusOK, "The file is printing in progress...")
}

func (server *Server) index(ctx *gin.Context) {
	log.Printf("mainpage")
	ctx.HTML(http.StatusOK, "uploadFile.html", gin.H{
		"title": "pRinter",
	})
}

func (server *Server) getFiles(ctx *gin.Context) {
	var fs http.FileSystem = gin.Dir("./files", true)
	ctx.FileFromFS("/files", fs)
}
