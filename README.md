# Introduction
This cli utility is meant to be used when in need of standardazing folder containing multiple files with diferent name conventions. When yout find yourself renaming files by hand and removing accents, choosing a separetor and a case for the names, it becomes very tiring and time consuming and error prone as time goes by. That being said, it is a good call to have an automated yet flexible way to rename your files.
## Prerequisites

For it to work properly is necessary to have go installed on your machine. Head to go website to download latest [go](https://go.dev/dl) version.

## Build

No magic here, just run `go build` inside the project's folder

## Usage

```
bulk rename files

Usage:
  rename bulk [flags]

Flags:
  -a, --all                        Rename all files
      --ascii                      Rename to ascii only characters
      --auto-separetor             Guess separetor and replace with _
  -c, --copy-path string           Copy files to destination folder with new names(recommended)
  -h, --help                       help for bulk
      --inplace                    Rename file inplace(possible loss of information)
      --limit int32                Limit of renamed files
  -l, --lower                      Rename to lowercase case
  -p, --path string                Path of your folder containing the files that sould be renamed
      --prefix string              Add at the begin of every file
  -r, --replace-separetor string   Separetor to put between words on file name (default "_")
      --suffix string              Add at the end of every file
  -t, --title                      Rename to title case
  -u, --upper                      Rename to uppercase case
```