package image

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/fsouza/go-dockerclient"
	"github.com/kube-ops/pops/image/login"
	"github.com/kube-ops/pops/properties"
	"github.com/sirupsen/logrus"
)

// DockerImage docker image type
type DockerImage struct {
	Name     string `yaml:"name"`
	Registry string `yaml:"registry"`
	Tag      string `yaml:"tag"`
	Path     string `yaml:"path,omitempty"`
}

// NewDockerImage create a new docker image
func NewDockerImage(name string, registry string, tag string, path string) *DockerImage {
	return &DockerImage{name, registry, tag, path}
}

// NewDockerImageFromPath create a new docker image from a docker path
func NewDockerImageFromPath(dockersDir string, name string) (*DockerImage, error) {
	dockerPath := path.Join(dockersDir, name)
	configPath := path.Join(dockerPath, "version")
	docker := &DockerImage{Path: dockerPath}
	err := properties.GetYAMLProperties(configPath, docker)
	return docker, err
}

// Print pretty display for a docker
func (dockerImage *DockerImage) Print() {
	fmt.Printf("%-55s%-23s%-12s\n", dockerImage.Registry, dockerImage.Name, dockerImage.Tag)
}

// Build build docker image
func (dockerImage *DockerImage) Build() {
	client := getClient()

	t := time.Now()
	inputbuf := bytes.NewBuffer(nil)
	tr := tar.NewWriter(inputbuf)
	defer func() {
		err := tr.Close()
		if err != nil {
			logrus.Fatal(err)
		}
	}()
	err := filepath.Walk(dockerImage.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logrus.Fatal(err)
		}
		if !info.IsDir() {
			bytes := getBytes(path)
			addToTar(tr, path[len(dockerImage.Path)+1:], t, bytes, int64(info.Mode()))
		}
		return nil
	})
	if err != nil {
		logrus.Fatal(err)
	}

	opts := docker.BuildImageOptions{
		Name:         dockerImage.Registry + "/" + dockerImage.Name + ":" + dockerImage.Tag,
		InputStream:  inputbuf,
		OutputStream: os.Stdout,
	}
	logrus.WithFields(
		logrus.Fields{
			"name":     dockerImage.Name,
			"tag":      dockerImage.Tag,
			"registry": dockerImage.Registry,
		}).Println("Building image")
	if err = client.BuildImage(opts); err != nil {
		logrus.Fatal(err)
	}
}

// Publish publish docker image to remote registry
func (dockerImage *DockerImage) Publish() {
	client := getClient()
	opts := docker.PushImageOptions{
		Name:         dockerImage.Registry + "/" + dockerImage.Name + ":" + dockerImage.Tag,
		OutputStream: os.Stdout,
	}
	logrus.WithFields(
		logrus.Fields{
			"name":     dockerImage.Name,
			"tag":      dockerImage.Tag,
			"registry": dockerImage.Registry,
		}).Println("Publishing image")
	if err := client.PushImage(opts, login.GetAWSCredentials()); err != nil {
		logrus.Fatal(err)
	}
}

// IsDockerDir return true if the directory contains a version file
func IsDockerDir(dir string) bool {
	configPath := path.Join(dir, "version")
	_, err := os.Stat(configPath)
	return err == nil
}

// ListImages return an array of dockers found in dockersDir
func ListImages(dockersDir string) []DockerImage {
	files, err := ioutil.ReadDir(dockersDir)
	if err != nil {
		logrus.Fatal(err)
	}
	var dockers []DockerImage
	for _, file := range files {
		if file.IsDir() {
			absPath := path.Join(dockersDir, file.Name())
			if IsDockerDir(absPath) {
				docker, err := NewDockerImageFromPath(dockersDir, file.Name())
				if err != nil {
					logrus.Warn(err)
				} else {
					dockers = append(dockers, *docker)
				}
			}
		}
	}
	return dockers
}

// PrintList print the list of dockers definitions found in dockersDir
func PrintList(dockersDir string) {
	dockers := ListImages(dockersDir)
	fmt.Printf("%-55s%-23s%-12s\n", "REGISTRY", "NAME", "TAG")
	for _, docker := range dockers {
		docker.Print()
	}
}

func addToTar(tr *tar.Writer, filename string, t time.Time, file []byte, mode int64) {
	err := tr.WriteHeader(&tar.Header{Name: filename, Size: int64(len(file)), ModTime: t, AccessTime: t, ChangeTime: t, Mode: mode})
	if err != nil {
		logrus.Fatal(err)
	}
	_, err = tr.Write(file)
	if err != nil {
		logrus.Fatal(err)
	}
}

func getClient() *docker.Client {
	client, err := docker.NewClient("unix:///var/run/docker.sock")
	if err != nil {
		logrus.Fatal(err)
	}
	return client
}

func getBytes(path string) []byte {
	fileReader, err := os.Open(path)
	if err != nil {
		logrus.Fatal(err)
	}
	readFile, err := ioutil.ReadAll(fileReader)
	if err != nil {
		logrus.Fatal(err)
	}
	return readFile
}
