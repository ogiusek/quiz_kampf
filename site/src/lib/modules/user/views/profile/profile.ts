import { Decode } from '../../../../scripts/decode';
import { ErrorDispacher, NotyDispacher } from '../../../noties/public_events';
import { fetchBackend, UnavailableBackendError as UnavailableBackendError } from "../../../backend/api";
import { refreshSession } from '../../scripts/account/session';

export class Profile {
  id: string;
  user_name: string;
  user_image: string;
}

export const GetProfile = async (): Promise<Profile | void> => {
  const response = await fetchBackend({
    endpoint: "api/v1/account/profile",
    method: "GET"
  });

  if (!response.success)
    return ErrorDispacher({ Message: UnavailableBackendError.message });

  return Decode<Profile>(response.message)
}

export const Rename = async (nick: string) => {
  const response = await fetchBackend({
    endpoint: "api/v1/account/rename",
    method: "POST",
    args: {
      new_name: nick
    }
  });

  if (!response.success)
    return ErrorDispacher({ Message: response.message });

  NotyDispacher({ Message: `succesfuly renamed to "${nick}"` });
  await refreshSession();
}