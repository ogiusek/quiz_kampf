<script lang="ts">
  let {
    value = $bindable(),
    options = {},
  }: {
    value?: any;
    options?: { [value in any]: string };
  } = $props();
  let optionsArr = $derived(Object.entries(options));
  const { valueKey, labelKey } = { valueKey: 0, labelKey: 1 };
  $effect(() => {
    if (options[value]) return;
    if (optionsArr.length === 0) return;
    value = optionsArr[0][valueKey];
  });
  value = `${value}`;
</script>

<select bind:value>
  {#each optionsArr as [val, label] (val)}
    <option value={val}>
      {label}
    </option>
  {/each}
</select>
