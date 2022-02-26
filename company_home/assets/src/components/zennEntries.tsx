import zennEntryList from "@contents/entry_list/Zenn.json"
import React from "react";

export const ZennEntries: React.VFC = () => {
    const zennListItem = zennEntryList.map((entry) =>
        <li key={entry.lastUpdatedAt}>
            <a href={entry.url}>{entry.title}</a>
        </li>
    )

    return (
        <div><ul>{zennListItem}</ul></div>
    )
}