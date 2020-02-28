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
              <MenuItem icon="far fa-envelope" description="Send Message"  to="send"/>
              <MenuItem icon="fas fa-file-code" description="Development" to="ide" />
              <MenuItem icon="fas fa-info-circle" description="About" to="info"/>
            </Router>
          </ul>
        </div>
      </aside>
    </div>
  </div>
</div>
