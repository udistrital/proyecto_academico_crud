Feature: Validate API responses
    PROYECTO_ACADEMICO_CRUD
    probe JSON responses

Scenario Outline: To probe route code response /enfasis
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"

    Examples:
    |method |route                |bodyreq                       |codres       |
    |GET    |/v1/enfasis          |./assets/requests/empty.json  |200 OK       |
    |GET    |/v1/enfasis/0        |./assets/requests/empty.json  |404 Not Found|
    |POST   |/v1/enfasis/0        |./assets/requests/empty.json  |404 Not Found|
    |POST   |/v1/enfasis          |./assets/requests/empty.json  |200 OK       |
    |PUT    |/v1/enfasi          |./assets/requests/empty.json  |404 Not Found|
    |DELETE |/v1/enfasi          |./assets/requests/empty.json  |404 Not Found|