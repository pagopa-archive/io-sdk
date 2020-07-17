<script>
import { onMount } from "svelte";
import ImportPreview from "./ImportPreview.svelte";
import ImportPreviewModal from "./ImportPreviewModal.svelte";
import { getContext } from 'svelte';
const { open } = getContext('simple-modal');
export let data;
export let previewMode;
import { createEventDispatcher } from 'svelte';
const dispatch = createEventDispatcher();

const onConfirm = () => {
    console.log('popup confirmed');
    dispatch('confirm');
}

onMount(() => {
    if(previewMode === 'modal')
        open(ImportPreviewModal, 
            {data, onConfirm},
            {
    		    closeOnEsc: false,
    		    closeOnOuterClick: true,
			});
}); 


</script>

{#if previewMode !== 'modal'}
    <ImportPreview data={data} />
{/if}