<script>
  export let key = "";

  import { onMount } from "svelte";
  import { formValidator } from './lib/validator/formValidator';

  let base = "http://localhost:3280/api/v1/web/guest/";

  let state = {};
  let action = "util/send";
  let data = {
    fiscal_code: {
      value: "",
      rules: "required|fiscalCode"
    },
    subject: {
      value: "",
      rules: "required|min:5|max:255"
    },
    markdown: {
      value: "",
      rules: "required"
    },
    amount: {
      value: "",
      rules: "required|numeric"
    },
    notice_number: {
      value: "",
      rules: "required"
    }
  };

  let isFormValid = formValidator.validateForm( data );

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
          //console.log(data[id])
          data.fiscal_code = message[id].fiscal_code;
          data.subject = message[id].subject;
          data.markdown = message[id].markdown;
          data.amount = message[id].amount;
          data.notice_number = message[id].notice_number;
        }
      })
      .catch(err => {
        state = { error: err };
      });
  }
  onMount(start);

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
    result = "";
    error = "";
    for (key in data) {
      data[key] = "";
    }
  }

  function onChangeFieldValue(field, value) {

    data[field].value = value;

    isFormValid = formValidator.validateForm( data )

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
        class="form-control"
        id="fiscal_code"
        on:input={e => onChangeFieldValue('fiscal_code', e.target.value) }
        bind:value={data.fiscal_code.value} />
    </div>
    <div class="form-group">
      <label class="active" for="subject">Subject</label>
      <input
        type="text"
        class="form-control"
        id="subject"
        on:input={e => onChangeFieldValue('subject', e.target.value) }
        bind:value={data.subject.value} />
    </div>
    <div class="form-group">
      <textarea 
        id="markdown" 
        rows="3" 
        bind:value={data.markdown.value} 
        on:input={e => onChangeFieldValue('markdown', e.target.value) } />
      <label class="active" for="markdown">Message (markdown)</label>
    </div>
    <div class="form-group">
      <label class="active" for="amount">Amount</label>
      <input
        type="text"
        class="form-control"
        id="amount"
        on:input={e => onChangeFieldValue('amount', e.target.value) }
        bind:value={data.amount.value} />
    </div>
    <div class="form-group">
      <label class="active" for="amount">Notice Number</label>
      <input
        type="text"
        class="form-control"
        id="notice_number"
        on:input={e => onChangeFieldValue('notice_number', e.target.value) }
        bind:value={data.notice_number.value} />
    </div>
    <div class="form-group">
      <div class="bootstrap-select-wrapper">
        <label for="select">Endpoint</label>
        <select id="select" bind:value={action} title="Select an endpoing">
          <option value="util/send">Development (Local)</option>
          <option value="iosdk/send">Production</option>
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