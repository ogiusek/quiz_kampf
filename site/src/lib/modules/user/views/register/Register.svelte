<script lang="ts">
  import View from "../../../../components/View.svelte";
  import Header from "../../../../modules_components/Header.svelte";
  import Form from "../../../../components/Form.svelte";
  import Submit from "../../../../components/Submit.svelte";
  import Links from "../../../../components/Links.svelte";
  import Section from "../../../../components/Section.svelte";
  import Main from "../../../../components/Main.svelte";
  import Footer from "../../../../modules_components/Footer.svelte";
  import { Nick } from "../../scripts/value_objects/nick";
  import {
    ConfirmPassword,
    Password,
  } from "../../scripts/value_objects/password";
  import Input from "../../../../components/Input.svelte";
  import H1 from "../../../../components/H1.svelte";
  import { Register } from "./register";
  import { onDestroy, onMount } from "svelte";
  import { IsLoggedIn } from "../user_views";
  import {
    AddLoggedInListener,
    RemoveLoggedInListener,
  } from "../../public_events";

  const {
    login,
    exit,
  }: {
    login: () => any;
    exit: () => any;
  } = $props();

  let nick: string = $state("");
  let password: string = $state("");
  let confirm_password: string = $state("");

  onMount(() => {
    AddLoggedInListener(exit);
    IsLoggedIn().then((is) => is && exit());
  });

  onDestroy(() => {
    RemoveLoggedInListener(exit);
  });
</script>

<View>
  <Header />
  <Main>
    <Section>
      <H1>register</H1>
      <Form
        onsubmit={() =>
          Register(
            new Nick(nick),
            new Password(password),
            new ConfirmPassword(confirm_password),
          )}
      >
        <Input
          type="text"
          label="nick"
          bind:value={nick}
          validator={(text) => new Nick(text).Error()}
        />
        <Input
          type="password"
          label="password"
          bind:value={password}
          validator={(text) => new Password(text).Error()}
        />
        <Input
          type="password"
          label="confirm password"
          bind:value={confirm_password}
          validator={(text) =>
            new ConfirmPassword(text).Error(new Password(password))}
        />
        <Submit>Register</Submit>
      </Form>
      <Links links={[{ onclick: login, text: "login" }]} />
    </Section>
  </Main>
  <Footer />
</View>
