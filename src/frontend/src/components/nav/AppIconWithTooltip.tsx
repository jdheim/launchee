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

import {Tooltip, TooltipContent, TooltipProvider, TooltipTrigger} from "@/components/ui/tooltip.tsx";
import {BrowserOpenURL} from "../../../wailsjs/runtime";
import {useState} from "react";
import {frontend} from "../../../wailsjs/go/models.ts";

export function AppIconWithTooltip({appIcon, iconSize, url}: Readonly<{
    appIcon: frontend.Icon | null,
    iconSize: number,
    url: string
}>) {
    const [open, setOpen] = useState(false);

    return (
        <>
            {appIcon?.Base64 && (
                <TooltipProvider delayDuration={0}>
                    <Tooltip open={open} onOpenChange={setOpen}>
                        <TooltipTrigger asChild>
                            <button onClick={() => BrowserOpenURL(url)}
                                    onMouseEnter={() => setOpen(true)}
                                    onMouseLeave={() => setOpen(false)}
                                    className={`active:scale-y-[0.85] transition-transform`}>
                                <img src={appIcon.Base64}
                                     width={iconSize}
                                     height={iconSize}
                                     alt=""/>
                            </button>
                        </TooltipTrigger>
                        <TooltipContent className="dark text-[11px] px-1.5 py-0.4 select-none"
                                        sideOffset={5}>
                            {url}
                        </TooltipContent>
                    </Tooltip>
                </TooltipProvider>
            )}
        </>
    )
}
