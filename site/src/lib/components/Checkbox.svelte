<script lang="ts">
  import type { Snippet } from "svelte";

  let {
    value = $bindable(),
    onchange = () => {},
    children,
  }: {
    value: boolean;
    onchange?: () => any;
    children?: Snippet;
  } = $props();
  let input: HTMLInputElement;
</script>

<button
  onclick={() => input.click()}
  class="wrapper {value ? 'checked' : 'unchecked'}"
>
  <label>
    <input
      bind:this={input}
      bind:checked={value}
      {onchange}
      class="input"
      type="checkbox"
    />
    {@render children?.()}
  </label>
</button>

<style lang="scss">
  @use "../styles/variables" as *;
  .wrapper {
    border-radius: 10rem;
    padding: 0.5rem 1rem;
    transition: 0.3s;

    &.checked {
      background-color: $add-color;
    }
    &.unchecked {
      background-color: $remove-color;
    }

    .input {
      display: none;
    }

    label {
      color: $secondary-text-color;
      font-size: 1rem;
    }
  }
</style>
