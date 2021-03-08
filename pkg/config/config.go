package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/mrxk/dev-env/pkg/constants"
)

type Config struct {
	Version string          `json:"version"`
	Envs    map[string]*Env `json:"envs"`
}

type Env struct {
	ContainerArgs  []string          `json:"containerArgs"`
	Name           string            `json:"name"`
	Options        map[string]string `json:"options"`
	dockerBuildDir string
}

func NewConfig() (Config, error) {
	c := Config{
		Version: "1",
		Envs: map[string]*Env{
			constants.DefaultEnvironment: NewEnv(constants.DefaultEnvironment),
		},
	}
	err := c.Read()
	return c, err
}

func NewConfigFor(name string) (Config, error) {
	c := Config{
		Version: "1",
		Envs: map[string]*Env{
			name: NewEnv(name),
		},
	}
	err := c.Read()
	return c, err
}

func NewEnv(name string) *Env {
	return &Env{
		ContainerArgs:  constants.DefaultContainerArgs,
		Name:           namesgenerator.GetRandomName(0),
		Options:        map[string]string{},
		dockerBuildDir: name,
	}
}

func (c *Config) AddEnvIfNotExistFor(name string) {
	_, ok := c.Envs[name]
	if ok {
		return
	}
	c.Envs[name] = NewEnv(name)
}

func (c *Config) Read() error {
	if c == nil {
		return nil
	}
	configDir, err := GetConfigDir()
	if err != nil {
		return err
	}
	configpaths, err := filepath.Glob(filepath.Join(configDir, "*", constants.ConfigFile))
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

func (c *Config) EnvFor(envName string) (*Env, error) {
	env, ok := c.Envs[envName]
	if !ok {
		return nil, fmt.Errorf("No such environment found: %s", envName)
	}
	return env, nil
}

func (c *Config) WriteIfNotExist() error {
	for name, e := range c.Envs {
		content, err := json.MarshalIndent(e, "", "    ")
		if err != nil {
			return err
		}
		writeConfigFileIfNotExist(name, constants.ConfigFile, content)
		writeConfigFileIfNotExist(
			e.dockerBuildDir,
			constants.DockerFile,
			[]byte(constants.DefaultDockerFile))

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
	return e.Name + ":" + constants.DevEnvTag
}

func (e *Env) ContainerName() string {
	return e.Name + "_" + constants.DevEnvTag
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
	path, err := GetProjectRoot()
	if err != nil {
		return "", nil
	}
	configDir := filepath.Join(path, constants.ConfigDir)
	err = os.MkdirAll(configDir, 0700)
	if err != nil {
		return "", err
	}
	return configDir, nil
}

func GetProjectRoot() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	candidatePath, err := filepath.Abs(cwd)
	if err != nil {
		return cwd, nil
	}
	candidatePath = filepath.Clean(candidatePath)
	for {
		candidateProjectRoot := filepath.Join(candidatePath, constants.ConfigDir)
		_, err = os.Stat(candidateProjectRoot)
		if err == nil {
			return candidatePath, nil
		}
		parentPath := filepath.Dir(candidatePath)
		if candidatePath == parentPath {
			break
		}
		candidatePath = parentPath
	}
	return cwd, nil
}

func writeConfigFileIfNotExist(dir, filename string, content []byte) error {
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
