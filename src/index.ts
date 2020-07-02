import * as config from "./config.json";
import * as cheerio from "cheerio";
import axios, { AxiosResponse } from "axios";
import { writeFileSync } from "fs";
import { getAnimeList } from "./scrapper";

/**
 * Anusic scrapper class
 */
export default class AnusicScrapper {

  //#region Properties
  //#endregion

  //#region Lifecycle

  constructor() { }

  //#endregion

  //#region Methods

  getAnimeList() {
    axios.get(`${config.endpoints.themes}/anime_index`)
      .then((response: AxiosResponse) => {
        const $ = cheerio.load(response.data);
        writeFileSync('dump.json', JSON.stringify(getAnimeList($)));
      });
  }

  //#endregion
}

new AnusicScrapper().getAnimeList();
