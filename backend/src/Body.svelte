<script>
  import { Router, Route } from "svelte-routing";
  import Home from "./Home.svelte";
  import Menu from "./Menu.svelte";
  import About from "./About.svelte";
  import Devel from "./Devel.svelte";
  import Send from "./Send.svelte";
</script>

<section class="row">
  <Menu />
  <div class="col-9">
    <div class="mr-4 pt-5 pb-5">
      <Router>
        <Route path="/">
          <Home />
        </Route>
        <Route path="send">
          <Send/>
        </Route>
        <Route path="devel">
          <Devel />
        </Route>
        <Route path="about">
          <About />
        </Route>
      </Router>
    </div>
  </div>
</section>
