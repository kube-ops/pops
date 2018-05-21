## Pops build stack

Build a stack artifact

### Synopsis

Build a stack artifact.
	Creates a tgz of the stack description.
  Only helm charts are supported for now

```
Pops build stack [flags]
```

### Options

```
  -s, --chart-dir string   Directory containing the Helm charts (default "charts")
  -h, --help               help for stack
```

### Options inherited from parent commands

```
  -b, --main-branch string   The main branch of the repository (default "master")
  -o, --out-dir string       Storage directory for artefacts (default ".out")
  -v, --verbose              Activates verbose mode
```

### SEE ALSO

* [Pops build](Pops_build.md)	 - Build an artifact to be deployed on Kubernetes

###### Auto generated by spf13/cobra on 21-May-2018