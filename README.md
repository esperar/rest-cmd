# rest-lic

**rest-lic** is an open-source RESTful API server built in Go, enabling users to execute Linux commands via HTTP requests. It processes commands, captures their output, and returns the results in JSON format, making it easy to integrate Linux command execution in other applications or workflows.

## Features

- Execute Linux commands through RESTful API endpoints.
- Receive JSON-formatted responses with command output and error messages.
- Simple setup and usage via Makefile.

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/rest-lic.git
   cd rest-lic
   ```

2. **Build and run using Make**:
   ```bash
   make run
   ```

   This command will start the server on port `8080` by default.

## Usage

Once the server is running, you can execute Linux commands via HTTP POST requests to the `/exec` endpoint.

### Example Request

```bash
curl -X POST http://localhost:8080/exec -d '{"command": "ls -l"}' -H "Content-Type: application/json"
```

### Example Response

```json
{
  "output": "total 4\n-rw-r--r-- 1 user group 0 Oct 1 12:34 file.txt\n",
  "error": ""
}
```

- **`output`**: Standard output of the command.
- **`error`**: Error message, if any.

### Configuration

- Customize server settings by modifying environment variables or config files (details can be added here as your project evolves).

## Contribution

Feel free to fork this repository, create issues, or submit pull requests! For major changes, please open an issue first to discuss your ideas.

## License

Distributed under the MIT License. See `LICENSE` for more information.

