import { Decode } from './../../../../scripts/decode';
import { LoggedInArgs } from '../../public_events';
import { ErrorDispacher } from '../../../noties/public_events';
import { Nick } from '../../scripts/value_objects/nick';
import { Password } from '../../scripts/value_objects/password';
import { fetchBackend } from '../../../backend/api'
import { setTokens } from '../../scripts/account/session_token';

export const Login = async (nick: Nick, password: Password) => {
  const response = await fetchBackend({
    endpoint: "api/v1/users/login",
    method: "POST",
    args: {
      username: nick.Value,
      password: password.Value,
    }
  });

  if (!response.success) ErrorDispacher({ Message: response.message });
  else setTokens(Decode<LoggedInArgs>(response.message));
}
