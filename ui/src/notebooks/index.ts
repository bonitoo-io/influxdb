import {FunctionComponent, ComponentClass, ReactNode} from 'react'
import {RemoteDataState} from 'src/types'
import {FromFluxResult} from '@influxdata/giraffe'

export interface BothResults {
  source: string
  parsed: FromFluxResult
  raw: string
  error?: string
}

export interface PipeContextProps {
  children?: ReactNode
  controls?: ReactNode
}

export type PipeData = any

export interface PipeProp {
  data: PipeData
  onUpdate: (data: PipeData) => void
  results?: BothResults
  loading: RemoteDataState

  Context:
    | FunctionComponent<PipeContextProps>
    | ComponentClass<PipeContextProps>
}

// NOTE: keep this interface as small as possible and
// don't take extending it lightly. this should only
// define what ALL pipe types require to be included
// on the page.
export interface TypeRegistration {
  type: string // a unique string that identifies a pipe
  priority?: number // 0 is lowest priority, equal priorities revert to string comparison
  disabled?: boolean // if you should show it or not
  featureFlag?: string // designates a flag that should enable the panel type
  component: FunctionComponent<PipeProp> | ComponentClass<PipeProp> // the view component for rendering the interface
  button: string // a human readable string for appending the type
  initial: any // the default state for an add
}

export interface TypeLookup {
  [key: string]: TypeRegistration
}

export const PIPE_DEFINITIONS: TypeLookup = {}

export function register(definition: TypeRegistration) {
  if (PIPE_DEFINITIONS.hasOwnProperty(definition.type)) {
    throw new Error(
      `Pipe of type [${definition.type}] has already been registered`
    )
  }

  PIPE_DEFINITIONS[definition.type] = {
    ...definition,
  }
}

// NOTE: this loads in all the modules under the current directory
// to make it easier to add new types
const context = require.context('./pipes', true, /index\.(ts|tsx)$/)
context.keys().forEach(key => {
  context(key)
})
