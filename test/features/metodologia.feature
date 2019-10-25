Feature: Validate API responses
    PROYECTO_ACADEMICO_CRUD Controlador metodologia
    probe JSON responses

Scenario Outline: To probe route code response /metodologia
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"

    Examples:
    |method |route                    |bodyreq                       |codres         |
    |GET    |/v1/metodologia          |./assets/requests/empty.json  |200 OK         |
    |GET    |/v1/metodologi           |./assets/requests/empty.json  |404 Not Found  |
    |POST   |/v1/metodologia/0        |./assets/requests/empty.json  |404 Not Found  |
    |POST   |/v1/metodologia          |./assets/requests/empty.json  |400 Bad Request|
    |PUT    |/v1/metodologi           |./assets/requests/empty.json  |404 Not Found  |
    |PUT    |/v1/metodologia          |./assets/requests/empty.json  |400 Bad Request|
    |DELETE |/v1/metodologi           |./assets/requests/empty.json  |404 Not Found  |
    |DELETE |/v1/metodologia          |./assets/requests/empty.json  |404 Not Found  |

   
Scenario Outline: To probe response route /metodologia      Probe method GET, POST, PUT, DELETE   
    When I send "<method>" request to "<route>" where body is json "<bodyreq>"
    Then the response code should be "<codres>"      
    And the response should match json "<bodyres>"

    Examples: 
    |method |route                     |bodyreq                          |codres           |bodyres                         |                                                 
    |GET    |/v1/metodologia           |./assets/requests/empty.json     |200 OK           |./assets/responses/Vok2.json    |
    |POST   |/v1/metodologia           |./assets/requests/empty.json     |400 Bad Request  |./assets/responses/Ierr6.json   |
    |POST   |/v1/metodologia           |./assets/requests/BodyGen1.json  |201 Created      |./assets/responses/Vok1.json    |
    |POST   |/v1/metodologia           |./assets/requests/Nt1.json       |400 Bad Request  |./assets/responses/Ierr1.json   |
    |POST   |/v1/metodologia           |./assets/requests/Nt2.json       |400 Bad Request  |./assets/responses/Ierr2.json   |
    |POST   |/v1/metodologia           |./assets/requests/Nt3.json       |400 Bad Request  |./assets/responses/Ierr3.json   |
    |POST   |/v1/metodologia           |./assets/requests/Nt4.json       |400 Bad Request  |./assets/responses/Ierr4.json   |
    |POST   |/v1/metodologia           |./assets/requests/Nt5.json       |400 Bad Request  |./assets/responses/Ierr5.json   | 
    |PUT    |/v1/metodologia           |./assets/requests/BodyRec2.json  |200 OK           |./assets/responses/Vok1.json    |
    |GETID  |/v1/metodologia           |./assets/requests/BodyGen1.json  |200 OK           |./assets/responses/Vok1.json    |
    |DELETE |/v1/metodologia           |./assets/requests/empty.json     |200 OK           |./assets/responses/Ino.json     |
