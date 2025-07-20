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

package cmd

import (
	"context"
	"os"
	"os/exec"

	"github.com/jdheim/launchee/internal/config/frontend"
	"github.com/jdheim/launchee/internal/config/yaml"
	"github.com/jdheim/launchee/internal/lctx"
	"github.com/jdheim/launchee/internal/util"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Launchee struct {
	Config *frontend.Config
}

var appVersion string

var buildForJdvm string

var customConfigPath string

func NewLaunchee() *Launchee {
	return &Launchee{}
}

type Window interface {
	SetTitle(title string)
	SetSize(width int, height int)
	SetMinSize(width int, height int)
	SetMaxSize(width int, height int)
	Quit()
}

type windowRuntime struct{}

func (windowRuntime) SetTitle(title string) {
	runtime.WindowSetTitle(lctx.GetContext(), title)
}

func (windowRuntime) SetSize(width int, height int) {
	runtime.WindowSetSize(lctx.GetContext(), width, height)
}

func (windowRuntime) SetMinSize(width int, height int) {
	runtime.WindowSetMinSize(lctx.GetContext(), width, height)
}

func (windowRuntime) SetMaxSize(width int, height int) {
	runtime.WindowSetMaxSize(lctx.GetContext(), width, height)
}

func (windowRuntime) Quit() {
	runtime.Quit(lctx.GetContext())
}

var windowImpl Window = windowRuntime{}

// Startup is called when the app starts. The context is saved so we can call the runtime methods
func (l *Launchee) Startup(ctx context.Context) {
	defer util.Measure("Startup")()
	lctx.SetContext(ctx)
	var config *frontend.Config
	var err error
	if customConfigPath != "" {
		config, err = yaml.UnmarshalCustomConfig(customConfigPath)
	} else {
		config, err = yaml.UnmarshalConfigs()
	}
	l.Config = config
	if err != nil {
		l.Config.Valid = false
		lctx.NewErrorMessageDialog("Error occurred during application startup", err)
		windowImpl.Quit()
		return
	}
	l.postStartup()
}

func (l *Launchee) postStartup() {
	width := l.Config.UI.Width()
	height := l.Config.UI.Height(len(l.Config.Shortcuts))
	windowImpl.SetTitle(l.Config.UI.Nav.Title)
	windowImpl.SetSize(width, height)
	windowImpl.SetMinSize(width, height)
	windowImpl.SetMaxSize(width, height)
}

func (l *Launchee) GetAppVersion() string {
	return appVersion
}

func (l *Launchee) IsBuildForJdvm() bool {
	return buildForJdvm == "true"
}

func (l *Launchee) GetCustomConfigPath() string {
	return customConfigPath
}

func (l *Launchee) SetCustomConfigPath(newCustomConfigPath string) {
	customConfigPath = newCustomConfigPath
}

func (l *Launchee) GetConfig() *frontend.Config {
	return l.Config
}

func (l *Launchee) RunCommand(command string, commandArgs []string) {
	cmd := exec.Command(command, commandArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		lctx.NewErrorMessageDialog("Error occurred when running a command", err)
		return
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			lctx.LogErrorf("Error occurred when finishing a command %v: %v", cmd, err)
		}
	}()
}
