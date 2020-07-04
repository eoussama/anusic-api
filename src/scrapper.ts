import { IAnime, Anime } from "./models/anime.model";

/**
 * Gets the category list
 * @param $ Cherrio object
 */
export function getCategories($: CheerioStatic): string[] {
  const categories: string[] = [];

  $('.toc li').each((index, element) => {
    categories.push($(element).text().toLowerCase());
  });

  return categories;
}

/**
 * Gets the anime list
 * @param $ Cherrio object
 */
export function getAnimeList($: CheerioStatic): any {
  const categories: any = getCategories($);
  const animeList: Anime[] = [];

  categories.forEach((category: string) => {
    $(`#wiki_${category} ~ p`).each((index: number, element: CheerioElement) => {
      animeList.push(new Anime({
        id: '',
        name: $(element).text(),
        altName: '',
        themes: []
      }));
    });
  });

  return animeList;
}
