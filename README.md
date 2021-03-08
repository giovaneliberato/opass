
![](opass_logo.png)
# OPass - Improved 1Password CLI for unix/pass lovers.

OPass is a wrapper CLI over official 1Password CLI. Its usage is very much inspired in the unix password store which has a killer usability for terminal users.

It aims to provide pass and 1Password users a familiar experience without learning (yet) another tool. 

Important note: OPass do not manage any sensitive information itself (beside the private key, which is saved on `$HOME/.opass/config`). The sign in is handled direct by the 1Password CLI and your data is fetched from the server and displayed directly in the terminal. 

:warning:⠀⠀`This projects has not been peer reviewed. Use at you own risk.`⠀⠀:warning: 


## Usage
```
Usage: opass <command>

Flags:
  -h, --help        Show context-sensitive help.
  -c, --copy        Copy password to clipboard.
  -a, --list-all    List all tags and items.

Commands:
  <tag-or-login>
    If a Tag name is given, list all logins under that tag. If an Item name is given, show details.

  list
    List all tags of account.

  config
    Initiate 1Password credentials configuration.

  signin
    Signin to 1Password using predefined credentials.

  flush
    Drop local list of items and sync with 1Password account. Useful after you update information on another device.

Run "opass <command> --help" for more information on a command.
```

## Getting started

#### List all tags

OPass uses the tags defined in the items to organize the sections in the tree. Untagged items go under the sescion `untagged`.

```
opass

1Password
└── finance
└── social
└── tech
└── untagged
```

#### List logins of a tag 
```
opass tech

1 Password
└── tech
    └── github
    └── gitlab
    └── vpn
```

#### Copy password to clipboard 
```
opass -c tech/vpn
Password copied to clipboard.
```

#### Get logins details 
```
opass tech/vpn
{
  "Username": "giovane",
  "Password": "<password>",
  "URL": "",
  "UpdatedAt": "2020-01-12T18:07:46Z",
  "ItemVersion": 2,
  "Tags": [
    "tech"
  ]
}
```

## Installing

1. Install the [1Password CLI](https://app-updates.agilebits.com/product_history/CLI)

2. Clone repo and compile
```
git clone git@github.com:giovaneliberato/opass.git
cd opass
make install
```
3. Setup account
```
opass config
Sign in Address: https://my.1password.com
Email Address: me@example.com
Private Key: <PRIVATE_KEY>
Configuration file created at $HOME/.opass/config
```
4. **You are good to go!**
