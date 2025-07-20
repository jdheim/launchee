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

import {ShortcutButtonWithTooltip} from "./ShortcutButtonWithTooltip.tsx";
import {HelpWithTooltip} from "@/components/content/HelpWithTooltip.tsx";
import {frontend} from "../../../wailsjs/go/models.ts";
import {useEffect, useState} from "react";
import {GetAppVersion} from "../../../wailsjs/go/cmd/Launchee";

export function ShortcutGrid({content, shortcuts}: Readonly<{
    content: frontend.Content | null,
    shortcuts: frontend.Shortcut[]
}>) {
    const defaultIconColumns = 1;
    const defaultIconSize = 8 * 4;
    const defaultMargin = 5;

    const iconColumnsClass = `grid-cols-${content?.IconColumns ?? defaultIconColumns}`;
    const iconSize = content?.IconSize ?? defaultIconSize;
    const marginClass = `m-${content?.Margin ?? defaultMargin}`;
    const gapClass = `gap-${content?.Margin ?? defaultMargin}`;

    const [appVersion, setAppVersion] = useState<string | null>(null);
    const appVersionText = appVersion ? ` ${appVersion}` : "";
    const appVersionTooltipText = `Launchee${appVersionText}`;

    useEffect(() => {
        GetAppVersion().then(setAppVersion);
    }, []);

    return (
        <div className={`grid ${iconColumnsClass} place-items-center-safe ${marginClass} ${gapClass}`}>
            {shortcuts.length > 0 ? (shortcuts.map((shortcut) => (
                <ShortcutButtonWithTooltip key={shortcut.Id}
                                           shortcut={shortcut}
                                           iconSize={iconSize}/>
            ))) : (
                <HelpWithTooltip/>
            )}
            <div className="absolute bottom-[2px] right-[4px] text-[#888a91] text-[9.5px] opacity-0 hover:opacity-100 transition-opacity duration-200 ease-in-out">
                <span>{appVersionTooltipText}</span>
            </div>
        </div>
    );
}
