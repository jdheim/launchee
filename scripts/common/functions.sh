#!/usr/bin/env bash
# COMMON FUNCTIONS

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

set -o errexit  # ABORT ON NON-ZERO EXIT STATUS
set -o nounset  # TREAT UNSET VARIABLES AS AN ERROR AND EXIT
set -o pipefail # DON'T HIDE ERRORS WITHIN PIPES

readonly INFO="[\e[1;34mINFO\e[0m]"
readonly WARNING="[\e[1;33mWARNING\e[0m]"
readonly ERROR="[\e[1;31mERROR\e[0m]"

# General functions

## step "Message"
step() {
  local step="[\e[1;96mSTEP\e[0m]"
  local line="\e[1;96m-----\e[0m"
  local message="${1}"
  echo -e "${step} ${line} ${message} ${line}"
}

## run command --argument
run() {
	echo -e "${INFO} \e[1m$\e[0m $*"; "$@"
}

getProjectVersion() {
  yq -r ".project.version" "jreleaser.yml" | sed "s/-.*//"
}

## updatePropertyInJsonFile "src/wails.json" ".info.productVersion" "0.0.1"
updatePropertyInJsonFile() {
  local file="${1}"
  local propertyName="${2}"
  local propertyValue="${3}"
  local oldValue
  oldValue=$(jq -r "${propertyName}" "${file}")
  if [[ "${oldValue}" != "${propertyValue}" ]]; then
    echo -e "${INFO} Updating ${file}: ${oldValue} -> ${propertyValue}"
    local tmpFile
    tmpFile=$(mktemp)
    cp "${file}" "${tmpFile}" && \
      jq "${propertyName} = \"${propertyValue}\"" "${tmpFile}" > "${file}" && \
      rm "${tmpFile}"
    isUpdated=true
  fi
}

## updateCopyrightInFile "README.md"
updateCopyrightInFile() {
  local file="${1}"
  local oldValue newValue
  copyrightsStartYear="$(grep -E "© [0-9]{4}-[0-9]{4} JDHeim.com" "${file}" | sed -E 's/.*([0-9]{4})-.*/\1/')"
  oldValue="$(grep -E "© [0-9]{4}-[0-9]{4} JDHeim.com" "${file}" | sed -E 's/.*-([0-9]{4}).*/\1/')"
  newValue="$(date +%Y)"
  if [[ "${oldValue}" != "${newValue}" ]]; then
    echo -e "${INFO} Updating ${file}: ${copyrightsStartYear}-${oldValue} -> ${copyrightsStartYear}-${newValue}"
    sed -i "s/\(©\).*\(JDHeim.com\)/\1 ${copyrightsStartYear}-${newValue} \2/" "${file}"
    isUpdated=true
  fi
}
