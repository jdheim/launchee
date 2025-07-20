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

import {useEffect, useState} from "react";
import {GetConfig} from "../wailsjs/go/cmd/Launchee";
import {WindowSetPosition, WindowShow} from "../wailsjs/runtime";
import {ShortcutGrid} from "@/components/content/ShortcutGrid.tsx";
import {TitleBar} from "@/components/nav/TitleBar.tsx";
import {frontend} from "../wailsjs/go/models";

function Launchee() {
    const [config, setConfig] = useState<frontend.Config | null>(null);
    const shortcuts = config?.Shortcuts ?? [];
    const ui = config?.UI ?? null;

    useEffect(() => {
        GetConfig().then(config => {
            if (config.Valid) {
                setConfig(config)
                WindowSetPosition(0, 0);
                WindowShow();
            }
        });
    }, []);

    return (
        <div className="grid grid-rows-[auto_1fr] h-screen w-screen bg-gradient-to-b from-[#48494C] to-[#2F3032] border-x-1 border-b-1 border-[#1e1f22] cursor-default select-none">
            <TitleBar nav={ui?.Nav ?? null}/>
            <ShortcutGrid content={ui?.Content ?? null}
                          shortcuts={shortcuts}/>
        </div>
    )
}

export default Launchee;
