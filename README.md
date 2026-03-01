# ccgen - AI-powered Conventional Commit Generator

`ccgen` is a lighweight CLI tool written in Go that automatically generates 
[Convenvtional Commits](https://www.conventionalcommits.org/en/v1.0.0/) messages based on your staged `git diff`

It leverages Google's Gemini API via the official Go SDK (`google.golang.org/genai`) and provied a simple, clean interface built with 
[Cobra](https://cobra.dev/)

## Why ccgen?
Writing good Conventional Commit messages takes time and discipline, `ccgen`helps you:
- Generate properly formatted Conventional Commit titles
- Optionally include a full commit body
- Work directly from your staged changes (`git diff --staged`)
- Exclude selected files from the generated diff
- Keep your history clean and consistent with minimal effort

## Instalation

### Option 1: Download a Release
1. Go to the [Releases page](https://github.com/b14ucky/ccgen/releases)
2. Download the latest version compatible with your OS.
3. Add the binary to your `$PATH`

After that, you can use `ccgen` from anywhere in your terminal.

### Option 2: Build from Source
Make sure you have Go installed \
Installation instructions: https://go.dev/doc/install

Then clone and install:
```bash
git clone https://github.com/b14ucky/ccgen
cd ccgen
go install
```

This will build and install `ccgen` into your Go binary directory (usually `$GOPATH/bin` or `$HOME/go/bin`)

## Configuration
`ccgen` requires a configuration file with your Gemini API credentials.

By default, it looks for:
```bash
$HOME/.ccgen.yaml
```
You can specify a custom config file using:
```bash
ccgen --config path/to/config.yaml
```

### Example Configuration
```yaml
api_key: your_api_key
model: gemini-2.5-flash
```
- `api_key` - Your Gemini API key \
  You can generate one here: https://aistudio.google.com/app/api-keys
- `model` - The Gemini model to use (e.g. `gemini-2.5-flash`)

## Usage

### Generate a commit title only
```bash
ccgen commit
```
This reads the output of:
```bash
git diff --staged
```
and generates a Conventional Commit title.

### Generate a full commit message (title + body)
```bash
ccgen commit -d
```
or
```bash
ccgen commit --description
```
### Exclude specific files from the diff
```bash
ccgen commit file1.go file2.go
```
The provided file names will be excluded from the staged diff before sending it to the AI model.
### Enable verbose mode
```bash
ccgen --verbose commit
```
Shows additional information such as the selected model and config file.

## Shell Completions
To generate the autocompletion script for your shell:
```bash
ccgen completion [bash|fish|powershell|zsh]
```
For detail's on how to use the generated script see sub-command's help.

## How It Works
1. `ccgen` reads your staged changes using `git diff --staged`.
2. The diff is sent to the configured Gemini model.
3. The AI generates a Conventional Commit message.
4. The result is printed to stdout.

By default, only the commit title is generated. \
With `-d`, a full message including description is produced.

## License
This project is licensed under the Apache License 2.0 \
See the [LICENSE](LICENSE) file for details.
