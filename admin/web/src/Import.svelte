<script>
  export let action;

  import { onDestroy, onMount } from 'svelte'
  import { formData } from "./store";
  import ImportForm from "./ImportForm.svelte";
  import ImportData from "./ImportData.svelte";

  const base = "http://localhost:3280/api/v1/web/guest/";
  const url = base + action;

  let message = "uploading...";
  let state
 
  async function start() {
    formData.set({})
    return fetch(url)
      .then((res) => res.json())
      .catch((err) => { return { "error": err }})
  }
  
  async function save(data) {
      await fetch(base + "util/store", {
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

onMount( () => { 
  state = start() 
})

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
    <ImportForm form={currState.form} {url} />
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
