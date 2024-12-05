<script lang="ts">
  import type { Snippet } from "svelte";

  type Direction = "row" | "col";
  type Size = "xs" | "sm" | "md" | "lg";
  type Max = "max";
  type None = "none";
  type JustifyContent =
    | "start"
    | "center"
    | "space-around"
    | "space-between"
    | "space-evenly";

  let {
    children,
    direction = "col",
    padding = "none",
    margin = "none",
    gap = "sm",
    border = false,
    rounded = "none",
    items_position = "start",
  }: {
    children?: Snippet;
    direction?: Direction;
    padding?: Size | None;
    margin?: Size | None;
    gap?: Size | None;
    border?: boolean;
    rounded?: Size | None | Max;
    items_position?: JustifyContent;
  } = $props();
</script>

<div
  class={`
    container
    direction-${direction}
    padding-${padding}
    margin-${margin}
    gap-${gap}
    rounded-${rounded}
    ${border ? "border" : ""}
  `}
  style="justify-content: {items_position};"
>
  {@render children?.()}
</div>

<style lang="scss">
  @use "../styles/variables" as *;
  .container {
    display: flex;
    width: 100%;
    height: 100%;
    align-items: center;

    &.direction-row {
      flex-direction: row;
    }
    &.direction-col {
      flex-direction: column;
    }

    &.padding-xs {
      padding: 0.25rem;
    }
    &.padding-sm {
      padding: 0.5rem;
    }
    &.padding-md {
      padding: 1rem;
    }
    &.padding-lg {
      padding: 3rem;
    }

    &.margin-xs {
      margin: 0.25rem;
    }
    &.margin-sm {
      margin: 0.5rem;
    }
    &.margin-md {
      margin: 1rem;
    }
    &.margin-lg {
      margin: 3rem;
    }

    &.gap-xs {
      gap: 0.25rem;
    }
    &.gap-sm {
      gap: 0.5rem;
    }
    &.gap-md {
      gap: 1rem;
    }
    &.gap-lg {
      gap: 3rem;
    }

    &.rounded-xs {
      border-radius: 0.25rem;
    }
    &.rounded-sm {
      border-radius: 0.5rem;
    }
    &.rounded-md {
      border-radius: 1rem;
    }
    &.rounded-lg {
      border-radius: 3rem;
    }
    &.rounded-max {
      border-radius: 999999999px;
    }

    &.border {
      // border: 1px solid $primary-text-color;
      -webkit-box-shadow: 0px 0px 18px -9px $primary-text-color;
      -moz-box-shadow: 0px 0px 18px -9px $primary-text-color;
      box-shadow: 0px 0px 18px -9px $primary-text-color;
    }
  }
</style>
