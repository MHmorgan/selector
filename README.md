Selector
========

Simple terminal application which lets the user select between its input arguments and prints out the the selected choice.

Intended to be used in shell functions and aliases, such as:

```Zsh
alias goto='cd $(selector $HOME/*)'

# or

function goto {
	local DIR=$(selector $HOME/*)/$1
	cd $DIR
	ls
}
```
