# Motorregister API - Statistik/data

**UNDER UDVIKLING - ENDNU IKKE STABIL**

Hostes på https://api.troendheim.dk og benyttes af https://statistik.troendheim.dk/

## Systemkrav
- Kompatibel med Linux og OSX. Med al sandsynlighed også Windows, men er ikke testet.
- Golang. Testet i `1.10.3`.
- MySQL. Testet i `5.7`.

## Setup
Kræver MySQL `5.7`.

1) Klon projektet
2) Opret konfigurationsfilen `config/app.yml`. Følgende felter findes:
    - dsn **(Påkrævet)**
    - port (Valgfri, default: 8999)

    Eksempelvis:
    ```
     dsn: "username:password@tcp(hostname.dk)/dbname"
     port: 8999
    ```
3) Byg med `go build`. Hent deps med `get get github.com/{dep-navn}`, hvis (når) den fejler. Jeg har ikke kompileret binaries endnu, men det er den langsigtede plan.
4) Kør import og patches. Kald den eksekverbare fil med dataImportFile parameteren. 
    
    `./motorregister-api -dataImportFile migration/data.json`.
     
     Data.json findes i migration mappen og er et snapshot fra motorregisterets eksport  2018-04-28. Vær opmærksom på, at patch-systemet ikke holder styr på, hvad der tidligere er kørt, så den rydder tabellerne først.
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
