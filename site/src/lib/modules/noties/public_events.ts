import { CreateEvent } from "../../scripts/events"

export class ErrorArgs {
  Message: string
}
export const {
  Dispatch: ErrorDispacher,
  AddListener: AddErrorListener,
  RemoveListener: RemoveErrorListener,
} = CreateEvent<ErrorArgs>("error");

export class NotyArgs {
  Message: string
}
export const {
  Dispatch: NotyDispacher,
  AddListener: AddNotyListener,
  RemoveListener: RemoveNotyListener,
} = CreateEvent<NotyArgs>("noty");