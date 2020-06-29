import * as config from "./config.json";
import axios from "axios";

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
      .then((response) => {
        console.log(response);
      });
  }

  //#endregion
}

new AnusicScrapper().getAnimeList();
