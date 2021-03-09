#/usr/bin/env bash


_opass_completions() {
  local tags=$(awk -F '=' '/\[tags\]/,0 {print $1}' $HOME/.opass/cache \
              | awk 'NR!=1' \
              | sort \
              | tr -d '[[:blank:]]' \
              | tr '\n' ' ')

  local tagged_items=$(awk -F '=' '/\[tags\]/,0 {tags[$1]=$2} END \
                                { \
                                  for (tag in tags) { \
                                    split(tags[tag], items, ",")
                                    for (i in items) { \
                                      print(sprintf("%s\/%s", tag, items[i])) \
                                    } \
                                  } \
                                }' $HOME/.opass/cache \
                      | sort \
                      | tr -d '[[:blank:]]' \
                      | tr '\n' ' ')


  if [[ "${COMP_WORDS[1]}" == *\/ ]]; then
    COMPREPLY+=($(compgen -S / -W "$tagged_items"))
  else
    COMPREPLY+=($(compgen -S / -W "$tags"))
  fi
}

complete -o nospace -F _opass_completions opass
