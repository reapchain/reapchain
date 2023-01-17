Thank you for considering making contributions to reapchain!

Contributing to this repo can mean many things such as participated in discussion or proposing code changes. To ensure a smooth workflow for all contributors, the general procedure for contributing has been established:

  1. either [open](https://github.com/irisnet/irishub/issues/new) or
     [find](https://github.com/irisnet/irishub/issues) an issue you'd like to help with,
  2. participate in thoughtful discussion on that issue,
  3. if you would then like to contribute code:
     1. if the issue is a proposal, ensure that the proposal has been accepted,
     2. ensure that nobody else has already begun working on this issue, if they have
       make sure to contact them to collaborate,
     3. if nobody has been assigned the issue and you would like to work on it
       make a comment on the issue to inform the community of your intentions
       to begin work,
     4. follow standard github best practices: fork the repo,
       if the issue if a bug fix, branch from the
       head of `main`, make some commits, and submit a PR to `main`; if the issue is a new feature,  branch from the tip of `feature/XXX`, make some commits, and submit a PR to `feature/XXX`
     5. include `WIP:` in the PR-title to and submit your PR early, even if it's
       incomplete, this indicates to the community you're working on something and
       allows them to provide comments early in the development process. When the code
       is complete it can be marked as ready-for-review by replacing `WIP:` with
       `R4R:` in the PR-title.

Note that for very small or blatantly obvious problems (such as typos) it is 
not required to an open issue to submit a PR, but be aware that for more complex
problems/features, if a PR is opened before an adequate design discussion has
taken place in a github issue, that PR runs a high likelihood of being rejected.

Other notes:

- Please make sure to run `make format` before every commit - the easiest way
  to do this is have your editor run it for you upon saving a file. Additionally
  please ensure that your code is lint compliant by running `golangci-lint run`.

## Pull Requests

To accommodate review process we suggest that PRs are categorically broken up.
Ideally each PR addresses only a single issue. Additionally, as much as possible
code refactoring and cleanup should be submitted as a separate PRs. And the feature branch `feature/XXX` should be synced with `develop` regularly.

The following PR structuring checklist can be used when submitting changes to the Gaia repository for review:

- [ ] Proto files: PR updating proto files. As a suggested next step, don't regenerate updated protobuf
  implementations using `protgen`, since this will break existing code.
- [ ] Broken code: If `protogen` is run, a PR disabling broken code
- [ ] Validation: PR with validation of types
- [ ] Functionality: PR integrating supporting functionality
- [ ] Servers: PR for `msgserver` and `queryserver`
- [ ] CLI: PR for CLI commands
- [ ] Orchestrators: PR for any orchestrators
- [ ] Genesis: PR for genesis
- [ ] Upgrades: PR for upgrades

## Testing

The `Makefile` defines `make test` and includes its continuous integration. For any new comming feature, the `test_unit` / `test_cli` and `test_lcd` must be provided.

We expect tests to use `require` or `assert` rather than `t.Skip` or `t.Fail`,unless there is a reason to do otherwise.

### PR Targeting

Ensure that you base and target your PR on the correct branch:

- `release/vxx.yy` for a merge into a release candidate
- `main` for a merge of a release
- `develop` in the usual case

All feature additions should be targeted against `feature/XXX`. Bug fixes for an outstanding release candidate
should be targeted against the release candidate branch. Release candidate branches themselves should be the
only pull requests targeted directly against master.

### Development Procedure

- the latest state of development is on `main`
- `main` must never fail `make test` or `make test_cli`
- `main` should not fail `make lint`
- no `--force` onto `main` (except when reverting a broken commit, which should seldom happen)
- create a development branch either on `https://github.com/reapchain/reapchain`, or your fork (using `git remote add origin`)
- before submitting a pull request, begin `git rebase` on top of `main`

### Pull Merge Procedure

- ensure pull branch is rebased on `main`
- run `make test` and `make test_cli` to ensure that all tests pass
- merge pull request

