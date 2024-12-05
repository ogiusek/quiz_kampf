<script lang="ts">
  import Icon from "./Icon.svelte";

  let {
    icon_src = "",
    icon_alt = "submit",
  }: {
    icon_src?: string;
    icon_alt?: string;
  } = $props();
</script>

<button type="submit" class="icon-btn">
  <Icon src={icon_src} alt={icon_alt} />
</button>

<style lang="scss">
  @use "../styles/variables" as *;

  .icon-btn {
    max-width: 100%;
    max-height: 100%;

    border-radius: 0.25rem;
    padding: 0.25rem;
  }
</style>
