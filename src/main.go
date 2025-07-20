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

package main

import (
	"embed"
	"log"
	"os"

	"github.com/jdheim/launchee/build"
	"github.com/jdheim/launchee/cmd"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	launchee := cmd.NewLaunchee()

	err := wails.Run(&options.App{
		Frameless:   true,
		StartHidden: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		LogLevel:           logger.INFO,
		LogLevelProduction: logger.ERROR,
		OnStartup:          launchee.Startup,
		Bind: []interface{}{
			launchee,
		},
		Linux: &linux.Options{
			Icon: build.AppIconBytes,
		},
	})

	if launchee.Config != nil && launchee.Config.Valid == false {
		os.Exit(1)
	}
	if err != nil {
		log.Fatalf("Unexpected Error: %v", err)
	}
}
