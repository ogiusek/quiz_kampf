import { AnswerMessage } from "../value_objects/answer_message";

export enum AnswerType {
  Options = 1,
  Text = 2,
}
export const AnswerTypes: { [value in AnswerType]: string } = {
  1: "options",
  2: "text"
};

export class Answer {
  Type: AnswerType;
  Answer: IAnswer;

  constructor(type: AnswerType, answer: IAnswer) {
    this.Type = type;
    this.Answer = answer
  }
}

export class RawAnswer {
  answer_type: number;
  answer_data: any;

  constructor(type: number, data: any) {
    this.answer_type = type;
    this.answer_data = data;
  }
}

interface IAnswer {
  IsCorrect(answer: AnswerMessage): boolean;
}
export class RawAnswerOptions {
  answers: string[] = [];
  correct: number = 0;
}
export class AnswerOptions implements IAnswer {
  Answers: AnswerMessage[];
  Correct: number;

  IsCorrect(answer: AnswerMessage): boolean {
    return this.Answers[this.Correct].Value == answer.Value;
  }

  constructor(answers: AnswerMessage[], correct: number) {
    this.Answers = answers;
    this.Correct = correct;
  }
}
export class RawAnswerText {
  correct_answers: string[] = [];
  case_sensitive: boolean = false;
}
export class AnswerText implements IAnswer {
  CorrectAnswers: AnswerMessage[];
  CaseSentitive: boolean;

  IsCorrect(answer: AnswerMessage): boolean {
    return this.CorrectAnswers.filter(e => this.CaseSentitive ?
      e.Value === answer.Value :
      e.Value.toLowerCase() === answer.Value.toLowerCase()
    ).length !== 0;
  }

  constructor(answers: AnswerMessage[], case_sensitive: boolean) {
    this.CorrectAnswers = answers;
    this.CaseSentitive = case_sensitive;
  }
}

export const RawAnswerToAnswer = (answer: RawAnswer): Answer => new Answer(
  answer.answer_type,
  answer.answer_type == AnswerType.Text ?
    new AnswerText(
      answer.answer_data["answers"].map(e => new AnswerMessage(e)),
      answer.answer_data["case_sensitive"]
    ) : answer.answer_type == AnswerType.Options ?
      new AnswerOptions(
        answer.answer_data["answers"].map(e => new AnswerMessage(e)),
        answer.answer_data["correct"]
      ) : new AnswerText([], false)
)