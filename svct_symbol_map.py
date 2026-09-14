import re
SYMBOL_MAP = {"::":"scope resolution","<<":"stream insertion",">>":"stream extraction","->":"arrow"}
def speak_symbols(text):
    toks = sorted(SYMBOL_MAP, key=len, reverse=True)
    pat = "|".join(re.escape(t) for t in toks)
    out = []
    for m in re.finditer(pat + r"|[A-Za-z0-9_]+", text):
        tok = m.group(0)
        out.append(SYMBOL_MAP[tok] if tok in SYMBOL_MAP else tok)
    return " ".join(out)
