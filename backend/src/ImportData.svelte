<script>
  import { Link } from "svelte-routing";
  export let data = [];
  console.log(data);
</script>

<div>
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
</div>
