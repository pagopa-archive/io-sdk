<script>
  export let url = ""
  import { Router, Route } from "svelte-routing";
  import Home from "./Home.svelte";
  import Menu from "./Menu.svelte";
  import About from "./About.svelte";
  import Devel from "./Devel.svelte";
  import Debug from "./Debug.svelte";
  import DebugData from "./DebugData.svelte";
  import Send from "./Send.svelte";
  import Ship from "./Ship.svelte";
  import Import from "./Import.svelte";
</script>

<section class="row">
  <Menu />
  <div class="col-9 col-md-offset-1 ">
    <div class="mr-1 pt-1 pb-1" >
      <Router url="{url}">
        <Route path="/" component="{Home}"/>
        <Route path="import" component="{Import}"/>
        <Route path="ship" component="{Ship}"/>
        <Route path="send" component="{Send}"/>
        <Route path="devel" component="{Devel}"/>
        <Route path="debug" component="{Debug}"/>
        <Route path="debug/:key" let:params><DebugData key="{params.key}"/></Route>
        <Route path="about" component="{About}"/>
      </Router>
    </div>
  </div>
</section>
