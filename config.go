package envutils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	tagName    = `env`
	defaultEnv = `.env`
)

func ParseConfig(input interface{}, opts ...Option) (err error) {
	config := &Config{Filenames: []string{defaultEnv}}
	for _, opt := range opts {
		opt(config)
	}
	for _, filename := range config.Filenames {
		err = loadEnv(filename)
		if err != nil {
			return err
		}
	}
	return parseEnvironmentConfig(input)
}

func loadEnv(fileName string) error {
	if fileName == "" {
		log.Println("no files to load")
		return nil
	}
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value := scanner.Text()
		if value == "" {
			continue
		}
		if !strings.Contains(value, "=") {
			return fmt.Errorf("missing the necessary characters '='")
		}
		err := os.Setenv(value[0:strings.Index(value, "=")], value[strings.Index(value, "=")+1:])
		if err != nil {
			return err
		}
	}
	return nil
}

func parseEnvironmentConfig(obj interface{}) error {
	objT := reflect.TypeOf(obj)
	objV := reflect.ValueOf(obj)
	switch {
	case isStruct(objT):
	case isStructPtr(objT):
		objT = objT.Elem()
		objV = objV.Elem()
	default:
		return fmt.Errorf("%v must be a struct or a struct pointer", obj)
	}
	for i := 0; i < objT.NumField(); i++ {
		field := objT.Field(i)
		tag := field.Tag.Get(tagName)
		if tag == "" {
			continue
		}
		envValue := os.Getenv(tag)
		fieldValue := reflect.ValueOf(obj).Elem()
		switch field.Type.Kind() {
		case reflect.String:
			fieldValue.FieldByName(field.Name).SetString(envValue)
		case reflect.Int:
			i, err := strconv.Atoi(envValue)
			if err != nil {
				return fmt.Errorf("format conversion error. [%s]", tag)
			}
			fieldValue.FieldByName(field.Name).SetInt(int64(i))
		case reflect.Int64:
			i, err := strconv.ParseInt(envValue, 10, 64)
			if err != nil {
				return fmt.Errorf("format conversion error. [%s]", tag)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Int32:
			i, err := strconv.ParseInt(envValue, 10, 32)
			if err != nil {
				return fmt.Errorf("format conversion error. [%s]", tag)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Int16:
			i, err := strconv.ParseInt(envValue, 10, 16)
			if err != nil {
				return fmt.Errorf("format conversion error. [%s]", tag)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Int8:
			i, err := strconv.ParseInt(envValue, 10, 8)
			if err != nil {
				return fmt.Errorf("format conversion error. [%s]", tag)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Bool:
			b, err := strconv.ParseBool(envValue)
			if err != nil {
				return fmt.Errorf("format conversion error. [%s]", tag)
			}
			fieldValue.FieldByName(field.Name).SetBool(b)
		case reflect.Float64:
			f, err := strconv.ParseFloat(envValue, 64)
			if err != nil {
				return fmt.Errorf("format conversion error. [%s]", tag)
			}
			fieldValue.FieldByName(field.Name).SetFloat(f)
		case reflect.Float32:
			f, err := strconv.ParseFloat(envValue, 32)
			if err != nil {
				return fmt.Errorf("format conversion error. [%s]", tag)
			}
			fieldValue.FieldByName(field.Name).SetFloat(f)
		default:
			return fmt.Errorf("not support %s", field.Type.Kind().String())
		}
	}
	return nil
}

func isStruct(t reflect.Type) bool {
	return t.Kind() == reflect.Struct
}

func isStructPtr(t reflect.Type) bool {
	return t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct
}
