# Manuale dell'utilizzatore di IO-SDK

In questa pagina descriviamo l'uso dell'IO-SDK per inviare messaggi alla app IO.

Dopo aver effettuato l'installazione come descritto nel [manuale dell'amministratore](/docs/amministratore.md) verrà lanciato il vostro browser predefinito che aprirà una pagina principale come quella che vedete di seguito:

![Main Menu](/docs/images/user-main-menu.png)

Le opzioni disponibili sono:

- **Import URL**
- **Custom Import**
- **Send Messages**
- **Single Message**
- **Debugging**
- **Development**
- **About**

## Import URL

Se è la prima volta che cliccate su questa funzione vedrete una pagina simile a questa:

![Import URL](/docs/images/user-import-url.png)

Da schermata è possibile inviare un singolo messaggio.

che vi invita a definire un connettore usando un file di esempio. Potete copiare dunque l'indirizzo proposto e premere il pulsante Import per caricare un esempio e proseguire nell'esplorazione delle funzionalità dell'IO-SDK.

![Import JSON Esempio](/docs/images/user-import-json-esempio.png)

## Send Messages

Il codice fiscale deve essere quello di un soggetto abilitato alla ricezione e dipende dalla API Key in uso. Consultate [il manuale amministratore](amministratore.md) per maggiori informazioni.

A questo punto è possibile inviare il messaggio al destinatario selezionado come "endpoint" Production, oppure Development. In quest'ultimo caso, il messaggio sarà visibile per l'ispezione in "Debugging", con il nome "sent:XXXX" dove XXXX è il codice fiscale del destinatario. Altrimenti, se il destinatario è abilitato, verrà ricevuto come email oppure nell'app.

![Send Message](/docs/images/user-send-message-first-time.png)

Se avete già importato almeno un messaggio, la schermata invece potrebbe essere simile a questa:

![Send Message After Import](/docs/images/user-send-message-after-import.png)

Trovate l'indicativo del messaggio o dei messaggi importati con alla sinistra di ciascun messaggio una casella di spunta che permette di indicare la selezione specifica del messaggio. Dopo avere selezionato almeno un messaggio potete scegliere l'endpoint di riferimento tra i valori proposti nella casella di scelta ad elenco e premere il pulsante Send Selected Messages per inviare il messaggio

![Send Selected Message](/docs/images/user-send-selected-message.png)

Il corretto invio verrà indicato da un messaggio di conferma come quello evidenziato nell'immagine sottostante:

![Send Selected Message OK](/docs/images/user-send-selected-message-ok.png)

## Single Message

![Send Single Message](/docs/images/user-send-single-message.png)

Selezionando l'invio di un singolo messaggio occorre specificare:

- Fiscal Code: il codice fiscale del destinatario;
- Subject: il soggetto del messaggio;
- Message: il testo del messaggio;
- Amout: l'importo;
- Notice Number: il numero dell'avviso.

Specificate anche in questo caso il tipo di Endpoint e premete sul pulsante Send per inviare. Qualora sceglieste "Production" come Endpoint, se il destinatario è abilitato correttamente, verrà spedito un messaggio e-mail nell'App IO.

**NOTE**

a) Per scrivere il testo del messaggio potete usare una sintassi del testo chiamata "Markdown", ovvero una codifica progettata in modo che possa essere convertita in HTML

b) Il codice fiscale deve essere quello di un soggetto abilitato alla ricezione ed è correlato alla API Key in uso. Consultate [il manuale dell'amministratore](amministratore.md) per maggiori informazioni.

## Debugging

Da questa sezione dell'applicazione è possibile effettuare i riscontri sul corretto funzionamento ed ottenere maggiori dettagli su eventuali anomalie. Nell'immagine qui sotto riportata vedete ad esempio il messaggio spedito prima come esempio.

![Debugging](/docs/images/user-debugging.png)

I pulsanti nella parte superiore permettono, nell'ordine, di mostrare tutte le informazioni di debug, visualizzare i messaggi importati, mostrare solo i messaggi inviati, cercare tra le informazioni. In particolare, per la ricerca potete sfruttare i pattern facendo riferimento a [questo formato](https://redis.io/commands/keys).

Cliccando sulla chiave corrispondente ad un certo messaggio visualizzerete il dettaglio (in formato JSON), come nell'esempio seguente:

![Debugging Details](/docs/images/user-debugging-dettaglio.png)

Il pulsante Clean Current Keys permette di pulire la cache.

## Development

Da questa funzione si accede all'ambiente di sviluppo integrato (IDE) [Eclipse Theia](https://theia-ide.org/), potete trovare maggiori informazioni nella [documentazione specifica di VSCode](https://code.visualstudio.com/docs) in quanto viene usata la stessa interfaccia utente. 

![Development](/docs/images/user-development.png)

Usando questo ambiente è possibile installare e modificare diversi connettori, come anticipato nel [manuale dell'amministratore](amministratore.md) per completare l'installazione del connettore scelto, bisogna aprire un terminale dei comandi partendo dal menu Terminal e successivamente cliccare su New Terminal 

![Dev New Terminal](/docs/images/user-dev-new-terminal.png)

quindi scrivere il comando seguente (premendo invio alla fine):

```
./build.sh
```

![Dev Build OK](/docs/images/user-dev-build-ok.png)

Ulteriori dettagli sui connettori li trovate nel [manuale dello sviluppatore](sviluppatore.md).


## Custom Import

Abbiamo tenuto per ultima la descrizione della funzionalità di Custom Import perché, se avete seguito l'esempio nel [manuale dell'amministratore](amministratore.md) e poi i passaggi descritti in questa pagina, a questo punto dovreste aver completato la configurazione per caricare un file in formato Excel.

![Custom Import](/docs/images/user-custom-import.png)

Un "connettore" è un programma (che può essere scritto in linguaggi di programmazione diversi secondo la necessità) che consente di collegarsi a varie fonti di dati: dai file in formato Excel, appunto, ai servizi REST.

Il file da caricare nel nostro caso di esempio deve trovarsi nella directory specificata dove è stato installato il connettore e potrebbe essere ad esempio:
```
test/data.xlsx
```

L'ambiente IDE è molto evoluto si rimanda a [documentazione specifica di VSCode](https://code.visualstudio.com/docs) in quanto l'editor usa la stessa interfaccia utente.  Notare che nell'IDE fornito sono installati alcuni plugin specifici per i vari linguaggi di programmazione.

![Import Excel Sample](/docs/images/user-import-excel-sample.png)
