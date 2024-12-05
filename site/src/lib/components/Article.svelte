<script lang="ts">
  import type { Snippet } from "svelte";
  import Container from "./Container.svelte";

  type Size = "auto" | "fixed";

  let {
    children,
    width = "auto",
    height = "auto",
  }: {
    children?: Snippet;
    width?: Size;
    height?: Size;
  } = $props();
</script>

<article class="width-{width} height-{height}">
  <Container direction="col">
    {@render children?.()}
  </Container>
</article>

<style lang="scss">
  @use "../styles/variables" as *;

  article {
    border-radius: 3rem;

    height: max-content;
    width: max-content;

    padding: 3rem;

    background-color: $secondary-bg-color;

    :global(*) {
      color: $secondary-text-color;
    }
  }

  .width-fixed {
    max-width: 95%;
    width: 40rem;
  }
  .height-fixed {
    max-height: 95%;
    height: 30rem;
  }
</style>
