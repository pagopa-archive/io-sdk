# Manuale dell'utilizzatore di IO-SDK

In questo manuale descrive l'uso del kit di sviluppo IO-SDK per invio di messaggi alla app IO.

Dopo aver effettuato l'installazione come descritto nel [manuale di amministrazione](/docs/amministratore.md) compare un menù principale come segue:

![Main Menu](/docs/images/user-main-menu.png)

Da questo menù è possibile selezionare le seguenti operazioni:

- **Single Message** apre la schermata per l'invio di  un singolo messaggio
- **Import Messages** apre la schermata per l'importazione di messaggi da varie fonte dati
- **Send Messages** apre la schermata per l'invio di messaggi multipli, che devono esser prima importati
- **Debugging** apre la schermata per il controllo della memoria intermedia e serve a fare verifiche di esecuzione
- **Development** apre un editor per lo sviluppo e la modifica degli importatori

## Invio Singolo Messaggio

![Send Message](/docs/images/user-send-message.png)

In questa schermata è possibile inviare un singolo messaggio.

Occorre specificare:

- Il codice fiscale del destinatario
- L'argomento del messaggio ("soggetto")
- Il corpo del messaggio in formato markdown

A questo punto è possibile mandare il messaggio direttamente a destinazione selezionado come endpoint Production, oppure Development. In quest'ultimo caso, il messaggio sarà visibile per l'ispezione in "Debugging", con il nome "sent:XXXX" dove XXXX è il codice fiscale del destinatario.

## Importazione Messaggi

La voce di menù "Import Messages" seleziona un connettore. Il connettore dipende da quello che è stato installato usando la funzione "Development".

Se non si installa nessun connettore (come specificato nella sezione), compare il connettore predefinito di esempio, che importa solamente dati di prova:

![Import Messages](/docs/images/user-import-messages.png)

Nel connettore predefinito si devono specificare come utente e password `demo` e questo genera dei dati fittizzi per provare il sistema.

Più interessante è l'importatore "excel" che viene installato a partire da  template `javascript` (vedi la sezione Development)

Una volta installato il connettore javascript presenta una form che chiede l'upload di un file:

![Import Excel](/docs/images/user-import-excel-form.png)

È possibile specificare un file nel formato dell'esempio (un foglio di test  si trova sotto test/data.xlsx).

![Import Excel Sample](/docs/images/user-import-excel-sample.png)

Usando questo importatore è possibile importare dati usando il formato Excel.

## Invio Messaggio Singolo

Selezionando "Single Message" è possibile inviare un singolo messaggio.

![Send Message](/docs/images/user-send-message.png)

Occorre riempire il campo Codice Fiscale del destinatario, il Soggetto del messaggio e il corpo del messaggio in formato markdown e selezionare la destinazione, che può essere "Development" o "Production".

Nel primo caso, il messaggio non verrà effettivamente inviato ma sarà visibile in "Debugging", nel secondo caso sarà effettivamente inviato a destinazione.

## Invio Messaggi

Una volta importati i messaggi sono pronti per l'invio e sono visibili sotto "Send Messages".

![Send Messages](/docs/images/user-send-messages.png)

I messaggi da inviare vanno selezionati. Possono essere selezionati uno ad uno cliccando checkbox. Possono essere selezionati tutti con "Select All", possono essere deleselezionati tutti con "Deselect All" e si può invertire la selezione.

Ogni messaggio è visibile come un link, seguendo il link si arriva alla funzione di invio di messaggio singolo con riempiti i campi.

Una volta selezionati i messaggi possono venire inviati. In basso esiste un pull down per selezionare se si inviano i messaggi "in produzione" oppure se vengono inviati a uno ricettore "Development" per testare cosa viene inviato. In quest'ultimo caso i messaggi inviati saranno visibili in "Debugging".

## Debugging

Questa schermata permette di esaminare la "cache", ovvero l'area di lavoro temporanea dove vengono memorizzati i dati per vari scopi

![Debugging](/docs/images/user-debugging.png)

Correntemente sono visibili due cose:

- I messaggi pronti per l'invio (in formato `message:XXX`) 
- i messaggi inviati a "Development" (`sent:XXX`).

È possibile filtrare i messaggi importati e i messaggi inviati.
È anche possibile filtrare usando generici pattern come in [questo formato](https://redis.io/commands/keys).

Cliccando su un elemento in cache è possibile vedere l'elemento memorizzato in formato JSON.

È infine possibile pulire la cache, eliminando i messaggi importati.

## Development

Cliccando su Development si accede ad un ambiente di sviluppo integrato, simile al [Microsoft VSCode](https://code.visualstudio.com/) ma eseguito in un web browser. Si tratta del [Eclipse Theia](https://theia-ide.org/).

![Development](/docs/images/user-development.png)

Usando l'editor è possibile installare e modificare generici connettori in vari linguaggi di programmazione. I connettori specifici sono scelti quando si installa l'SDK come [descritto nel manuale amministratore](amministratore.md).

Per installare un connettore, occorre aprire un terminale (Menù "Terminal" seguito da "New Terminal") ed eseguire il comando `./build.sh`.

I dettagli dei singoli connettori variano per ogni linguaggio di programmazione e fonte dati. In generale sono delle funzioni scritte nei vari linguaggi di programmazione che devono "accedere a fonti dati diverse" e ritornare i dati in un formato JSON predefinito. 

Il formano dei connettori [è descritto qui](sviluppatore.md)

L'ambiente è molto evoluto e non può essere descritto qui, si rimanda a [documentazione specifica di VSCode](https://code.visualstudio.com/docs) in quanto l'editor usa la stessa interfaccia utente. 



