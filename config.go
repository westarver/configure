package config

import (
	"errors"
	"fmt"
	"os"

	apath "github.com/rhysd/abspath"
	"gopkg.in/ini.v1"
)

//=================== getIniLoadOptions ===================

func getIniLoadOptions() ini.LoadOptions {
	return ini.LoadOptions{
		Loose:                       true,
		Insensitive:                 false,
		InsensitiveSections:         false,
		InsensitiveKeys:             false,
		IgnoreContinuation:          false,
		IgnoreInlineComment:         false,
		SkipUnrecognizableLines:     true,
		ShortCircuit:                false,
		AllowBooleanKeys:            true,
		AllowShadows:                false,
		AllowNestedValues:           false,
		AllowPythonMultilineValues:  true,
		SpaceBeforeInlineComment:    false,
		UnescapeValueDoubleQuotes:   false,
		UnescapeValueCommentSymbols: false,
		UnparseableSections:         []string{},
		KeyValueDelimiters:          "=",
		KeyValueDelimiterOnWrite:    "",
		ChildSectionDelimiter:       ".",
		PreserveSurroundedQuote:     false,
		DebugFunc:                   func(message string) {},
		ReaderBufferSize:            0,
		AllowNonUniqueSections:      false,
		AllowDuplicateShadowValues:  false,
	}
}

//===============================| ReadIni |===============================
// reads a key-value pair text file and populates an ini.File structure.
// see documentation for the ini package at http://gopkg.in/ini.v1
// function will fail if
// 1. file cant be opened or created
// 2. file is created but the default string is empty
func ReadIni(filename string) (*ini.File, error) {

	if len(filename) == 0 {
		return nil, errors.New("ReadIni: File name was empty")
	}

	cfg, err := apath.ExpandFrom(filename)
	if err != nil { //can't get a good rooted path to cwd
		return nil, errors.New("ReadIni: Configuration file " + filename + " is not valid.\n Make sure the file name passed to ReadIni is relative to the current working directory\n")
	}
	iniOptions := getIniLoadOptions()
	return ini.LoadSources(iniOptions, cfg.String())
}

//=================== SetEnvarsFromIni ===================

func SetEnvarsFromIni(inif *ini.File, env string) error {
	sec, err := inif.GetSection(env)
	if err == nil {
		for _, k := range sec.Keys() {
			os.Setenv(k.Name(), k.Value())
			fmt.Println("Set environment variable " + k.Name() + " to " + k.Value() + "\n")
		}
		return nil
	}
	return err
}
