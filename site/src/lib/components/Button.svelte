<script lang="ts">
  import type { Snippet } from "svelte";

  type Accents = "default" | "add" | "remove" | "update";

  let {
    onclick,
    children,
    disabled = false,
    style = "default",
  }: {
    onclick: () => any;
    children?: Snippet;
    disabled?: boolean;
    style?: Accents;
  } = $props();
</script>

<button {disabled} type="button" {onclick} class="btn btn-{style}">
  {@render children?.()}
</button>

<style lang="scss">
  @use "../styles/variables" as *;

  .btn {
    font-size: 1rem;
    border-radius: 3rem;
    margin: 0.5rem;
    padding: 0.5rem 1rem;

    box-sizing: content-box;
    color: $secondary-text-color;
    transition: 0.3s;

    &:disabled {
      opacity: 0.5;
    }

    &:hover {
      opacity: 0.8;
    }
  }

  .btn-default {
    background-color: $accent-color;
  }

  .btn-add {
    background-color: $add-color;
  }

  .btn-remove {
    background-color: $remove-color;
  }

  .btn-update {
    background-color: $update-color;
  }
</style>
