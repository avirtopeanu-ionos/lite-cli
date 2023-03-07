# lite-cli
Generate CLIs at runtime based on swagger spec

### Current State

Currently it generates some commands and subcommands based on the paths:

```bash
 % ./litectl api cloud --help
Use the Compute API (Ionos Cloud V6)

Usage:
  litectl api cloud [flags]

Flags:
  -h, --help   help for cloud

Additional help topics:
  litectl api cloud      Handle  operations
  litectl api cloud backupunits Handle backupunits operations
  litectl api cloud contracts Handle contracts operations
  litectl api cloud datacenters Handle datacenters operations
  litectl api cloud images Handle images operations
  litectl api cloud ipblocks Handle ipblocks operations
  litectl api cloud k8s  Handle k8s operations
  litectl api cloud labels Handle labels operations
  litectl api cloud locations Handle locations operations
  litectl api cloud pccs Handle pccs operations
  litectl api cloud requests Handle requests operations
  litectl api cloud snapshots Handle snapshots operations
  litectl api cloud targetgroups Handle targetgroups operations
  litectl api cloud templates Handle templates operations
  litectl api cloud um   Handle um operations
```

The commands themselves don't do anything. But feel free to contribute or fork this repo if you think it's useful.

This CLI is likely not going to be worked upon further
