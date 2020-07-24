// Libraries
import React, {FC, useEffect, useContext, useCallback} from 'react'

// Components
import {
  DapperScrollbars,
  TechnoSpinner,
  ComponentSize,
  RemoteDataState,
} from '@influxdata/clockface'
import SelectorListItem from 'src/notebooks/pipes/Data/SelectorListItem'
import {BucketContext} from 'src/notebooks/context/buckets'
import {PipeContext} from 'src/notebooks/context/pipe'

// Types
import {Bucket} from 'src/types'

const BucketSelector: FC = () => {
  const {data, update} = useContext(PipeContext)
  const {buckets, loading} = useContext(BucketContext)

  const selectedBucketName = data.bucketName

  const updateBucket = useCallback(
    (updatedBucket: Bucket): void => {
      update({bucketName: updatedBucket.name})
    },
    [update]
  )

  useEffect(() => {
    // selectedBucketName will only evaluate false on the initial render
    // because there is no default value
    if (!!buckets.length && !selectedBucketName) {
      updateBucket(buckets[0])
    }
  }, [buckets, selectedBucketName, updateBucket])

  let body

  if (loading === RemoteDataState.Loading) {
    body = (
      <div className="data-source--list__empty">
        <TechnoSpinner strokeWidth={ComponentSize.Small} diameterPixels={32} />
      </div>
    )
  }

  if (loading === RemoteDataState.Error) {
    body = (
      <div className="data-source--list__empty">
        <p>Could not fetch Buckets</p>
      </div>
    )
  }

  if (loading === RemoteDataState.Done && selectedBucketName) {
    body = (
      <DapperScrollbars className="data-source--list">
        {buckets.map(bucket => (
          <SelectorListItem
            key={bucket.name}
            value={bucket}
            onClick={updateBucket}
            selected={bucket.name === selectedBucketName}
            text={bucket.name}
          />
        ))}
      </DapperScrollbars>
    )
  }

  return (
    <div className="data-source--block">
      <div className="data-source--block-title">Bucket</div>
      {body}
    </div>
  )
}

export default BucketSelector
