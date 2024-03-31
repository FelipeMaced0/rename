# Introduction
This cli utility is meant to be used when in need of standardazing folder containing multiple files with diferent name conventions. When yout find yourself renaming files by hand and removing accents, choosing a separetor and a case for the names, it becomes very tiring and time consuming and error prone as time goes by. That being said, it is a good call to have an automated yet flexible way to rename your files.
## Prerequisites

For it to work properly is necessary to have go installed on your machine. Head to go website to download latest [go](https://go.dev/dl) version. If you pretends to use metadata to add to your filenames it is necessary to install `exiftool` and have it on the machines `PATH`. You can run `apt install exiftool`
or use your package manager of choice.

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
      --attach-title               Attach Title(Generally in text file)
      --author                     Attach Author(Generally in text file)
      --auto-separetor             Guess separetor and replace with _
  -c, --copy-path string           Copy files to destination folder with new names(recommended)
      --create-date                Attach create date
      --dimension                  Attach image size to file name WidthxHeight
      --duration                   Attach video duration
      --file-size                  Attach file size to file name
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

# Examples

`rename bulk --path="your/path/folder/to/rename" --inplace -u`
Here `--path` is the path of the folder containing the files to be renamed, `--inplace` files will be renamed in the current folder(see`--copy-path`) and `-u` is
the flag responsible for upper case the names


`rename bulk --path="your/path/folder/to/rename" --copy-path="your/new/folder/location" -l`
Here `--path` is the path of the folder containing the files to be renamed, `--copy-path` is the path of destination of the files(see`--inplace`) and `-l` is
the flag responsible for lower case the names

`rename bulk --path="your/path/folder/to/rename" --copy-path="your/new/folder/location" -l --auto-separetor`
Here `--path` is the path of the folder containing the files to be renamed, `--copy-path` is the path of destination of the files(see`--inplace`) and `-l` is
the flag responsible for lower case the names, `--auto-separetor` will guess the separetor used on the file name "your_name_is" or "your-name-is" or "your name is" and replace by one occurence of "_".