<script lang="ts">
  import type { Snippet } from "svelte";

  let {
    children,
  }: {
    children?: Snippet;
  } = $props();
</script>

<div class="view">
  {@render children?.()}
</div>

<style lang="scss">
  @use "../styles/variables" as *;

  .view {
    width: 100vw;
    height: 100vh;

    min-height: max-content;

    display: flex;
    flex-direction: column;

    overflow-y: scroll;
  }
</style>
