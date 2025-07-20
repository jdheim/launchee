---
sidebar_position: 2
---

# Configuration

Launchee uses a simple YAML file (`launchee.yml`) to define your shortcuts.

Each shortcut can launch an app (binary, script, alias, etc.) or open a URL in your default browser.

## Example

```yaml title="launchee.yml"
title: "My Launchee"
shortcuts:
  - name: "Terminal"
    icon: "/opt/kitty/lib/kitty/logo/kitty-128.png"
    command: "kitty"
  - name: "Weather"
    icon: "/usr/share/icons/Yaru/48x48@2x/apps/weather-app.png"
    url: "https://www.windy.com"
```
## Fields

- `title` - the title of the Launchee window
- `shortcuts` - a list of shortcuts to display in the Launchee window
  - `name` - a unique name for the shortcut
  - `icon` - a path to the icon image
  - Action (on click):
    - `command` - a command to run (binary, script, alias, etc.)
    - `url` - a URL to open in your default browser
