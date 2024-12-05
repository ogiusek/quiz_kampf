import { mount } from 'svelte'
import App from './App.svelte'
import UserViews from './lib/modules/user/views/UserViews.svelte'
import NotiesViews from './lib/modules/noties/views/NotiesViews.svelte'
import "./lib/styles/style.scss";

const app = mount(App, {
  target: document.getElementById('app'),
})

const users = mount(UserViews, {
  target: document.getElementById('users')
})

const errors = mount(NotiesViews, {
  target: document.getElementById('noties')
})

export default [app, users, errors]
