import { getSession } from './../../user/scripts/account/session';
import { push } from 'svelte-spa-router';
import { Link } from './../../../scripts/links';
import NewQuestion from "./new_question/NewQuestion.svelte";
import { wrap } from 'svelte-spa-router/wrap';
import GetQuestions from './select_questions/GetQuestions.svelte';

export const question_routes = {
  "/question/add": wrap({ // @ts-ignore
    component: NewQuestion,
  }),
  "/question/search": wrap({ // @ts-ignore
    component: GetQuestions,
  })
};

export const question_links: Link[] = [
  { text: "add question", redirect: () => push("/question/add"), canRender: () => !!getSession() },
  { text: "search questions", redirect: () => push("/question/search"), canRender: () => true }
];