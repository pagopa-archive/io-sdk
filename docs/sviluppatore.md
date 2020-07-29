Connettori IO-SDK  
_________________

![connector](/docs/images/UseCase-IO-SDK-Connector.png)

I vari connettori, nei vari linguaggi, devono tornare un json di tipo "form", se non ci sono parametri nell'invocazione del connettore, di tipo "body" quando i parametri ci sono  

# Indice

[Json FORM](#json-form)  
- [Proprietà dei descrittori](#proprietà-dei-descrittori)  

[Json BODY](#json-body)  
- [Gestione errore](#gestione-errore)
- [Proprietà dei dati di ritorno](#proprietà-dei-dati-di-ritorno)  

[Sviluppo contenitore già esistente](#sviluppo-contenitore-già-esistente)
- [Modifica del repository](#modifica-del-repository)
  - [Forkare il repository sul proprio GitHub](#forkare-il-repository-sul-proprio-github)
  - [Clonare il repository in locale](#clonare-il-repository-in-locale)
  - [Creare un branch feature del repository in locale](#creare-un-branch-feature-del-repository-in-locale)
  - [Modifica al codice e verifica con unit test](#modifica-al-codice-e-verifica-con-unit-test)
  - [Provare il connettore integrandolo con IO-SDK](#provare-il-connettore-integrandolo-con-io-sdk)
  - [Al termine dello sviluppo pubblicare la feature nel proprio GitHub](#al-termine-dello-sviluppo-pubblicare-la-feature-nel-proprio-github)
  - [Creare una Pull Request verso NoiOpen](#creare-una-pull-request-verso-noiopen)
- [Modifica Live](#modifica-live)  
  - [Installare IO-SDK](#installare-io-sdk)
    - [Modalità SNAPSHOT](#modalità-snapshot)
    - [Modalità Sviluppo](#modalità-sviluppo)
      - [Clonare il repository IO-SDK in locale](#clonare-il-repository-io-sdk-in-locale)
      - [Avviare IO-SDK in sviluppo](#avviare-io-sdk-in-sviluppo)
  - [Aprire interfaccia Development](#aprire-interfaccia-development)
    - [Linguaggio Javascript](#linguaggio-javascript)
      - [Installare / Aggiungere reference da ide](#installare-/-aggiungere-reference-da-ide)
      - [Eseguire Unit Test da ide](#eseguire-unit-test-da-ide)
      - [Eseguire Build da ide](#eseguire-build-da-ide)
      - [Test di integrazione](#test-di-integrazione)  
[Sviluppo contenitore linguaggio non ancora presente](#sviluppo-contenitore-linguaggio-non-ancora-presente)
  - [Creazione repository](#creazione-repository)
    - [Linguaggio PHP](#linguaggio-php)
    - [Linguaggio Phyton](#linguaggio-phyton)
    - [Linguaggio CSharp Dotnet Core](#linguaggio-csharp-dotnet-core)
    - [Linguaggio JAVA](#linguaggio-java)

## Json FROM  

Il json form descrive i campi che saranno visualizzati nella pagina "Import" dell'io-sdk, di seguito un esempio:


```
{ "body": 
  {
    "form": [
      {
        "type": "message",
        "name": "note",
        "description": "Connect to SQLServer db to import messages"
      },
      {
        "name": "connectionstring",
        "description": "ConnectionString",
        "type": "string",
        "required": true
      }
  ]
}
```

### Proprietà dei descrittori
[ToDo]  


## Json BODY  

Se ritornano i messaggi:

```
{
    "body": 
    {
        "data":
        [
            {
                "amount": 0,
                "due_date": "2020-06-01",
                "fiscal_code": "CodiceFiscale",
                "invalid_after_due_date": false,
                "markdown": "testo messaggio con formattazione markdown",
                "notice_number": 1,
                "subject": "titolo messaggio"
            }
        ]
    }
}

```

### Gestione errore
oppure, se c'è un errore  

```
{
    "body": 
    {
        "error": "messaggio di errore"
    }
}
```

### Proprietà dei dati di ritorno
[ToDo]  

## Sviluppo contenitore già esistente
(Es.: io-sdk-javascript)

### Modifica del repository
Sviluppo in TDD e poi integrazione nell'IO-SDK (vedi Modifica LIVE)  
Esempio con io-sdk-javascript  

#### Forkare il repository sul proprio GitHub
1. Loggarsi su GitHub.  
2. Andare su https://github.com/noiopen/io-sdk-javascript  
3. Cliccare sul bottone "Fork" in alto a destra  

#### Clonare il repository in locale
```
git clone https://github.com/<tuo_username>/io-sdk-javascript
```
(sostituire <tuo_username> con il proprio username)  

#### Creare un branch feature del repository in locale
```
git checkout -b <nome_feature>
```
(sostituire <nome_feature> con il nome della propria feature)    

#### Modifica al codice e verifica con unit test  
1. Creare/modificare il test  
2. Implementare/modificare il codice
3. Da console lanciare il test
```
npm test
```
4. Continuare ad iterare fino al termine dell'implementazione (nel classico iter del TDD)  

#### Provare il connettore integrandolo con IO-SDK
Vedi paragrafo [Modifica Live](#modifica-live) e in particolare il paragrafo [Test di Integrazione](#test-di-integrazione).  

#### Al termine dello sviluppo pubblicare la feature nel proprio GitHub
```
git add .
git commit -m"<tuo_messaggio>"
git push -u origin <nome_feature>
```
(sostituire <nome_feature> con il nome della propria feature)  
(sostituire <tuo_messaggio> con il messaggio del commit)  

#### Creare una Pull Request verso NoiOpen
1. Loggarsi su GitHub.  
2. Andare su https://github.com/<tuo_username>/io-sdk-javascript   
3. Cliccare sul bottone "Pull request"  
4. Indicare, a sinistra, il repository NoiOpen e il branch al quale si vuole contribuire, a destra, il nostro repository e il branch dove si è sviluppata la feature.  
5. Indicare un messaggio per la Pull request, e se presente indicare la Issue relativa
```
See #numero_issue
```
6. Confermare

(sostituire <tuo_username> con il proprio username)  

*Grazie per il tuo contributo! :)*


### Modifica Live

#### Installare IO-SDK

##### Modalità SNAPSHOT
Consulta le informazioni sulla documentazione del repo [IO-SDK](https://github.com/noiopen/io-sdk/blob/master/README.md)  

In particolare una volta installato (o se hai già installato DOCKER) [queste istruzioni](https://github.com/noiopen/io-sdk/blob/master/README.md#using-io-sdk)  

##### Modalità Sviluppo
###### Clonare il repository IO-SDK in locale
```
git clone https://github.com/noiopen/io-sdk-javascript
```
(non abbiamo bisogno di forkare il repo perché non dovremo ad andarlo a modificare)  

##### Avviare IO-SDK in sviluppo
Consulta la documentazione di [sviluppo](https://github.com/noiopen/io-sdk/blob/master/DEVEL.md)  

In particolare una volta installato (o se hai già installato DOCKER) [queste istruzioni](https://github.com/noiopen/io-sdk/blob/master/DEVEL.md#setup-the-development-environmewnt)

#### Aprire interfaccia Development
Esempio con io-sdk-javascript (impostato di default in IO-SDK)  

Dall'UI IO-SDK, menù di sinistra cliccare sulla voce 'Development'  
![Menù Sinistra](/docs/images/Menu-IO-SDK.png)  

Si apre una nuova finestra con un ide online  
![UI Development](/docs/images/UI-IO-SDK-Development.png)


##### Linguaggio Javascript

###### Installare / Aggiungere reference da ide
- da interfaccia
  1. Tasto destro sulla voce **'package.json'** del menù *'NPM SCRIPTS'* e sul menù contestuale selezionare *Open*
  2. Aggiungere la riga della reference da installare e salvare
  3. Tasto destro sulla voce **'package.json'** del menù *'NPM SCRIPTS'* e sul menù contestuale selezionare *Run Install*  
![Menù NPM SCRIPTS](/docs/images/Menu-NPM-SCRIPTS.png)

- da terminale
Aprire il terminale dal menù in altro *'Terminal'* -> **'New terminal'**  
Lanciare il compando
```
npm install
```
![New Terminal](/docs/images/New-Terminal.png)  


###### Eseguire Unit Test da ide
- da interfaccia
  1. Tasto destro sulla voce **'test'** del menù *'NPM SCRIPTS'* e sul menù contestuale selezionare *Run* oppure usare il simbolo *play* sulla destra  
![Menù Run Test](/docs/images/IDE-Test-Run.png)  

- da terminale
Aprire il terminale dal menù in altro *'Terminal'* -> **'New terminal'**  
Lanciare il compando
```
npm test
```


###### Eseguire Build da ide
- da interfaccia
  1. Tasto destro sulla voce **'build'** del menù *'NPM SCRIPTS'* e sul menù contestuale selezionare *Run* oppure usare il simbolo *play* sulla destra  
![Menù Run Build](/docs/images/IDE-Build-Run.png)  

- da terminale
Aprire il terminale dal menù in altro *'Terminal'* -> **'New terminal'**  
Lanciare il compando
```
npm run-script build
```

###### Test di integrazione
Se i test automatici terminano con successo, avviare IO-SDK, cliccare sulla voce **'Import Messages'** del menù di sinistra e testare il comportamento atteso.  
![Import Page](/docs/images/IO-SDK-Import-Page.png)  

##### Linguaggio [nome_linguaggio]
[ToDo - Template per aggiungere nuovi linguaggi]

###### Installare / Aggiungere reference da ide

###### Buildare da ide

###### Test da ide

###### Test di integrazione

## Sviluppo contenitore linguaggio non ancora presente

[ToDo]

### Creazione repository
(TDD e poi integrazione nell'IO-SDK - vedi Modifica LIVE)

[ToDo]

#### Linguaggio PHP

[ToDo]

#### Linguaggio Python

[ToDo]

#### Linguaggio CSharp Dotnet Core

[ToDo]

#### Linguaggio JAVA

[ToDo]

