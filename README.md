# Motorregister API - Statistik/data

Hostes på https://api.troendheim.dk og benyttes af https://statistik.troendheim.dk/

## Systemkrav
- Golang @ `1.11`.
- MySQL @ `5.7`.

## Setup
1) Klon projektet
2) Opret konfigurationsfilen `config/app.yml`. Følgende felter findes:
    - dsn **(Påkrævet)**
    - port (Valgfri, default: 8999)

    Eksempelvis:
    ```
     dsn: "username:password@tcp(hostname.dk)/dbname"
     port: 8999
    ```
2) Hent dependencies med `go get -d ./...` når du står i projektets rod.
3) Byg med `go build`. 
4) Kør import og patches. Kald den eksekverbare fil med normalizedDataImportFile parameteren. 
    
    `./motorregister-api -normalizedDataImportFile migration/data.json`.
     
     `data.json` findes i migration mappen og er et snapshot fra motorregisterets eksport  2018-09-03. Vær opmærksom på, at patch-systemet ikke holder styr på, hvad der tidligere er kørt, så den rydder tabellerne først. De rå filer direkte fra motorregisteret kan normaliseres ved at køre `-rawDataImportFile migration/ESStatistikListeModtag.xml`. Resultatet af denne giver data.json til brug af førstnævnte kommando. Det er ikke nødvendigt at køre `rawDataImportFile` med mindre der skal importeres en specifik version af motorregisteret udover den jeg har bundlet.
5) Start applikationen `./motorregister-api`. Der skrives i konsollen, hvilken port den lytter på. Efterfølgende kan `hostname:port` besøges.

## Endpoints
- **/statistics/model/{mærke}/{model}**

    Hent statistik baseret på mærke og model. 
    
    Returnerer json i følgende format:
    `[{"brandName":"VW","latitude":55.680172,"longtitude":12.585372,"modelName":"CADDY","totalCount":3,"zipCode":1050} , {...} ]` 

- **/models/{brand}**

    Hent alle modeller med indregistrerede køretøjer for et givent mærke.
    
    Returnerer json i følgende format:
    `[{"brand_id":2,"id":7,"name":"X3 Xdrive 30d"} , {...} ]`
    
- **/brands**

    Hent alle mærker med indregistrerede køretøjer.
    
    Returnerer json i følgende format:

    `[{"id":21,"name":"ALFA ROMEO"} , {...} ]`
