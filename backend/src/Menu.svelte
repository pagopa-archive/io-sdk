<script>
  import MenuItem from "./MenuItem.svelte";
  import { Router } from "svelte-routing";
</script>

<div class="col-3">
  <div class="container my-4">
    <div class="ml-4 pt-5 pb-5">
      <aside>
        <div class="link-list-wrapper">
          <ul class="link-list">
            <Router>
              <MenuItem />
              <MenuItem icon="fas fa-file-import" description="Import Messages" to="import"/>
              <MenuItem icon="fas fa-shipping-fast" description="Send Messages" to="ship"/>
              <MenuItem icon="far fa-envelope" description="Send One Message" to="send"/>
              <MenuItem icon="fas fa-file-code" description="Development" to="devel" />
              <MenuItem icon="fas fa-bug" description="Debugging" to="debug"/>
              <MenuItem icon="fas fa-info-circle" description="About" to="about"/>
            </Router>
          </ul>
        </div>
      </aside>
    </div>
  </div>
</div>
