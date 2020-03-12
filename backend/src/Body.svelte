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
  <div class="col-9 col-md-offset-1 ">
    <div class="mr-1 pt-1 pb-1" >
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
