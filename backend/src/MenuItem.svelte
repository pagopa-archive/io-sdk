<script>
  import { Link } from "svelte-routing";
 
  export let icon = "fas fa-home";
  export let description = "Home";
  export let to = "/";

  function getProps(args) {
      return { 
          class: "large list-item" 
      }
  }
</script>

<li>
    <Link {to} getProps="{getProps}">
      <i class={icon} /> {description}
    </Link>
</li>
