import hashlib, json
from datetime import datetime, timezone

class SVCTBloom:
    """Append-only, hash-linked constant store.
    This is the BACKUP validator — not a substitute for the primary logic.
    When drift enters, the chain lets us detect and correct our own mistakes."""

    def __init__(self, seed="SVCT-GENESIS"):
        self.chain = []
        self._append_row({"type": "genesis", "value": seed,
                          "note": "SVCT bloom chain initialized"})

    def _hash(self, obj):
        return hashlib.sha256(
            json.dumps(obj, sort_keys=True).encode()
        ).hexdigest()

    def _append_row(self, payload):
        prev = self.chain[-1]["hash"] if self.chain else "GENESIS"
        row = {
            "ts": datetime.now(timezone.utc).isoformat(),
            "prev": prev,
            "payload": payload,
            "hash": None,
        }
        row["hash"] = self._hash(row)
        self.chain.append(row)
        return row

    def add_constant(self, kind, value, note=""):
        for r in self.chain:
            p = r["payload"]
            if p.get("kind") == kind and p.get("value") == value:
                return {"status": "duplicate", "value": value}
        row = self._append_row({"kind": kind, "value": value,
                               "note": note, "op": "ADD"})
        return {"status": "chained", "row": row["hash"][:16]}

    def validate(self, kind, value):
        for r in self.chain:
            p = r["payload"]
            if p.get("kind") == kind and p.get("value") == value:
                return True
        return False

    def export(self):
        return json.dumps(self.chain, indent=2)


# --- Build the chain ---
bloom = SVCTBloom()
bloom.add_constant("email", "timnorman730@gmail.com", "primary Gmail")
bloom.add_constant("email", "timothy_h_norman@yahoo.com", "primary Yahoo")
bloom.add_constant("github", "GitTim2Day", "primary GitHub handle")

if __name__ == "__main__":
    print("Validate timnorman730:", bloom.validate("email", "timnorman730@gmail.com"))
    print("Validate gittim2day:", bloom.validate("email", "gittim2day@gmail.com"))
    print("\n--- SVCT BLOOM CHAIN ---")
    print(bloom.export())
