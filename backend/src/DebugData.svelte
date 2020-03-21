<script>
  export let key = "";
  import { Link } from "svelte-routing";
  import { onMount } from "svelte";

  const url = "http://localhost:3280/api/v1/web/guest/util/cache";
  let entry = window.atob(key);
  
  let state = {};

  async function cache(action, entry) {
    let u = url + "?"+action+"="+encodeURI(entry)
    console.log(u)
    const res = await fetch(u);
    state[action] = await res.json();
  }
  
  onMount(() => cache("get", entry));
  let clean = () => {cache("del", entry)}
</script>

<h1>Cache Data</h1>

{#if state.get}
<big>Key:</big>
<p>{entry}</p>
<big>Value:</big>
<pre>{JSON.stringify(state.get, null, 2)}</pre>
<br>
<button type="button" class="btn btn-primary" on:click={clean}>Clean Key</button>
{:else if state.del}
<big>Result</big>
<pre>{JSON.stringify(state.data, null, 2)}</pre>
{:else if state.error}
<h1>Error</h1>
<p>{state.error}</p>
{/if}
<br>
<Link to="debug">Back</Link>
