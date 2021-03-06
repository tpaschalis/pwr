package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"pwr/pwr_functions"
)

type config struct {
	storagePath string
	editor      string
}

func main() {
	setupFlags(flag.CommandLine)
	confFlag := flag.String("c", "", "-c					Change path for config.json")
	renderFlag := flag.Bool("v", false, "-v                    Render notes so they're accessible by a browser")
	newPageFlag := flag.String("new", "", "-new		create a new empty note page")
	delPageFlag := flag.String("del", "", "-del		delete a page and all of its contents")

	flag.Parse()
	args := flag.Args()

	confFn := os.Getenv("PWR_PATH_TO_CONFIG")
	if confFn == "" {
		confFn = "config.json"
	}
	byt, err := ioutil.ReadFile(confFn)
	check(err)

	var conf map[string]interface{}
	err = json.Unmarshal(byt, &conf)
	check(err)
	storagePath := conf["storagePath"].(string)
	templatesPath := conf["templatesPath"].(string)
	editor := conf["editor"].(string)
	editorArgs := conf["editorArgs"].(string)

	//fmt.Println(editor, editorArgs)
	//err = checkEarlyExit(args, *confFlag)
	//check(err)
	if *renderFlag == true {
		pwrf.BuildIndex(storagePath, templatesPath)
		pwrf.RenderPages(storagePath, templatesPath)
		os.Exit(0)
	}

	if *newPageFlag != "" {
		createPage := strings.ToLower(*newPageFlag)
		pwrf.CreateEmptyPage(storagePath, createPage)
		os.Exit(0)
	}


	if *delPageFlag != "" {
		deletePage := strings.ToLower(*delPageFlag)
		pwrf.DeleteNamedPage(storagePath, deletePage)
		os.Exit(0)
	}

	if len(args) == 0 || strings.ToLower(args[0]) == "today" {
		pwrf.OpenTodayPage(storagePath, editor, editorArgs)
	}

	if len(args) != 0 {
		switch strings.ToLower(args[0]) {
		case "yesterday":
			pwrf.OpenYesterdayPage(storagePath, editor, editorArgs)
		default:
			pwrf.OpenNamedPage(storagePath, strings.ToLower(args[0]), editor, editorArgs)
		}
	}

	_ = confFlag
	fmt.Println("")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}



func setupFlags(f *flag.FlagSet) {
	f.Usage = func() {
		fmt.Println("\nSome Message")

		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		f.PrintDefaults()

		fmt.Println("\nExplanatory stuff")
	}
}

func checkEarlyExit(args []string, c string) (err error) {
	if len(args) > 1 {
		fmt.Println("Couldn't understand what you're trying to do")
		os.Exit(1)
		flag.Usage()
	}
	//validArg := strings.ToLower(args[0])
	//switch validArg {
	//case
	//    "today",
	//    "tomorrow",
	//    "yesterday",
	//    "todo":
	//    return nil
	//}
	return fmt.Errorf("No such argument is available")
}

