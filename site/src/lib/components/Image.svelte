<script lang="ts">
  let {
    src,
    alt,
  }: {
    src: string;
    alt: string;
  } = $props();
</script>

<div class="wrapper">
  <img class="image" {src} {alt} />
</div>

<style>
  .wrapper {
    display: flex;
    width: 100%;
    height: 100%;
  }
  .image {
    flex-grow: 1;
  }
</style>
