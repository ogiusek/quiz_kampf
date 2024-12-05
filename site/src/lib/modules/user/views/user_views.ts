import { LoggedIn } from "../public_events";
import { getSession, refreshSession } from "../scripts/account/session";
import { getRefreshToken, getSessionToken } from "../scripts/account/session_token";

export const IsLoggedIn = async (): Promise<boolean> => {
  if (getSession()) {
    LoggedIn({
      session_token: getSessionToken(),
      refresh_token: getRefreshToken(),
    });
    return true;
  }
  return await refreshSession();
}
