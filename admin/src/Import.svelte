<script>
  import { formData } from "./store";
  import { onMount } from "svelte";
  import ImportForm from "./ImportForm.svelte";
  import ImportData from "./ImportData.svelte";
  import Modal from 'svelte-simple-modal';
  import DispatchImportPreview from "./DispatchImportPreview.svelte";
  export let action;

  const base = "http://localhost:3280/api/v1/web/guest/";

  let url = base + action;
  let loading = true;
  let state = {};
  let message = "uploading...";
  let isPreview = false;

  async function start() {
    const res = await fetch(url);
    state = await res.json();
    console.log(state);
    loading = false;
  }

  function confirm() {
    console.log('confirm');
    loading = true;
    isPreview = false;
		save(state.data);
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
    if ("data" in state && !isPreview) 
       save(state["data"]);
  });
</script>

<h2>Import</h2>
<div>
  {#if loading}
    <div>loading...</div>
  {:else if state.form}
    <ImportForm form={state.form} {url} />
    <input type=checkbox bind:checked={isPreview}> Mostra anteprima
  {:else if state.error}
    <div class="alert alert-danger" role="alert">Error: {state.error}</div>
    <div>
      <button type="button" class="btn btn-primary" on:click={start}>
        Restart
      </button>
    </div>
  {:else if state.data}
    {#if !isPreview}
      <ImportData data={state.data} />
      <big><b>Import status: </b>{message}</big>
    {:else}
      <Modal>
        <DispatchImportPreview previewMode="modal" data={state.data} on:confirm={confirm} />
      </Modal>
      {#if loading}
        <big><b>Import status: </b>{message}</big>
      {/if}
    {/if}
  {:else}
    <p>Cannot contact importer, please redeploy it.</p>
    <button type="button" class="btn btn-primary" on:click={start}>
      Retry
    </button>
  {/if}
</div>
