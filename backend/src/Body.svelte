<script>
  import { Router, Route } from "svelte-routing";
  import Lorem from "./Lorem.svelte";
  import Card from "./Card.svelte";
  import Form from "./Form.svelte";
  import Profilo from "./Profilo.svelte";
  import Menu from "./Menu.svelte";
</script>

<section class="row">
  <Menu />
  <Router>
    <Route path="/">
      <Profilo />
    </Route>
    <Route path="send">
      <Form />
    </Route>
    <Route path="ide">
      <Card />
    </Route>
    <Route path="info">
      <Lorem />
    </Route>
  </Router>
</section>
