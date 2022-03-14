package main

import (
	"flag"
	"log"
	"os"
)

type TargetType int64

const (
	FileObj TargetType = iota
	Directory
)

type arguments struct {
	target     string
	recursive  bool
	targetMode os.FileMode
}

func initArgs() arguments {

	target_flag := flag.String("target", "", "File or directory to inspect.")
	recursive_flag := flag.Bool("R", false, "Inspect directory recursively. Does nothing when target is a file.")

	flag.Parse()
	if *target_flag == "" {
		log.Fatal("No target file or directory specified.")
	}

	fi, err := os.Stat(*target_flag)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("target:", *target_flag)

	return arguments{
		target:     *target_flag,
		recursive:  *recursive_flag,
		targetMode: fi.Mode(),
	}
}
