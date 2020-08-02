# Manuale dell'utilizzatore di IO-SDK

In questa pagina descriviamo l'uso dell'IO-SDK per inviare messaggi alla app IO.

Dopo aver effettuato l'installazione come descritto nel [manuale dell'amministratore](/docs/amministratore.md) verrà lanciato il vostro browser predefinito che aprirà una pagina principale come quella che vedete di seguito:

![Main Menu](/docs/images/user-main-menu.png)

Le opzioni disponibili sono:

- **Import URL** ...;
- **Custom Import** ...;
- **Send Messages** apre la schermata per l'invio di messaggi multipli, che devono esser prima importati;
- **Single Message** apre la schermata per l'invio di  un singolo messaggio;
- **Single Message** apre la schermata per l'invio di un singolo messaggio;
- **Debugging** apre la schermata per il controllo della memoria intermedia e serve a fare verifiche di esecuzione;
- **Development** apre un editor per lo sviluppo e la modifica degli importatori;
- **About**: fornisce informazioni sullo scopo dell'applicazione e sul team di sviluppatori.

## Import URL

![Import URL](/docs/images/user-import-url.png)

## Custom Import

![Custom Import](/docs/images/user-custom-import.png)


## Send Messages

Quando userete per la prima volta l'applicazione IO-SDK potreste non avere ancora messaggi da inviare, in questo caso apparirà una pagina simile a quella seguente che vi informerà sulle operazioni da eseguire per importare i messaggi:

![Send Message](/docs/images/user-send-message-first-time.png)

...
...


![Send Message](/docs/images/user-send-message.png)

In questa schermata è possibile inviare un singolo messaggio.

Occorre specificare:

- Il codice fiscale del destinatario
- Il soggetto del messaggio
- Il corpo del messaggio in formato markdown

Il codice fiscale deve essere quello di un soggetto abilitato alla ricezione e dipende dalla API Key in uso. Consultare [il manuale amministratore](amministratore.md) per maggiori informazioni.

A questo punto è possibile mandare il messaggio a destinazione selezionado come "endpoint" Production, oppure Development. In quest'ultimo caso, il messaggio sarà visibile per l'ispezione in "Debugging", con il nome "sent:XXXX" dove XXXX è il codice fiscale del destinatario. Altrimenti, se il destinatario è abilitato, verrà ricevuto come email oppure nell'app.

## Importazione Messaggi

La voce di menù "Import Messages" permette di importare i messaggi utilizzando un "connettore". Un "connettore" è un programma (che può essere scritto in molti linguaggi di programmazione, e modificato per l'occorrenza) che consente di collegarsi a varie fonte dati (fogli excel, database, servizi REST) 

Il connettore preasente dipende da quello che è stato installato usando la funzione "Development".

Se non si installa nessun nuovo connettore (come specificato nella sezione), compare il connettore predefinito di esempio, che importa solamente dati di prova:

![Import Messages](/docs/images/user-import-messages.png)

Nel connettore predefinito si devono specificare come utente e password `demo` e questo genera dei dati fittizzi per provare il sistema.

Più interessante è l'importatore "excel" che viene installato a partire dal template `javascript`. Vedi la sezione Development per informazioni su come installare un nuovo connettore.

Una volta installato il connettore "excel" presenta una form che chiede l'upload di un file:

![Import Excel](/docs/images/user-import-excel-form.png)

È possibile specificare un file in un formato predefinito (un esempio di formato è `test/data.xlsx` nella directory specificata dove è installato il connettore stesso). 

![Import Excel Sample](/docs/images/user-import-excel-sample.png)

Usando questo importatore è possibile importare dati usando il formato Excel. Dovete uploadare il un file Excel (in formato `xslsx` e l'importatore farà il resto).

## Invio Messaggio Singolo

Selezionando "Single Message" è possibile inviare un singolo messaggio.

![Send Message](/docs/images/user-send-message.png)

Occorre riempire il campo codice fiscale del destinatario, il soggetto del messaggio e il corpo del messaggio in formato markdown e selezionare l'endpoint, che può essere "Development" o "Production".

Nel primo caso, il messaggio non verrà effettivamente inviato ma sarà visibile in "Debugging", nel secondo caso sarà effettivamente inviato a destinazione, generando una email oppure una notifica nella app IO.

## Invio Messaggi

Una volta importati, i messaggi sono pronti per l'invio e sono visibili sotto "Send Messages".

![Send Messages](/docs/images/user-send-messages.png)

I messaggi da inviare vanno selezionati. Possono essere selezionati uno ad uno con le checkbox. Possono essere selezionati tutti con "Select All". Possono essere deleselezionati tutti con "Deselect All". Infine  si può invertire la selezione.

Ogni messaggio è visibile come un linkp; seguendo il link si arriva alla funzione di invio di messaggio singolo con riempiti i campi di quel messagio.

Una volta selezionati i messaggi possono venire inviati. In basso esiste una selezione dell'endpoint per selezionare se si inviano i messaggi in "Production" oppure se vengono inviati a un ricettore "Development" per testare cosa viene inviato. In quest'ultimo caso i messaggi inviati saranno visibili in "Debugging".

## Debugging

Questa schermata permette di esaminare la "cache", ovvero l'area di lavoro temporanea dove vengono memorizzati i dati per vari scopi

![Debugging](/docs/images/user-debugging.png)

Correntemente sono visibili due cose:

- I messaggi pronti per l'invio (in formato `message:XXX`) 
- I messaggi inviati a "Development" (`sent:XXX`) per verificare l'invio.

È possibile filtrare i messaggi importati e i messaggi inviati.
È anche possibile filtrare usando generici pattern in [questo formato](https://redis.io/commands/keys).

Cliccando su un elemento in cache è possibile vedere l'elemento memorizzato in formato JSON.

È infine possibile pulire la cache, eliminando i messaggi importati.

## Development

Cliccando su Development si accede ad un ambiente di sviluppo integrato, simile al [Microsoft VSCode](https://code.visualstudio.com/) ma eseguito in un web browser. Si tratta di [Eclipse Theia](https://theia-ide.org/).

![Development](/docs/images/user-development.png)

Usando l'editor è possibile installare e modificare generici connettori in vari linguaggi di programmazione. I connettori specifici sono scelti quando si installa l'SDK come [descritto nel manuale amministratore](amministratore.md).

Per installare un connettore, occorre aprire un terminale (Menù "Terminal" seguito da "New Terminal") ed eseguire il comando `./build.sh`.

I dettagli dei singoli connettori variano per ogni linguaggio di programmazione e fonte dati. In generale sono delle funzioni scritte nei vari linguaggi di programmazione che devono "accedere a fonti dati diverse" e ritornare i dati in un formato JSON predefinito. 

Il formato dati dei connettori [è descritto qui](sviluppatore.md).

L'ambiente IDE è molto evoluto  si rimanda a [documentazione specifica di VSCode](https://code.visualstudio.com/docs) in quanto l'editor usa la stessa interfaccia utente.  Notare che nell'IDE fornito sono installati alcuni plugin speicifi per i vari linguaggi di programmazione.

## About

![About](/docs/images/user-manual-about.png)
