import hatenaEntryList from "@contents/entry_list/Hatena.json"
import React from "react"

export const HatenaEntries: React.VFC = () => {
	const hatenaListItem = hatenaEntryList.map((entry) =>
		<li key={entry.lastUpdatedAt}>
			<a href={entry.url}>{entry.title}</a>
		</li>
	)

	return (
		<div><ul>{hatenaListItem}</ul></div>
	)
}