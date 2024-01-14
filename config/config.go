package config

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// redis 配置管理记录
// 比如 端口 ip 等一些redis所需的配置信息 都在config.go中体现

// 常量
var Properties *ServerProperties

// 服务端配置 数据结构
type ServerProperties struct {
	Bind       string `cfg:"bind"`
	Port       int    `cfg:"port"`
	Dir        string `cfg:"dir"`
	MaxClients int    `cfg:"maxclients"`
	Databases  int    `cfg:"databases"`
}

// 初始化默认的服务端配置
func init() {
	// 默认的配置
	Properties = &ServerProperties{
		Bind: "127.0.0.1",
		Port: 6379,
	}
}

// 配置文件的解析
func parse(src io.Reader) *ServerProperties {
	config := new(ServerProperties)
	rawMap := make(map[string]string)
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && strings.TrimLeft(line, " ")[0] == '#' {
			continue
		}
		pivot := strings.IndexAny(line, " ")
		if pivot > 0 && pivot < len(line)-1 { // separator found
			key := line[0:pivot]
			value := strings.Trim(line[pivot+1:], " ")
			rawMap[strings.ToLower(key)] = value
		}
	}
	if err := scanner.Err(); err != nil {
		errors.New("读取配置文件异常")
	}
	// 通过反射获取config的 cfg名称 在通过反射赋值  完成配置文件的解析
	t := reflect.TypeOf(config)
	v := reflect.ValueOf(config)
	n := t.Elem().NumField()
	for i := 0; i < n; i++ {
		field := t.Elem().Field(i)
		fieldVal := v.Elem().Field(i)
		key, ok := field.Tag.Lookup("cfg")
		if !ok || strings.TrimLeft(key, " ") == "" {
			key = field.Name
		}
		value, ok := rawMap[strings.ToLower(key)]
		if ok {
			switch field.Type.Kind() {
			case reflect.String:
				fieldVal.SetString(value)
			case reflect.Int:
				intValue, err := strconv.ParseInt(value, 10, 64)
				if err == nil {
					fieldVal.SetInt(intValue)
				}
			case reflect.Bool:
				boolValue := "yes" == value
				fieldVal.SetBool(boolValue)
			case reflect.Slice:
				if field.Type.Elem().Kind() == reflect.String {
					slice := strings.Split(value, ",")
					fieldVal.Set(reflect.ValueOf(slice))
				}
			}
		}
	}
	return config
}

func ReadFilePath(filePath string) *ServerProperties {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("读取file path error")
	}
	return parse(file)
}
