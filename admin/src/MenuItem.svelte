<script>
  import { Link } from "svelte-routing";

  export let icon = "fas fa-home";
  export let description = "Home";
  export let to = "/";
  export let url = "";

  function getProps(args) {
    return {
      class: "large list-item"
    };
  }
  function visit(event) {
    event.preventDefault()
    window.open(url, to)
  }
</script>

<li>
  {#if url == ''}
    <Link {to} {getProps}>
      <i class={icon} />
      {description}
    </Link>
  {:else}
    <a href={url} class="large list-item" rel="noopener noreferrer" on:click={visit}>
      <i class={icon} />
      {description}
    </a>
  {/if}
</li>
