set -o vi

parse_git_branch() {
    git branch 2>/dev/null | grep '^*' | sed 's/^* //'
}

export PS1="\[\033[1;37m\]\h\[\033[1;90m\]|\[\033[31m\]\u \[\033[90m\]\w \[\033[35m\]\$(parse_git_branch)\[\033[00m\] \n\[\033[90m\]\t \[\033[00m\]$ "

