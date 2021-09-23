Selector
========

Simple terminal application which lets the user select between its input arguments and prints out the the selected choice.

Intended to be used in shell functions and aliases, such as:

```Zsh
alias goto='cd $(selector $HOME/*)'

function goto {
	local path=$(selector $HOME/*)/$1
	cd $path
	ls
}
```
