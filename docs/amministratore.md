## Prequisiti

Il Software Development Kit (SDK) è disponibile per i seguenti sistemi operativi:

- Windows 10 Professional
- Mac OS Catalina
- Distribuzioni Linux basate su `dpkg` (Debian, Ubuntu)
- Distribuzioni Linux basate su `rpm` (Fedora, CentOS, RedHat)

Per il suo funzionamento richiede l'installazione di Docker e di una "API KEY" per inviare messaggi ad IO. 

### Docker

L'SDK richiede come prerequisito Docker versione 18.06.3 o superiore. Sui sistemi Windows e MacOS potete installare [Docker Desktop](https://www.docker.com/products/docker-desktop).

Sui sistemi Linux seguite le istruzioni fornite per [il vostro sistema come descritto qui](https://docs.docker.com/engine/install/)

### API KEY IO

Per ottenere una API KEY occorre registrarsi al [backoffice di IO](https://developer.io.italia.it/). Le informazioni per ottenere la API Key [sono descritte in questa pagina](https://developer.io.italia.it/openapi.html).

## Installazione

Partiamo dall'installazione della Command Line Interface (CLI): accedete al sito [https://github.com/pagopa/io-sdk/releases](https://github.com/pagopa/io-sdk/releases), selezionate e scaricate una "release" (versione rilasciata) valida per il vostro sistema operativo.

![Releses](/docs/images/admin-releases.png)

Ricordate che: 
- il file `.exe` è l'installer per Windows
- il file `.pkg` è l'installer per Mac OS
- il file `.deb` è l'installer per Ubuntu Linux, Debian Linux e similari
- il file `.rpm` è l'installer per Redhat Linux, Fedora Linux e similari

Installate ed eseguite il programma di installazione come fate normalmente con gli altri progetti sul vostro sistema operativo. Tenete presente che gli installer per MacOS e Windows non sono attualmente firmati digitalmente, quindi occorre autorizzarne l'installazione. Per MacOS potete trovare le istruzioni in [questa pagina](https://support.apple.com/it-it/HT202491).

## Configurazione

Dopo che avrete installato la Command Line Interface (CLI), aprendo una finestra dei comandi sul vostro computer, sarà disponibile il nuovo comando `iosdk`. Nell'esempio seguente sono indicati i comandi che potete digitare per verificare se l'installazione è corretta. Il prefisso `$` potrebbe cambiare in funzione del vostro sistema operativo, in ogni caso non dovete scriverlo prima dei comandi. 

Per esempio, per verificare che la CLI sia disponibile e della versione giusta scrivete (premendo invio alla fine):

```
$ iosdk --version
```

Dovreste ottenere una risposta simile a questa, ricordandovi che potrebbe cambiare in funzione della versione di CLI che avete scaricato:
```
2020.0723.1102-snapshot
```

Potete adesso proseguire inizializzando il sistema scrivendo il comando:

`iosdk init`

Il vostro computer vi chiederà una serie di informazioni per configurare il sistema. La prima informazione richiesta è relativa alla la directory di lavoro:

```
Work Directory (can already exists)
Enter a value (Default is importer):
```

Dovete specifare una directory di lavoro dove il sistema installerà l'importer corrente.

**NOTA:** Tenete presente che Docker, sia su Windows che su MacOS, richiede che la directory si trovi sotto la vostra "home" directory (quella dell'utente che state usando per accedere al vostro computer).

Il secondo passo riguarda la selezionare di uno dei connettori tra quelli disponinibili:

```
Which template:

1. javascript
2. php
3. python
4. github

Enter a number (Default is 1):
```

In particolare:

- `javascript` è un importer per i file in formato Microsoft Excel scritto in Javascript;
- `php` è un importer per caricare dati dai database PostgreSQL scritto in PHP;
- `python` è un importer per caricare dati da fonti GraphQL scritto in Python;
- `github` permette di selezionare generici importer disponibili su GitHub.

Il programma quindi collocherà il codice sorgente dell'SDK nella directory specificata e potrete personalizzarlo successivamente.

**NOTA**: Il connettore specificato non viene automaticamente installato. Deve essere "deployato", vedere il manuale utente per dettagli.

L'ultimo passo è fornire la API Key di IO:

```
IO Api Key
Enter a value: 
```

Occorre specificare il valore ottenuto dal backoffice.

**NOTA**: Per ottenere una API key occorre correntemente registrarsi nel sito `developers.io.italia.it` con la vostra email e da lì è possibile ottenere una API per inviare messaggi al vosto indirizzo emaail tramite un codice fiscale  predefinito e con importo zero. Per ottenere l'invio di messaggi con importi occorre richiederli per email a `onboarding@io.italia.it`.

## Avvio e gestione dei servizi

A questo punto è possibile avviare il sistema con il comando: `iosdk start`

Il comando provvederà a scaricare le immagini docker del kit di sviluppo, configurarlo e lanciarlo in esecuzione, mostrando l'interfaccia utente web descritta nel [manuale utente](/docs/utente.md).


Se non si disponone di sufficiente memoria si può evitare di avviare l'ambiente di sviluppo (IDE) con il comando: `iosdk start --skip-ide`.

### Controllo di stato

Per verificare lo stato di esecuzione dei servizi si può utilizzare `iosdk status`:

```
$ iosdk status
iosdk-openwhisk: running
iosdk-redis: running
iosdk-theia: running
```

Normalmente devono essere in esecuzione 3 servizi (o due se non si è abilitato l'IDE).

### Stop

Per fermare l'esecuzione dei servizi utilizzare `iosdk stop`:

```
$ iosdk stop
Destroying IDE...

Destroying Whisk...

Destroying Redis...
```

### Riavvio

Utilizzare il comando `iosdk restart` per stoppare e riavviare i servizi.
