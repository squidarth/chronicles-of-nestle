kchronicles-of-nestle
====================

A command-line tool and web app to annotate your bash history.

### Bash vs. ZSH

This works in zsh out of the box, but for bash, requires that you
do 

`shopt -s histappend PROMPT_COMMAND="history -a;$PROMPT_COMMAND"`
