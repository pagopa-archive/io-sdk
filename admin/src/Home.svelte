<script>
  import { onMount } from 'svelte'
  let error = "error"

async function test(url) {
    //try {
      return fetch(url, {
        headers: { 'Accept': 'application/json'}
      }).then((res) => {
        console.log("then")
        console.log(res)
        error = res.statusText
        return res.json()
      }).catch((err) => {
        console.log("catch")
        console.log(err)
        error = err
      })
    //} catch(ex) { error = ex }
}

async function start() {
  let r = await test("https://httpstat.us/429")
  console.log(r)
}

onMount(start)


</script>

<div class="it-hero-wrapper">
  <div class="container-fluid">
    <div class="row">
        <div class="col-9">
          <div class="it-hero-text-wrapper bg-dark">
              <span class="it-category">Welcome</span>
              <h1 class="no_toc">IO-SDK</h1>
              <p class="d-none d-lg-block">Welcome to the development kit to easily build IO importers.</p>
              <br>
              <p>ERROR: {error}</p>
          </div>
        </div>
    </div>
  </div>
</div>
