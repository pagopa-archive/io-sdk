<script>
  import { Router, Route } from "svelte-routing";
  import Home from "./Home.svelte";
  import Menu from "./Menu.svelte";
  import About from "./About.svelte";
  import Devel from "./Devel.svelte";
  import Send from "./Send.svelte";
  import Import from "./Import.svelte";
  export let url = ""
</script>

<section class="row">
  <Menu />
  <div class="col-9 col-md-offset-1 ">
    <div class="mr-1 pt-1 pb-1" >
      <Router url="{url}">
        <Route path="/" component="{Home}"/>
        <Route path="send" component="{Send}"/>
        <Route path="import" component="{Import}"/>
        <Route path="devel" component="{Devel}"/>
        <Route path="about" component="{About}"/>
      </Router>
    </div>
  </div>
</section>
