import * as config from "./config.json";
import * as cheerio from "cheerio";
import axios, { AxiosResponse } from "axios";
import { writeFileSync } from "fs";

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
        const categories: string[] = [];

        $('.toc li').each((index, element) => {
          categories.push($(element).text().toLowerCase());
        });

        writeFileSync('dump.json', JSON.stringify(categories));
      });
  }

  //#endregion
}

new AnusicScrapper().getAnimeList();
