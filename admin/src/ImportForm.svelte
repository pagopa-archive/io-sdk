<script>
  let loading = false;
  let uploadUrl = "http://localhost:3280/api/v1/web/guest/util/upload";

  import { formData } from "./store";
  export let form = {};
  export let url;

  async function parseForm() {
    let submit = await fetch(uploadUrl, {
      method: "POST",
      body: new FormData(document["_theForm_"])
    })
    let data = await submit.json()
    return data
  }

  async function importer() {
    // gather data from form
    let data = await parseForm()
    // perform import
    if(data) {
      loading = true;
      fetch(url, {
        method: "POST",
        body: JSON.stringify(data),
        headers: {
          "Content-Type": "application/json"
        }
      }).then(
        async function(res) {
          console.log(res);
          if(res.status===204){
            let error = "data format not valid";
            formData.set({ error });
            loading = false;
          }
          else{
            let data = await res.json();
            formData.set(data);
            loading = false;
          }
        },
        function(err) {
          error = err.message;
          formData.set({ error: err.message });
          loading = false;
        }
      );
    } 
    else 
      error = "cannot parse form";
  }


</script>

<form name="_theForm_" action={uploadUrl} method="post" enctype="multipart/form-data">
  {#each form as field}
    {#if field.type == 'message'}
      <div class="callout">
        <div class="callout-title">{field.name}</div>
        <p>{field.description}</p>
      </div>
      <br>
    {:else if field.type == 'string'}
      <div class="form-group">
        <input
          type="text"
          class="form-control"
          id={field.name}
          name={field.name}
          bind:value={field.value} />
        <label class="active" for={field.name}>{field.description}</label>
      </div>
    {:else if field.type == 'password'}
      <div class="form-group">
        <label class="active" for={field.name}>{field.description}</label>
        <input
          type="password"
          class="form-control"
          id={field.name}
          name={field.name}
          bind:value={field.value} />
      </div>
    {:else if field.type == 'upload'}
      <div class="form-group">
        <input
          type="file"
          class="form-control"
          id={field.name}
          name={field.name}
          bind:files={field.value} />
        <label class="active" for={field.name}>{field.description}</label>
      </div>
    {:else if field.type == 'textarea'}
      <div class="form-group">
        <textarea 
        id={field.name}
        name={field.name}
        rows="3" 
        bind:value={field.value} />
        <label for={field.name}>{field.description}</label>
      </div>
      <br />
    {:else}
      <div class="alert alert-danger" role="alert">
        Error: Unknown or empty type {field.type}
      </div>
    {/if}
  {:else}
    <div class="alert alert-danger" role="alert">Error: no form data</div>
  {/each}
  <div>
    {#if loading}
      <div>loading...</div>
    {:else}
      <button type="button" class="btn btn-primary" on:click={importer}>
        Import
      </button>
    {/if}
  </div>
</form>
