# Connettori io-sdk  

![connector](/docs/images/UseCase-IO-SDK-Connector.png)

I vari connettori, nei vari linguaggi, devono tornare un json di tipo "form", se non ci sono parametri nell'invocazione del connettore, di tipo "body" quando i parametri ci sono  

## Json FROM  

Il json form descrive i campi che saranno visualizzati nella pagina "Import" dell'io-sdk, di seguito un esempio:


```
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

oppure, se c'è un errore  

```
{
    "body": 
    {
        "data": "messaggio di errore"
    }
}
```

