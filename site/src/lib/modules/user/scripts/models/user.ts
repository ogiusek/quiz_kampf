import { Nick } from './../value_objects/nick';

export class User {
  id: string;
  user_name: Nick;
  user_image: string;
}

export class RawUser {
  id: string;
  user_name: string;
  user_image: string;
  ToUser(): User {
    return {
      id: this.id,
      user_name: new Nick(this.user_name),
      user_image: this.user_image
    }
  }
}