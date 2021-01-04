# Anusic-API

## Description

API that scraps `https://www.reddit.com/r/AnimeThemes`.

## Examples

*Anime list*: `https://anusic-api.herokuapp.com/api/v1/anime`

*Anime list filtered by name*: `https://anusic-api.herokuapp.com/api/v1/anime?name=titan`

*Anime list filtered by year*: `https://anusic-api.herokuapp.com/api/v1/anime?year=2017`

*Anime list filtered by multiple years*: `https://anusic-api.herokuapp.com/api/v1/anime?year=2003,2017`

*Anime list filtered by multiple name and years*: `https://anusic-api.herokuapp.com/api/v1/anime?name=titan&year=2013,2019`

*Anime info + themes*: `https://anusic-api.herokuapp.com/api/v1/anime/:id`

---

In order to access the following routes, you need to send an authentication code in the `x-access-token` header

*Logs*: `https://anusic-api.herokuapp.com/api/v1/logs`

*Log content*: `https://anusic-api.herokuapp.com/api/v1/logs/:id`
