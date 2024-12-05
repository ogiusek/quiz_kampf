<script lang="ts">
  import { onDestroy, onMount, type Component } from "svelte";
  import {
    AddRequireLoginListener,
    RemoveRequireLoginListener,
    RequireLoginArgs,
  } from "../public_events";
  import Login from "./login/Login.svelte";
  import Register from "./register/Register.svelte";
  import Overlay from "../../../components/Overlay.svelte";
  import { IsLoggedIn } from "./user_views";

  let show: boolean = $state(false);
  let Content: Component = $state(Login);

  const RequireLoginListener = (_: RequireLoginArgs) => {
    IsLoggedIn().then((isLoggedIn) => {
      if (isLoggedIn) return;
      show = true;
    });
  };
  const exit = () => {
    show = false;
  };

  onMount(() => {
    AddRequireLoginListener(RequireLoginListener);
  });

  onDestroy(() => {
    RemoveRequireLoginListener(RequireLoginListener);
  });

  const props = {
    login: () => (Content = Login),
    register: () => (Content = Register),
    forgotPassword: () => {},
    exit: exit,
  };
</script>

{#if show}
  <Overlay>
    <Content {...props} />
  </Overlay>
{/if}
