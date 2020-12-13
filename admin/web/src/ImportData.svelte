<script>

export let data = [];

</script>

<div>
  <table class="table">
    <thead>
      <tr>
        <th scope="col">FiscalCode</th>
        <th scope="col">Subject</th>
        <th scope="col">Message</th>
        <th scope="col">Amount</th>
      </tr>
    </thead>
    <tbody>
      {#each data as rec}
        <tr>
          <td>{rec.fiscal_code}</td>
          <td>{rec.subject}</td>
          <td>{rec.markdown}</td>
          <td>{rec.amount || ''}</td>
        </tr>
      {:else}
        <tr>
          <td colspan="4">Error: no data</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
