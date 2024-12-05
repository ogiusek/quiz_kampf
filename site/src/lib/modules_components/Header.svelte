<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import Container from "../components/Container.svelte";
  import Icon from "../components/Icon.svelte";
  import { getSession } from "../modules/user/scripts/account/session";
  import {
    AddSessionChangedListener,
    RemoveSessionChangedListener,
  } from "../modules/user/public_events";
  import type { Link } from "../scripts/links";
  import Links from "../components/Links.svelte";
  import { user_links } from "../modules/user/views/routes";
  import { question_links } from "../modules/question/views/routes";

  const getRenderedLinks = () =>
    Object.entries(links).map(([section, links]) => ({
      section: section,
      links: links
        .filter((link) => link.canRender())
        .map((link) => ({ text: link.text, onclick: link.redirect })),
    }));

  const links: { [section in string]: Link[] } = {
    user: user_links,
    question: question_links,
  };

  let renderedLinks = $state(getRenderedLinks());
  let session = $state(getSession());

  const onSessionChange = () => {
    session = getSession();
    renderedLinks = getRenderedLinks();
  };

  onMount(() => {
    AddSessionChangedListener(onSessionChange);
  });

  onDestroy(() => {
    RemoveSessionChangedListener(onSessionChange);
  });
</script>

<header class="header">
  <Container direction="row" padding="sm" items_position="space-around">
    <Icon
      src="https://www.svgrepo.com/show/475647/facebook-color.svg"
      alt="logo"
    />
    {#each renderedLinks as links}
      {#if links.links.length > 0}
        <h2>{links.section}</h2>
        <Links links={links.links} />
      {/if}
    {/each}
    <Icon src="https://www.svgrepo.com/show/532195/menu.svg" alt="options" />
  </Container>
</header>

<style lang="scss">
  @use "../styles/variables" as *;
  .header {
    position: sticky;
    top: 0;

    width: 100%;
    height: 5rem;

    background-color: $primary-bg-color;
    color: $primary-text-color;

    & > :global(*) {
      flex-grow: 1;
    }
  }
</style>
