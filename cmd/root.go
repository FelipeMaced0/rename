/*
Copyright © 2024 Felipe Macedo <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rename [options] [flags]",
	Short: "Rename your files to a certain pattern and be a happy programmer",
	Long:  `Renamer is meant to make your life easier when renaming files in bulk to conform to a certain pattern`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var toLower bool
	var toUpper bool
	var toTitle bool
	var ascii bool
	var all bool

	var inplace bool

	var prefix string
	var suffix string
	var path string
	var copyPath string
	var replaceSeparetor string
	var currentSeparetor string

	var limit *int32

	var rename = &cobra.Command{
		Use:   "bulk",
		Short: "bulk rename files",
		Long:  ``,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			currentSeparetor = ""
			var count int32 = 0

			err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				count++

				if err != nil {
					return err
				}
				// Check if it's a regular file
				if !info.Mode().IsRegular() {
					return nil
				}

				if !all && (count == (*limit)) {
					return nil
				}

				// Extract the directory and file name
				dir, file := filepath.Split(path)
				ext := filepath.Ext(path)

				currentSeparetor = guessSeparetor(file)

				var oldShortname string = file

				// remove extension from file name
				file = strings.Replace(file, ext, "", 1)

				if prefix != "" {
					file = prefix + replaceSeparetor + file
				}

				if suffix != "" {
					file = file + replaceSeparetor + suffix
				}

				if toLower {
					file = strings.ToLower(file)
				}

				if toUpper {
					file = strings.ToUpper(file)
				}

				if toTitle {
					temp := strings.ReplaceAll(file, currentSeparetor, " ")
					temp = cases.Title(language.English).String(temp)
					file = strings.ReplaceAll(temp, " ", currentSeparetor)
				}

				if ascii {
					file = removeAccents(file)
				}

				fmt.Print(replaceSeparetor)
				file = strings.ReplaceAll(file, currentSeparetor, replaceSeparetor)

				/*TODO
				check the final file name length before attempt to rename
				*/

				newFileName := filepath.Join(dir, file) + ext

				// Should pass a flag --inline to rename files
				// the default behaviour should be copying to another folder
				// Rename the file
				if inplace {
					if err := os.Rename(path, newFileName); err != nil {
						return err
					}
					fmt.Printf("from %s to %s\n", oldShortname, file)
				}

				if copyPath != "" {

					if err := os.MkdirAll(copyPath, os.ModePerm); err != nil {
						fmt.Println(err)
					}

					source, err := os.ReadFile(filepath.Join(dir, oldShortname))

					if err != nil {
						fmt.Println(err)
						return nil
					}

					newDir := filepath.Join(copyPath, file) + ext
					err = os.WriteFile(newDir, source, os.ModePerm)

					if err != nil {
						fmt.Println("Could not copy file: ", err)
						return nil
					}

					fmt.Printf("%s Copied to %s\n", oldShortname, newDir)

				}

				return nil
			})

			if err != nil {
				fmt.Println("Error:", err)
			}
		},
	}
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rename.PersistentFlags().BoolVarP(&toLower, "lower", "l", false, "Tells whether or not files names should be on lowercase letters")
	rename.PersistentFlags().BoolVarP(&toUpper, "upper", "u", false, "Tells whether or not files names should be on uppercase letters")
	rename.PersistentFlags().BoolVarP(&toTitle, "title", "t", false, "Tells whether or not files names should be title case")
	rename.PersistentFlags().BoolVar(&inplace, "inplace", false, "Rename file inplace(possible loss of information)")
	rename.PersistentFlags().BoolVar(&ascii, "ascii", false, "Rename to ascii only characters")
	rename.PersistentFlags().BoolVarP(&all, "all", "a", false, "Process all files")
	rename.PersistentFlags().StringVar(&copyPath, "copy-path", "", "Copy files to destination folder with new names")
	rename.PersistentFlags().StringVarP(&path, "path", "p", "", "Path of your folder containing the files that sould be renamed")
	rename.PersistentFlags().StringVarP(&replaceSeparetor, "replace-separetor", "r", "_", "Separetor to put between words on file name")
	rename.PersistentFlags().StringVarP(&currentSeparetor, "current-separetor", "c", "", "Separetor to search between words on file name and replace with [replace-separetor]")
	rename.PersistentFlags().StringVar(&prefix, "prefix", "", "Prefix to add at the begin of every file")
	rename.PersistentFlags().StringVar(&suffix, "suffix", "", "Suffix to add at the end of every file")
	limit = rename.PersistentFlags().Int32("limit", 5, "Limit to processed files")

	rename.MarkPersistentFlagRequired("path")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.AddCommand(rename)
}

func removeAccents(s string) string {

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)

	sanitized, _, _ := transform.String(t, s)

	return sanitized
}

func guessSeparetor(s string) string {
	var pattern string = "[^a-zA-Z-0-9]+"

	regex, err := regexp.Compile(pattern)

	if err != nil {
		fmt.Println("Error guessing")
		return "asss"
	}

	matches := regex.FindAllString(s, -1)

	mapMatches := make(map[string]int)

	for _, letter := range matches {
		mapMatches[letter]++
	}

	var mostOccurringChar string
	maxFrequency := 0

	// Find the most occurring character
	for char, frequency := range mapMatches {
		if frequency > maxFrequency {
			maxFrequency = frequency
			mostOccurringChar = char
		}
	}

	return mostOccurringChar
}
