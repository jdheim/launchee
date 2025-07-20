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
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jdheim/launchee/internal/config/frontend"
)

func TestNewConfigWithoutShortcuts(t *testing.T) {
	in := "Test Title"
	got := newConfigWithoutShortcuts(in)
	want := &config{Title: in}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("newConfigWithoutShortcuts(%q) = diff -want +got\n%s", in, diff)
	}
}

func TestTrim(t *testing.T) {
	testCases := map[string]struct {
		in   *config
		want *config
	}{
		"nil":          {nil, nil},
		"empty":        {&config{}, &config{}},
		"title only":   {&config{Title: "  T  "}, &config{Title: "T"}},
		"nil shortcut": {&config{Shortcuts: []*shortcut{nil}}, &config{Shortcuts: []*shortcut{nil}}},
		"full": {
			&config{
				Title: "  Test Title  ",
				Shortcuts: []*shortcut{{
					Name:        "  Name  ",
					Icon:        "  Icon  ",
					Command:     "  Command  ",
					CommandArgs: "  Args  ",
					Url:         "  Url  ",
					Patch:       "  Replace  ",
				}},
			},
			&config{
				Title: "Test Title",
				Shortcuts: []*shortcut{{
					Name:        "Name",
					Icon:        "Icon",
					Command:     "Command",
					CommandArgs: "Args",
					Url:         "Url",
					Patch:       "Replace",
				}},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testCase.in.trim()
			got := testCase.in
			if diff := cmp.Diff(testCase.want, got); diff != "" {
				t.Errorf("trim() = diff -want +got\n%s", diff)
			}
		})
	}
}

func TestToFrontendConfig(t *testing.T) {
	defaultUI := frontend.NewUI(0)
	testTitle := "Test Title"
	defaultUIOverrideTitleNoShortcuts := frontend.NewUI(0)
	defaultUIOverrideTitleNoShortcuts.Nav.Title = testTitle
	defaultUIOverrideTitle := frontend.NewUI(1)
	defaultUIOverrideTitle.Nav.Title = testTitle

	testCases := map[string]struct {
		in   *config
		want *frontend.Config
	}{
		"nil": {nil, &frontend.Config{
			UI:    defaultUI,
			Valid: true,
		}},
		"title only": {
			&config{
				Title: testTitle,
			},
			&frontend.Config{
				UI:    defaultUIOverrideTitleNoShortcuts,
				Valid: true,
			},
		},
		"full no icon": {
			&config{
				Title: testTitle,
				Shortcuts: []*shortcut{{
					Name:        "Name",
					Icon:        "",
					Command:     "Command",
					CommandArgs: "Arg1 Arg2",
					Url:         "Url",
					Patch:       "Replace",
				}},
			},
			&frontend.Config{
				UI: defaultUIOverrideTitle,
				Shortcuts: []*frontend.Shortcut{{
					Id:   0,
					Name: "Name",
					Icon: &frontend.Icon{
						Path:   "",
						Bytes:  nil,
						Base64: "data:image/png;base64,",
					},
					Command:     "Command",
					CommandArgs: []string{"Arg1", "Arg2"},
					Url:         "Url",
				}},
				Valid: true,
			},
		},
		"full with icon": {
			&config{
				Title: testTitle,
				Shortcuts: []*shortcut{{
					Name:        "Name",
					Icon:        "../../../build/appicon.png",
					Command:     "Command",
					CommandArgs: "Arg1 Arg2",
					Url:         "Url",
					Patch:       "Replace",
				}},
			},
			&frontend.Config{
				UI: defaultUIOverrideTitle,
				Shortcuts: []*frontend.Shortcut{{
					Id:          0,
					Name:        "Name",
					Icon:        frontend.NewIcon("../../../build/appicon.png"),
					Command:     "Command",
					CommandArgs: []string{"Arg1", "Arg2"},
					Url:         "Url",
				}},
				Valid: true,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := testCase.in.toFrontendConfig()
			if diff := cmp.Diff(testCase.want, got); diff != "" {
				t.Errorf("toFrontendConfig() = diff -want +got\n%s", diff)
			}
		})
	}
}

func TestParseCommandArgs(t *testing.T) {
	testCases := map[string]struct {
		in   *shortcut
		want []string
	}{
		"nil":                {nil, nil},
		"empty":              {&shortcut{CommandArgs: ""}, nil},
		"2 args":             {&shortcut{CommandArgs: "foo bar"}, []string{"foo", "bar"}},
		"2 args with spaces": {&shortcut{CommandArgs: "'foo bar' 'baz'"}, []string{"foo bar", "baz"}},
		"4 args":             {&shortcut{CommandArgs: "-d /tmp/dummy -f foo.txt"}, []string{"-d", "/tmp/dummy", "-f", "foo.txt"}},
		"4 args with spaces": {&shortcut{CommandArgs: "-s \"space test\" -f foo.txt"}, []string{"-s", "space test", "-f", "foo.txt"}},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := testCase.in.parseCommandArgs()
			if diff := cmp.Diff(testCase.want, got); diff != "" {
				t.Errorf("parseCommandArgs() = diff -want +got\n%s", diff)
			}
		})
	}
}

func TestIsPatchMode(t *testing.T) {
	testCases := map[string]struct {
		in   *shortcut
		want bool
	}{
		"nil":     {nil, false},
		"empty":   {&shortcut{Patch: ""}, false},
		"delete":  {&shortcut{Patch: "delete"}, true},
		"merge":   {&shortcut{Patch: "merge"}, true},
		"replace": {&shortcut{Patch: "replace"}, false},
		"other":   {&shortcut{Patch: "other"}, false},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := testCase.in.isPatchMode()
			if got != testCase.want {
				t.Errorf("isPatchMode() = %t, want %t", got, testCase.want)
			}
		})
	}
}
