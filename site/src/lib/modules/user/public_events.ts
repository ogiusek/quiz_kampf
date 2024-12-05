import { CreateEvent } from "../../scripts/events"

export class LoggedInArgs {
  session_token: string;
  refresh_token: string;
}
export const {
  Dispatch: LoggedIn,
  AddListener: AddLoggedInListener,
  RemoveListener: RemoveLoggedInListener,
} = CreateEvent<LoggedInArgs>("logged_in");

export class RequireLoginArgs { }
export const {
  Dispatch: RequireLogin,
  AddListener: AddRequireLoginListener,
  RemoveListener: RemoveRequireLoginListener
} = CreateEvent<RequireLoginArgs>("require_login");

export class SessionChangedArgs { }
export const {
  Dispatch: SessionChanged,
  AddListener: AddSessionChangedListener,
  RemoveListener: RemoveSessionChangedListener
} = CreateEvent<SessionChangedArgs>("session_changed");