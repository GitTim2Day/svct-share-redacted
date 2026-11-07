# Ollama No-Sudo Install (aarch64 / Pi 5 shepherd)

**Date:** 2026-11-07
**Device:** Raspberry Pi 5 8GB (shepherd), aarch64
**Rule:** No sudo. Static binary only. Local-first.
**Pointer skill:** ollama-nosudo-install-2026-11-07
**Ledger:** 1a1b7016d43befc3

## Steps (validated on shepherd)

1. Download static arm64 binary:
   ```bash
   curl -fsSL -o ollama https://ollama.com/download/ollama-linux-arm64
   ```

2. Make executable:
   ```bash
   chmod +x ollama
   ```

3. Move to user-writable bin:
   ```bash
   mkdir -p ~/.local/bin
   mv ollama ~/.local/bin/
   ```

4. Append PATH to bashrc:
   ```bash
   echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
   ```

5. Source and verify:
   ```bash
   source ~/.bashrc
   ollama --version
   ```

## Model location

- `OLLAMA_MODELS=~/.ollama/models`
- USB-copied llama3.2:3b (2.0 GB Q4_K_M) loaded from A15 via `_RADIOACTIV` USB stick.
- Transfer path: A15 `~/.ollama/models` → USB F:\ollama-models → Pi `/media/shepherd/_RADIOACTIV/ollama-models` → `~/.ollama/models`

## Notes

- Do not use `curl -fsSL https://ollama.com/install.sh | sh` — that requires root.
- Sibling of pi5-model-adapter-gosub.
- Install script avoided per no-sudo standing rule.
- Validated: binary runs, model list shows llama3.2:3b.

## Related

- pi5_model_adapter_gosub.py (SHA-256 e6f5a63b5a9f47c4e8a5848a9bf55c5edc00eb5c172812f523c561d9f387984e)
- SESSION_PATTERN_MAP_2026-11-07_PI5_MODEL_ADAPTER_GOSUB.md (SHA-256 87f283b6caa192183fd5d61d0dc2c0c0da2ab90ac79349512e362ba4972aa812)
- Pointer skill: session-map-2026-11-07-pi5-model-adapter-gosub