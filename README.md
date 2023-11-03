# Intro
- A simple CLI pomodoro tool written in Go

# Usage
## CLI-mode
### Start a pomodoro
`godoro start`
- When used without time flags, `godoro start` begins a stopwatch. This stopwatch keeps running until manually stopped by the user
- When used with time flags, `godoro start` initiates a timer with a predefined duration. The timer will run for the specified duration, after which it will automatically stop

#### Flags
`-s [arg]`
- Specify seconds

`-m [arg]`
- Specify minutes

`-h [arg]`
- Specify hours

`-t [arg1 arg2 ...]`
- Specify tags for the current pomodoro session

### View history
`godoro history`

## TUI mode
`godoro tui`

### View all sessions
- Type 'a'

### Create a new pomodoro
- Type 'n'
- Enter the name and any tag info

### Pause a pomodoro
- Type 's'
- Enter the name of the session you want to pause

### Delete a pomodoro
- Type 'd'
- Enter the name of the session you want to delete

