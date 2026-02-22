When creating a commit message, ensure that it includes a precise and informative subject line that succinctly summarizes the crux of the changes in under 50 characters. If necessary, follow with an explanatory body providing insight into the nature of the changes, the reasoning behind them, and any significant consequences or considerations arising from them. Conclude with any relevant issue references at the end of the message. 

Make sure to include the scope, as per Conventional Commits specification! 

Make sure the body is separated from the subject line by a blank line. For more information, refer to the following resources:
- https://chris.beams.io/posts/git-commit/
- https://www.conventionalcommits.org/en/v1.0.0/.
Also, ensure the length of each line for the commit message is less than 100 characters.

**Breaking changes:**
- Do NOT use the exclamation mark (!) in the header (e.g., `feat(scope)!:`) as a breaking change indicator.
- Instead, add a line in the body starting with `BREAKING CHANGE:` followed by a description of the breaking change and its impact.
- Example:

  feat(my-scope): some breaking feature
  
  This feature changes the API in a way that is not backward compatible.
  
  BREAKING CHANGE: The API for X is now Y and requires Z.

  [JIRA: FES-1234]

Make sure to include the Jira issue number, taken from the branch name (for example, KD-1234 or KLD-12345), wrapped in square brackets and separated by one space.

Ensure each line for the commit message, including body, is less than 100 characters and is separated by a blank line.

Make sure the subject line is in the imperative mood, as if you are giving a command. For example, use "Fix" instead of "Fixed" or "Fixes". 

Make sure the subject line is in lowercase, as per the Conventional Commits specification.