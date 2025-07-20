#!/usr/bin/env bash
# UPDATES ALL DEPENDENCIES

#
# © 2025-2025 JDHeim.com
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

[[ -f "$(dirname "${BASH_SOURCE[0]}")/common/functions.sh" ]] && . "$(dirname "${BASH_SOURCE[0]}")/common/functions.sh"

usage() {
  cat << EOF
Usage: $(basename "$0")

Updates all dependencies
EOF
  exit 1
}

main() {
  cd ..
  readOptions "$@"
  cd src
  update
}

readOptions() {
  while getopts ":h" option; do
    case "${option}" in
      h|?) usage ;;
    esac
  done
}

update() {
  step "Update all dependencies"
  echo -e "${INFO} Updating Go modules"
  go get -u . ./build/... ./cmd/... ./internal/...
  echo -e "${INFO} Add missing and remove unused Go modules"
  go mod tidy
  go install github.com/wailsapp/wails/v2/cmd/wails@latest
  echo -e "${INFO} Updating Frontend dependencies"
  (cd frontend && pnpm update --latest)
  echo -e "${INFO} Updating Docs dependencies"
  (cd ../docs && pnpm update --latest)
}

main "$@"
