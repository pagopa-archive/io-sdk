<script>
  export let data = [];
  console.log(data);
</script>

<div>
  <table class="table">
    <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">FiscalCode</th>
        <th scope="col">Subject</th>
        <th scope="col">Message</th>
        <th scope="col">Amount</th>
      </tr>
    </thead>
    <tbody>
      {#each data as rec}
        <tr>
          <th>{rec.notice_number}</th>
          <td>{rec.fiscal_code}</td>
          <td>{rec.subject}</td>
          <td>{rec.markdown}</td>
          <td>{rec.amount}</td>
        </tr>
      {:else}
        <tr>
          <td colspan="5">Error: no data</td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
