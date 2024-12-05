import { RawQuestion } from "../../scripts/models/question";
import { fetchBackend } from "../../../backend/api";
import { ErrorDispacher, NotyDispacher } from "../../../noties/public_events";
import { Decode } from "../../../../scripts/decode";

class SelectQuestionArgs {
  search_question: string;
  search_question_case_sensitive: boolean;
  search_answer: string;
  search_answer_case_sensitive: boolean;
  search_nick: string;
  search_nick_case_sensitive: boolean;
  limit: number;
  page: number;
}

export const GetQuestions = async (args: SelectQuestionArgs) => {
  const res = await fetchBackend({
    endpoint: "api/v1/questions",
    method: "GET",
    args: args
  });
  if (!res.success)
    return ErrorDispacher({ Message: res.message });
  return Decode<RawQuestion[]>(res.message)
}

class RemoveQuestionArgs {
  question_id: string;
}

export const RemoveQuestion = async (args: RemoveQuestionArgs) => {
  const res = await fetchBackend({
    endpoint: "api/v1/questions",
    method: "DELETE",
    args: args
  });
  if (!res.success)
    return ErrorDispacher({ Message: res.message })
  return NotyDispacher({ Message: "succesfully removed" })
} 