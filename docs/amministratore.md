## Prequisiti

L'sdk è disponibile per i seguenti sistemi operativi:

- Windows 10
- Mac OS Catalina
- Distribuzioni Linux basate su `dpkg` (Debian, Ubuntu)
- Distribuzioni Linux basate su `rpm` (Fedora, CentOS, RedHat)

Richiede l'installazione di Docker e una "API KEY" per inviare messaggi ad IO. 

### Docker

L'SDK richiede come prerequisito Docker, versione 18.06.3 o superiore.

Sui sistemi Windows e Mac installare [Docker Desktop](https://www.docker.com/products/docker-desktop)

Sui sistemi Linux seguire le installazioni per [il vostro sistema come descritto qui](https://docs.docker.com/engine/install/)

### API KEY IO

Per ottenere una API KEY occorre registrarsi al [backoffice di IO](https://developer.io.italia.it/)

Informazioni per ottenere la API Key [sono qui](https://developer.io.italia.it/openapi.html).

## Installazione

Per installare l'SDK occorre installare la CLI (command line interface) di controllo.

Andare andare sul sito [https://github.com/pagopa/io-sdk/releases](https://github.com/pagopa/io-sdk/releases), selezionare una "release" (versione rilasciata) e scaricare l'installer corrispondente per il vostro sistema operativo.

![Releses](/docs/images/admin-releases.png)

Notare che 
- il file `.exe` è l'installer per Windows
- il file `.pkg` è l'installer per Mac OS
- il file `.deb` è l'installer per Ubuntu Linux, Debian Linux e similari
- il file `.rpm` è l'installer per Redhat Linux, Fedora Linux e similari

Installare ed eseguire l'installer come da prassi per il vostro sistema operativo.

**NOTA:** gli installer per MacOS e Windows non sono attualmente firmati, quindi occorre autorizzarne l'installazione.

## Configurazione

Una volta installata la CLI sarà disponibile nel PATH il comando `iosdk`.

**NOTA:** nel seguito il prefisso `$` indica comandi che dovreste digitare alla riga di comando, mentre il resto sono output di esempio, che possono cambiare a seconda dei casi.

Verificare che il comando sia disponibile e della versione giusta con:

```
$ iosdk --version
pagopa0.5.1
```

A questo punto inizializzate il sistema con il comando: `iosdk init`.

Il sistema vi chiederà una serie di informazioni per configurare il sistema.

La prima informazione richiesta è la directory di lavoro:


```
Work Directory (can already exists)
Enter a value (Default is importer):
```

Dovete specifare una directory di lavoro dove il sistema installerà l'importer corrente.

**NOTA:**  per vincoli tecnici di Docker su Windows e MacOS, è necessario che la directory si trovi sotto la vostra "home" directory.

Il secondo step è selezionare uno dei connettori esistenti.

```
Which template:

1. javascript
2. php
3. python
4. github

Enter a number (Default is 1):
```

Sono disponibili correntemente:

- `javascript` è un importer per importare file excel scritto in Javascript
- `php` è un importer per importare dati da database Postgresql scritto in PHP
- `python` è un importer per importare dati da fonti dati GraphQL scritto in Python
- `github` permette di selezionare generici importer disponibili su GitHub

L'SDK importerà il connettore selezionato.

L'ultimo passo è fornire la API Key di IO

```
IO Api Key
Enter a value: 
```

Occorre specificare il valore ottenuto dal backoffice.

## Avvio e gestione dei servizi

A questo punto è possibile avviare il sistema con il comando: `iosdk start`

Il comando provvederà a scaricare le immagini docker del kit di sviluppo, configurarlo e lanciarlo in esecuzione, mostrando l'interfaccia utente web descritta nel [manuale utente](/docs/utente.md)


### Controllo di stato

Per verificare lo stato di esecuzione dei servizi si può utilizzare `iosdk status`:

```
$ iosdk status
iosdk-openwhisk: running
iosdk-redis: running
iosdk-theia: running
```

Normalmente devono essere in esecuzione 3 servizi.

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

