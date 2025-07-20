/*
 * © 2025-2025 JDHeim.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package yaml

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidate(t *testing.T) {
	testCases := map[string]struct {
		in   func() *config
		want bool
	}{
		"invalid title": {func() *config {
			validConfig := newValidConfig()
			validConfig.Title = "Te"
			return validConfig
		}, false},
		"invalid shortcut name": {func() *config {
			validConfig := newValidConfig()
			validConfig.Shortcuts[0].Name = "Te"
			return validConfig
		}, false},
		"invalid shortcut patch": {func() *config {
			validConfig := newValidConfig()
			validConfig.Shortcuts[1].Patch = "other"
			return validConfig
		}, false},
		"invalid shortcut icon": {func() *config {
			validConfig := newValidConfig()
			validConfig.Shortcuts[2].Icon = "not-exists.png"
			return validConfig
		}, false},
		"invalid shortcut command and url": {func() *config {
			validConfig := newValidConfig()
			validConfig.Shortcuts[0].Command = "echo"
			validConfig.Shortcuts[0].Url = "https://example.com"
			return validConfig
		}, false},
		"invalid shortcut command": {func() *config {
			validConfig := newValidConfig()
			validConfig.Shortcuts[1].Command = "invalid"
			return validConfig
		}, false},
		"invalid shortcut command args": {func() *config {
			validConfig := newValidConfig()
			validConfig.Shortcuts[2].Command = ""
			validConfig.Shortcuts[2].CommandArgs = "-d /tmp/dummy -f foo.txt"
			validConfig.Shortcuts[2].Patch = "merge"
			return validConfig
		}, false},
		"invalid shortcut url": {func() *config {
			validConfig := newValidConfig()
			validConfig.Shortcuts[0].Command = ""
			validConfig.Shortcuts[0].Url = "www.example.com"
			return validConfig
		}, false},
		"nil": {func() *config {
			return nil
		}, true},
		"valid": {func() *config {
			return newValidConfig()
		}, true},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validate(testCase.in())
			if got := err == nil; got != testCase.want {
				t.Errorf("validate() = %t, want %t", got, testCase.want)
			}
		})
	}
}

func newValidConfig() *config {
	return &config{
		Title: "Test Title",
		Shortcuts: []*shortcut{
			{Name: "Test1", Icon: "../../../build/appicon.png", Command: "echo"},
			{Name: "Test2", Icon: "../../../build/appicon.png", Command: "echo"},
			{Name: "Test3", Icon: "../../../build/appicon.png", Command: "echo"},
		},
	}
}

func TestValidateTitle(t *testing.T) {
	testCases := map[string]struct {
		in   *config
		want bool
	}{
		"empty":          {&config{Title: ""}, true}, // The default will be "Launchee"
		"1 char":         {&config{Title: "T"}, false},
		"2 chars":        {&config{Title: "Te"}, false},
		"3 chars":        {&config{Title: "Tes"}, true},
		"15 chars":       {&config{Title: "Test Title Test"}, true},
		"30 chars":       {&config{Title: "Test Title TestTest Title Test"}, true},
		"30 chars utf-8": {&config{Title: "😎ąćęłńóśźżαβγδεζηθικλμνξοπρστ😎"}, true},
		"31 chars":       {&config{Title: "Test Title Test Test Title Test"}, false},
		"31 chars utf-8": {&config{Title: "😎ąćęłńóśźżαβγδεζηθικλμνξοπρστυ😎"}, false},
		"45 chars":       {&config{Title: "Test Title TestTest Title TestTest Title Test"}, false},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validateTitle(testCase.in)
			if got := err == nil; got != testCase.want {
				t.Errorf("validateTitle(%q) = %t, want %t", testCase.in.Title, got, testCase.want)
			}
		})
	}
}

func TestValidateShortcutName(t *testing.T) {
	testCases := map[string]struct {
		in   *shortcut
		want bool
	}{
		"empty":          {&shortcut{Name: ""}, false},
		"1 char":         {&shortcut{Name: "T"}, false},
		"2 chars":        {&shortcut{Name: "Te"}, false},
		"3 chars":        {&shortcut{Name: "Tes"}, true},
		"15 chars":       {&shortcut{Name: "Test Title Test"}, true},
		"30 chars":       {&shortcut{Name: "Test Title TestTest Title Test"}, true},
		"30 chars utf-8": {&shortcut{Name: "😎ąćęłńóśźżαβγδεζηθικλμνξοπρστ😎"}, true},
		"31 chars":       {&shortcut{Name: "Test Title Test Test Title Test"}, false},
		"31 chars utf-8": {&shortcut{Name: "😎ąćęłńóśźżαβγδεζηθικλμνξοπρστυ😎"}, false},
		"45 chars":       {&shortcut{Name: "Test Title TestTest Title TestTest Title Test"}, false},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validateShortcutName(testCase.in)
			if got := err == nil; got != testCase.want {
				t.Errorf("validateShortcutName(%q) = %t, want %t", testCase.in.Name, got, testCase.want)
			}
		})
	}
}

func TestValidateShortcutPatch(t *testing.T) {
	testCases := map[string]struct {
		in   *shortcut
		want bool
	}{
		"empty":   {&shortcut{Patch: ""}, true}, // The default will be "replace"
		"replace": {&shortcut{Patch: "replace"}, true},
		"merge":   {&shortcut{Patch: "merge"}, true},
		"delete":  {&shortcut{Patch: "delete"}, true},
		"other":   {&shortcut{Patch: "other"}, false},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validateShortcutPatch(testCase.in)
			if got := err == nil; got != testCase.want {
				t.Errorf("validateShortcutPatch(%q) = %t, want %t", testCase.in.Patch, got, testCase.want)
			}
		})
	}
}

func TestValidateShortcutIcon(t *testing.T) {
	dirWithSpaces := filepath.Join(t.TempDir(), "dir with spaces")
	icon1MBPath := filepath.Join(dirWithSpaces, "1MB.png")
	iconLargerThan1MBPath := filepath.Join(t.TempDir(), "largerThan1MB.png")

	testCases := map[string]struct {
		in   *shortcut
		want bool
	}{
		"empty":                    {&shortcut{Icon: ""}, false},
		"empty with patch empty":   {&shortcut{Icon: "", Patch: ""}, false},
		"empty with patch replace": {&shortcut{Icon: "", Patch: "replace"}, false},
		"empty with patch merge":   {&shortcut{Icon: "", Patch: "merge"}, true},
		"empty with patch delete":  {&shortcut{Icon: "", Patch: "delete"}, true},
		"empty with patch other":   {&shortcut{Icon: "", Patch: "other"}, false},
		"not exists":               {&shortcut{Icon: "not-exists.png"}, false},
		"not valid":                {&shortcut{Icon: "validator.go"}, false},
		"1MB":                      {&shortcut{Icon: icon1MBPath}, true},
		"larger than 1MB":          {&shortcut{Icon: iconLargerThan1MBPath}, false},
		"valid":                    {&shortcut{Icon: "../../../build/appicon.png"}, true},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			if "1MB" == name {
				if err := os.Mkdir(dirWithSpaces, 0755); err != nil {
					t.Errorf("Failed to create %s: %v", dirWithSpaces, err)
				}
				if err := os.WriteFile(icon1MBPath, make([]byte, 1<<20), 0644); err != nil {
					t.Errorf("Failed to write %s: %v", icon1MBPath, err)
				}
			} else if "larger than 1MB" == name {
				if err := os.WriteFile(iconLargerThan1MBPath, make([]byte, (1<<20)+1), 0644); err != nil {
					t.Errorf("Failed to write %s: %v", iconLargerThan1MBPath, err)
				}
			}
			err := validateShortcutIcon(testCase.in)
			if got := err == nil; got != testCase.want {
				t.Errorf("validateShortcutIcon(%q) = %t, want %t", testCase.in.Icon, got, testCase.want)
			}
		})
	}
}

func TestValidateShortcutCommandAndUrl(t *testing.T) {
	testCases := map[string]struct {
		in   *shortcut
		want bool
	}{
		"both empty with patch empty":   {&shortcut{Command: "", Url: "", Patch: ""}, false},
		"both empty with patch replace": {&shortcut{Command: "", Url: "", Patch: "replace"}, false},
		"both empty with patch merge":   {&shortcut{Command: "", Url: "", Patch: "merge"}, true},
		"both empty with patch delete":  {&shortcut{Command: "", Url: "", Patch: "delete"}, true},
		"both empty with patch other":   {&shortcut{Command: "", Url: "", Patch: "other"}, false},
		"both empty":                    {&shortcut{Command: "", Url: ""}, false},
		"command not empty":             {&shortcut{Command: "echo", Url: ""}, true},
		"url not empty":                 {&shortcut{Command: "", Url: "https://example.com"}, true},
		"both not empty":                {&shortcut{Command: "echo", Url: "https://example.com"}, false},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validateShortcutCommandAndUrl(testCase.in)
			if got := err == nil; got != testCase.want {
				t.Errorf("validateShortcutCommandAndUrl(%q, %q) = %t, want %t", testCase.in.Command, testCase.in.Url, got, testCase.want)
			}
		})
	}
}

func TestValidateShortcutCommand(t *testing.T) {
	dirWithSpaces := filepath.Join(t.TempDir(), "dir with spaces")
	execPath := filepath.Join(dirWithSpaces, "echo")

	testCases := map[string]struct {
		in   *shortcut
		want bool
	}{
		"invalid":                {&shortcut{Command: "invalid"}, false},
		"empty":                  {&shortcut{Command: ""}, true}, // We can have other actions
		"valid command":          {&shortcut{Command: "echo"}, true},
		"valid path":             {&shortcut{Command: "/usr/bin/echo"}, true},
		"valid path with spaces": {&shortcut{Command: execPath}, true},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			if "valid path with spaces" == name {
				if err := os.Mkdir(dirWithSpaces, 0755); err != nil {
					t.Errorf("Failed to create %s: %v", dirWithSpaces, err)
				}
				if err := os.WriteFile(execPath, []byte("#!/bin/sh\necho hi\n"), 0755); err != nil {
					t.Errorf("Failed to write %s: %v", execPath, err)
				}
			}
			err := validateShortcutCommand(testCase.in)
			if got := err == nil; got != testCase.want {
				t.Errorf("validateShortcutCommand(%q) = %t, want %t", testCase.in.Command, got, testCase.want)
			}
		})
	}
}

func TestValidateShortcutCommandArgs(t *testing.T) {
	testCases := map[string]struct {
		in   *shortcut
		want bool
	}{
		"empty without command": {&shortcut{Command: "", CommandArgs: ""}, true},
		"empty with command":    {&shortcut{Command: "command", CommandArgs: ""}, true},
		"invalid":               {&shortcut{Command: "", CommandArgs: "-d /tmp/dummy -f foo.txt"}, false},
		"valid":                 {&shortcut{Command: "command", CommandArgs: "-d /tmp/dummy -f foo.txt"}, true},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validateShortcutCommandArgs(testCase.in)
			if got := err == nil; got != testCase.want {
				t.Errorf("validateShortcutCommandArgs(%q) = %t, want %t", testCase.in.CommandArgs, got, testCase.want)
			}
		})
	}
}

func TestValidateShortcutUrl(t *testing.T) {
	testCases := map[string]struct {
		in   *shortcut
		want bool
	}{
		"empty":           {&shortcut{Url: ""}, true}, // We can have other actions
		"invalid":         {&shortcut{Url: "invalid"}, false},
		"invalid www":     {&shortcut{Url: "www.example.com"}, false},
		"invalid http":    {&shortcut{Url: "http:/example.com"}, false},
		"valid http":      {&shortcut{Url: "http://example.com"}, true},
		"valid http www":  {&shortcut{Url: "http://www.example.com"}, true},
		"invalid https":   {&shortcut{Url: "https:/example.com"}, false},
		"valid https":     {&shortcut{Url: "https://example.com"}, true},
		"valid https www": {&shortcut{Url: "https://www.example.com"}, true},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			err := validateShortcutUrl(testCase.in)
			if got := err == nil; got != testCase.want {
				t.Errorf("validateShortcutUrl(%q) = %t, want %t", testCase.in.Url, got, testCase.want)
			}
		})
	}
}
