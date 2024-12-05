<script lang="ts">
  import type { Snippet } from "svelte";

  type Size = "xs" | "sm" | "md" | "lg";

  let {
    children,
    show_sizes,
  }: {
    children?: Snippet;
    show_sizes: Size[];
  } = $props();
</script>

<div class={`container ${show_sizes.join(" ")}`}>
  {@render children?.()}
</div>

<style lang="scss">
  @use "../styles/variables" as *;

  .container {
    display: none;
  }
</style>
