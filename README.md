# API til motorregisteret

**UNDER UDVIKLING - IKKE STABIL ELLER BRUGBAR ENDNU**

## Setup
Er pt. baseret på MySQL og fungerer i version `5.7`.

- Opret konfigurationsfilen `config/app.yml`.
    Følgende felter findes:
    - dsn **(Påkrævet)**
    - port (Valgfri, default: 8999)

    Eksempelvis:
    ```
    dsn: "username:password@tcp(hostname.dk)/dbname"
    port: 8999
    ```
