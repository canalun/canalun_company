import hatenaEntryList from "@contents/entry_list/Hatena.json"
import { ErrorBoundary, FallbackProps } from 'react-error-boundary'
import React, { useState } from "react";

const ErrorFallback: React.FC<FallbackProps> = ({ error, resetErrorBoundary }) => {
  return (
    <div role="alert">
      <p>Oh...cannot get entries...</p>
      <pre>{error.message}</pre>
      <button onClick={resetErrorBoundary}>Try again???</button>
    </div>
  )
}

export const JapaneseEntryList = () => {
  console.log('sssss');

  const [hatenaListOn, setHatenaListState] = useState(true)

  const listItems = hatenaEntryList.map((entry) =>
    <li key={entry.lastUpdatedAt}>
      <a href={entry.url}>{entry.title}</a>
    </li>
  )

  return (

    <ul>
      <a href="https://canalundayo.hatenablog.com/"
        title="https://canalundayo.hatenablog.com/"><strong>やほほ村</strong></a>:
      かなるんカンパニーのメンバーが毎日の暮らしの中で考えたことや思ったこと<br />
      <ErrorBoundary
        FallbackComponent={ErrorFallback}
        onReset={() => {
          console.log('reset');
          setHatenaListState(false)
        }}
      >
        {hatenaListOn &&
          <div>
            <ul>{listItems}</ul>
          </div>}
      </ErrorBoundary>
      <br />
      <a href="https://zenn.dev/canalun" title="https://zenn.dev/canalun"><strong>Zenn</strong></a>:
      かなるんカンパニーのエンジニアが技術について考えたことや調べたこと<br />
      <ErrorBoundary
        FallbackComponent={ErrorFallback}
        onReset={() => {
          console.log('reset');
          setHatenaListState(false)
        }}
      >
        {hatenaListOn &&
          <div>
            <ul>{listItems}</ul>
          </div>}
      </ErrorBoundary>
    </ul>

  )
}