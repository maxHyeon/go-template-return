package main

import "C"
import (
	"fmt"
	"os"
	"io"
	"path"
	"gopkg.in/yaml.v2"
	"text/template"
	"io/ioutil"
        "bytes"
)

type store struct{
	FileName string ;
	Values map[string]interface{};
}

func tpl(fileName string, vals interface{}, output string) (string,error) {
	name := path.Base(fileName)
	tmpl, err := template.New(name).ParseFiles(fileName)
	if err != nil {
		return "",err
	}

	var file io.Writer
	if output != "" && output != "return" {
		f, _ :=os.Create(output)
		defer f.Close()
		file = f
	} else if output == "return" {
           
        }  else {
		file = os.Stdout
	}
        
        var rslt bytes.Buffer
        if output == "return" {
            err = tmpl.Execute(&rslt, vals)
            return rslt.String(), nil
        } else {
	    err = tmpl.Execute(file, vals)
        }
	if err != nil {
		return "", err
	}
	return "", nil
}

func (s *store)getValues() {
    yamlFile, err := ioutil.ReadFile(s.FileName)
    if err != nil {
        fmt.Printf("yamlFile.Get err   #%v ", err)
	}
    err = yaml.Unmarshal(yamlFile, &s.Values)
    if err != nil {
		panic(err)
    }
}

//export RenderTemplate
func RenderTemplate(template, fileName, output string) *C.char {
	s := store{FileName: fileName}
	s.getValues()
	rslt, err := tpl(template, s.Values, output)
	if err != nil {
		panic(err)
    }
    return C.CString(rslt)
}

func main(){
	//RenderTemplate("sample.tmpl", "values.yml", "")
}
