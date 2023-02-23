# Anusic-API

## Description

API that scraps `https://www.reddit.com/r/AnimeThemes`.

## Routes

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

---

## Deployment

Anusic-API provides a Docker image that is readily available for use. This makes it easy to quickly deploy the API to any environment that supports Docker.

Quick bootstrap with the following Docker image `eoussama/anusic-api:0.0.3`

## Incident report

This project relied on the [r/AnimeThemes](https://www.reddit.com/r/AnimeThemes/) subreddit's wiki for its data, but the wiki was removed, leaving the API with only cached data that stops at late 2022. As long as the API continues to depend on this cache, the content provided will be limited to that timeframe.

## Deprecation Notice

All Docker tags prior to 0.0.3 are now deprecated as they do not fall back to the cached data. To ensure that you are using the latest version of the API, please use the Docker image tag 0.0.3 or later.