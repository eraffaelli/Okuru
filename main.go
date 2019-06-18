package main

//https://echo.labstack.com/guide
//https://github.com/verybluebot/echo-server-tutorial/

import (
	"github.com/eraffaelli/Okuru/router"
	. "github.com/eraffaelli/Okuru/utils"
	"github.com/flosch/pongo2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"math/rand"
	"os"
	"time"
)

var DebugLevel bool

func Flags() {
	pflag.BoolVar(&DebugLevel, "debug", false, "--debug")
	defer pflag.Parse()
	return
}

func init() {
	Flags()

	pool := NewPool()
	c := pool.Get()
	defer c.Close()
	if Ping(c) == false {
		log.Panic("Redis problem")
	}

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)


	if DebugLevel == true {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	go CleanFileWatch()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	e:= router.New()
	err := pongo2.DefaultLoader.SetBaseDir("views")
	if err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":" + APP_PORT))
}