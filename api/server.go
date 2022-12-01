package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ruifeng377/pRinter/db"
	"github.com/ruifeng377/pRinter/printer"
)

type Server struct {
	store   db.Store
	router  *gin.Engine
	printer printer.Printer
}

func NewServer(store db.Store, printer printer.Printer) (*Server, error) {
	server := &Server{
		store:   store,
		printer: printer,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/www/css", "./www/css")
	router.Static("/www/js", "./www/js")

	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.GET("/", server.index)
	router.POST("/uploadFile", server.uploadFile)
	router.POST("/printFile", server.printFile)
	router.GET("/files", server.getFiles)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
