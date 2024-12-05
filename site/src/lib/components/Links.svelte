<script lang="ts">
  type Link = {
    onclick: () => any;
    text: string;
  };

  let {
    links,
  }: {
    links: Link[];
  } = $props();
</script>

<ul class="list row">
  {#if links.length === 1}
    <li></li>
  {/if}
  {#each links as button}
    <li class="element">
      <button type="button" class="btn" onclick={button.onclick}
        >{button.text}</button
      >
    </li>
  {/each}
</ul>

<style lang="scss">
  @use "../styles/variables" as *;
  .list {
    width: 100%;
    display: flex;

    justify-content: space-between;
    align-items: center;

    &.row {
      flex-direction: row;
    }

    .element {
      .btn {
        background-color: transparent;
        font-size: 1rem;
        text-decoration: underline;
        color: $link-color;
        transition: 0.3s;

        &:hover {
          cursor: pointer;
          opacity: 0.7;
        }
      }
    }
  }
</style>
