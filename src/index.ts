import { writeFileSync, mkdirSync, existsSync } from "fs";
import { resolve } from "path";

import * as config from "./config.json";
import * as cheerio from "cheerio";
import axios, { AxiosResponse } from "axios";


import { getAnimeList } from "./scrapper";
import { Anime } from "./models/anime.model";
import { IAnusicConfig } from "./models/config.model";

/**
 * Anusic scrapper class
 */
export default class AnusicScrapper {

  //#region Properties

  dumpLocation: string = '';

  anime: Anime[] = [];

  //#endregion



  //#region Lifecycle

  /**
   * Creates an Anusic instance
   * @param config The configuration object
   */
  constructor(config?: IAnusicConfig) {

    // Checking if the config object is provided
    if (config) {

      // Sanitizing the dump location property
      this.dumpLocation = config.dumpLocation ?? '';
    }
  }

  //#endregion



  //#region Methods

  /**
   * Gets the Anime list
   */
  async getAnimeList(): Promise<Anime[]> {
    return new Promise((res, rej) => {

      // Getting the scrapped data
      axios.get(`${config.endpoints.themes}/anime_index`)
        .then((response: AxiosResponse) => {

          // Loading the scrapped data
          const $ = cheerio.load(response.data);

          // Formatting the data
          this.anime = getAnimeList($);

          // Resolving the promise
          res(this.anime);
        })
        .catch(err => rej(err));
    });
  }

  /**
   * Creates a dump file
   */
  createDump(): void {

    // Constructing the dump object
    const dump = {
      lastUpdate: new Date().getTime(),
      anime: [...this.anime]
    };

    // Constructing the dump file path
    const path = resolve(this.dumpLocation, 'dump.json');

    // Treating the recursive sub-foldering
    if (!existsSync(this.dumpLocation) && this.dumpLocation) {
      mkdirSync(this.dumpLocation, { recursive: true });
    }

    // Dumping the object into a file
    writeFileSync(path, JSON.stringify(dump, null, 2));
  }

  //#endregion
}

const client = new AnusicScrapper({ dumpLocation: 'dir1/dir2' });

client.getAnimeList()
  .then(() => {
    client.createDump();
  });
