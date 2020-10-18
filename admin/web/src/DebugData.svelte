<script>
  import { onMount } from "svelte";
  import { navigate } from './url.js'

  export let key = "";
  export let api = "";

  let state = {};

  const url = api + "/util/cache";
  const entry = window.atob(key);

  async function cache(action, entry) {
    let u = url + "?" + action + "=" + encodeURI(entry);
    console.log(u);
    const res = await fetch(u);
    state[action] = await res.json();
  }

  onMount(() => cache("get", entry));
  function back() {
    navigate("debug", "")
  }

  function clean() {
    if(confirm("Are you sure?")) {
      cache("del", entry);
      back();     
    }
  }
</script>

<h1>Cache Data</h1>

{#if state.get}
  <big>Key:.</big>
  <p>{entry}</p>
  <big>Value:</big>
  <pre>{JSON.stringify(state.get, null, 2)}</pre>
  <br />
    <button type="button" class="btn btn-primary" on:click={clean}>
      Clean Key
    </button>
 {:else if state.del}
  <big>Result</big>
  <pre>{JSON.stringify(state.data, null, 2)}</pre>
{:else if state.error}
  <h1>Error</h1>
  <p>{state.error}</p>
{/if}
<br>
<button type="button" class="btn btn-primary" on:click={back}>
  Back to Debugging
</button>

