<script lang="ts">
  import type { Snippet } from "svelte";

  let {
    children,
  }: {
    children?: Snippet;
  } = $props();
</script>

<button class="btn" type="submit">
  {@render children?.()}
</button>

<style lang="scss">
  @use "../styles/variables" as *;

  .btn {
    width: 22rem;
    padding: 0.5rem 1rem;
    border-radius: 2px;
    background-color: $secondary-bg-color;
    color: $secondary-text-color;

    font-size: 1rem;
    font-weight: 400;

    transition: 0.3s;

    :global(form:invalid) & {
      opacity: 0.6;
    }

    :global(form:valid) & {
      &:hover,
      &:focus {
        background-color: $accent-color;
        outline: none;
      }
    }
  }
</style>
