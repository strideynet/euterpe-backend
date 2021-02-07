package main

import (
	"euterpe/lib/log"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

var l = log.Package("gendockercompose")

type Service struct {
	Name string
}

func getServices(dir string) ([]Service, error) {
	contents, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	services := []Service{}
	for _, cont := range contents {
		if cont.IsDir() {
			if strings.Contains(cont.Name(), "service.") {
				name := strings.TrimPrefix(cont.Name(), "service.")
				services = append(services, Service{Name: name})
			}
		}
	}

	return services, nil
}

func gen(writer io.Writer, svcs []Service) error {
	// TODO: be less dumb about working directory
	tpl, err := template.ParseFiles("./cmd/gendockercompose/docker-compose.yml.tmpl")
	if err != nil {
		return err
	}

	return tpl.Execute(writer, svcs)
}

// TODO: error handle rather than panic lol
func main() {
	svcs, err := getServices("./")
	if err != nil {
		panic(err)
	}
	l.Info("svces", zap.Any("svces", svcs))

	f, err := os.Create("./docker-compose.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = gen(f, svcs)
	if err != nil {
		panic(err)
	}
}
