package celeritas

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
}

func (c *Celeritas) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath:    rootPath,
		folderNames: []string{"handers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}
	err := c.Init(pathConfig)
	if err != nil {
		return err
	}
	err = c.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	// read .env
	err = godotenv.Load(rootPath)
	if err != nil {
		return err
	}

	// create loggers
	infoLog, errorLog := c.startLoggers(rootPath)
	c.InfoLog = infoLog
	c.ErrorLog = errorLog
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version

	return nil
}

func (c *Celeritas) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := c.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Celeritas) checkDotEnv(path string) error {
	err := c.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

func (c *Celeritas) startLoggers(path string) (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}
