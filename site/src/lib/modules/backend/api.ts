import { api } from './../../scripts/env';
import { NewApi } from "../../scripts/fetch_api";
import { getRefreshToken, getSessionToken, setTokens } from '../user/scripts/account/session_token';
import { AddLoggedInListener, LoggedInArgs, RemoveLoggedInListener } from '../user/public_events';
import { Decode } from '../../scripts/decode';

const {
  getHeaders: getBackendHeaders,
  setHeaders: setBackendHeaders,
  fetch: fetchBackend
} = NewApi({
  api: api,
  headers: {},
  retries: 1,
  beforeRetry: (options, _, result) => new Promise(async (resolve, reject) => {
    if (result.statusCode !== 401) return resolve(false);
    const res = await fetch(`${options.api}/api/v1/users/refresh`, {
      method: "POST",
      body: JSON.stringify({
        session_token: getSessionToken(),
        refresh_token: getRefreshToken()
      })
    });
    if (!res.ok) {
      setTokens({
        session_token: "",
        refresh_token: "",
      })
      return resolve(false);
    }
    const message = await res.text();
    const tokens = Decode<LoggedInArgs>(message);
    const fn = () => {
      RemoveLoggedInListener(fn);
      resolve(true);
    }
    AddLoggedInListener(fn);
    setTokens(tokens);
  })
});

export { fetchBackend };
export const UnavailableBackendError = new Error("api is unavailable");

setBackendHeaders({
  ...getBackendHeaders(),
  "Authorization": getSessionToken(),
});

AddLoggedInListener((session) => {
  setBackendHeaders({
    ...getBackendHeaders(),
    "Authorization": session.session_token
  });
})