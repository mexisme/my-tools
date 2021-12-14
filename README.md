# My Tools Monorepo

Welcome!

This repo is a Monorepo.

In [docs/MONOREPO.md](docs/MONOREPO.md) is some more details about what a Monorepo is, and how we're (currently) managing it.

## The basic structure

NOTE: the `//` symbol is "sugar" to designate the root of the entire repo, independently of where it's been checked-out.

### Source Code

`//third_party/`
- Source code and tools from Third Parties that we use. This should normally be source code, as storing binary data in here will bloat the repo rapidly.
- There is a [//third_party/README.md](third_party/README.md) in the root of this directory detailing the directory hierarchy, all the Third Party code, where it came from, versioning, etc.
- Code may be imported from external sources with tools like Git Subrepo, Subtree, Strees, Submodule, etc, and should also go in here.


`//my/`
- This is the root of all source code.
- There is a [//my/README.md](my/README.md) in the root of this directory detailing the directory hierarchy, and where various apps and libraries can be found.

### Documentation:

`//docs/`
- Information and documentation describing the repo --- not documentation for our code or apps, but describing how to use the Monorepo, and tools for making it work for you.

`//third_party/README.md`
`//third_party/.../README.${COMPONENT}.md`
In the root of `//third_paty` is a `README.md` describing all the third party tools, where they came
from, and the command(s) used to get them into the repo.
Some tools require more details about how they're built for our use, but you want to keep the tool's
directory consistent with upstream.
For those, a `README.${COMPONENT}.md` file in the directory just above should be created with this
additional detail.

`//my/docs/`
- The root documentation. Any generalised or non-app / non-library specific documentation should go in here.

`//my/.../${COMPONENT}/README.md`
- Every standalone component --- e.g. an app or library --- should have a `README.md` file in its root describing the component.
  By preference

`//my/.../${COMPONENT}/`
`//my/.../${COMPONENT}/docs/`
- Wherever possible, try to use in-code "doc comment" style for documentation (e.g. JSDoc, GoDoc, etc).
  This style allows for long-form human-readable documentation as well as function-level, etc.
- Any documentation that can't be stored in code should be stored in the component's `README.md`.
  It's not usually needed to subdivide into multiple files, as that won't necessarily help findability in the same way it does for code.
- Use the doc-generation tooling to extract doc comments to either HTML or Markdown formats.
  The post-generated documentation can be stored in the `docs/` directory of and checked-in for others to read, if necessary, but please make sure not to update it in-place, to avoid any new comments being overwritten by another generation step.

### Build and Release

`//.github/CODEOWNERS`
- This is a GitHub metadata file describing the "owners" for various parts of the code hierarchy.

`//.gitlab-ci.yml`
- The configuration for the Gitlab CI Pipeline(s).

`//.editorconfig`
- An editor and IDE-agnostic config for describing basic code-structure for various languages, incl. whether to use Tabs or Spaces.
