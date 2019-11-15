Feature: Validate API responses
    PROYECTO_ACADEMICO_CRUD Controlador tipo_titulacion
    probe JSON responses

Scenario Outline: To probe route code response /tipo_titulacion
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"

    Examples:
    |method |route                        |bodyreq                       |codres         |
    |GET    |/v1/tipo_titulacion          |./assets/requests/empty.json  |200 OK         |
    |GET    |/v1/tipo_titulacio           |./assets/requests/empty.json  |404 Not Found  |
    |POST   |/v1/tipo_titulacion/0        |./assets/requests/empty.json  |404 Not Found  |
    |POST   |/v1/tipo_titulacion          |./assets/requests/empty.json  |400 Bad Request|
    |PUT    |/v1/tipo_titulacio           |./assets/requests/empty.json  |404 Not Found  |
    |PUT    |/v1/tipo_titulacion          |./assets/requests/empty.json  |400 Bad Request|
    |DELETE |/v1/tipo_titulacio           |./assets/requests/empty.json  |404 Not Found  |
    |DELETE |/v1/tipo_titulacion          |./assets/requests/empty.json  |404 Not Found  |

   
Scenario Outline: To probe response route /tipo_titulacion      Probe method GET, POST, PUT, DELETE   
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"      
    And the response should match json "<bodyres>"

    Examples: 
    |method |route                         |bodyreq                          |codres           |bodyres                         |                                                 
    |GET    |/v1/tipo_titulacion           |./assets/requests/empty.json     |200 OK           |./assets/responses/Vok2.json    |
    |POST   |/v1/tipo_titulacion           |./assets/requests/empty.json     |400 Bad Request  |./assets/responses/Ierr6.json   |
    |POST   |/v1/tipo_titulacion           |./assets/requests/BodyGen4.json  |201 Created      |./assets/responses/Vok1.json    |
    |POST   |/v1/tipo_titulacion           |./assets/requests/Nt1.json       |400 Bad Request  |./assets/responses/Ierr1.json   |
    |POST   |/v1/tipo_titulacion           |./assets/requests/Nt2.json       |400 Bad Request  |./assets/responses/Ierr2.json   |
    |POST   |/v1/tipo_titulacion           |./assets/requests/Nt3.json       |400 Bad Request  |./assets/responses/Ierr3.json   |
    |POST   |/v1/tipo_titulacion           |./assets/requests/Nt4.json       |400 Bad Request  |./assets/responses/Ierr4.json   |
    |POST   |/v1/tipo_titulacion           |./assets/requests/Nt5.json       |400 Bad Request  |./assets/responses/Ierr5.json   | 
    |PUT    |/v1/tipo_titulacion           |./assets/requests/BodyRec2.json  |200 OK           |./assets/responses/Vok1.json    |
    |GETID  |/v1/tipo_titulacion           |./assets/requests/BodyGen4.json  |200 OK           |./assets/responses/Vok1.json    |
    |DELETE |/v1/tipo_titulacion           |./assets/requests/empty.json     |200 OK           |./assets/responses/Ino.json     |