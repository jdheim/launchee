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
	"strings"

	"github.com/google/shlex"
	"github.com/jdheim/launchee/internal/config/frontend"
)

type config struct {
	Title     string
	Shortcuts []*shortcut
}

type shortcut struct {
	Name        string
	Icon        string
	Command     string
	CommandArgs string `yaml:"commandArgs"`
	Url         string
	Patch       string `yaml:"$patch"`
}

// Creates a new config without shortcuts.
func newConfigWithoutShortcuts(title string) *config {
	return &config{
		Title: title,
	}
}

// Trims all strings in the config.
func (yc *config) trim() {
	if yc == nil {
		return
	}
	yc.Title = strings.TrimSpace(yc.Title)
	for _, shortcut := range yc.Shortcuts {
		if shortcut == nil {
			continue
		}
		shortcut.Name = strings.TrimSpace(shortcut.Name)
		shortcut.Icon = strings.TrimSpace(shortcut.Icon)
		shortcut.Command = strings.TrimSpace(shortcut.Command)
		shortcut.CommandArgs = strings.TrimSpace(shortcut.CommandArgs)
		shortcut.Url = strings.TrimSpace(shortcut.Url)
		shortcut.Patch = strings.TrimSpace(shortcut.Patch)
	}
}

// Converts the config to a frontend.Config.
func (yc *config) toFrontendConfig() *frontend.Config {
	if yc == nil {
		return frontend.NewConfig(0)
	}
	frontendConfig := frontend.NewConfig(len(yc.Shortcuts))
	yc.overrideFrontendConfig(frontendConfig)
	return frontendConfig
}

// Overrides the frontend.Config with the config values.
func (yc *config) overrideFrontendConfig(config *frontend.Config) {
	if yc.Title != "" {
		config.UI.Nav.Title = yc.Title
	}
	if frontendShortcuts := yc.toFrontendShortcuts(); frontendShortcuts != nil {
		config.Shortcuts = frontendShortcuts
	}
}

// Converts the shortcut(s) to a frontend.Shortcut(s).
func (yc *config) toFrontendShortcuts() []*frontend.Shortcut {
	shortcutCount := len(yc.Shortcuts)
	if shortcutCount == 0 {
		return nil
	}
	frontendShortcuts := make([]*frontend.Shortcut, shortcutCount)
	for i := range yc.Shortcuts {
		frontendShortcuts[i] = yc.toFrontendShortcut(i)
	}
	return frontendShortcuts
}

// Converts the shortcut to a frontend.Shortcut.
func (yc *config) toFrontendShortcut(i int) *frontend.Shortcut {
	return &frontend.Shortcut{
		Id:          i,
		Name:        yc.Shortcuts[i].Name,
		Icon:        frontend.NewIcon(yc.Shortcuts[i].Icon),
		Command:     yc.Shortcuts[i].Command,
		CommandArgs: yc.Shortcuts[i].parseCommandArgs(),
		Url:         yc.Shortcuts[i].Url,
	}
}

// Parses the command arguments.
func (s *shortcut) parseCommandArgs() []string {
	if s == nil || s.CommandArgs == "" {
		return nil
	}
	commandArgParts, _ := shlex.Split(s.CommandArgs)
	return commandArgParts
}

// Checks if the shortcut is in patch mode.
func (s *shortcut) isPatchMode() bool {
	if s == nil || s.Patch == "" {
		return false
	}
	return s.Patch == patchDelete || s.Patch == patchMerge
}
