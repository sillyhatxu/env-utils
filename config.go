package envutils

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const (
	tagName  = `env`
	localEnv = `.env`
)

func ParseConfig(input interface{}, opts ...Option) (err error) {
	config := &Config{}
	for _, opt := range opts {
		opt(config)
	}
	err = loadEnv(localEnv, false)
	if err != nil {
		return err
	}
	err = loadEnv(config.FileName, true)
	if err != nil {
		return err
	}
	return parseEnvironmentConfig(input)
}

func loadEnv(fileName string, required bool) error {
	if fileName == "" {
		logrus.Infof("No local file to load. [%s]", fileName)
		return nil
	}
	logrus.Infof("Load env file : %s", fileName)
	file, err := os.Open(fileName)
	if err != nil && required {
		panic(err)
	} else if err != nil && !required {
		logrus.Infof("No local file to load. [%s]", fileName)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		array := strings.Split(scanner.Text(), "=")
		if array == nil || len(array) != 2 {
			return fmt.Errorf("split error. %s", scanner.Text())
		}
		err := os.Setenv(array[0], array[1])
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
		logrus.Debugf("tag : %s; value : %s", tag, envValue)
		fieldValue := reflect.ValueOf(obj).Elem()
		switch field.Type.Kind() {
		case reflect.String:
			fieldValue.FieldByName(field.Name).SetString(envValue)
		case reflect.Int:
			i, err := strconv.Atoi(envValue)
			if err != nil {
				return fmt.Errorf("strconv.Atoi(%s) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetInt(int64(i))
		case reflect.Int64:
			i, err := strconv.ParseInt(envValue, 10, 64)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Int32:
			i, err := strconv.ParseInt(envValue, 10, 32)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Int16:
			i, err := strconv.ParseInt(envValue, 10, 16)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Int8:
			i, err := strconv.ParseInt(envValue, 10, 8)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetInt(i)
		case reflect.Bool:
			b, err := strconv.ParseBool(envValue)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetBool(b)
		case reflect.Float64:
			f, err := strconv.ParseFloat(envValue, 64)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
			}
			fieldValue.FieldByName(field.Name).SetFloat(f)
		case reflect.Float32:
			f, err := strconv.ParseFloat(envValue, 32)
			if err != nil {
				return fmt.Errorf("strconv.ParseInt(%v, 10, 64) error. %v", envValue, err)
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
