<script>
  export let key = "";

  import { Link } from "svelte-routing";
  import { formData } from "./store";
  import { onMount } from "svelte";

  let base = "http://localhost:3280/api/v1/web/guest/";

  let state = {};
  let action = "util/send";
  let data = {
    fiscal_code: "",
    subject: "",
    markdown: "",
    amount: "",
    notice_number: ""
  };

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

  function submitForm() {
    let url = base + action;
    console.log(url);
    fetch(url, {
      method: "POST",
      body: JSON.stringify(data),
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
        bind:value={data.fiscal_code} />
    </div>
    <div class="form-group">
      <label class="active" for="subject">Subject</label>
      <input
        type="text"
        class="form-control"
        id="subject"
        bind:value={data.subject} />
    </div>
    <div class="form-group">
      <textarea id="markdown" rows="3" bind:value={data.markdown} />
      <label class="active" for="markdown">Message (markdown)</label>
    </div>
    <div class="form-group">
      <label class="active" for="amount">Amount</label>
      <input
        type="text"
        class="form-control"
        id="amount"
        bind:value={data.amount} />
    </div>
    <div class="form-group">
      <label class="active" for="amount">Notice Number</label>
      <input
        type="text"
        class="form-control"
        id="notice_number"
        bind:value={data.notice_number} />
    </div>
    <div class="form-group">
      <div class="bootstrap-select-wrapper">
        <label>Endpoint</label>
        <select bind:value={action} title="Select an endpoing">
          <option value="util/send">Development (Local)</option>
          <option value="iosdk/send">Production</option>
        </select>
      </div>
    </div>
    <div class="form-group">
      <button type="button" class="btn btn-primary" on:click={submitForm}>
        Send
      </button>
      <button type="button" class="btn btn-primary" on:click={resetForm}>
        Reset
      </button>
    </div>
  </div>
{/if}
