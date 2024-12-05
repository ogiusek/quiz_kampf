import { Decode } from './../../../../scripts/decode';
import { LoggedInArgs } from './../../public_events';
import { fetchBackend } from "../../../backend/api";
import { Nick } from '../value_objects/nick';
import { getRefreshToken, getSessionToken, setTokens } from "./session_token";
import { jwtDecode } from "jwt-decode";

class RawSession {
  exp: number;
  user_id: string;
  user_name: string;
  user_photo: string;
}

class Session {
  UserId: string;
  UserName: Nick;
  UserPhoto: string;
}

export const getSession = (): Session | void => {
  try {
    const session = jwtDecode<RawSession>(getSessionToken());
    // if (Date.now() >= session.exp * 1000) return;
    return {
      UserId: session.user_id,
      UserName: new Nick(session.user_name),
      UserPhoto: session.user_photo,
    }
  } catch (_) { }
}

/**
 * @returns {boolean} true on success
 */
export const refreshSession = async (): Promise<boolean> => {
  const sessionToken = getSessionToken();
  const refreshToken = getRefreshToken();
  if (!sessionToken || !refreshToken) {
    setTokens({
      refresh_token: "",
      session_token: ""
    });
    return false;
  }
  const response = await fetchBackend({
    endpoint: "api/v1/users/refresh",
    method: "POST",
    args: {
      session_token: sessionToken,
      refresh_token: refreshToken
    }
  });

  if (!response.success) {
    setTokens({
      refresh_token: "",
      session_token: ""
    });
    return false;
  }

  setTokens(Decode<LoggedInArgs>(response.message));

  return true;
}