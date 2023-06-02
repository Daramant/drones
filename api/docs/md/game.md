## The rules

You have a maze with 2 drones, one of them is yours.
The drones start in the start point, for example {1,1} (🛸).
The winner is a player who leads his/her drone first to the goal, for example {22,1} (🥅).

```
 22↑ ████████████████████████
 21│ ██     █          █   ██
 20│ ██ ███ █ ██████ █ █ ████
 19│ ██   █     ██   █ █   ██
 18│ ████ █████ ██ ███████ ██
 17│ ██   █     ██ █       ██
 16│ ██████ ██████ █ █ ██████
 15│ ██   █     ██   █     ██
 14│ ██ █ ███ █ ██████████ ██
 13│ ██ █     █ ██     █ █ ██
 12│ ██ █████ ████ █ █ █ █ ██
 11│ ██     █   ██ █ █ █   ██
 10│ ██████ ███ ████ █ ███ ██
  9│ ██   █   █ ██   █     ██
  8│ ██ █ ███ █ ██ █████ ████
  7│ ██ █     █ ██   █   █ ██
  6│ ██ ███████ ████ █ ███ ██
  5│ ██ █       ██   █     ██
  4│ ████ ████████ ██████████
  3│ ██   █ █   ██ █     █ ██
  2│ ██ █ █ █ █ ██ █████ █ ██
  1│ █🛸█     █ ██        🥅█
  0│ ████████████████████████
   └────────────────────────→
     0123456789...
```

## How to start a game

1. Call the method **[join](#/RPC%20methods/join_v1)**.
2. Call the method **[wait_turn](#/RPC%20methods/wait_turn_v1)** to receive the current game state.
3. Call the action method (you have 5 seconds).
4. Go to step 2.

For more information see methods description below.

## The source code and bot examples

You can find them in the [GitHub repository](https://github.com/bot-games/drones).

## LocalRunner

See [README](https://github.com/bot-games/drones/tree/master/cmd/localrunner)