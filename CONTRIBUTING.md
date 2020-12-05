# How to contribute

Welcome to the Anusic API repository, we're glad that you're offering your contributions.

Issues first, PRs later. That's our motto here unless it's for fixing silly typos, we love documenting the process in which this API evolves.

## Testing

More testing is always good, dare we say, more appreciated.

## Submitting changes

Please amend your commit messages with a clear list of what you've done.
When you send a pull request, we will love you forever if you include RSpec examples.
We can always use more test coverage.
Please follow our coding conventions (below) and make sure all of your commits are atomic (one feature per commit).

Always write a clear log message for your commits. One-line messages are fine for small changes, but bigger changes should look like this:

    $ git commit -m "A brief summary of the commit
    > 
    > A paragraph describing what changed and its impact."

## Coding conventions

Start reading our code and you'll get the hang of it. We optimize for readability:

  * We indent using 1 tab (2 spaces)
  * Always prettify your code
  * Local functions and variables should be camelCased
  * Exported functions and variables should be PascalCased
  * Don't be shy to add comments, on top of structs, functions, variable declarations...
  * Code readability over complexity, it's okay to deliver semantic code over clever or hacky snippets.
  * This is open-source software. Consider the people who will read your code, and make it look nice for them.

Thanks,
EOussama.
