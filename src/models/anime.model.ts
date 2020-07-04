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
