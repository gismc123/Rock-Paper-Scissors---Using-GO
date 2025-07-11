# Rock Paper Scissors - Go CLI Game

A simple command-line Rock Paper Scissors game written in Go.  
Play against the computer (bot) with the first to 5 wins declared the champion!

---

## Table of Contents

- [About](#about)  
- [Features](#features)  
- [Installation](#installation)  
- [Usage](#usage)  
- [How It Works](#how-it-works)  
- [Contributing](#contributing)  
- [License](#license)  

---

## About

This project is a beginner-friendly Go application that demonstrates basic user input handling, random number generation, and game logic implementation. It’s a fun way to practice Go programming and CLI interaction.

I developed this as part of my ongoing exploration of Go and to create a simple, interactive game that can be run on any system with Go installed. It’s also a neat example for anyone learning Go or interested in game programming basics.

---

## Features

- Interactive command-line interface  
- Input validation (accepts only "rock", "paper", or "scissors")  
- Randomized bot choices  
- Score tracking for player and bot  
- Game ends when either player or bot reaches 5 wins  
- Option to exit anytime by typing `exit`  

---

## Installation

1. Make sure you have [Go](https://golang.org/dl/) installed (version 1.18+ recommended).  
2. Clone this repository:  
   ```bash
   git clone https://github.com/gismc123/Rock-Paper-Scissors---Using-GO.git
   ```  
3. Navigate into the project directory:  
   ```bash
   cd rock-paper-scissors
   ```  
4. Run the game:  
   ```bash
   go run main.go
   ```  

---

## Usage

- When prompted, type your choice: `rock`, `paper`, or `scissors`.  
- The bot will randomly pick its choice.  
- The winner of the round is displayed, and scores are updated.  
- The game continues until either you or the bot reaches 5 points.  
- Type `exit` anytime to quit the game early.  

---

## How It Works

- The program seeds the random number generator with the current time to ensure different bot choices each run.  
- User input is read from the command line, trimmed, and converted to lowercase for consistency.  
- The bot randomly selects from the three options.  
- Game logic compares choices to determine the winner of each round.  
- Scores are tracked and displayed after each round.  
- The game loop continues until a player reaches 5 points or the user exits.  

---

## Contributing

Contributions are welcome! If you want to add features, fix bugs, or improve documentation:

1. Fork the repository  
2. Create a new branch (`git checkout -b feature-name`)  
3. Commit your changes (`git commit -m 'Add some feature'`)  
4. Push to the branch (`git push origin feature-name`)  
5. Open a Pull Request  

Please ensure your code follows Go best practices and includes comments where necessary.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

*Happy coding and enjoy the game!* 🎮🎸

