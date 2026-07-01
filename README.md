# Go REST API for Notification serivce

```bash
# --< Used commands >--

# Start server
go run ./cmd/app

# Pad image with white
ffmpeg -i abs.png -vf "pad=width='max(iw,ih)':height='max(iw,ih)':x='(oh-iw)/2':y='(ow-ih)/2':color=white" icon.png
```
