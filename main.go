package main

import (
	"filehandlers"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"zinc_handler"

	"github.com/joho/godotenv"
)

func uploadFilesIntoZinc(path string) {

	defer log.Default().Printf("Files were uploaded.")

	indexHandler, err := zinc_handler.NewIndexHandler("./static/standard_index_structure.json")
	if err != nil {
		log.Fatal(err)
	}
	indexHandler.CreateIndex()
	err = filehandlers.FolderInsert(path)
	if err != nil {
		log.Default().Printf("Error trying to insert, %v", err)
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()

	path_arg := 1
	if *cpuprofile != "" {
		path_arg++
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	godotenv.Load()
	godotenv.Load(".env")

	if len(os.Args) > path_arg {
		log.Default().Println("Start to insert")
		uploadFilesIntoZinc(os.Args[path_arg])
	}
}
