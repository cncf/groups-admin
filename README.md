# groups-admin CLI

A command-line tool for administering groups on a groups.io server. It authenticates as a user, lists subscriptions, exports members, inspects users, transfers ownership permissions, and shows pending messages.

## Installation
- Prerequisite: Go 1.22+
- Build with Makefile: `make build` (outputs `go/groups-admin`)
- Or manually: `cd go && go build -o groups-admin .`

## Quick Start
```
cd go
./groups-admin --baseUrl https://groups.io \
  --srcEmail you@example.com --srcPass 'your-password' \
  --cmd srcUserSubs --filter 'k8s-.*'
```

Using the Makefile:
```
make build
make run ARGS="--baseUrl https://groups.io \
  --srcEmail you@example.com --srcPass 'your-password' \
  --cmd srcUserSubs --filter 'k8s-.*'"

# Discover targets and examples
make help
```

## Commands
- `srcUserSubs`: Print the authenticated user’s group subscriptions (optionally filtered by `--filter` regex).
- `members`: Export members of all subscribed groups to `members.psv` (pipe-separated: group | name | email).
- `getUser`: Lookup a user by `--destEmail` and print details.
- `xferSubs`: Grant Owner permissions in your groups to `--destEmail` (optionally filter with `--filter`).
- `pendMsgs`: List pending messages for moderation on the main group.
- `remove-member`: Remove a member from all your groups (optionally filtered by `--filter` regex).

## Flags
- `--baseUrl`: Base URL of your groups.io instance (e.g., `https://groups.io`).
- `--srcEmail`, `--srcPass`: Credentials of the source (authenticated) user.
- `--cmd`: One of `srcUserSubs|getUser|xferSubs|members|pendMsgs`.
- `--destEmail`: Target user email for `getUser`/`xferSubs`/`removeMember`.
- `--filter`: Regular expression to select groups (e.g., `'k8s-.*'`).

## Examples
- Export members to file:
  `./groups-admin --baseUrl https://groups.io --srcEmail you@example.com --srcPass '...' --cmd members`
- Transfer owner perms for filtered groups:
  `./groups-admin --baseUrl https://groups.io --srcEmail you@example.com --srcPass '...' --cmd xferSubs --destEmail new-owner@example.com --filter 'k8s-.*'`

## Notes & Safety
- Never share real credentials in issue reports or examples.
- Operations like `xferSubs` change permissions; review the group list (use `--filter`) before proceeding.
- Default flag help: run `./groups-admin -h` for available flags.
