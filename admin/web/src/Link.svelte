<script>
  import { navigate } from './url.js'
  export let icon = "";
  export let description = "TODO";
  export let to = "";
  export let target = "";
  export let key = "";
  
  function open(event) {
    event.preventDefault()
    window.open(target, "editor", )
  }
  function jump(event) {
    event.preventDefault()
    navigate(to,key)
  }
</script>

{#if target == ''}
<a href="#{to}" class="large list-item" on:click={jump}>
 {#if icon}<i class={icon} />{/if}
 {description}
</a>
{:else}
<a href={target} class="large list-item" rel="noopener noreferrer" on:click={open}>
 {#if icon}<i class={icon} />{/if}
 {description}
</a>
{/if}
