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

che vi invita a definire un connettore usando un file di esempio. Potete copiare dunque l'indirizzo proposto e premere il pulsante Import per caricare un esempio e proseguire nell'esplorazione delle funzionalità dell'IO-SDK.

![Import JSON Esempio](/docs/images/user-import-json-esempio.png)

## Custom Import

![Custom Import](/docs/images/user-custom-import.png)

## Send Messages

Quando userete per la prima volta l'applicazione IO-SDK potreste non avere ancora messaggi da inviare, in questo caso apparirà una pagina simile a quella seguente che vi informerà sulle operazioni da eseguire per importare i messaggi:

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

I pulsanti nella parte superiore permettono, nell'ordine, di mostrare tutte le informazioni di debug, visualizzare i messaggi importati, mostrare solo i messaggi inviati, cercare tra le informazioni (qualora l'elenco fosse molto lungo). Per la ricerca potete sfruttare i pattern definendoli in [questo formato](https://redis.io/commands/keys).

Cliccando sulla chiave corrispondente ad un certo messaggio visualizzerete il dettaglio (in formato JSON), come nell'esempio seguente:

![Debugging Details](/docs/images/user-debugging-dettaglio.png)

Il pulsante Clean Current Keys permette di pulire la cache.

## Development

Da questa funzione si accede all'ambiente di sviluppo integrato (IDE) [Eclipse Theia](https://theia-ide.org/), potete trovare maggiori informazioni nella [documentazione specifica di VSCode](https://code.visualstudio.com/docs) in quanto viene usata la stessa interfaccia utente. 

![Development](/docs/images/user-development.png)

Usando questo ambiente è possibile installare e modificare diversi connettori, come anticipato nel [manuale dell'amministratore](amministratore.md) per completare l'installazione del connettore scelto, bisogna aprire un terminale dei comandi partendo dal menu Terminal e successivamente cliccando su New Terminal 

![Dev New Terminal](/docs/images/user-dev-new-terminal.png)

quindi scrivere il comando seguente (premendo invio alla fine):

```
./build.sh
```

![Dev Build OK](/docs/images/user-dev-build-ok.png)

Ulteriori dettagli sui connettori li trovate nel [manuale dello sviluppatore](sviluppatore.md).

## About

![About](/docs/images/user-manual-about.png)



### questa diventa una nuova sezione importare un excel

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
