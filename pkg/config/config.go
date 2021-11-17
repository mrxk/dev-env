package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/docker/docker/pkg/namesgenerator"
	"github.com/mrxk/dev-env/pkg/constants"
)

type Config struct {
	Version string          `json:"version"`
	Envs    map[string]*Env `json:"envs"`
}

type Env struct {
	dockerBuildDir string
	envData        envJSON
}

type envJSON struct {
	ContainerArgs []string          `json:"containerArgs"`
	Name          string            `json:"name"`
	Options       map[string]string `json:"options"`
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
		dockerBuildDir: name,
		envData: envJSON{
			ContainerArgs: constants.DefaultContainerArgs,
			Name:          "",
			Options:       map[string]string{},
		},
	}
}

func (c *Config) EnvFor(envName string) (*Env, error) {
	env, ok := c.Envs[envName]
	if !ok {
		return nil, fmt.Errorf("No such environment found: %s", envName)
	}
	return env, nil
}

func (c *Config) EnvNames() []string {
	names := []string{}
	for name := range c.Envs {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
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
		envData := envJSON{}
		err = json.Unmarshal(data, &envData)
		if err != nil {
			return err
		}
		name := filepath.Base(filepath.Dir(configpath))
		e := Env{
			dockerBuildDir: name,
			envData:        envData,
		}
		c.Envs[name] = &e
	}
	return nil
}

func (c *Config) WriteIfNotExist() error {
	for name, e := range c.Envs {
		content, err := json.MarshalIndent(e.envData, "", "    ")
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

func (e *Env) Name() string {
	if e.envData.Name != "" {
		return e.envData.Name
	}
	e.envData.Name = generateName(e.dockerBuildDir)
	return e.envData.Name
}

func (e *Env) ContainerArgs() []string {
	return e.envData.ContainerArgs
}

func (e *Env) Options() map[string]string {
	return e.envData.Options
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
	return e.Name() + ":" + constants.DevEnvTag
}

func (e *Env) ContainerName() string {
	return e.Name() + "_" + constants.DevEnvTag
}

func (e *Env) WithName(name string) *Env {
	newEnv := &Env{
		dockerBuildDir: e.dockerBuildDir,
		envData: envJSON{
			ContainerArgs: make([]string, len(e.ContainerArgs())),
			Name:          name,
			Options:       make(map[string]string),
		},
	}
	copy(newEnv.envData.ContainerArgs, e.ContainerArgs())
	for k, v := range e.Options() {
		newEnv.envData.Options[k] = v
	}
	return newEnv
}

func (e *Env) WithContainerArgs(containerArgs []string) *Env {
	newEnv := &Env{
		dockerBuildDir: e.dockerBuildDir,
		envData: envJSON{
			ContainerArgs: containerArgs,
			Name:          e.Name(),
			Options:       make(map[string]string),
		},
	}
	for k, v := range e.Options() {
		newEnv.envData.Options[k] = v
	}
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

func generateName(prefix string) string {
	path, err := GetProjectRoot()
	if err != nil {
		return namesgenerator.GetRandomName(0)
	}
	path = filepath.Join(prefix, path)
	return sanitizePath(path)
}

func sanitizePath(path string) string {
	path = strings.ReplaceAll(path, string(os.PathSeparator), "_")
	path = strings.ReplaceAll(path, string(":"), "_")
	path = strings.Trim(path, "_")
	path = strings.ToLower(path)
	return path
}
