/*
 * © 2025-2025 JDHeim
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
	"runtime"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"

	"github.com/pkg/errors"

	"github.com/jdheim/launchee/internal/config/frontend"
)

const (
	configFileTestSystem = "/home/dev/projects/launchee/src/test/config/launchee-7.yml"
	configFileTestUser   = "/home/dev/projects/launchee/src/test/config/launchee-override.yml"
	configFileDir        = "launchee"
	configFileYML        = "launchee.yml"
	configFileYAML       = "launchee.yaml"
)

type unmarshalResult struct {
	config *config
	err    error
}

func UnmarshalConfigs() (*frontend.Config, error) {
	systemConfigResult, userConfigResult := unmarshalConfigsAsync()
	if systemConfigResult.err != nil {
		return frontend.NewConfig(0), systemConfigResult.err
	}
	if userConfigResult.err != nil {
		return frontend.NewConfig(0), userConfigResult.err
	}
	if systemConfigResult.config != nil {
		return systemConfigResult.config.sanitize().merge(userConfigResult.config).toFrontendConfig(), nil
	} else if userConfigResult.config != nil {
		return userConfigResult.config.sanitize().toFrontendConfig(), nil
	}
	return frontend.NewConfig(0), nil
}

func unmarshalConfigsAsync() (*unmarshalResult, *unmarshalResult) {
	var systemConfigResult *unmarshalResult
	var userConfigResult *unmarshalResult
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		systemConfigResult = unmarshalConfigFile(getSystemConfigPath())
	}()
	go func() {
		defer wg.Done()
		userConfigResult = unmarshalConfigFile(getUserConfigPath())
	}()
	wg.Wait()
	return systemConfigResult, userConfigResult
}

func unmarshalConfigFile(configFile string) *unmarshalResult {
	if bytes, err := os.ReadFile(configFile); err == nil {
		var config config
		if err = yaml.Unmarshal(bytes, &config); err != nil {
			return &unmarshalResult{nil, errors.WithMessagef(err, "Could not parse %s", configFile)}
		}
		return &unmarshalResult{&config, validate(config.trim())}
	}
	return &unmarshalResult{nil, nil}
}

func (c *config) trim() *config {
	c.Title = strings.TrimSpace(c.Title)
	for _, shortcut := range c.Shortcuts {
		shortcut.Name = strings.TrimSpace(shortcut.Name)
		shortcut.Icon = strings.TrimSpace(shortcut.Icon)
		shortcut.Command = strings.TrimSpace(shortcut.Command)
		shortcut.Url = strings.TrimSpace(shortcut.Url)
		shortcut.Patch = strings.TrimSpace(shortcut.Patch)
	}
	return c
}

func getSystemConfigPath() string {
	var systemConfigPath string
	switch runtime.GOOS {
	case "windows":
		if path := os.Getenv("PROGRAMDATA"); path != "" {
			systemConfigPath = path
		}
	case "linux":
		systemConfigPath = "/etc"
	case "darwin":
		systemConfigPath = "/Library/Application Support"
	}
	return findConfigFile(systemConfigPath)
}

func getUserConfigPath() string {
	if userConfigPath, _ := os.UserConfigDir(); userConfigPath != "" {
		return findConfigFile(userConfigPath)
	}
	return ""
}

func findConfigFile(dir string) string {
	configFileYmlPath := filepath.Join(dir, configFileDir, configFileYML)
	if _, err := os.Stat(configFileYmlPath); err == nil {
		return configFileYmlPath
	}
	configFileYamlPath := filepath.Join(dir, configFileDir, configFileYAML)
	if _, err := os.Stat(configFileYamlPath); err == nil {
		return configFileYamlPath
	}
	return ""
}
