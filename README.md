# IO-SDK

IO-SDK è uno strumento di sviluppo che serve a semplificare l'invio di messaggi all'app IO, facilitando l'integrazione con le molteplici fonti di dati esistenti nella Pubblica Amministrazione.

L'app IO è sviluppata dal Governo italiano come unico punto di accesso per interagire con i servizi pubblici locali e nazionali.

Tutti i servizi pubblici italiani possono integrarsi con l'app collegandosi tramite REST per l'invio di comunicazioni al cittadino.

Per facilitare queste operazioni il Software Development Kit per IO (IO-SDK) fornisce una serie di strumenti:

- un server per l'esecuzione di micro-servizi;
- un'interfaccia utente per la gestione dei messaggi;
- un'interfaccia a riga di comando (CLI) per il controllo del server;
- una serie di componenti aggiuntivi (PLUGIN) modificabili per connettersi alle varie fonti dati;
- un ambiente di sviluppo integrato (IDE).

Per maggior informazioni potete consultare la documentazione in lingua italiana:

- [il manuale dell'utilizzatore](/docs/utente.md) indirizzato all'utilizzatore finale;
- [il manuale dell'amministratore](/docs/amministratore.md) che fornisce le informazioni per l'installazione;
- [il manuale dello sviluppatore](/docs/sviluppatore.md) rivolto agli sviluppatori che desiderano contribuire con nuove integrazioni.

Il [documento di sviluppo dell'SDK](DEVEL.md) è invece scritto in lingua inglese, ed è rivolto a tutti coloro che desiderano contribuire allo sviluppo dell'SDK stesso.

Per incontrare la community di IO-SDK potete visitare il topic IO-SDK della [Community NoiOpen](https://noiopen.discourse.group/c/progetti/io-sdk/11) che gestisce il progetto.

# FAQ

## Come lo installo?

- Ti serve Docker, meglio se la versione Desktop, ma almeno la versione Toolbox.
- Scaricati l’installatore dell’ultima [release](https://github.com/pagopa/io-sdk/releases). 
- Leggi la [documentazione di installazione](/docs/amministratore.md).
- Ottieni una chiave dal [backend di io](https://developers.io.italia.it).

## Come lo uso?

Consiglio vivamente di [leggere il manuale](/docs/utente.md).

## Perché non riesco a mandare messaggi alla mia app e al mio codice fiscale?

La chiave che crei nel backend ti permette di mandare messaggi email solo al codice fiscale mostrato nel backend.

Questo codice corrisponde all'indirizzo email con cui ti sei loggato al backend.

Per ottenere maggiori permessi devi contattare l'onboarding di IO.

## Sono uno sviluppatore e voglio contribuire, che mi serve?

Prima di tutto un computer con almeno otto giga di memoria. Se ne hai meno, è dura...

- Se hai un Mac, ti serve una versione recente di OSX a 64 bit e Docker per Mac ([leggi le linee guida per l'installazione](docs/Prerequisites/Mac/OSX.md)). 

- Se hai Linux, ti serve Ubuntu versione 18 oppure 20. Può funzionare anche su altre versioni, ma devi adattare il codice ([leggi le linee guida per l'installazione](docs/Prerequisites/Linux/Ubuntu.md)).

-  Se hai Windows devi assolutamente abilitare la virtualizzazione per poter eseguire Docker su Windows. Devi, inoltre, installare una versione di Windows 10 Professional che supporta `WSL2` e installare Ubuntu 18 o 20, e collegare Docker a WSL2 ([leggi le linee guida per l'installazione](docs/Prerequisites/Windows/10.md)). 

Maggiori istruzioni dettagliate sono nel file [DEVEL.md](/DEVEL.md) (in lingua inglese).

## Come si sviluppa?

Prima di tutto dovrai clonare il repository, compilare il codice e testarlo, questi i comandi da eseguire:

```
git clone https://github.com/pagopa/io-sdk
cd io-sdk
./setup.sh
source source-me-first
make
```

Nel dettaglio ogni comando effettua, rispettivamente, queste attività:
- scarica i sorgenti da github
- accede alla cartella del repository appena scaricato
- installa le dipendenze con `./setup.sh`
- attiva e verifica l'ambiente con `source source-me-first`
- compila tutto con `make` 

**Se questa procedura non funziona in una delle configurazioni supportate è un bug, per favore [riportalo](https://github.com/pagopa/io-sdk/issues).**

A questo punto per sviluppare l'applicazione frontend accedi alla cartella `admin\web`, le action si trovano in `admin\packages`, mentre il launcher è sotto `iosdk`. Vedi dopo.

## Che cosa devo conoscere per contribuire allo sviluppo di IO-SDK?

- Il frontend è scritto in [Javascript moderno](https://javascript.info) con il framework [Svelte](https://svelte.dev).
- Il backend è scritto in [Python 3](https://www.python.org) come azioni (action) [OpenWhisk](https://github.com/apache/openwhisk/blob/master/docs/actions-python.md).
- Il launcher è scritto in [Go](https://golang.org/) e usa [Docker](https://www.docker.com/).

## Come si sviluppa l'applicazione (frontend e action)?

Per sviluppare l'applicazione dovrai lanciare sia il frontend che il backend con i seguenti comandi, dalla root del repository:

```
cd admin
make start
make devel
```

Nel dettaglio i seguenti comandi svolgono le seguenti attività:  
- Entra nella cartella `admin`,
- Lancia il server in modalità sviluppo con `make start`.
- Lancia una build di sviluppo con `make devel`.

### Se vuoi sviluppare il frontend
- Editi il codice javascript/svelte sotto `admin/web/src`.

### Se vuoi sviluppare le action backend
- Editi il codice python delle azioni sotto `admin/web/packages` e le deploy con `make deploy`.

### Per eseguire i test
- Esegui i test con `make test`.

### Compila e installa per l'uso
- Una volta che tutto funziona, puoi tornare nella cartella principale del repository e lanciare il comando `make` che impacchetta il tutto ed esegue tutti i test automatici.

### Aggiorna la documentazione
- Non dimenticarti di [aggiornare la documentatione](/docs) apportando le modifiche introdotte dal tuo sviluppo.

### Effettua una PR per contribuire
- Se tutto funziona, verifica di avere l'ultima versione del codice [*upstream*](https://stackoverflow.com/questions/7244321/how-do-i-update-a-github-forked-repository) se no recupera le modifiche ed effettua un *rebase* 
```
git rebase master
git push origin feature --force
```
e quindi crea una [Pull Request](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/about-pull-requests).

## Come si fa una release?

Devi avere i permessi di accesso al repo per farlo. 
Anche se li hai, per favore non fare release alla leggera...

`git tag <x>.<y>.<z>[<t>] ; git push --tags`

e la build crea automaticamente una release (se passano tutti i test!).

## Come si sviluppa il launcher?

- entra nella cartella `iosdk`
- `make` per compilare
- `make test` per testare

Se un test fallisce c'è uno script che ti permette di capire quale output differisce dalle attese.

- `python3 difftest.py` ti da una lista dei test falliti (0, 1, 2)
- `python3 difftest.py 2` ti mostra i dettagli del test fallito nr 2

## Come si scrivono i test?

- i test del backend sono fatti con [bats](https://github.com/sstephenson/bats)
- i test del launcher sono fatti in Go come [Testable Examples](https://blog.golang.org/examples)
- i test del frontend sono fatti in [Playwright](https://playwright.dev/)


## Il master non funziona!

Il master riceve le pull requests e può essere instabile.  Se vuoi usare una versione stabile ma senza le ultime novità guarda i tag con `git tag` e usa:

```
git checkout v<tag>
```

## Come vi contatto?

Usa in preferenza il [forum Discourse di NoiOpen, Categoria IO-SDK](https://noiopen.discourse.group/c/progetti/io-sdk/11).
