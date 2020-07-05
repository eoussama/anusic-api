import { ITheme } from './theme.model';

/**
 * The Anime
 */
export interface IAnime {

  /**
   * The ID of the Anime
   */
  id: string;

  /**
   * The name of the Anime
   */
  name: string;

  /**
   * ALT name of the Anime
   */
  altName?: string;

  /**
   * The themes of the Anime
   */
  themes: ITheme[];
}

export class Anime implements IAnime {

  //#region Properties

  id: string;
  name: string;
  year: number;
  altName?: string;
  themes: ITheme[];

  //#endregion

  //#region Lifecycle

  constructor(anime: IAnime) {
    const yearIndex: number = anime.name.lastIndexOf('(');

    this.id = anime.id;
    this.name = anime.name.substring(0, anime.name.lastIndexOf('(')).trim();
    this.year = parseInt(anime.name.substring(yearIndex).match(/\d+/gi).join(''), 10);
    this.altName = anime.altName;
    this.themes = anime.themes;
  }

  //#endregion

  //#region Methods

  private fetchInfo(): void { }

  //#endregion
}