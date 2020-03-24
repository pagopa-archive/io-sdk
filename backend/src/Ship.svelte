<script>
  import { Link } from "svelte-routing";
  import { formData } from "./store";
  import { onMount } from "svelte";

  const url =
    "http://localhost:3280/api/v1/web/guest/util/cache?scan=message:*";

  let state = {};


  let selection = [];
  let sent = {}

  async function start() {
    fetch(url)
      .then(async res => {
        let data = await res.json();
        if (data && "scan" in data) {
          state = { list: data["scan"] };
        } else {
          state = { error: "cannot retrieve data" };
        }
        console.log(state);
      })
      .catch(err => {
        state = { error: err };
      });
  }
  onMount(start);

  function invertSelection() {
    console.log("invert")
    let newSelection = []
    for(let i of state.list)
      if(selection.indexOf(i) == -1)
        newSelection.push(i)
    selection = newSelection
    console.log(selection)
  }

  function sendSelected() {
    for(let i of selection) 
      sent[i] = true
  }
</script>

<h1>Imported Messages</h1>
<div>
  {#if state.list}
    {#if state.list.length > 0}
      <div class="form-group">
        <button type="button" class="btn btn-secondary" on:click={() => selection = state.list}>Select All</button>
        <button type="button" class="btn btn-secondary" on:click={() => selection = []}>Deselect All</button>
        <button type="button" class="btn btn-secondary" on:click={invertSelection}>
          Invert Selection
        </button>
      </div>
      <ul>
        {#each state.list as key}
          <div class="form-check">
            {#if !sent[key]}
              <input
                id={window.btoa(key)}
                type="checkbox"
                bind:group={selection}
                value={key} />
              <label for={window.btoa(key)}>
                <Link to="send/{window.btoa(key)}">{key}</Link>
              </label>
            {:else}
             <del>{key}</del> sent!
            {/if}
          </div>
        {/each}
      </ul>
      <div class="form-group">
        <big>
          Total Messages: {state.list.length} - Selected Messages: {selection.length}
        </big>
      </div>

      <div class="form-group">
        <button type="button" class="btn btn-primary" on:click={sendSelected}>
          Send Selected Messages
        </button>
      </div>
    {:else}
      <p>
        <big>No messages to send.</big>
      </p>
      <p />
      <p>
        Please
        <Link to="/import">import some messages</Link>
        .
      </p>
    {/if}
  {:else if state.error}
    <div class="alert alert-danger" role="alert">Error: {state.error}</div>
  {:else}
    <p>Loading...</p>
    <button type="button" class="btn btn-primary" on:click={start}>
      Reload
    </button>
  {/if}
</div>
