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

	"github.com/coreos/go-semver/semver"
	"github.com/fsouza/go-dockerclient"
	"github.com/kube-ops/pops/image/login"
	"github.com/kube-ops/pops/properties"
	log "github.com/sirupsen/logrus"
)

// DockerImage docker image type
type DockerImage struct {
	Name     string `yaml:"name"`
	Registry string `yaml:"registry"`
	Version  string `yaml:"version"`
	Path     string `yaml:"path,omitempty"`
}

// NewDockerImage create a new docker image
func NewDockerImage(name string, registry string, version string, path string) *DockerImage {
	return &DockerImage{name, registry, version, path}
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
	fmt.Printf("%-55s%-23s%-12s\n", dockerImage.Registry, dockerImage.Name, dockerImage.Version)
}

// Create create the layout directory in dockersDir
func (dockerImage *DockerImage) Create(dockersDir string) {
	dockerImage.SaveToFile(dockersDir)
	dockerPath := path.Join(dockersDir, dockerImage.Name)
	dockerfilePath := path.Join(dockerPath, "Dockerfile")
	err := ioutil.WriteFile(dockerfilePath, []byte("FROM alpine:3.6"), os.FileMode(uint32(0664)))
	if err != nil {
		log.Fatal(err)
	}
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
			log.Fatal(err)
		}
	}()
	err := filepath.Walk(dockerImage.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			bytes := getBytes(path)
			addToTar(tr, path[len(dockerImage.Path)+1:], t, bytes, int64(info.Mode()))
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	opts := docker.BuildImageOptions{
		Name:         dockerImage.Registry + "/" + dockerImage.Name + ":" + dockerImage.Version,
		InputStream:  inputbuf,
		OutputStream: os.Stdout,
	}
	log.WithFields(
		log.Fields{
			"name":     dockerImage.Name,
			"version":  dockerImage.Version,
			"registry": dockerImage.Registry,
		}).Println("Building image")
	if err = client.BuildImage(opts); err != nil {
		log.Fatal(err)
	}
}

// Publish publish docker image to remote registry
func (dockerImage *DockerImage) Publish() {
	client := getClient()
	opts := docker.PushImageOptions{
		Name:         dockerImage.Registry + "/" + dockerImage.Name + ":" + dockerImage.Version,
		OutputStream: os.Stdout,
	}
	log.WithFields(
		log.Fields{
			"name":     dockerImage.Name,
			"version":  dockerImage.Version,
			"registry": dockerImage.Registry,
		}).Println("Publishing image")
	if err := client.PushImage(opts, login.GetAWSCredentials()); err != nil {
		log.Fatal(err)
	}
}

// SaveToFile save dockerImage to its version file
func (dockerImage *DockerImage) SaveToFile(dockersDir string) {
	dockerPath := path.Join(dockersDir, dockerImage.Name)
	configPath := path.Join(dockerPath, "version")
	_, err := os.Stat(dockerPath)
	if os.IsNotExist(err) {
		err = os.Mkdir(dockerPath, os.FileMode(uint32(0775)))
		if err != nil {
			log.Fatal(err)
		}
	}
	// Do not save path
	dockerImage.Path = ""
	err = properties.SetYAMLProperties(configPath, dockerImage)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
	var dockers []DockerImage
	for _, file := range files {
		if file.IsDir() {
			absPath := path.Join(dockersDir, file.Name())
			if IsDockerDir(absPath) {
				docker, err := NewDockerImageFromPath(dockersDir, file.Name())
				if err != nil {
					log.Warn(err)
				} else {
					dockers = append(dockers, *docker)
				}
			}
		}
	}
	return dockers
}

// BumpVersion bump the right version number according to importance
func (dockerImage *DockerImage) BumpVersion(importance string) {
	version := semver.New(dockerImage.Version)
	switch importance {
	case "major":
		version.BumpMajor()
	case "minor":
		version.BumpMinor()
	default:
		version.BumpPatch()
	}
	dockerImage.Version = version.String()
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
		log.Fatal(err)
	}
	_, err = tr.Write(file)
	if err != nil {
		log.Fatal(err)
	}
}

func getClient() *docker.Client {
	client, err := docker.NewClient("unix:///var/run/docker.sock")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func getBytes(path string) []byte {
	fileReader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	readFile, err := ioutil.ReadAll(fileReader)
	if err != nil {
		log.Fatal(err)
	}
	return readFile
}
