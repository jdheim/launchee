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

package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/jdheim/launchee/build"
	"github.com/jdheim/launchee/cmd"
	flag "github.com/spf13/pflag"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	processCommandLineFlags()
	startGUI()
}

var parse = flag.Parse

func processCommandLineFlags() {
	help := flag.BoolP("help", "h", false, "Show help")
	version := flag.BoolP("version", "v", false, "Show version")
	customConfigPath := flag.StringP("config", "c", "", "Set custom config path, e.g. `/tmp/launchee.yml`")

	parse()

	switch {
	case *help:
		fmt.Printf("Launchee - clean, minimalist dock for launching your essential shortcuts\n\n")
		flag.Usage()
		os.Exit(0)
	case *version:
		fmt.Println("Launchee version:", cmd.NewLaunchee().GetAppVersion())
		os.Exit(0)
	case *customConfigPath != "":
		if _, err := os.Stat(*customConfigPath); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		cmd.NewLaunchee().SetCustomConfigPath(*customConfigPath)
	}
}

func startGUI() {
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
			Icon: build.GetAppIconBytes(),
		},
	})

	assertLauncheeConfigValid(launchee)
	assertError(err)
}

func assertLauncheeConfigValid(launchee *cmd.Launchee) {
	if launchee.Config != nil && launchee.Config.Valid == false {
		log.Fatalf("Invalid Config: %+v", launchee.Config)
	}
}

func assertError(err error) {
	if err != nil {
		log.Fatalf("Unexpected Error: %v", err)
	}
}
