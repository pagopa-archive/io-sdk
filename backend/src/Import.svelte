<script>

  import { formData } from "./store";
  import { onMount } from "svelte";
  import ImportForm from "./ImportForm.svelte";
  import ImportData from "./ImportData.svelte";

  const base = "http://localhost:3280/api/v1/web/guest/";
  let action = "util/import"

  let url
  $: url = base + action
  
  let state = {};

  async function start() {
    const res = await fetch(url);
    state = await res.json();
    console.log(state);
  }

  onMount(start);
  formData.subscribe(value => {
    state = value;
    console.log(typeof(state));
  });
</script>

<h2>Import</h2>

<div>
  {#if state.form}
    <ImportForm form={state.form} {url} />
  {:else if "error" in state}
    <div class="alert alert-danger" role="alert">Error: {state.error}</div>
    <div>
      <button type="button" class="btn btn-primary" on:click={start}>
        Restart
      </button>
    </div>
  {:else if "data" in state}
    <ImportData data={state.data} />
  {:else}
    <p>Cannot contact importer, please redeploy it.</p>
    <button type="button" class="btn btn-primary" on:click={start}>
      Retry
    </button>
  {/if}
</div>
<!--
<div class="alert alert-success" role="alert">
  {@html result}
</div>
-->
