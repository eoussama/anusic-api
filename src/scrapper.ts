
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
