 <h2>Esempio di Form</h2>
  <div class="my-5">
    <div class="form-row">
      <div class="form-group col-md-6">
        <input type="email" class="form-control" id="inputEmail4" placeholder="inserisci il tuo indirizzo email">
        <label for="inputEmail4">Email</label>
      </div>
      <div class="form-group col-md-6">
        <input type="password" class="form-control input-password" id="exampleInputPassword" placeholder="inserisci la tua password">
        <span class="password-icon" aria-hidden="true">
        <svg class="password-icon-visible icon icon-sm"><use xlink:href="/svg/sprite.svg#it-password-visible"></use></svg>
        <svg class="password-icon-invisible icon icon-sm d-none"><use xlink:href="/svg/sprite.svg#it-password-invisible"></use></svg>
      </span>
        <label for="exampleInputPassword">Password</label>
      </div>
    </div>
    <div class="form-row">
      <div class="form-group col">
        <input type="text" class="form-control" id="inputAddress" placeholder="Via Roma, 1">
        <label for="inputAddress">Indirizzo</label>
      </div>
    </div>
    <div class="form-row">
      <div class="form-group col-md-6">
        <input type="text" class="form-control" id="inputCity">
        <label for="inputCity">Comune</label>
      </div>
      <div class="form-group col-md-2">
        <input type="text" class="form-control" id="inputCAP">
        <label for="inputCAP">CAP</label>
      </div>
      <div class="col-md-4">
        <div class="bootstrap-select-wrapper">
          <label>Provincia</label>
          <select title="Scegli un'opzione">
            <option value="Value 1">Opzione 1</option>
            <option value="Value 2">Opzione 2</option>
            <option value="Value 3">Opzione 3</option>
            <option value="Value 4">Opzione 4</option>
            <option value="Value 5">Opzione 5</option>
          </select>
        </div>
      </div>
    </div>
    <div class="form-row">
      <div class="form-group col-md-6">
        <div class="toggles">
          <label for="toggleEsempio1a">
            Label dell'interruttore 1
            <input type="checkbox" id="toggleEsempio1a">
            <span class="lever"></span>
          </label>
        </div>
      </div>
    </div>
    <div class="form-row">
      <div class="form-group col text-center">
        <button type="button" class="btn btn-outline-primary">Annulla</button>
        <button type="submit" class="btn btn-primary">Conferma</button>
      </div>
    </div>
  </div>