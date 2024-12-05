import { ConfirmPassword, Password } from '../../scripts/value_objects/password';
import { Nick } from '../../scripts/value_objects/nick';
import { ErrorDispacher, NotyDispacher } from '../../../noties/public_events';
import { fetchBackend } from '../../../backend/api';

const Register = async (nick: Nick, password: Password, _: ConfirmPassword) => {
  const response = await fetchBackend({
    endpoint: "api/v1/users/register",
    method: "POST",
    args: {
      username: nick.Value,
      password: password.Value,
    }
  });

  if (!response.success) ErrorDispacher({ Message: response.response });
  else NotyDispacher({ Message: "Success you can login" });
}

export { Register };