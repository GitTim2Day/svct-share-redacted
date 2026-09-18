# RUN — svct-share-redacted

**Fail-closed make-functional note (2026-09-18).** Docs + small demos only. No package lock. Sealed marker `EIGHT_PILLARS_SEALED_2026-09-13.md` is left sealed (not opened for bypass).

## Requirements

See `requirements.txt`: **no third-party pip packages** earned from `*.py` imports (stdlib only).

## Python demos (stdlib)

```bash
python3 svct_boot.py
python3 svct_bloom_chain.py
python3 -c "from svct_symbol_map import speak_symbols; print(speak_symbols('a::b->c'))"
```

## Go self-test (optional)

```bash
go run validation_module.go
```

Requires a Go toolchain. Module file (`go.mod`) is **ABSENT** — fail-closed, not invented here.

## C++ stubs

`svct_boot.cpp` / `svct_symbol_map.cpp` are stubs. No build recipe claimed in-repo.

## Docs entry

Start at `YOUR_CHECKLIST.md` and `docs/SVCT_Share_Bundle_xAI_REDACTED.txt`.

## Not in scope

- Do not invent `kbld9_core_r4` or Base360 CSV here.
- Do not send mail from this pass.
