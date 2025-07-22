# Contributing to Snake Game

First off, thank you for considering contributing to the Snake Game project! It's people like you that make this project better.

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues to see if the problem has already been reported. When you are creating a bug report, please include as many details as possible:

- **Use a clear and descriptive title** for the issue
- **Describe the exact steps to reproduce the problem**
- **Provide specific examples to demonstrate the steps**
- **Describe the behavior you observed after following the steps**
- **Explain which behavior you expected to see instead and why**
- **Include details about your configuration and environment**

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, please include:

- **Use a clear and descriptive title** for the issue
- **Provide a step-by-step description of the suggested enhancement**
- **Provide specific examples to demonstrate the steps**
- **Describe the current behavior** and **explain which behavior you expected to see instead**
- **Explain why this enhancement would be useful**

### Pull Requests

The process described here has several goals:

- Maintain the project's quality
- Fix problems that are important to users
- Engage the community in working toward the best possible Snake Game
- Enable a sustainable system for maintainers to review contributions

Please follow these steps to have your contribution considered by the maintainers:

1. Follow all instructions in the template
2. Follow the styleguides
3. After you submit your pull request, verify that all status checks are passing

## Development Process

### Setting Up Development Environment

1. Fork the repository on GitHub
2. Clone your fork locally:
   ```bash
   git clone https://github.com/yourusername/snake.git
   cd snake
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Create a topic branch:
   ```bash
   git checkout -b feature/my-new-feature
   ```

### Making Changes

1. Make your changes in your feature branch
2. Add or update tests as needed
3. Run the test suite to ensure nothing is broken:
   ```bash
   make test
   ```
4. Run formatting and linting:
   ```bash
   make fmt
   make vet
   ```
5. Test the game manually to ensure it works as expected

### Committing Changes

1. Write good commit messages:

   - Use the present tense ("Add feature" not "Added feature")
   - Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
   - Limit the first line to 72 characters or less
   - Reference issues and pull requests liberally after the first line

2. Example commit message:

   ```
   Add multiplayer support for Snake game

   - Implement network communication between players
   - Add lobby system for matchmaking
   - Update game loop to handle multiple snakes

   Closes #123
   ```

### Submitting Changes

1. Push your changes to your fork:
   ```bash
   git push origin feature/my-new-feature
   ```
2. Create a pull request on GitHub
3. Wait for review and address any feedback

## Styleguides

### Go Styleguide

- Follow the official [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Use `gofmt` to format your code
- Run `go vet` to catch common errors
- Use meaningful variable and function names
- Write comments for exported functions and types
- Keep functions small and focused on a single responsibility

### Git Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests liberally after the first line
- Consider starting the commit message with an applicable emoji:
  - 🎨 `:art:` when improving the format/structure of the code
  - 🐎 `:racehorse:` when improving performance
  - 📝 `:memo:` when writing docs
  - 🐛 `:bug:` when fixing a bug
  - 🔥 `:fire:` when removing code or files
  - ✅ `:white_check_mark:` when adding tests
  - 🔒 `:lock:` when dealing with security
  - ⬆️ `:arrow_up:` when upgrading dependencies
  - ⬇️ `:arrow_down:` when downgrading dependencies

## Code Architecture Guidelines

### Package Organization

- Follow the existing package structure
- Keep packages focused on a single responsibility
- Use the `internal/` directory for packages that shouldn't be imported by other projects
- Place the main application entry point in `cmd/`

### Error Handling

- Return errors rather than panicking
- Wrap errors with context using `fmt.Errorf`
- Handle errors at the appropriate level
- Use sentinel errors for expected error conditions

### Concurrency

- Use channels for communication between goroutines
- Avoid shared memory where possible
- Use proper synchronization primitives when shared memory is necessary
- Always clean up goroutines properly

### Testing

- Write tests for new functionality
- Use table-driven tests where appropriate
- Test error conditions
- Aim for good test coverage
- Use meaningful test names that describe what is being tested

## Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [How to Write a Git Commit Message](https://chris.beams.io/posts/git-commit/)

Thank you for contributing to the Snake Game project!
