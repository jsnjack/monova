monova
=====

[![CircleCI](https://circleci.com/gh/jsnjack/monova.svg?style=svg)](https://circleci.com/gh/jsnjack/monova)

### What is it?
monova automatically calculates version of the application based on the commit messages

### How to use?
monova uses format `Major.Minor.Patch` (f.e. `1.11.3`) to calculate the version. To automatically increase the version of the application add one of the reserved words (by default `:major:`, `:minor:` or `:patch:`) to the commit message

#### Overview
```
Usage of ./monova:
  -checkpoint
    	Create checkpoint [version]
  -debug
    	Enable extra logging
  -history
    	Print version history
  -info
    	Print old and new version
  -reset
    	Recalculate version
  -version
    	Print version information
```

#### Change reserved words
To change reserved words, create file `.monova.config` in the root folder of your project with the following content:
```json
{
    "MajorKeys": [
        ":major:", ":M:"
    ],
    "MinorKeys": [
        ":minor:", ":m:"
    ],
    "PatchKeys": [
        ":patch:", ":p:"
    ]
}
```

#### Get commit ID from the version and vice versa
monova keeps track of versions. Use the following command to print history:
```bash
monova -history
```
Use `grep` or [kazy](https://github.com/jsnjack/kazy-go) to filter the results
```bash
$ ./monova -history | grep 0.6.0
936fa02f8a4d975df4a40be329a09aaf5d2cdbea :m: Update version handles checkpoint              0.6.0
```
```bash
$ ./monova -history | kazy -i f1b7018e5c9fcb7be551f142b1582948cbceec26
f1b7018e5c9fcb7be551f142b1582948cbceec26 :m: Add config                                     0.4.0
```

#### Using checkpoints
You can force any version by creating a checkpoint. Checkpoint is an empty commit with specially formated message:
```bash
$ ./monova -checkpoint 1.0.0
$ git log
01de65a Version:1.0.0 generated by monova
010a56a :m: Add progress bar
73b6a72 :m: Add history flag
270f422 :p: Update help command
```

### How to install
