import { LoggedIn, LoggedInArgs, SessionChanged } from "../../public_events";

const sessionTokenId = "session_token";
export const getSessionToken = (): string => localStorage.getItem(sessionTokenId) ?? "";
const setSessionToken = (token: string | void): void => token ?
  localStorage.setItem(sessionTokenId, token) :
  localStorage.removeItem(sessionTokenId);


const refreshTokenId = "refresh_token";
export const getRefreshToken = (): string => localStorage.getItem(refreshTokenId) ?? "";
const setRefreshToken = (token: string | void): void => token ?
  localStorage.setItem(refreshTokenId, token) :
  localStorage.removeItem(refreshTokenId);


export const setTokens = (args: LoggedInArgs) => {
  setSessionToken(args.session_token);
  setRefreshToken(args.refresh_token);
  if (args.refresh_token && args.session_token) LoggedIn(args);
  SessionChanged({});
}