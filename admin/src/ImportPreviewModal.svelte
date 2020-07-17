<script>
  import { Link } from "svelte-routing";
  export let data = [];
  export let onConfirm = () => {};
  import { createEventDispatcher } from 'svelte';
  const dispatch = createEventDispatcher();
  console.log(data);

	function _onConfirm() {
		onConfirm();
		close();
  }
    
</script>

<div>
  <h3>Anteprima importazione</h3>
  <table class="table">
    <thead>
      <tr>
        <th scope="col">FiscalCode</th>
        <th scope="col">Subject</th>
        <th scope="col">Message</th>
      </tr>
    </thead>
    <tbody>
      {#each data as rec}
        <tr>
          <td>{rec.fiscal_code}</td>
          <td>{rec.subject}</td>
          <td>{rec.markdown}</td>
        </tr>
      {:else}
        <tr>
          <td colspan="3">Error: no data</td>
        </tr>
      {/each}
    </tbody>
  </table>
    <button type="button" class="btn btn-primary" on:click={_onConfirm}>
        Conferma importazione
    </button>
</div>
