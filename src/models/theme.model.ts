import { MediaFormat } from '../enums/media-format.enum';

/**
 * The theme
 */
export interface ITheme {

  /**
   * The name of the theme
   */
  name: string;

  /**
   * The link to the theme
   */
  link: string;

  /**
   * The episode(s) the theme is used in
   */
  episodes: string[];

  /**
   * The media format
   */
  format: MediaFormat;

  /**
   * The video resolution
   */
  resolution: string;

  /**
   * Whether the theme contains spoilers
   */
  hasSpoilers: boolean;

  /**
   * Whether the theme is NSFW
   */
  isNSFW: boolean;

  /**
   * Whether the theme has credits
   */
  isCreditless: boolean;

  /**
   * Whether the theme has lyrics
   */
  hasLyrics: boolean;

  /**
   * Whether the theme transitions from the episode
   */
  isTransition: boolean;

  /**
   * Whether the them overlays the episode
   */
  isOver: boolean;
}
