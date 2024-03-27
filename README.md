# Introduction
This cli utility is meant to be used when in need of standardazing folder containing multiple files with diferent name conventions. When yout find yourself renaming files by hand and removing accents, choosing a separetor and a case for the names, it becomes very tiring and time consuming and error prone as time goes by. That being said, it is a good call to have an automated yet flexible way to rename your files.
## Prerequisites

For it to work properly is necessary to have go installed on your machine. Head to go website to download latest [go](https://go.dev/dl) version.

## Build

No magic her, just run `go build` on inside folder

## Usage

```
bulk rename files

Usage:
  rename bulk [flags]

Flags:
  -a, --all                        Process all files
      --ascii                      Rename to ascii only characters
      --copy-path string           Copy files to destination folder with new names
  -c, --current-separetor string   Separetor to search between words on file name and replace with [replace-separetor]
  -h, --help                       help for bulk
      --inplace                    Rename file inplace(possible loss of information)
      --limit int32                Limit to processed files (default 5)
  -l, --lower                      Tells whether or not files names should be on lowercase letters
  -p, --path string                Path of your folder containing the files that sould be renamed
      --prefix string              Prefix to add at the begin of every file
  -r, --replace-separetor string   Separetor to put between words on file name (default "_")
      --suffix string              Suffix to add at the end of every file
  -t, --title                      Tells whether or not files names should be title case
  -u, --upper                      Tells whether or not files names should be on uppercase letters
```