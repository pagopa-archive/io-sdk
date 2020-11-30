<script>
  export let key = "";

  import { onMount } from "svelte";
  import { formValidator } from './lib/validator/formValidator';

  let base = "http://localhost:3280/api/v1/web/guest/";

  let state = {};
  let data = formatFormData();
  let action = "util/send";
  let isFormValid = false;

  function formatFormData( data = {} ) {

    return {
      fiscal_code: {
        value: data.fiscal_code || "",
        rules: "required|fiscalCode"
      },
      subject: {
        value: data.subject || "",
        rules: "required|min:5|max:255"
      },
      markdown: {
        value: data.markdown || "",
        rules: "required"
      },
      amount: {
        value: parseInt(data.amount) || 0,
        rules: "required|numeric"
      },
      notice_number: {
        value: data.notice_number || "",
        rules: "required"
      }
    };

  }

  function validateForm( ) {
    const res = formValidator.validateForm( data );
    data        = res.validatedData;
    isFormValid = res.isFormValid;
  }

  async function start() {
    if (key == "") return;
    let id = window.atob(key);
    console.log("start");
    fetch(base + "util/cache", {
      method: "POST",
      body: JSON.stringify({ get: id }),
      headers: { "Content-Type": "application/json" }
    })
      .then(async res => {
        if (res.ok) {
          let message = await res.json();

          data = formatFormData(message[id]);
          validateForm();

        }
      })
      .catch(err => {
        state = { error: err };
      });
  }
  
  onMount(() => {
    start();
  });

  function formatData( data ) {
    let formattedData = {};
    Object.keys(data).map((key, index) => {
        formattedData[key] = data[key].value;
    })
    return formattedData;
  }

  function submitForm() {

    let url = base + action;
    
    fetch(url, {
      method: "POST",
      body: JSON.stringify(formatData(data)),
      headers: {
        "Content-Type": "application/json"
      }
    })
      .then(async function(res) {
        state.result = await res.text();
      })
      .catch(function(err) {
        state.error = err.message;
      });
  }

  function resetForm() {

    state.result = "";
    state.error = "";

    Object.keys(data).forEach((key, index) => {
      data[key].value = "";
    });
    validateForm();
  }

  function onChangeFieldValue(field, value) {

    data[field].value = value;

    validateForm( );

  }

</script>

<h2>Send a Single Message</h2>
<br />
{#if state.result}
  <div class="alert alert-success" role="alert">
    <pre>{state.result}</pre>
  </div>
  <div>
    <button type="button" class="btn btn-primary" on:click={resetForm}>
      Nuovo Messaggio
    </button>
  </div>
{:else if state.error}
  <div class="alert alert-danger" role="alert">{state.error}</div>
  <div>
    <button type="button" class="btn btn-primary" on:click={resetForm}>
      Nuovo Messaggio
    </button>
  </div>
{:else}
  <div>
    <div class="form-group">

        <label class="active" for="fiscal_code">
          Fiscal Code
        </label>
        <input
          type="text"
          class={data.fiscal_code.valid === true ? 'form-control is-valid' : 'form-control is-invalid'}
          id="fiscal_code"
          value={data['fiscal_code'].value}
          on:input={e => onChangeFieldValue('fiscal_code', e.target.value) } />
    </div>
    <div class="form-group">
      <label class="active" for="subject">Subject</label>
      <input
        type="text"
        class={data.subject.valid === true ? 'form-control is-valid' : 'form-control is-invalid'}
        id="subject"
        value={data['subject'].value}
        on:input={e => onChangeFieldValue('subject', e.target.value) } />
    </div>
    <div class="form-group">
      <textarea 
        id="markdown" 
        rows="3" 
        value={data['markdown'].value}
        on:input={e => onChangeFieldValue('markdown', e.target.value) } />
      <label class="active" for="markdown">Message (markdown)</label>
    </div>
    <div class="form-group">
      <label class="active" for="amount">Amount</label>
      <input
        type="text"
        class={data.amount.valid === true ? 'form-control is-valid' : 'form-control is-invalid'}
        id="amount"
        value={data.amount.value}
        on:input={e => onChangeFieldValue('amount', e.target.value) } />
    </div>
    <div class="form-group">
      <label class="active" for="amount">Notice Number</label>
      <input
        type="text"
        class={data.notice_number.valid === true ? 'form-control is-valid' : 'form-control is-invalid'}
        id="notice_number"
        value={data['notice_number'].value}
        on:input={e => onChangeFieldValue('notice_number', e.target.value) } />
    </div>
    <div class="form-group">
      <div class="bootstrap-select-wrapper">
        <label for="select">Endpoint</label>
        <select id="select" bind:value={action} title="Select an endpoint">
          <option value="/util/send">Development (Local)</option>
          <option value="/iosdk/send">Production</option>
        </select>
      </div>
    </div>
    <div class="form-group">
      <button type="button" class="btn btn-primary" disabled={!isFormValid} on:click={submitForm}>
        Send
      </button>
      <button type="button" class="btn btn-primary" on:click={resetForm}>
        Reset
      </button>
    </div>
  </div>
{/if}
