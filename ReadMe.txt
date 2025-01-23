To Seralize the struct we are using : clean env package

Link: go get -u github.com/ilyakaznacheev/cleanenv

# Config Loader for Go Applications

This project provides a configuration loader for Go applications using the `cleanenv` library. The loader supports configuration via YAML files and environment variables, ensuring a clean and flexible setup for different environments.

## Features
- **YAML Config Parsing**: Reads configuration from a YAML file.
- **Environment Variables**: Overrides or supplements config values with environment variables.
- **Validation**: Ensures required fields are provided.
- **Command-Line Support**: Accepts configuration file path via the `--config` flag.

## Configuration Structure

### YAML File Example (`config.yaml`)
```yaml
env: production
storage_path: /data/storage
http_server:
  address: 0.0.0.0:8080
