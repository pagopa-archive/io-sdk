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

![Import Messages](/docs/images/user-import-messages.png)

![Import Excel](/docs/images/user-import-excel-form.png)

![Import Excel Sample](/docs/images/user-import-excel-sample.png)

## Invio Messaggi

![Send Message](/docs/images/user-send-messages.png)

## Debugging

![Debugging](/docs/images/user-debugging.png)

## Development

![Development](/docs/images/user-development.png)
