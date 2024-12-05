<script lang="ts">
  import type { Snippet } from "svelte";
  import Container from "./Container.svelte";

  let {
    children,
    onsubmit,
    validate = () => null,
  }: {
    children?: Snippet;
    onsubmit: () => any;
    validate?: () => string | void;
  } = $props();

  let input: HTMLInputElement;
  $effect(() => {
    if (!input) return;
    input.setCustomValidity(`${validate() ?? ""}`);
  });
</script>

<form
  onsubmit={(e) => {
    e.preventDefault();
    onsubmit();
  }}
>
  <Container direction="col">
    {@render children?.()}
    <input bind:this={input} class="input" />
  </Container>
</form>

<style lang="scss">
  form {
    display: block;
  }

  .input {
    display: none;
  }
</style>
