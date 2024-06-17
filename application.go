package summergo_web

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"path/filepath"
)

// AppConfig 是配置结构体
type Application struct {
	Server struct {
		Port int    `yaml:"port"`
		Name string `yaml:"name"`
	} `yaml:"server"`
}

// LoadConfigFromDir 从指定文件夹及其子文件夹中加载配置信息
// 如果 dirPath 为空，则默认为当前项目根目录
func loadApplicationYamlFromDir(dirPath string) (*Application, error) {
	if dirPath == "" {
		dirPath = "./"
	}

	var cfg *Application
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == "application.yaml" {
			// 找到 application.yaml 文件，读取配置信息
			yamlFile, err := os.ReadFile(path)
			if err != nil {
				log.Fatal("寻找yaml文件失败", err)
				return err
			}
			var a Application
			err = yaml.Unmarshal(yamlFile, &a)
			if err != nil {
				return err
			}
			cfg = &a
			return nil
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func getServerName() string {

	// 提取根目录的名称
	rootDirName := "DefaultServiceName"

	log.Println("根目录名称:", rootDirName)
	return rootDirName
}
func ApplicationInit() (*Application, error) {
	dir, err := loadApplicationYamlFromDir("")
	if err != nil {
		log.Fatal("初始化application.yaml报错:", err)
	}
	if dir != nil {
		log.Print("成功初始化application.yaml：", dir)
		if dir.Server.Port == 0 {
			dir.Server.Port = 8080
		}
		if dir.Server.Name != "" {
			log.Print("服务名为：", dir.Server.Name)
		}
		if dir.Server.Name == "" {
			dir.Server.Name = getServerName()
		}
	}
	return dir, err
}
