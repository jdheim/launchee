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

import {useState} from "react";
import {Tooltip, TooltipContent, TooltipProvider, TooltipTrigger} from "@/components/ui/tooltip.tsx";
import {RunCommand} from "../../../wailsjs/go/cmd/Launchee";
import {frontend} from "../../../wailsjs/go/models.ts";
import {BrowserOpenURL} from "../../../wailsjs/runtime";

export function ShortcutButtonWithTooltip({shortcut, iconSize}: Readonly<{
    shortcut: frontend.Shortcut,
    iconSize: number
}>) {
    const [open, setOpen] = useState(false);

    return (
        <TooltipProvider key={shortcut.Id} delayDuration={0}>
            <Tooltip open={open} onOpenChange={setOpen}>
                <TooltipTrigger asChild>
                    {shortcut?.Command?.length > 0 ? (
                        <button onClick={() => RunCommand(shortcut.Command, shortcut.CommandArgs)}
                                onMouseEnter={() => setOpen(true)}
                                onMouseLeave={() => setOpen(false)}
                                className={`active:scale-y-[0.85] transition-transform`}>
                            <img src={shortcut?.Icon?.Base64}
                                 width={iconSize}
                                 height={iconSize}
                                 alt={shortcut.Name}/>
                        </button>
                    ) : (
                        <button onClick={() => BrowserOpenURL(shortcut.Url)}
                                onMouseEnter={() => setOpen(true)}
                                onMouseLeave={() => setOpen(false)}
                                className={`active:scale-y-[0.85] transition-transform`}>
                            <img src={shortcut?.Icon?.Base64}
                                 width={iconSize}
                                 height={iconSize}
                                 alt={shortcut.Name}/>
                        </button>
                    )}
                </TooltipTrigger>
                <TooltipContent className="dark text-[11px] px-1.5 py-0.4 select-none" side="bottom" sideOffset={1}>
                    {shortcut.Name}
                </TooltipContent>
            </Tooltip>
        </TooltipProvider>
    );
}
