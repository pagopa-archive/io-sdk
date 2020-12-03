<script>

import { onMount } from 'svelte';
import { githubService } from './services/github.service.js';

let connectors = [];

const getConnectorsList = async () => {

    try{

        const response = await githubService.getIoGetawayConnectors();

        connectors = response.map( e => ({...e, checked: false}));
        

    } catch(e) {

        console.log(e);

    }

}

onMount(() => {
    getConnectorsList();
})

const handleCheckBox = ( connectorName, event ) => {

    const connectorIndex = connectors.findIndex( connector => connector.name === connectorName );

    if(connectorIndex !== -1) {

        connectors[connectorIndex].checked = event.target.checked;

    }

}

const onSubmit = () => {

    //TODO: Implement POST to submit selected connectors
    console.log(connectors);

}

</script>

<div>

    <h1> Connectors </h1>
    <form>
        {#each connectors as connector}

            <div class="form-check">

                <input
                    checked={connectors.find( c => c.name === connector.name && c.checked )} 
                    on:change={ e => handleCheckBox( connector.name, e )}
                    class="form-check-input"
                    id={"checkbox_"+connector.name}
                    type="checkbox"/>
                <label 
                    class="form-check-label"
                    for={"checkbox_"+connector.name}>
                    {connector.name}
                </label>

            </div>

        {/each}
        <button 
            type="button" 
            class="btn btn-primary" 
            on:click={onSubmit}>
            Submit
        </button>
    </form>

</div>