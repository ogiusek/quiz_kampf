<script lang="ts">
  import { onMount } from "svelte";

  type Type = "text" | "password";
  type Value = string;
  let {
    type = "text",
    label,
    value = $bindable(),
    validator = () => {},
    onChange = () => {},
  }: {
    type?: Type;
    label: string;
    value: Value;
    validator?: (input: Value) => Error | void;
    onChange?: () => void;
  } = $props();

  let input: HTMLInputElement;
  let error: string = $state("");

  const Validate = () => {
    const err = validator(value);
    const error_message: string = err ? err.message : "";
    if (input.validationMessage != error_message)
      input.setCustomValidity(error_message);
    error = error_message;
  };

  onMount(Validate);
  $effect(Validate);
</script>

<button type="button" onclick={() => input.focus()} class="btn" tabindex="-1">
  <label class="label">
    <p class="label-text">{label}</p>
    <input
      placeholder={label}
      onchange={() => {
        Validate();
        onChange();
      }}
      bind:this={input}
      class="input"
      {type}
      bind:value
    />
  </label>
</button>

<style lang="scss">
  @use "../styles/variables" as *;

  .btn {
    background-color: transparent;

    .label {
      display: flex;
      justify-content: end;
      align-items: center;
      flex-direction: column;

      padding: 0.5rem 1rem;
      border-radius: 2px;
      height: 2.5rem;
      border: 1px solid $primary-text-color;

      margin-top: 0.25rem;
      text-align: left;
      text-decoration: underline;
      position: relative;

      * {
        color: $primary-text-color;
      }

      .label-text {
        // display: none;
        z-index: -1;

        position: absolute;
        top: 0%;
        left: 0.25rem;

        transition: 0.2s;
        transition-delay: 0.1s;
        background-color: $primary-bg-color;
        padding: 0 2px;

        font-weight: 100;
        color: transparent;
      }

      .input {
        width: 20rem;
        outline: none;
      }

      &:has(:focus) {
        outline: $accent-color 1px solid;
        input::placeholder {
          color: transparent;
        }
        &,
        * {
          color: $accent-color !important;
        }
      }

      &:has(:focus),
      &:not(:has(.input:placeholder-shown)) {
        .label-text {
          display: block;
          z-index: 1;
          transform: translateY(-50%);
          color: $primary-text-color;
        }
      }
    }
  }
</style>
