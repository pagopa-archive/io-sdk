<script>
  import { onDestroy, onMount } from 'svelte'
  import { formData } from "./store";
  import ImportForm from "./ImportForm.svelte";
  import ImportData from "./ImportData.svelte";

  export let api;
  export let action;

  let storeURL = api + "/util/store"
  let importURL = api + action;

  let message = "uploading...";
  let state

  onMount( () => { 
    state = start() 
  })

      /**
     * Called only when api or action prop is updated by the parent
    */
    $: {
        importURL = api + action;
        state = start();
    }

  async function start() {
    formData.set({})
    return fetch(importURL)
      .then((res) => res.json())
      .catch((err) => { return { "error": err }})
  }
  
  async function save(data) {
      await fetch(storeURL, {
        method: "POST",
        body: JSON.stringify({"data":data}),
        headers: { "Content-Type": "application/json" }
      }).then(async function(res) {
          console.log(res)
          if(res.ok) {
            message = "OK";
          } else {
            message = "ERROR: "+res.statusText+" "+res.status
          }
      }).catch(function(err) {
          message = "ERROR: "+err.message;
      });
}


let unsubscribe = formData.subscribe(value => {
    console.log("subscribe")
    state = value;
    if(state.data) {
     save(state.data)
    }
});

onDestroy(unsubscribe)

</script>

<h2>Import</h2>
<div>
  {#await state}
    <div>loading...</div>
  {:then currState}
    {#if currState.form}
    <ImportForm form={currState.form} {api} {action} />
    {:else if currState.data}
    <ImportData data={state.data} />
    <big><b>Import status: </b>{message}</big>
    {/if}
  {:catch err}
   <div class="alert alert-danger" role="alert">Error: {err.error}</div>
   <div>
      <button type="button" class="btn btn-primary" on:click={start}>
        Restart
      </button>
   </div>
  {/await}
</div>
