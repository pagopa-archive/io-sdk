<script>
  import { formData } from "./store";
  export let form = {};
  export let url;

  function importer() {
    // gather data from form
    let data = {};
    for (const field of form) {
      if (field.value) data[field.name] = field.value;
    }
    //console.log(data)
    // perform import
    fetch(url, {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Content-Type": "application/json"
      }
    }).then(
      async function(res) {
        let data = await res.json()
        formData.set(data);
      },
      function(err) {
        error = err.message;
        formData.set({ error: err.message });
      }
    );
  }
</script>

<div>
  {#each form as field}
    {#if field.type == 'message'}
      <div class="callout">
        <div class="callout-title">{field.name}</div>
        <p>{field.description}</p>
      </div>
    {:else if field.type == 'string'}
      <div class="form-group">
        <input
          type="text"
          class="form-control"
          id={field.name}
          bind:value={field.value} />
        <label for={field.name}>{field.description}</label>
      </div>
    {:else if field.type == 'password'}
      <div class="form-group">
        <input
          type="password"
          class="form-control"
          id={field.name}
          bind:value={field.value} />
        <label for={field.name}>{field.description}</label>
      </div>
    {:else if field.type == 'textarea'}
      <div class="form-group">
        <textarea id={field.name} rows="3" bind:value={field.value} />
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
    <button type="button" class="btn btn-primary" on:click={importer}>
      Import
    </button>
  </div>
</div>
