<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import {
    ErrorArgs,
    AddErrorListener,
    RemoveErrorListener,
    NotyArgs,
    AddNotyListener,
    RemoveNotyListener,
  } from "../public_events";

  const OnError = (error: ErrorArgs) => {
    alert(`error:\n${error.Message}`);
  };

  const OnNoty = (noty: NotyArgs) => {
    alert(`notification:\n${noty.Message}`);
  };

  onMount(() => {
    AddErrorListener(OnError);
    AddNotyListener(OnNoty);
  });
  onDestroy(() => {
    RemoveErrorListener(OnError);
    RemoveNotyListener(OnNoty);
  });
</script>
