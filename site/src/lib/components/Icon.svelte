<script lang="ts">
  let {
    src,
    alt,
  }: {
    src: string;
    alt: string;
  } = $props();
</script>

<img class="img" {src} {alt} />

<style lang="scss">
  .img {
    max-height: 100%;
    max-width: 100%;
    height: 3rem;
  }
</style>
