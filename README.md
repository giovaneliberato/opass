
![](opass_logo.png)
# OPASS - A Unix/Pass-like interface for the 1Password CLI.


:warning:⠀⠀`This projects has not been peer reviewed. Use at you own risk.`⠀⠀:warning: 


## Usage
```
Usage: opass <command>

Flags:
  -h, --help    Show context-sensitive help.
  -c, --copy    Copy password to clipboard.

Commands:
  <tag-or-login>
    If a Tag name is given, list all logins. If a Login is given, show details.

  list
    List all tags of account.

  config
    Initiate 1Password credentials configuration.

  signin
    Signin to 1Password using predefined credentials.

Run "opass <command> --help" for more information on a command.
```

## Getting started

To organize the sections in the tree, OPASS uses the tags defined in the login item. Untagged items go under the sescion `untagged`.

#### List all tags
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
make build
```

3. Copy executable to lib folder
```
cp opass /usr/local/bin
``` 

4. Setup account
```
opass config
Sign in Address: https://my.1password.com
Email Address: me@example.com
Private Key: <PRIVATE_KEY>
Configuration file created at $HOME/.opass/config
```

5. **You are good to go!**
