import { Answer, RawAnswer } from "./answer";
import { QuestionText } from "../value_objects/question_text";
import { RawUser, User } from "../../../user/scripts/models/user";

// type Question struct {
// 	QuestionId   id.ID               `json:"question_id"`
// 	CreatorId    id.ID               `json:"creator_id"`
// 	QuestionText models.QuestionText `json:"question_text"`
// 	Answer       Answer              `json:"answer"`
// 	CreatedAt    time.Time           `json:"created_at"`
// }

export class Question {
  questionId: string;
  creator: User;
  questionText: QuestionText;
  answer: Answer;
  createdAt: Date;
}

export class RawQuestion {
  question_id: string;
  creator: RawUser;
  question: string;
  answer: RawAnswer;
  created_at: string;

  constructor(question: string, answer: RawAnswer) {
    this.question = question;
    this.answer = answer;
  }
}

export function NewQuestion() { }