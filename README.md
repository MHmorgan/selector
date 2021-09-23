Selector
========

Simple terminal application which lets the user select between its input arguments and prints out the the selected choice.  
The usage is simple and intuitive: use the arrow keys to move between values; press enter to select a value;
type anything to filter the values; press backspace to remove the last filter character, or delete to clear the entire filter text.

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
