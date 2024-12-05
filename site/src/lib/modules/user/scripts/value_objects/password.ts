export class Password {
  Value: string;
  constructor(value: string) {
    this.Value = value;
  }

  Error(): Error | void {
    if (this.Value == "") {
      return new Error("password cannot be empty")
    }
    return null
  }
}

export class ConfirmPassword {
  Value: string;
  constructor(value: string) {
    this.Value = value;
  }

  Error(password: Password): Error | void {
    return this.Value == password.Value ? null : new Error("password must equal")
  }
}