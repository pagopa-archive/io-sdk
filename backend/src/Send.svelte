<script>

  let url = "http://localhost:3280/api/v1/web/guest/backend/send"

  let result = ""
  let error = ""

  let data = {
     "CodFiscDest": "",
     "ApiKeyIO": "",
     "Message": ""
  }

  function submitForm() {
    fetch(url, {
       method: "POST",
       body: JSON.stringify(data),
       headers: {
         "Content-Type": "application/json"
       }
    }).then(async function(res){
      result = await res.text()
    }, function(err) {
       error = err.message 
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
<h2>Send Message</h2>
<br>

{#if result == ""}
  <div>
    <div class="form-group">
      <input type="text" class="form-control" id="codFiscDest" bind:value={data.CodFiscDest} />
      <label for="codFiscDest">Codice Fiscale Destinatario</label>
    </div>
    <div class="form-group">
      <input type="text" class="form-control" id="messageSubject" bind:value={data.MessageSuject}>
      <label for="messageSubject">Soggetto del messaggio</label>
    </div>
    <div class="form-group">
      <textarea id="message" rows="3" bind:value={data.Message}></textarea>
      <label for="message">Markdown del messaggio</label>
    </div>
    <div class="form-group">
    
    <div class="bootstrap-select-wrapper">
      <label>Endpoint</label>
      <select title="Scegli una opzione">
        <option value="devel">Development (Local)</option>
        <option value="prod">Production (IO API - richiede API Key)</option>
      </select>
    </div>
  </div>
    <div class="form-group">
    <button type="button" class="btn btn-primary" on:click={submitForm}>Invia</button>
    </div>
  </div>
{/if}

{#if result != ""}
<div class="alert alert-success" role="alert">
  {@html result}
</div>
<div>
    <button type="button" class="btn btn-primary" on:click={resetForm}>Nuovo Messaggio</button>
</div>
{/if}

{#if error != ""}
<div class="alert alert-danger" role="alert">
  {error}
</div>
<div>
    <button type="button" class="btn btn-primary" on:click={resetForm}>Nuovo Messaggio</button>
</div>
{/if}
