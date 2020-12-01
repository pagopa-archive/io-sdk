<script>

import { onMount } from 'svelte';
import { githubService } from './services/github.service.js';

let release = null;

const getLatestRelease = async () => {

    try{

        release = await githubService.getIoGetawayConnectorLatestRelease();
        console.log(release.assets);

    } catch(e) {

        console.log(e);

    }

}

onMount(() => {
    getLatestRelease();
})

</script>

<div>

    <h1> Connectors </h1>
    {#if release !== null }
        {#each release.assets as asset}
            <li>{asset.name}</li>
        {/each}
    {/if}

</div>