import { ErrorBoundary, FallbackProps } from 'react-error-boundary'
import React, { useState } from "react";
import { HatenaEntries } from './hatenaEntries';
import { ZennEntries } from './zennEntries';

const ErrorFallback: React.FC<FallbackProps> = ({ error, resetErrorBoundary }) => {
  return (
    <div role="alert">
      <p>Oh...cannot get entries...</p>
      <pre>{error.message}</pre>
      <button onClick={resetErrorBoundary}>Try again???</button>
    </div>
  )
}

export const JapaneseEntries: React.VFC = () => {
  const [hatenaEntriesOn, setHatenaEntriesState] = useState(true)
  const [zennEntriesOn, setZennEntriesState] = useState(true)

  return (
    <ul>
      <a href="https://canalundayo.hatenablog.com/"
        title="https://canalundayo.hatenablog.com/"><strong>やほほ村</strong></a>:
      かなるんカンパニーのメンバーが毎日の暮らしの中で考えたことや思ったこと<br />
      <ErrorBoundary
        FallbackComponent={ErrorFallback}
        onReset={() => {
          console.log('reset');
          setHatenaEntriesState(false)
        }}
      >
        {hatenaEntriesOn &&
          <HatenaEntries />}
      </ErrorBoundary>
      <br />
      <a href="https://zenn.dev/canalun" title="https://zenn.dev/canalun"><strong>Zenn</strong></a>:
      かなるんカンパニーのエンジニアが技術について考えたことや調べたこと<br />
      <ErrorBoundary
        FallbackComponent={ErrorFallback}
        onReset={() => {
          console.log('reset');
          setZennEntriesState(false)
        }}
      >
        {zennEntriesOn &&
          <ZennEntries />}
      </ErrorBoundary>
    </ul>
  )
}