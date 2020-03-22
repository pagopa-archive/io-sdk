<script>

  let base = "http://localhost:3280/api/v1/web/guest/"

  let state = {}

  let action = "util/send"
  let data = {
     "dest": "",
     "subject": "",
     "markdown": ""
  }
  function submitForm1() {
    console.log(action, data)
  }

  function submitForm() {
    fetch(base+action, {
       method: "POST",
       body: JSON.stringify(data),
       headers: {
         "Content-Type": "application/json"
       }
    }).then(async function(res){
      state.result = await res.text()
    }).catch(function(err) {
       state.error = err.message 
    })
  }
  function resetForm() {
    result=""
    error=""
    for(key in data) {
      data[key] = ""
    }
  }
</script>
<h2>Send a Single Message</h2>
<br>
{#if state.result}
<div class="alert alert-success" role="alert">
  <pre>{state.result}</pre>
</div>
<div>
  <button type="button" class="btn btn-primary" 
          on:click={resetForm}>Nuovo Messaggio</button>
</div>
{:else if state.error}
<div class="alert alert-danger" role="alert">
  {state.error}
</div>
<div>
    <button type="button" class="btn btn-primary" 
            on:click={resetForm}>Nuovo Messaggio</button>
</div>
{:else}
  <div>
    <div class="form-group">
      <input type="text" class="form-control" 
             id="codFiscDest" bind:value={data.dest} />
      <label for="codFiscDest">Codice Fiscale Destinatario</label>
    </div>
    <div class="form-group">
      <input type="text" class="form-control" 
             id="messageSubject" bind:value={data.subject}>
      <label for="messageSubject">Soggetto del messaggio</label>
    </div>
    <div class="form-group">
      <textarea id="message" rows="3" 
       bind:value={data.markdown}></textarea>
      <label for="message">Markdown del messaggio</label>
    </div>
    <div class="form-group">
    
    <div class="bootstrap-select-wrapper">
      <label>Endpoint</label>
      <select value={action} title="Scegli una opzione">
        <option value="util/send">Development (Local)</option>
        <option value="iosdk/send">Production</option>
      </select>
    </div>
  </div>
    <div class="form-group">
    <button type="button" class="btn btn-primary" 
            on:click={submitForm}>Invia</button>
    </div>
  </div>
{/if}

