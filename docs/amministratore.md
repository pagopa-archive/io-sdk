## Prerequisiti

Il Software Development Kit (SDK) è disponibile per i seguenti sistemi operativi:

- Windows 10 Professional
- Mac OS Catalina
- Distribuzioni Linux basate su `dpkg` (Debian, Ubuntu)
- Distribuzioni Linux basate su `rpm` (Fedora, CentOS, RedHat)

Per il suo funzionamento richiede l'installazione di Docker e di una "API KEY" per inviare messaggi ad IO. 

### Docker

L'SDK richiede come prerequisito Docker versione 18.06.3 o superiore.
Per Windows e MacOS potete installare [Docker Desktop](https://www.docker.com/products/docker-desktop). Per Linux seguite le istruzioni fornite per [il vostro sistema come descritto qui](https://docs.docker.com/engine/install/)

### API KEY IO

Per ottenere una API Key occorre registrarsi sul [backoffice di IO](https://developer.io.italia.it/). Le informazioni per ottenere la API Key [sono descritte in questa pagina](https://developer.io.italia.it/openapi.html).

**IMPORTANTE** 
In questo modo sarà possibile ottenere una API Key per inviare messaggi al vosto indirizzo e-mail tramite un codice fiscale fittizio predefinito e con importo zero. Per ottenere l'invio di messaggi con importi diversi, occorre farne espressa richiesta scrivendo all'indirizzo di posta elettronica `onboarding@io.italia.it`.

## Installazione

Partiamo dall'installazione della Command Line Interface (CLI): accedete al sito [https://github.com/pagopa/io-sdk/releases](https://github.com/pagopa/io-sdk/releases), selezionate e scaricate una versione rilasciata (release) valida per il vostro sistema operativo.

![Releases](/docs/images/admin-releases.png)


Ricordate che: 
- il file `.exe` è l'installer per Windows
- il file `.pkg` è l'installer per Mac OS
- il file `.deb` è l'installer per Ubuntu Linux, Debian Linux e similari
- il file `.rpm` è l'installer per Redhat Linux, Fedora Linux e similari

Installate ed eseguite il programma di installazione come fate normalmente con gli altri programmi sul vostro sistema operativo. Tenete presente che gli installer per MacOS e Windows non sono al momento firmati digitalmente, quindi occorre autorizzarne l'installazione. Per MacOS potete trovare le istruzioni in [questa pagina](https://support.apple.com/it-it/HT202491).

## Configurazione

Dopo che avrete installato la CLI, aprite un terminale dei comandi sul vostro computer e verificate scrivendo il comando seguente (premendo invio alla fine):

```
iosdk --version
```

Dovreste ottenere una risposta simile a questa, ricordatevi che potrebbe cambiare in funzione della versione di CLI che avrete installato:

```
2020.0723.1102-snapshot
```

Sempre dal terminale dei comandi potete proseguire scrivendo il comando:

```
iosdk init
```

Il programma vi chiederà una serie di informazioni per completare la prima configurazione. La prima domanda è quella relativa alla la directory di lavoro dell'SDK:

```
Work Directory (can already exists)
Enter a value (Default is importer):
```

Dovete specificare una directory di lavoro nella quale il programma installerà l'importer che sceglierete in segiuto.

**IMPORTANTE** 
Docker, sia su Windows che su MacOS, richiede obbligatoriamente che la directory di lavoro sia collocata sotto la vostra "home" directory (quella dell'utente che state usando per accedere al vostro computer).

### Selezione dei connettori

Il secondo passo riguarda la selezionare di uno dei connettori tra quelli disponibili:

```
Which template:

1. javascript
2. php
3. python
4. github

Enter a number (Default is 1):
```

In particolare:

- `javascript` serve per importare i file in formato Microsoft Excel ed è scritto in Javascript;
- `php` per caricare dati dai database PostgreSQL potete scegliere il connettore scritto in PHP;
- `python` è un importer per caricare dati da fonti GraphQL scritto in linguaggio Python;
- `github` permette di selezionare generici importer disponibili su GitHub.

Per effettuare la scelta il numero corrispondente, da 1 a 4.

Il programma quindi collocherà il codice sorgente dell'SDK nella directory specificata. 

Poiché il connettore specificato non viene automaticamente installato, deve essere implementato. Potrete farlo successivamente seguendo le indicazioni nel [manuale dell'utilizzatore](/docs/utente.md).

L'ultimo passo è fornire la API Key di IO:

```
IO Api Key
Enter a value: 
```

Occorre inserire la stringa alfanumerica che si trova nella sezione [Sottoscrizioni del vostro profilo sul "backoffice" di IO](https://developer.io.italia.it/profile).

## Avvio e gestione dei servizi

A questo punto è possibile avviare il sistema con il comando:

```
iosdk start
```

**IMPORTANTE**
Poiché su Docker dovranno essere disponibili almeno 4 Gb di memoria RAM, se non si disponete di memoria sufficiente, potete evitare di avviare l'ambiente di sviluppo (IDE) scrivendo questo comando: 

Se non si dispone di sufficiente memoria si può evitare di avviare l'ambiente di sviluppo (IDE) con il comando: `iosdk start --skip-ide`.

Il comando provvederà a scaricare le immagini docker del kit di sviluppo, configurarlo e lanciarlo in esecuzione

![Releses](/docs/images/iosdk-docker-start.png)

Al termine verrà visualizzata l'interfaccia utente come una pagina web (verrà lanciata nel vostro browser predefiniti). Le funzionalità sono descritte nel [manuale dell'utilizzatore](/docs/utente.md)

![Main Menu](/docs/images/user-main-menu.png)

### Controllo dello stato dei servizi

Per verificare lo stato di esecuzione dei servizi potete scrivere, sempre da un terminale dei comandi:

```
iosdk status
```

che vi darà come risultato qualcosa di simile:

```
iosdk-openwhisk: running
iosdk-redis: running
iosdk-theia: running
```

Normalmente devono essere in esecuzione questi tre servizi (oppure soltanto due se non è stato lanciato l'IDE).

### Fermare l'esecuzione dei servizi

Per fermare l'esecuzione dei servizi scrivete:

```
iosdk stop
```

nel terminale leggerete:

```
Destroying IDE...

Destroying Whisk...

Destroying Redis...
```

### Riavviare i servizi

Per fermare e riavviare i servizi, scrivete, sempre da un terminale dei comandi:

```
iosdk restart
```
