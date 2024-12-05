<script lang="ts">
  import type { Snippet } from "svelte";

  let {
    children,
  }: {
    children?: Snippet;
  } = $props();
</script>

<div class="wrapper">
  {@render children?.()}
</div>

<style lang="scss">
  .wrapper {
    position: fixed;
    top: 0;
    left: 0;

    width: 100vw;
    height: 100vh;

    z-index: 10;
  }
</style>
