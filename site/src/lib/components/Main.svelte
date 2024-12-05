<script lang="ts">
  import type { Snippet } from "svelte";

  let {
    children,
    changesHeight = false,
  }: {
    children?: Snippet;
    changesHeight?: boolean;
  } = $props();
</script>

<main class="main {changesHeight ? 'changes' : ''}">
  <div class="children">
    {@render children?.()}
  </div>
</main>

<style lang="scss">
  @use "../styles/variables" as *;

  main {
    flex-grow: 1;

    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    background-color: $primary-bg-color;

    min-height: max-content;
    overflow-y: scroll;
    padding: 1rem;

    &.changes {
      justify-content: start;
    }

    .children {
      padding: 0 1rem;

      display: flex;
      justify-content: center;
      align-items: center;
      flex-direction: column;
    }
  }
</style>
