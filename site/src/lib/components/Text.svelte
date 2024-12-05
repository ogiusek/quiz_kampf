<script lang="ts">
  import type { Snippet } from "svelte";

  let {
    children,
  }: {
    children?: Snippet;
  } = $props();
</script>

<p>{@render children?.()}</p>

<style lang="scss">
  p {
    text-wrap: wrap;
    word-break: break-all;
  }
</style>
