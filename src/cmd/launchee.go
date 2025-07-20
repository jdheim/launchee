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

package cmd

import (
	"context"
	"fmt"
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

func NewLaunchee() *Launchee {
	return &Launchee{}
}

func newErrorMessageDialog(message string, err error) {
	if _, err := runtime.MessageDialog(lctx.GetContext(), runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "Oops... Something Went Wrong",
		Message: fmt.Sprintf("%s:\n%v", message, err),
	}); err != nil {
		runtime.LogErrorf(lctx.GetContext(), "Error while displaying message dialog: %v", err)
	}
}

// Startup is called when the app starts. The context is saved so we can call the runtime methods
func (l *Launchee) Startup(ctx context.Context) {
	defer util.Measure(ctx, "Startup")()
	lctx.SetContext(ctx)
	config, err := yaml.UnmarshalConfigs()
	l.Config = config
	if err != nil {
		l.Config.Valid = false
		newErrorMessageDialog("Error occurred during application startup", err)
		runtime.Quit(ctx)
		return
	}
	l.postStartup()
}

func (l *Launchee) postStartup() {
	l.Config.Valid = true
	width := l.Config.UI.Width()
	height := l.Config.UI.Height(len(l.Config.Shortcuts))
	ctx := lctx.GetContext()
	runtime.WindowSetTitle(ctx, l.Config.UI.Nav.Title)
	runtime.WindowSetSize(ctx, width, height)
	runtime.WindowSetMinSize(ctx, width, height)
	runtime.WindowSetMaxSize(ctx, width, height)
}

func (l *Launchee) GetAppVersion() string {
	return appVersion
}

func (l *Launchee) IsBuildForJdvm() bool {
	return buildForJdvm == "true"
}

func (l *Launchee) GetConfig() *frontend.Config {
	return l.Config
}

func (l *Launchee) RunCommand(command []string) {
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		newErrorMessageDialog("Error occurred when running a command", err)
		return
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			runtime.LogErrorf(lctx.GetContext(), "Error occurred when finishing a command %v: %v", cmd, err)
		}
	}()
}
