# **Pokedex CLI in Go** 🐾

This project is a Command-Line Interface (CLI) Pokedex built using the Go programming language. It interacts with the Pokémon API (PokeAPI) to fetch data about Pokémon, their locations, abilities, and more. With this project I learned essential Go concepts such as HTTP requests, data parsing, and handling user inputs in a REPL (Read-Eval-Print Loop) environment.

---

## **Features**

- **Browse map locations**  
  Browse over 1000 locations in the pokemon world.

- **Catch Pokémon**  
  Simulate catching Pokémon with a random chance based on their base experience.

- **Pokedex**  
  Display your catch pokemon in your pokedex.
  
- **Interactive REPL (Read-Eval-Print Loop)**  
  Engage with the Pokedex in an interactive shell environment.

- **Caching**  
  Speed up subsequent searches by caching results locally.

- **HTTP Integration**  
  Fetch live Pokémon data from a public API.

- **Custom Commands**  
  Use specialized commands for advanced queries or help options.

---

## **Goals**

By the end of this project, you will:

1. Understand and make **HTTP requests** in Go.
2. Implement a **REPL interface**.
3. Learn how to **cache API responses**.
4. Enhance your understanding of Go's **standard library**.
5. Build a fully functional command-line tool.

---

## **Setup**

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/pokedex-cli.git
   ```
2. Navigate to the project directory:
   ```bash
   cd pokedex-cli
   ```
3. Install dependencies (if any):
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```

---

## **Usage**

- **Start the Pokedex CLI:**  
  Run the application and follow the interactive prompts.

- **Available Commands:**  
  - `map`: Get the next page of locations.
  - `mapb`: Get the previous page of locations.
  - `explore`: Get the next page of locations.
  - `catch`: Attempt to catch a pokemon.
  - `inspect`: View details about a caught Pokemon.
  - `pokedex`: See all the pokemon you've caught.
  - `help`: List all available commands.
  - `exit`: Exit the application.

---

## **Technologies Used**

- **Go**: The primary programming language.
- **HTTP Requests**: To fetch Pokémon data.
- **REPL**: For interactive command handling.
- **Caching**: To optimize performance.

---

## **Contributing**

This is a personal project to learn go, but feel free to clone the repository.

---

### **Author**

Developed by **Jaime**  
*“Gotta code ‘em all!”*