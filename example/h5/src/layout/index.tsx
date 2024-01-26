import { Loading } from '@/components/ui'
import { profile } from '@/services/api'
import { appAtom } from '@/store/app'
import { globalStore, useGlobalStoreAtom } from '@/store/global'
import { produce } from 'immer'
import { Provider } from 'jotai'
import React, { Suspense, useEffect } from 'react'
import {
  useRoutes
} from 'react-router-dom'
import routes from '~react-pages'

const LayoutPage = React.memo(() => {
  const [ appStore, setAppStore ] = useGlobalStoreAtom(appAtom)
  const pages = useRoutes(routes)

  useEffect(() => {
    profile().then(user => {
      if (user) {
        setAppStore(produce((state) => {
          state.userInfo = user
        }))
      }
    })
  }, [])

  if (!appStore.userInfo) {
    return <Loading />
  }

  return (
    <Suspense fallback={<Loading />}>
      <Provider store={globalStore}>
        {pages}
      </Provider>
    </Suspense>
  )
})

export default LayoutPage
