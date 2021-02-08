package main

import (
	"euterpe/lib/log"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"
)

var l = log.Package("gendockercompose")

type TemplateData struct {
	Time     string
	Services []Service
}

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

func gen(writer io.Writer, data TemplateData) error {
	// TODO: be less dumb about working directory
	tpl, err := template.ParseFiles("./cmd/gendockercompose/docker-compose.yml.tmpl")
	if err != nil {
		return err
	}

	return tpl.Execute(writer, data)
}

func run() error {
	svcs, err := getServices("./")
	if err != nil {
		return err
	}
	l.Info("detected services", zap.Any("services", svcs))

	f, err := os.Create("./docker-compose.yml")
	if err != nil {
		return err
	}
	defer f.Close()

	return gen(f, TemplateData{
		Time:     time.Now().UTC().String(),
		Services: svcs,
	})
}

// TODO: error handle rather than panic lol
func main() {
	err := run()
	if err != nil {
		l.Fatal("an error occurred generating the docker-compose", zap.Error(err))
	}
}
