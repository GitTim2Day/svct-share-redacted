America/New_York 2026-11-06T18:36:08-05:00

# SHEPHERD EDGE NODE — Usable Product Consolidation

## Product
Shepherd Edge Node (Pi 5)

One-sentence: A local, no-sudo, offline-first edge node on Raspberry Pi 5 that turns any sensor or spoken phrase into a captured frame, OCR text, spoken reply, or actuator command, and only leaves the board when it needs an instruction or a place to store the record.

## Three Layers
1. Triggers — Cron (time), Voice/Vosk (intent), Motion/PIR (activity). All call the same GOSUB. Features off until detection of need.
2. Edge Core — ffmpeg MJPEG + local Tesseract (TESSDATA_PREFIX) → verified/quarantine/reject. Piper TTS. Optional no-sudo actuator. Reflex path.
3. Upstream Bridge — messaging/phone/API for instructions or shelves.

## Working Pair
ffmpeg -f v4l2 -input_format mjpeg -video_size 1920x1080 -i /dev/video0 -frames:v 1 -update 1 -y page.jpg
TESSDATA_PREFIX=~/vosk/tessdata ~/vosk/tesseract page.jpg page -l eng --psm 6 && cat page.txt

## Constraints
No-sudo, local tools, on-detection-of-need, selective bidirectional index, misses archived.

Pointer: session-map-2026-11-06-binary-pipeline-ocr-pi5

SHA-256 89c9021f4615a03fa609b4239c8da20512acfc0357325d880eb57b12e81e604a size 2752

Fresh time check: America/New_York 2026-11-06T18:36:08-05:00
