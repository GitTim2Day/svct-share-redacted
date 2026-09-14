import json, hashlib
from datetime import datetime, timezone
PILLARS = ["content","context","data","time","validation","provenance","integrity","custody"]
CONSTANTS = {"email":["timnorman730@gmail.com","timothy_h_norman@yahoo.com"],"github":["GitTim2Day"]}
BLOOM_FILE = "svct_bloom_chain.json"
print("SVCT boot loaded", len(PILLARS), "pillars")
