<script>
  import Link from "./Link.svelte";
</script>

<div class="col-3">
  <div class="container my-4">
    <div class="ml-4 pt-5 pb-5">
      <aside>
        <div class="link-list-wrapper">
          <ul class="link-list">
              <li><Link icon="fas fa-file-import" description="Import URL" to="import"/></li>
              <li><Link icon="fa fa-wrench" description="Custom Import" to="custom"/></li>
              <li><Link icon="fas fa-shipping-fast" description="Send Messages" to="ship"/></li>
              <li><Link icon="far fa-envelope" description="Single Message" to="send"/></li>
              <li><Link icon="fas fa-bug" description="Debugging" to="debug"/></li>
              {#if location.hostname == "localhost"}
              <li><Link icon="fas fa-file-code" description="Development" target="http://localhost:3000"/></li>
              {/if}
              <li><Link icon="fas fa-info-circle" description="About" to="about"/></li>
              <li>
          </ul>
        </div>
      </aside>
    </div>
  </div>
</div>
