package generator

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"
)

func GenMapper(m map[string]interface{}, dir, tplfile, fp string) error {
	if err := MkdirIfNotExist(dir); err != nil {
		return err
	}

	content, err := ioutil.ReadFile(tplfile)
	if err != nil {
		return err
	}

	t := template.Must(template.New("").Parse(string(content)))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, m)
	if err != nil {
		return err
	}

	code := formatCode(buffer.String())


	_, err = os.Stat(fp)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsNotExist(err) {
		f, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE, 0664)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(code)

		return err
	}

	data, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fp, []byte(data), os.FileMode(0600))

	return err
}


func formatCode(code string) string {
	ret, err := format.Source([]byte(code))
	if err != nil {
		return code
	}

	return string(ret)
}

func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}