<script>
  export let action;

  import { formData } from "./store";
  import { onMount } from "svelte";
  import ImportForm from "./ImportForm.svelte";
  import ImportData from "./ImportData.svelte";

  const base = "http://localhost:3280/api/v1/web/guest/";

  let url = base + action;
  let loading = true;
  let state = {};
  let message = "uploading...";

  async function start() {
    const res = await fetch(url);
    state = await res.json();
    console.log(state);
    loading = false;
  }

  async function save(data) {
    let u = base + "util/store";
    console.log("save", u, data);
    fetch(u, {
      method: "POST",
      body: JSON.stringify({"data":data}),
      headers: { "Content-Type": "application/json" }
    })
      .then(async function(res) {
        console.log(res)
        if(res.ok) {
          message = "OK";
        } else {
          message = res.statusText
        }
        loading = false;
      })
      .catch(function(err) {
        message = err.message;
        loading = false;
      });
  }

  onMount(start);
  formData.subscribe(value => {
    state = value;
    if ("data" in state) 
       save(state["data"]);
  });
</script>

<h2>Import</h2>
<div>
  {#if loading}
    <div>loading...</div>
  {:else if state.form}
    <ImportForm form={state.form} {url} />
  {:else if state.error}
    <div class="alert alert-danger" role="alert">Error: {state.error}</div>
    <div>
      <button type="button" class="btn btn-primary" on:click={start}>
        Restart
      </button>
    </div>
  {:else if state.data}
    <ImportData data={state.data} />
    <big><b>Import status: </b>{message}</big>
  {:else}
    <p>Cannot contact importer, please redeploy it.</p>
    <button type="button" class="btn btn-primary" on:click={start}>
      Retry
    </button>
  {/if}
</div>
