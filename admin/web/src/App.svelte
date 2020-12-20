<script>
 import { parseURL, navigate, url } from './url.js'
 import { onMount, onDestroy } from 'svelte'

 import Cookiebar from './Cookiebar.svelte'
 import Header from './Header.svelte'
 import Footer from './Footer.svelte'
 import Menu from "./Menu.svelte";
 import Home from "./Home.svelte";
 import About from "./About.svelte";
 import Devel from "./Devel.svelte";
 import Debug from "./Debug.svelte";
 import Send from "./Send.svelte";
 import Ship from "./Ship.svelte";
 import Import from "./Import.svelte";
 import Connectors from "./Connectors.svelte";
 
 let menu = ""
 let key = ""
 let api = ""

onMount(() => {
	[menu, key, api] = parseURL(location.href)
	navigate(menu, key)
})

onDestroy(url.subscribe((href) => {	
	[menu, key, api] = parseURL(href)
	console.log("menu=",menu, "key=", key, "api=", api)
}))

</script>

<main>
<div id="root">
	<Header/>
	<section class="row">
	    <Menu />
		<div class="col-9 col-md-offset-1 ">
		  <div class="mr-1 pt-1 pb-1">
			{#if menu == ""}
			  <Home/>
			{:else if menu=="about"}
			  <About/>
			{:else if menu=="ship"}
			  <Ship {api}/>
			{:else if menu=="send"}
			  <Send {api} {key}/>
			{:else if menu=="debug"}
			  <Debug {api} {key}/>
			{:else if menu=="import"}
			  <Import {api} action="/util/import"/>
			{:else if menu=="custom"}
			  <Import {api} action={key ? `/util/${key}` : "/iosdk/import"} />
			{:else if menu=="devel"}
			  <Devel {api} {key}/>
			{:else if menu=="connectors"}
			  <Connectors {api} {key} />
			{/if}
		  </div>
		</div>
	</section>	  
	<Cookiebar/>
	<Footer/>
</div>

</main>

