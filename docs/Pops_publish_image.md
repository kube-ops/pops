## Pops publish image

Publish a container image to its repository

### Synopsis

Publish a container image to its repository
  Only docker images are supported for now

```
Pops publish image IMAGE [flags]
```

### Options

```
  -h, --help               help for image
  -i, --image-dir string   Directory containing the Docker images (default "dockers")
```

### Options inherited from parent commands

```
  -b, --main-branch string   The main branch of the repository (default "master")
  -o, --out-dir string       Storage directory for artefacts (default ".out")
  -v, --verbose              Activates verbose mode
```

### SEE ALSO

* [Pops publish](Pops_publish.md)	 - Publish an artifact to its repository

###### Auto generated by spf13/cobra on 21-May-2018