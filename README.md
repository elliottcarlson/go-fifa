# go-fifa Library
This libary allows for interactions with the FIFA API. It's currently in alpha phase, feel free to submit PR's or file bugs.

### Usage
Create a new client using the `NewClient()` function. You will need to pass in an `Options` parameter, with optional values for an `HTTPClient`, an `APIBaseURL` and a `UserAgent`. If any values are not provIded, the defaults will be used.

### Currently Supported
The following endpoints are currently supported:

> **Warning**
> There may be issues with the new v3 API

| API Endpoint                                                | Function               |
| ----------------------------------------------------------- | ---------------------- |
| `/competitions`                                             | `GetCompetitions()`    |
| `/competitions/{competitionId}`                             | `GetCompetition()`     |
| `/timelines/{competitionId}/{seasonId}/{stageId}/{matchId}` | `GetMatchEvents()`     |
| `/live/football/now`                                        | `GetCurrentMatches()`  |
| `/calendar/matches`                                         | `GetTodaysMatches()`   |
| `/teams/{teamId}`                                           | `GetTeam()`            |
| `/players/{playerId}`                                       | `GetPlayer()`          |
| `/seasons/{seasonId}`                                       | `GetSeason()`          |
| `/calendar/{competitionId}/{seasonId}/{stageId}`            | `GetSeasonStandings()` |