import { BaseProps } from '@/utils/props'
import { Icon } from '@iconify/react'
import React from 'react'

export type LoadingProps = BaseProps

export const Loading: React.FC<LoadingProps> = ({ className, style }) => {
  return (
    <Icon
      icon="eos-icons:three-dots-loading"
      className={className}
      style={style}
    />
  )
}