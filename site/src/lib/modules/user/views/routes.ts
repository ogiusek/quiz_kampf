import { Link } from './../../../scripts/links';
import wrap from "svelte-spa-router/wrap";
import Login from "./login/Login.svelte";
import Register from "./register/Register.svelte";
import { push } from "svelte-spa-router";
import Profile from "./profile/Profile.svelte";
import { getSession } from "../scripts/account/session";


const loggedOutProps = {
  login: () => push("/user/login"),
  register: () => push("/user/register"),
  exit: () => push("/")
}

const loggedInProps = {
  exit: () => push("/")
}

export const user_routes = {
  "/user/login": wrap({ // @ts-ignore
    component: Login,
    props: loggedOutProps,
  }),
  "/user/register": wrap({ // @ts-ignore
    component: Register,
    props: loggedOutProps
  }),
  "/account/profile": wrap({ // @ts-ignore
    component: Profile,
    props: loggedInProps
  })
}

export const user_links: Link[] = [
  { text: "login", redirect: () => push("/user/login"), canRender: () => !getSession() },
  { text: "register", redirect: () => push("/user/register"), canRender: () => !getSession() },
  { text: "profile", redirect: () => push("/account/profile"), canRender: () => !!getSession() },
]