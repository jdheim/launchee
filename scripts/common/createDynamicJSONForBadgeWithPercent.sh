#!/usr/bin/env bash
# CREATES DYNAMIC JSON FOR BADGE WITH PERCENT

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

[[ -f "$(dirname "${BASH_SOURCE[0]}")/functions.sh" ]] && . "$(dirname "${BASH_SOURCE[0]}")/functions.sh"

# createDynamicJSONForBadgeWithPercent.sh "go-test-coverage" "go" "Test Coverage" "73.9"
main() {
  step "Create dynamic JSON for badge with percent"

  local jsonFilename="${1}"
  local namedLogo="${2}"
  local logoColor="white"
  local label="${3}"
  local message="${4}"

  local testCoverage="${message%%%}"
  local color="red"
  if (( $(echo "${testCoverage} >= 90" | bc -l) )); then
    color="brightgreen"
  elif (( $(echo "${testCoverage} >= 80" | bc -l) )); then
    color="green"
  elif (( $(echo "${testCoverage} >= 70" | bc -l) )); then
    color="yellowgreen"
  elif (( $(echo "${testCoverage} >= 50" | bc -l) )); then
    color="yellow"
  elif (( $(echo "${testCoverage} >= 30" | bc -l) )); then
    color="orange"
  fi

  jq -n \
    --arg namedLogo "${namedLogo}" \
    --arg logoColor "${logoColor}" \
    --arg label "${label}" \
    --arg message "${message}" \
    --arg color "${color}" \
    '{
       schemaVersion: 1,
       namedLogo: $namedLogo,
       logoColor: $logoColor,
       label: $label,
       message: $message,
       color: $color
     }' > "${jsonFilename}.json"
}

main "$@"
