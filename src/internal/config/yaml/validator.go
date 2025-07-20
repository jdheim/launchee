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
	"os/exec"
	"strings"
	"unicode/utf8"

	"github.com/jdheim/launchee/internal/config/frontend"
	"github.com/pkg/errors"
)

const maxIconSize = 1 << 20 // 1 MB

func validate(config *config) error {
	if err := validateTitle(config); err != nil {
		return err
	}
	if err := validateShortcuts(config); err != nil {
		return err
	}
	return nil
}

func validateTitle(config *config) error {
	titleLength := utf8.RuneCountInString(config.Title)
	if titleLength != 0 && (titleLength < 3 || titleLength > 30) {
		return errors.Errorf("Title \"%s\" must be between 3 and 30 characters long (got %d)", config.Title, titleLength)
	}
	return nil
}

func validateShortcuts(config *config) error {
	for _, shortcut := range config.Shortcuts {
		if err := validateShortcut(shortcut); err != nil {
			return err
		}
	}
	return nil
}

func validateShortcut(shortcut *shortcut) error {
	if err := validateShortcutName(shortcut); err != nil {
		return err
	}
	if err := validateShortcutPatch(shortcut); err != nil {
		return err
	}
	if err := validateShortcutIcon(shortcut); err != nil {
		return err
	}
	if err := validateShortcutCommandAndUrl(shortcut); err != nil {
		return err
	}
	if err := validateShortcutCommand(shortcut); err != nil {
		return err
	}
	if err := validateShortcutCommandArgs(shortcut); err != nil {
		return err
	}
	if err := validateShortcutUrl(shortcut); err != nil {
		return err
	}
	return nil
}

func validateShortcutName(shortcut *shortcut) error {
	nameLength := utf8.RuneCountInString(shortcut.Name)
	if nameLength < 3 || nameLength > 30 {
		return errors.Errorf("Name of \"%s\" Shortcut must be between 3 and 30 characters long (got %d)", shortcut.Name, nameLength)
	}
	return nil
}

func validateShortcutPatch(shortcut *shortcut) error {
	if shortcut.Patch != "" && shortcut.Patch != patchReplace && shortcut.Patch != patchMerge && shortcut.Patch != patchDelete {
		return errors.Errorf("Patch of \"%s\" Shortcut must be either \"%s\", \"%s\" or \"%s\" (got \"%s\")",
			shortcut.Name, patchReplace, patchMerge, patchDelete, shortcut.Patch)
	}
	return nil
}

func validateShortcutIcon(shortcut *shortcut) error {
	if !shortcut.isPatchMode() && shortcut.Icon == "" {
		return errors.Errorf("Icon of \"%s\" Shortcut must be set", shortcut.Name)
	}
	if shortcut.Icon != "" {
		if !fileExists(shortcut.Icon) {
			return errors.Errorf("Icon of \"%s\" Shortcut does not exist under: \"%s\"", shortcut.Name, shortcut.Icon)
		}
		if !frontend.IsValidIcon(shortcut.Icon) {
			return errors.Errorf("Icon of \"%s\" Shortcut is not a valid Icon: \"%s\". Supported extensions: %s",
				shortcut.Name, shortcut.Icon, frontend.SupportedExtensions())
		}
		if isFileLargerThan1MB(shortcut.Icon) {
			return errors.Errorf("Icon of \"%s\" Shortcut is larger than 1 MB", shortcut.Name)
		}
	}
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func isFileLargerThan1MB(path string) bool {
	info, _ := os.Stat(path)
	return info.Size() > maxIconSize
}

func validateShortcutCommandAndUrl(shortcut *shortcut) error {
	if !shortcut.isPatchMode() && shortcut.Command == "" && shortcut.Url == "" {
		return errors.Errorf("Either Command or URL of \"%s\" Shortcut must be set", shortcut.Name)
	}
	if shortcut.Command != "" && shortcut.Url != "" {
		return errors.Errorf("\"%s\" Shortcut cannot have both a Command and a URL set - choose one (got Command: \"%s\" and URL: \"%s\")",
			shortcut.Name, shortcut.Command, shortcut.Url)
	}
	return nil
}

func validateShortcutCommand(shortcut *shortcut) error {
	if shortcut.Command != "" && !isExec(shortcut.Command) {
		return errors.Errorf("Command of \"%s\" Shortcut is not a valid Command (got \"%s\")", shortcut.Name, shortcut.Command)
	}
	return nil
}

func isExec(command string) bool {
	_, pathErr := exec.LookPath(command)
	return pathErr == nil
}

func validateShortcutCommandArgs(shortcut *shortcut) error {
	if shortcut.CommandArgs != "" && shortcut.Command == "" {
		return errors.Errorf("Command Args of \"%s\" Shortcut not allowed without a Command (got \"%s\")", shortcut.Name, shortcut.CommandArgs)
	}
	return nil
}

func validateShortcutUrl(shortcut *shortcut) error {
	if shortcut.Url != "" {
		shortcutUrl := strings.ToLower(shortcut.Url)
		if !strings.HasPrefix(shortcutUrl, "https://") && !strings.HasPrefix(shortcutUrl, "http://") {
			return errors.Errorf("URL of \"%s\" Shortcut must start with \"http://\" or \"https://\" (got \"%s\")", shortcut.Name, shortcut.Url)
		}
	}
	return nil
}
