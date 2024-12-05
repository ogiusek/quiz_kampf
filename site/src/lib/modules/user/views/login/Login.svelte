<script lang="ts">
  import Main from "../../../../components/Main.svelte";
  import Section from "../../../../components/Section.svelte";
  import View from "../../../../components/View.svelte";
  import Footer from "../../../../modules_components/Footer.svelte";
  import Header from "../../../../modules_components/Header.svelte";
  import Form from "../../../../components/Form.svelte";
  import H1 from "../../../../components/H1.svelte";
  import Input from "../../../../components/Input.svelte";
  import Submit from "../../../../components/Submit.svelte";
  import Links from "../../../../components/Links.svelte";
  import { Nick } from "../../scripts/value_objects/nick";
  import { Password } from "../../scripts/value_objects/password";
  import { Login } from "./login";
  import { IsLoggedIn } from "../user_views";
  import {
    AddLoggedInListener,
    RemoveLoggedInListener,
  } from "../../public_events";
  import { onDestroy, onMount } from "svelte";

  let {
    register,
    exit,
  }: {
    register: () => any;
    exit: () => any;
  } = $props();

  onMount(() => {
    AddLoggedInListener(exit);
    IsLoggedIn().then((is) => is && exit());
  });
  onDestroy(() => RemoveLoggedInListener(exit));

  let raw_nick: string = $state("");
  let nick: Nick = $derived(new Nick(raw_nick));
  let raw_password: string = $state("");
  let password: Password = $derived(new Password(raw_password));
</script>

<View>
  <Header />
  <Main>
    <Section>
      <H1>login</H1>
      <Form onsubmit={() => Login(nick, password)}>
        <Input
          type="text"
          label="nick"
          bind:value={raw_nick}
          validator={(text) => new Nick(text).Error()}
        />
        <Input
          type="password"
          label="password"
          bind:value={raw_password}
          validator={(text) => new Password(text).Error()}
        />
        <Submit>Login</Submit>
      </Form>
      <Links links={[{ onclick: register, text: "register" }]} />
    </Section>
  </Main>
  <Footer />
</View>

<!-- submit uses form.checkValidity() to determine style -->
