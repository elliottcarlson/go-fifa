# go-fifa Library
This libary allows for interactions with the FIFA API. It's currently in alpha phase, feel free to submit PR's or file bugs.

> **Warning**  
> This library uses an undocumented API from the FIFA web servers. It may stop functioning correctly at any point
> due to this

### Usage
Create a new client using the `NewClient()` function. You will need to pass in an `Options` parameter, with optional values for an `HTTPClient`, an `APIBaseURL` and a `UserAgent`. If any values are not provIded, the defaults will be used.

### Currently Supported
The following endpoints are currently supported

| API Endpoint                                                | Function              |
| ----------------------------------------------------------- | --------------------- |
| `/competitions`                                             | `GetCompetitions()`   |
| `/competitions/{competitionId}`                             | `GetCompetition()`    |
| `/timelines/{competitionId}/{seasonId}/{stageId}/{matchId}` | `GetMatchEvents()`    |
| `/live/football/now`                                        | `GetCurrentMatches()` |
| `/calendar/matches`                                         | `GetMatches()`        |
| `/teams/{teamId}`                                           | `GetTeam()`           |
| `/players/{playerId}`                                       | `GetPlayer()`         |
| `/seasons/{seasonId}`                                       | `GetSeason()`         |

### Use of `any`
Because this uses an undocumented API, it is not always possible to find all the proper values for building the types. Because of this, we use the `any` object
for any object we don't have an example of. As we discover new object types, we can update accordingly.