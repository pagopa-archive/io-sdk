# IO-SDK

IO-SDK è uno strumento di sviluppo che serve a semplificare l'invio di messaggi all'app IO, facilitando l'integrazione con le molteplici fonti di dati esistenti nella Pubblica Amministrazione.

L'app IO è sviluppata dal Governo italiano come unico punto di accesso per interagire con i servizi pubblici locali e nazionali.

Tutti i servizi pubblici italiani possono integrarsi con l'app collegandosi tramite REST per l'invio di comunicazioni al cittadino.

Per facilitare queste operazioni il Software Development Kit per IO (IO-SDK) fornisce una serie di strumenti:

- un server per l'esecuzione di micro-servizi;
- una interfaccia utente per la gestione dei messaggi;
- una interfaccia a riga di comando (CLI) per il controllo del server;
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

## Perchè non riesco a mandare messaggi alla mia app e al mio codice fiscale?

La chiave che crei nel backend ti permette di mandare solo messaggi di email al codice fiscale mostrato nel backend.

Questo codice  corrisponde all'indirizzo di email con cui ti sei loggato al backend.

Per ottenere maggiori permessi devi contattare l'onboarding di IO.

## Sono uno sviluppatore e voglio contribuire, che mi serve?

Prima di tutto un computer con almeno otto giga di memoria. Se ne hai meno, è dura...

- Se hai un Mac, ti serve una versione recente di OSX a 64 bit e Docker per Mac. 

- Se hai Linux, ti serve Ubuntu versione 18 oppure 20. Può funzionare su altre versioni ma devi adattare il codice.

-  Se hai Windows devi assolutamente abilitare la virtualizzazione per poter eseguire Docker  per Windows. Poi devi istallare una versione di Windows 10 che supporta `WSL2` e innstallare Ubuntu 18 o 20, e collegare Docker a WSL2. 

Istruzioni dettagliate sono nel file [DEVEL.md](/DEVEL.md).

## Come si sviluppa?

Prima di tutto compilalo e testalo:

- scarica i sorgenti da github
- installa le dipendenze con `./setup.sh`
- attiva l'ambiente con `source source-me-first`
- builda tutto con `make` 

In codice, se hai una delle configurazioni supportate fai:

```
git clone https://github.com/pagopa/io-sdk
cd io-sdk
./setup.sh
source source-me-first
make
```

**Se questa procedura non funziona in una delle configurazioni supportate è un bug, per favore [riportalo](https://github.com/pagopa/io-sdk/issues).**

A questo punto per sviluppare l'applicazione è sotto `admin` mentre il launcher è sotto `iosdk`. Vedi dopo.

## Che cosa devo conoscere per sviluppare per IO-SDK?

- Il front-end è scritto in [Javascript moderno](https://javascript.info) con il framework [Svelte](https://svelte.dev).
- Il backend è scritto in [Python 3](https://www.python.org) come azioni [OpenWhisk](https://github.com/apache/openwhisk/blob/master/docs/actions-python.md).
- Il launcher è scritto in [Go](https://golang.org/) e usa [Docker](https://www.docker.com/).

## Come si sviluppa l'applicazione?

- Entri nella cartella `admin`,
- Lanci il server in modalità sviluppo con `make start`.
- Lanci una build di sviluppl con `make devel`.
- Editi il codice javascript/svelte sotto `admin/web/src`.
- Editi il codice python delle azioni sotto `admin/web/packages` e le deploy con `make deploy`.
- Esegui i test con `make test`.
- Una volta che tutto funziona, puoi tornare al top level e fare `make` che impacchetta il tutto e testa.
- Non dimenticarti di [aggiornare la documentatione](/docs).
- Se tutto funziona fai una [Pull Request](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/about-pull-requests).

## Come si fa una release?

Devi avere i permessi di accesso al repo per farlo. 
E se li hai, per favore non fare release alla leggera...

`git tag <x>.<y>.<z>[<t>] ; git push --tags`

e la build crea automaticamente una release (se passa i test!).

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

## Come vi contatto?

Usa in preferenza il [forum Discourse di NoiOpen, Categoria IO-SDK](https://noiopen.discourse.group/c/progetti/io-sdk/11).
