<script>
  import { Link } from "svelte-routing";
  import { reload } from "./store";
  import { onMount } from "svelte";

  const url = "http://localhost:3280/api/v1/web/guest/util/cache?scan=*";
  let state = {};

  async function start() {
    fetch(url).then(async (res) => {
        let data = await res.json()
        if(data && "scan" in data) {
          state = { "list": data["scan"]}
        } else {
          state = { "error": "cannot retrieve data"}
        }
        console.log(state)
    }).catch(err => {
        state = { "error": err}
    }) 
  }
  onMount(start);
  reload.subscribe(() => {
    console.log("reload!!!")
    start()
  })
</script>

<h1>Cache</h1>
<div>
 {#if state.list}
    <ul>
     {#each state.list as key}
       <li><Link to="debug/{window.btoa(key)}">{key}</Link></li>
     {/each}
    </ul>
    <p>Elements in cache: #{state.list.length}</p>
  {:else if state.error}
     <div class="alert alert-danger" role="alert">
        Error: {state.error}
      </div>
 {:else}
   <p>Loading...</p>
   <button type="button" class="btn btn-primary" on:click={start}>Reload</button>
 {/if}
</div>
