# aka demo — generated with `vhs demo/demo.tape`
#
# Requires: https://github.com/charmbracelet/vhs
#
# `make demo` runs this file and produces demo.gif.

Output demo.gif

Set Shell "zsh"
Set FontSize 14
Set Width  1100
Set Height  640
Set Theme "Dracula"
Set Padding 20

Type "aka version"
Enter
Sleep 1s

Type "aka add cc claude"
Enter
Sleep 500ms

Type "aka add -d 'auto-approve' ccauto 'claude --dangerously-skip-permissions'"
Enter
Sleep 500ms

Type "aka add -t git gs 'git status'"
Enter
Sleep 500ms

Type "aka add -t git gl 'git log --oneline -20'"
Enter
Sleep 500ms

Type "aka ls"
Enter
Sleep 1500ms

Type "aka"
Enter
Sleep 2s

# filter
Type "/"
Sleep 300ms
Type "git"
Sleep 1500ms
Escape
Sleep 500ms

# help overlay
Type "?"
Sleep 2s
Enter
Sleep 500ms

Type "q"
Sleep 500ms

Type "aka rm -y gs"
Enter
Sleep 500ms

Type "aka ls"
Enter
Sleep 1500ms
