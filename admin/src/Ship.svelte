<script>
  import { Link } from "svelte-routing";
  import { formData } from "./store";
  import { onMount } from "svelte";

  let action = "util/send";
  const baseUrl = "http://localhost:3280/api/v1/web/guest/";
  const scanUrl = baseUrl + "util/cache?scan=message:*";
  const getUrl = baseUrl + "util/cache?get=";
  const delUrl = baseUrl + "util/cache?del=";

  let state = {};
  let selection = [];
  let sent = {};

  async function start() {
    fetch(scanUrl)
      .then(async res => {
        let data = await res.json();
        if (data && "scan" in data) {
          state = { list: data["scan"] };
          sent = {}
          selection = []
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
    console.log("invert");
    let newSelection = [];
    for (let i of state.list)
      if (selection.indexOf(i) == -1) newSelection.push(i);
    selection = newSelection;
    console.log(selection);
  }

  function findNext() {
    for (let i of selection) {
      if (sent[i]) continue;
      return i;
    }
    return undefined;
  }

  async function cache(url, key) {
    let u = url + encodeURI(key);
    //console.log("cache", u);
    let res = await fetch(u);
    //console.log(key, "res", res)
    if(res.ok) { 
     let data = await res.json()
     //console.log(key, "data", data)
     return data
    }
    return { "error": res.statusText }
  }

  async function send(data) {
    let u = baseUrl + action
    console.log("send", u, data)
    let res = await fetch(u, {
      method: "POST",
      body: JSON.stringify(data),
      headers: { "Content-Type": "application/json" }
    });
    if (res.ok) {
      return await res.json();
    }
    return { "error": res.statusText };
  }

  async function sendSelected() {
    let key = findNext()
    if (!key) 
      return;
    sent[key] = "sending..."
    let msg = await cache(getUrl, key);
    //console.log("cache=", msg);
    if (key in msg) {
      let res = await send(msg[key]);
      console.log("send=", res)
      if (res["id"]) {
        await cache(delUrl, key);
        sent[key] = "OK: " + res.id;
      } else if (res["detail"]) {
        sent[key] = "ERROR: " + res["detail"];
      } else if (res.error) {
        sent[key] = "ERROR: " + res.error;
      } else {
        send[key] = "unknow error, check logs";
        console.log(key, res);
      }
    } else {
      send[key] = key + " not found in cache, check logs";
    }
    sendSelected();
  }
</script>

<h1>Imported Messages</h1>
<div>
  {#if state.list}
    {#if state.list.length > 0}
      <div class="form-group">
        <button
          type="button"
          class="btn btn-secondary"
          on:click={() => (selection = state.list)}>
          Select All
        </button>
        <button
          type="button"
          class="btn btn-secondary"
          on:click={() => (selection = [])}>
          Deselect All
        </button>
        <button
          type="button"
          class="btn btn-secondary"
          on:click={invertSelection}>
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
              <del>{key}</del>
              {sent[key]}
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
        <div class="bootstrap-select-wrapper">
          <label>Endpoint</label>
          <select bind:value={action} title="Scegli una opzione">
            <option value="util/send">Development (Local)</option>
            <option value="iosdk/send">Production</option>
          </select>
        </div>
      </div>
      <div class="form-group">
        <button type="button" class="btn btn-primary" on:click={sendSelected}>
          Send Selected Messages
        </button>
        <button type="button" class="btn btn-primary" on:click={start}>
         Restart
       </button>
      </div>
    {:else}
      <p>
        <big>No messages to send.</big>
      </p>
      <p />
      <p>
        Please, <Link to="/import">import some messages</Link>.
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
