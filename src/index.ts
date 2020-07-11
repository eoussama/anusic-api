import * as config from "./config.json";
import * as cheerio from "cheerio";
import axios, { AxiosResponse } from "axios";
import { writeFileSync } from "fs";
import { getAnimeList } from "./scrapper";
import { Anime } from "./models/anime.model";

/**
 * Anusic scrapper class
 */
export default class AnusicScrapper {

  //#region Properties

  anime: Anime[] = [];

  //#endregion

  //#region Lifecycle

  constructor() { }

  //#endregion

  //#region Methods

  /**
   * Gets the Anime list
   */
  getAnimeList(): void {
    axios.get(`${config.endpoints.themes}/anime_index`)
      .then((response: AxiosResponse) => {
        const $ = cheerio.load(response.data);
        this.anime = getAnimeList($);
      });
  }

  /**
   * Creates a dump file
   */
  createDump(): void {
    const dump = {
      lastUpdate: new Date().getTime(),
      anime: [...this.anime]
    };

    writeFileSync('dump.json', JSON.stringify(dump, null, 2));
  }

  //#endregion
}

const client = new AnusicScrapper();

client.getAnimeList();
client.createDump();
