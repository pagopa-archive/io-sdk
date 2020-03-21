<script>
  import { Link } from "svelte-routing";
  import { formData } from "./store";
  import { onMount } from "svelte";

  const url = "http://localhost:3280/api/v1/web/guest/util/cache?scan=*";
  let state = {};

  async function start() {
    let res = await fetch(url);
    res = await res.json();
    if("scan" in res) {
        state = { "list": res["scan"]}
    }
  }
  onMount(start);
</script>

<h1>Cache</h1>
<div>
 {#if state.list}
    <ul>
     {#each state.list as key}
       <li><Link to="debug/{window.btoa(key)}">{key}</Link></li>
     {/each}
    </ul>
 {:else}
   <p>Loading...</p>
   <button type="button" class="btn btn-primary" on:click={start}>Reload</button>
 {/if}
</div>
