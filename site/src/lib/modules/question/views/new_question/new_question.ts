import { RawQuestion } from "../../scripts/models/question";
import { fetchBackend } from "../../../backend/api";
import { ErrorDispacher, NotyDispacher } from "../../../noties/public_events";

export const AddQuestion = async (question: RawQuestion) => {
  const res = await fetchBackend({
    endpoint: "api/v1/questions",
    method: "POST",
    args: question
  });
  if (!res.success) return ErrorDispacher({
    Message: res.message
  });
  NotyDispacher({
    Message: "added question succesfuly"
  })
}