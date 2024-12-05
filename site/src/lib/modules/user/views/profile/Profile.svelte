<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import { GetProfile, Profile, Rename } from "./profile";
  import {
    AddSessionChangedListener,
    RemoveSessionChangedListener,
    RequireLogin,
  } from "../../public_events";
  import { getSession } from "../../scripts/account/session";
  import View from "../../../../components/View.svelte";
  import Header from "../../../../modules_components/Header.svelte";
  import Main from "../../../../components/Main.svelte";
  import H1 from "../../../../components/H1.svelte";
  import Section from "../../../../components/Section.svelte";
  import Footer from "../../../../modules_components/Footer.svelte";
  import Image from "../../../../components/Image.svelte";
  import Form from "../../../../components/Form.svelte";
  import Input from "../../../../components/Input.svelte";
  import Container from "../../../../components/Container.svelte";
  import SubmitIcon from "../../../../components/SubmitIcon.svelte";
  import Submit from "../../../../components/Submit.svelte";
  import Icon from "../../../../components/Icon.svelte";
  import { getImageSrc } from "../../../user/scripts/value_objects/image_src";

  let {
    // exit,
  }: {
    exit: () => any;
  } = $props();

  let profile: Profile = $state();

  function RefreshProfile() {
    if (!getSession()) {
      RequireLogin({});
      return;
    }
    GetProfile().then((new_profile) => new_profile && (profile = new_profile));
  }

  onMount(() => {
    RefreshProfile();
    AddSessionChangedListener(RefreshProfile);
  });

  onDestroy(() => {
    RemoveSessionChangedListener(RefreshProfile);
  });

  let nick = $state("");
  $effect(() => {
    if (!profile) return;
    nick = profile.user_name;
  });
</script>

<View>
  <Header />
  <Main>
    <Section>
      <H1>profile</H1>
      <Form onsubmit={() => profile?.user_name != nick && Rename(nick)}>
        <Icon src={getImageSrc(profile?.user_image)} alt="profile" />
        <Input label="nick" bind:value={nick} />
        <Submit>Save changes</Submit>
      </Form>
    </Section>
  </Main>
  <Footer />
</View>
