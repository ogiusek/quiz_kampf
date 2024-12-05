export class Nick {
  Value: string
  constructor(email: string) {
    this.Value = email;
  }

  Error(): Error | void {
    if (this.Value == "") {
      return new Error("nick cannot be empty")
    }
  }
}
