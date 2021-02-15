package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/docker/docker/pkg/namesgenerator"
)

const (
	ConfigDir  = ".dev-env"
	ConfigFile = "dev-env.json"
	DockerFile = "Dockerfile"
	devEnvTag  = "dev_env"
)

type Config struct {
	Version string          `json:"version"`
	Envs    map[string]*Env `json:"envs"`
}

type Env struct {
	ContainerArgs  []string `json:"containerArgs"`
	Name           string   `json:"name"`
	dockerBuildDir string
}

func NewConfig() (Config, error) {
	c := Config{
		Version: "1",
		Envs: map[string]*Env{
			"main": {
				Name:           namesgenerator.GetRandomName(0),
				dockerBuildDir: "main",
			},
		},
	}
	err := c.Read()
	return c, err
}

func (c *Config) Read() error {
	if c == nil {
		return nil
	}
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}
	configpaths, err := filepath.Glob(filepath.Join(configDir, "*", ConfigFile))
	if err != nil {
		return err
	}
	for _, configpath := range configpaths {
		data, err := ioutil.ReadFile(configpath)
		if err != nil {
			return err
		}
		e := Env{}
		err = json.Unmarshal(data, &e)
		if err != nil {
			return err
		}
		name := filepath.Base(filepath.Dir(configpath))
		e.dockerBuildDir = name
		c.Envs[name] = &e
	}
	return nil
}

func (c *Config) Write() error {
	return c.write(true)
}

func (c *Config) WriteIfNotExist() error {
	return c.write(false)
}

func (c *Config) EnvFor(envName string) (*Env, error) {
	env, ok := c.Envs[envName]
	if !ok {
		return nil, fmt.Errorf("No such environment found: %s", envName)
	}
	return env, nil
}

func (c *Config) write(force bool) error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}
	for name, e := range c.Envs {
		dirpath := filepath.Join(configDir, name)
		err = os.MkdirAll(dirpath, 0700)
		if err != nil {
			continue
		}
		datapath := filepath.Join(dirpath, ConfigFile)
		_, err = os.Stat(datapath)
		if err == nil { // file exists
			if !force { // if not forcing then done
				continue
			}
		} else if !os.IsNotExist(err) {
			return err
		}
		// either err == nil && force == true or err == NotExist
		data, err := json.MarshalIndent(e, "", "    ")
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(datapath, data, 0600)
		if err != nil {
			return err
		}

	}
	return nil
}

func (e *Env) DockerBuildDir() (string, error) {
	configDir, err := GetConfigDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(configDir, e.dockerBuildDir)
	err = os.MkdirAll(path, 0700)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (e *Env) ImageNameAndTag() string {
	return e.Name + ":" + devEnvTag
}

func (e *Env) ContainerName() string {
	return e.Name + "_" + devEnvTag
}

func (e *Env) WithName(name string) *Env {
	newEnv := &Env{
		ContainerArgs:  make([]string, len(e.ContainerArgs)),
		Name:           name,
		dockerBuildDir: e.dockerBuildDir,
	}
	copy(newEnv.ContainerArgs, e.ContainerArgs)
	return newEnv
}

func (e *Env) WithContainerArgs(containerArgs []string) *Env {
	newEnv := &Env{
		ContainerArgs:  make([]string, 0, len(e.ContainerArgs)+len(containerArgs)),
		Name:           e.Name,
		dockerBuildDir: e.dockerBuildDir,
	}
	newEnv.ContainerArgs = append(newEnv.ContainerArgs, e.ContainerArgs...)
	newEnv.ContainerArgs = append(newEnv.ContainerArgs, containerArgs...)
	return newEnv
}

func GetConfigDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", nil
	}
	path := GetProjectRoot(cwd)
	configDir := filepath.Join(path, ConfigDir)
	err = os.MkdirAll(configDir, 0700)
	if err != nil {
		return "", err
	}
	return configDir, nil
}

func WriteConfigFileIfNotExist(dir, filename string, content []byte) error {
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}
	configFileDir := filepath.Join(configDir, dir)
	err = os.MkdirAll(configFileDir, 0700)
	if err != nil {
		return err
	}
	configFilePath := filepath.Join(configFileDir, filename)
	_, err = os.Stat(configFilePath)
	if err == nil || os.IsNotExist(err) == false {
		return err
	}
	return ioutil.WriteFile(configFilePath, content, 0600)
}

func GetProjectRoot(path string) string {
	candidatePath, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	candidatePath = filepath.Clean(candidatePath)
	for {
		candidateProjectRoot := filepath.Join(candidatePath, ConfigDir)
		_, err = os.Stat(candidateProjectRoot)
		if err == nil {
			return candidatePath
		}
		parentPath := filepath.Dir(candidatePath)
		if candidatePath == parentPath {
			break
		}
		candidatePath = parentPath
	}
	return path
}
