<script>
  import { Link } from "svelte-routing";
  import { reload } from "./store";
  import { onMount } from "svelte";

  const cacheUrl = "http://localhost:3280/api/v1/web/guest/util/cache?";
  const scanUrl = cacheUrl + "scan="
  const cleanUrl = cacheUrl + "clean="

  let pattern = "*";
  let state = {};

  async function scan() {
    let res = await fetch(scanUrl + pattern);
    if (!res.ok) return { error: res.statusText };
    return await res.json();
  }

  async function start() {
    let res = await fetch(scanUrl + pattern);
    if (!res.ok) {
      state.error = res.statusText;
      return;
    }
    let data = await res.json();
    if (data && data.scan) {
      state.list = data.scan.sort();
    } else {
      state.error = "cannot retrieve data";
    }
  }
  
  function search(what) {
    pattern = what
    start()
  }

  async function clean() {
    if(!confirm("Are you sure you want to remove?"))
      return;
    let res = await fetch(cleanUrl + pattern);
    if (!res.ok) return { error: res.statusText };
    search("*")
  }

  onMount(start);
  reload.subscribe(() => {
    start();
  });
</script>

<h1>Cache</h1>
<div class="form-group">
  <div class="input-group">
    <div class="input-group-prepend">
      <div class="input-group-text">
        <i class="fas fa-search"/>
      </div>
    </div>
    <input
      type="text"
      class="form-control"
      id="search"
      name="search"
      bind:value={pattern} />
  </div>
  <small
    id="formGroupExampleInputWithHelpDescription"
    class="form-text text-muted">
    Search pattern for keys in cache
  </small>
</div>
<div class="form-group">
  <button class="btn btn-primary" type="button" id="button-3" on:click={() => search("*")}>Show All</button>
  <button class="btn btn-primary" type="button" id="button-3" on:click={() => search("message:*")}>Imported Messages</button>
  <button class="btn btn-primary" type="button" id="button-3" on:click={() => search("send:*")}>Sent Messages</button>
  <button class="btn btn-primary" type="button" id="button-3" on:click={start}>Search</button>
</div>
<div>
  {#if state.list}
    <ul>
      {#each state.list as key}
        <li>
          <Link to="debug/{window.btoa(key)}">{key}</Link>
        </li>
      {/each}
    </ul>
    <p>Elements Found: #{state.list.length}</p>
    <button type="button" class="btn btn-primary" on:click={clean}>Clean Current Keys</button>
  {:else if state.error}
    <div class="alert alert-danger" role="alert">Error: {state.error}</div>
  {:else}
    <p>Loading...</p>
    <button type="button" class="btn btn-primary" on:click={start}>
      Reload
    </button>
  {/if}
</div>
