# Функция для отображения ветки Git
parse_git_branch() {
  if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    branch=$(git symbolic-ref --short HEAD 2>/dev/null || git describe --tags --exact-match 2>/dev/null)
    # Фиолетовый цвет и жирный текст для ветки
    echo -e "\[\033[1;35m\] $branch\[\033[0m\]"
  fi
}

# Настройка PS1 (без export, чтобы работало в WSL)
PS1="\[\033[1;37m\]\h\[\033[1;90m\]|\[\033[31m\]\u \[\033[90m\]\w \$(parse_git_branch) \n\[\033[90m\]\t \[\033[0m\]\$ "
