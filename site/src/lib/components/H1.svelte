<script lang="ts">
  import type { Snippet } from "svelte";

  let {
    children,
  }: {
    children?: Snippet;
  } = $props();
</script>

<h1>{@render children?.()}</h1>

<style lang="scss">
  h1 {
    font-size: 1.5rem;
    font-weight: 600;
    text-transform: capitalize;
  }
</style>
