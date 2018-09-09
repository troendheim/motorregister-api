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
