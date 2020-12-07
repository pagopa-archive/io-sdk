<script>

import { onMount } from 'svelte';
import { githubService } from './services/github.service.js';
import { ioSDKService } from './services/iosdk.service';

let connectors = [];
let loading = false;

const init = async () => {

    await getConnectorsList(
        await getCustomConnectors()
    );

}

const getConnectorsList = async ( custom_connectors = []) => {

    try{

        const response = await githubService.getIoGetawayConnectors();

        connectors = response.map( e => ({...e, checked: custom_connectors.includes(e.name)}));
        

    } catch(e) {

        console.log(e);

    }

}

const getCustomConnectors = async () => {

    try{

        const response = await ioSDKService.getCustomConnectors();

        return response?.details?.connectors || []

    } catch(e) {

        console.log(e);

    }

}

onMount( async () => {
    await init()
})

const handleCheckBox = ( connectorName, event ) => {

    const connectorIndex = connectors.findIndex( connector => connector.name === connectorName );

    if(connectorIndex !== -1) {

        connectors[connectorIndex].checked = event.target.checked;

    }

}

const onSubmit = async () => {

    loading = true;
    try{

        const response = await ioSDKService.createConnectors( 
            connectors.filter( connector => connector.checked).map( connector => connector.name)
        );

    } catch( e ) {

        console.error( e )

    }

    init();

    loading = false
    
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
            disabled={loading}
            type="button" 
            class="btn btn-primary" 
            on:click={onSubmit}>
            Submit
        </button>
    </form>

</div>