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
	"github.com/jdheim/launchee/internal/lctx"
	"github.com/jdheim/launchee/internal/test/stub"
)

func TestSanitize(t *testing.T) {
	testCases := map[string]struct {
		input *config
		want  *config
	}{
		"nil":   {nil, nil},
		"empty": {&config{}, &config{}},
		"no duplicates": {&config{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Terminal",
		}, {
			Name:        "Text Editor",
			Icon:        "internal/test/stub/stub_config/icons/accessories-text-editor.png",
			Command:     "echo",
			CommandArgs: "Text Editor",
		}}}, &config{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Terminal",
		}, {
			Name:        "Text Editor",
			Icon:        "internal/test/stub/stub_config/icons/accessories-text-editor.png",
			Command:     "echo",
			CommandArgs: "Text Editor",
		}}}},
		"one duplicate": {&config{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Terminal",
		}, {
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Terminal",
		}, {
			Name:        "Text Editor",
			Icon:        "internal/test/stub/stub_config/icons/accessories-text-editor.png",
			Command:     "echo",
			CommandArgs: "Text Editor",
		}}}, &config{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Terminal",
		}, {
			Name:        "Text Editor",
			Icon:        "internal/test/stub/stub_config/icons/accessories-text-editor.png",
			Command:     "echo",
			CommandArgs: "Text Editor",
		}}}},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := testCase.input.sanitize()
			if diff := cmp.Diff(testCase.want, got); diff != "" {
				t.Errorf("sanitize() = diff -want +got\n%s", diff)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	testCases := map[string]struct {
		input []*config
		want  *config
	}{
		"nil":       {[]*config{nil, nil}, nil},
		"nil empty": {[]*config{nil, {}}, nil},
		"empty":     {[]*config{{}, {}}, &config{}},
		"empty nil": {[]*config{{}, nil}, &config{}},
		"replace": {[]*config{{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Terminal",
		}}}, {Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Replaced Terminal",
		}}}}, &config{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Replaced Terminal",
		}}}},
		"replace explicitly": {[]*config{{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Terminal",
		}}}, {Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/terminal-app.png",
			Command:     "echo",
			CommandArgs: "Replaced Terminal",
			Patch:       patchReplace,
		}}}}, &config{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/terminal-app.png",
			Command:     "echo",
			CommandArgs: "Replaced Terminal",
			Patch:       patchReplace,
		}}}},
		"merge icon": {[]*config{{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Terminal",
		}}}, {Shortcuts: []*shortcut{{
			Name:  "Terminal",
			Icon:  "internal/test/stub/stub_config/icons/terminal-app.png",
			Patch: patchMerge,
		}}}}, &config{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/terminal-app.png",
			Command:     "echo",
			CommandArgs: "Terminal",
			Patch:       patchMerge,
		}}}},
		"merge command": {[]*config{{Shortcuts: []*shortcut{{
			Name: "Terminal",
			Icon: "internal/test/stub/stub_config/icons/kitty-128.png",
			Url:  "https://example.com",
		}}}, {Shortcuts: []*shortcut{{
			Name:    "Terminal",
			Command: "echo",
			Patch:   patchMerge,
		}}}}, &config{Shortcuts: []*shortcut{{
			Name:    "Terminal",
			Icon:    "internal/test/stub/stub_config/icons/kitty-128.png",
			Command: "echo",
			Patch:   patchMerge,
		}}}},
		"merge command args with command": {[]*config{{Shortcuts: []*shortcut{{
			Name: "Terminal",
			Icon: "internal/test/stub/stub_config/icons/kitty-128.png",
			Url:  "https://example.com",
		}}}, {Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Command:     "echo",
			CommandArgs: "Terminal",
			Patch:       patchMerge,
		}}}}, &config{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Terminal",
			Patch:       patchMerge,
		}}}},
		"merge command args without command": {[]*config{{Shortcuts: []*shortcut{{
			Name: "Terminal",
			Icon: "internal/test/stub/stub_config/icons/kitty-128.png",
			Url:  "https://example.com",
		}}}, {Shortcuts: []*shortcut{{
			Name:        "Terminal",
			CommandArgs: "Terminal",
			Patch:       patchMerge,
		}}}}, &config{Shortcuts: []*shortcut{{
			Name:  "Terminal",
			Icon:  "internal/test/stub/stub_config/icons/kitty-128.png",
			Url:   "https://example.com",
			Patch: patchMerge,
		}}}},
		"merge url": {[]*config{{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Terminal",
		}}}, {Shortcuts: []*shortcut{{
			Name:  "Terminal",
			Url:   "https://example.com",
			Patch: patchMerge,
		}}}}, &config{Shortcuts: []*shortcut{{
			Name:  "Terminal",
			Icon:  "internal/test/stub/stub_config/icons/kitty-128.png",
			Url:   "https://example.com",
			Patch: patchMerge,
		}}}},
		"delete": {[]*config{{Shortcuts: []*shortcut{{
			Name:        "Terminal",
			Icon:        "internal/test/stub/stub_config/icons/kitty-128.png",
			Command:     "echo",
			CommandArgs: "Terminal",
		}}}, {Shortcuts: []*shortcut{{
			Name:  "Terminal",
			Patch: patchDelete,
		}}}}, &config{Shortcuts: []*shortcut{}}},
	}

	lctx.LoggerImpl = stub.LoggerStub{}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			got := testCase.input[0].merge(testCase.input[1])
			if diff := cmp.Diff(testCase.want, got); diff != "" {
				t.Errorf("merge() = diff -want +got\n%s", diff)
			}
		})
	}
}
