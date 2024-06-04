#!/bin/bash

_evi_completions() {
	local cur prev opts
	COMPREPLY=()
	cur="${COMP_WORDS[COMP_CWORD]}"
	prev="${COMP_WORDS[COMP_CWORD-1]}"
	opts="-k -key -no-decrypt -no-edit -no-encrypt -help"
	doubleopts="--key --no-decrypt --no-edit --no-encrypt --help"

	if [[ ${cur} == --* ]]; then
		COMPREPLY=( $(compgen -W "${doubleopts}" -- ${cur}) )
		return 0
	fi
	if [[ ${cur} == -* ]]; then
		COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
		return 0
	fi

	COMPREPLY=( $(compgen -f -- ${cur}) )
	return 0
}

complete -F _evi_completions evi
