<script lang="ts">
  import type { Snippet } from "svelte";
  import Container from "./Container.svelte";

  let {
    children,
  }: {
    children?: Snippet;
  } = $props();
</script>

<section class="section">
  <Container direction="col">
    {@render children?.()}
  </Container>
</section>

<style lang="scss">
  @use "../styles/variables" as *;

  .section {
    background-color: $primary-bg-color;

    width: 100%;
    height: 100%;
  }
</style>
