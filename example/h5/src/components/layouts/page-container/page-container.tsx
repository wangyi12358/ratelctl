import { BaseProps } from '@/utils/props'
import React, { PropsWithChildren } from 'react'

export interface PageContainerProps extends BaseProps {
  title?: React.ReactNode
}

export const PageContainer: React.FC<PropsWithChildren<PageContainerProps>> = ({ title, children, className, style }) => {
  return (
    <div className={className} style={style}>
      {title && (
        <div className='w-full h-10 text-center leading-10 bg-slate-100'>
          {title}
        </div>
      )}
      {children}
    </div>
  )
}