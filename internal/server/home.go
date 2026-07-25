package server

import "net/http"

func shouldShowHome(r *http.Request) bool {
	return r.Method == http.MethodGet && r.URL.Path == "/" && r.URL.Query().Get("text") == ""
}

func handleHome(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(homeHTML))
}

const homeHTML = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>qrtool</title>
  <style>
    :root {
      color-scheme: light dark;
      --bg: #f6f8fb;
      --panel: #ffffff;
      --text: #172033;
      --muted: #5c667a;
      --border: #d9deea;
      --accent: #1463ff;
      --accent-strong: #0b4bd3;
      --preview: #eef2f8;
    }
    @media (prefers-color-scheme: dark) {
      :root {
        --bg: #10131a;
        --panel: #181d27;
        --text: #eef2ff;
        --muted: #aab4c8;
        --border: #303747;
        --accent: #75a7ff;
        --accent-strong: #9abdff;
        --preview: #111722;
      }
    }
    * { box-sizing: border-box; }
    body {
      margin: 0;
      min-height: 100vh;
      background: var(--bg);
      color: var(--text);
      font: 16px/1.45 system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
    }
    main {
      width: min(960px, calc(100% - 32px));
      margin: 0 auto;
      padding: 48px 0;
      display: grid;
      grid-template-columns: minmax(0, 1fr) 320px;
      gap: 28px;
      align-items: start;
    }
    header {
      grid-column: 1 / -1;
    }
    h1 {
      margin: 0 0 8px;
      font-size: clamp(32px, 6vw, 56px);
      line-height: 1;
      letter-spacing: 0;
    }
    p {
      margin: 0;
      max-width: 640px;
      color: var(--muted);
    }
    form, .preview {
      background: var(--panel);
      border: 1px solid var(--border);
      border-radius: 8px;
    }
    form {
      padding: 22px;
      display: grid;
      gap: 16px;
    }
    label {
      display: grid;
      gap: 7px;
      color: var(--muted);
      font-size: 14px;
      font-weight: 600;
    }
    input, select, textarea {
      width: 100%;
      border: 1px solid var(--border);
      border-radius: 6px;
      background: transparent;
      color: var(--text);
      font: inherit;
      padding: 10px 11px;
    }
    textarea {
      min-height: 132px;
      resize: vertical;
    }
    .grid {
      display: grid;
      grid-template-columns: repeat(2, minmax(0, 1fr));
      gap: 14px;
    }
    button {
      border: 0;
      border-radius: 6px;
      background: var(--accent);
      color: #fff;
      cursor: pointer;
      font: inherit;
      font-weight: 700;
      min-height: 44px;
      padding: 10px 14px;
    }
    button:hover {
      background: var(--accent-strong);
    }
    .preview {
      padding: 18px;
      display: grid;
      gap: 14px;
      position: sticky;
      top: 24px;
    }
    .preview-frame {
      display: grid;
      place-items: center;
      min-height: 282px;
      border-radius: 6px;
      background: var(--preview);
      overflow: hidden;
    }
    .preview img {
      width: 256px;
      height: 256px;
      image-rendering: pixelated;
    }
    .preview a {
      color: var(--accent);
      overflow-wrap: anywhere;
      text-decoration: none;
    }
    .preview a:hover {
      text-decoration: underline;
    }
    @media (max-width: 760px) {
      main {
        grid-template-columns: 1fr;
        padding: 32px 0;
      }
      .preview {
        position: static;
      }
    }
  </style>
</head>
<body>
  <main>
    <header>
      <h1>qrtool</h1>
      <p>Generate QR code images from text.</p>
    </header>

    <form id="qr-form">
      <label>
        Text
        <textarea name="text" required autofocus>hello</textarea>
      </label>

      <div class="grid">
        <label>
          File name
          <input name="name" value="qrcode" autocomplete="off">
        </label>
        <label>
          Format
          <select name="format">
            <option value="png">PNG</option>
            <option value="svg">SVG</option>
          </select>
        </label>
        <label>
          Size
          <input name="size" type="number" min="64" max="2048" step="1" value="256">
        </label>
        <label>
          Level
          <select name="level">
            <option value="l">L</option>
            <option value="m" selected>M</option>
            <option value="q">Q</option>
            <option value="h">H</option>
          </select>
        </label>
        <label>
          Margin
          <input name="margin" type="number" min="0" max="20" step="1" value="4">
        </label>
      </div>

      <button type="submit">Open QR Code</button>
    </form>

    <section class="preview" aria-label="Preview">
      <div class="preview-frame">
        <img id="preview-image" alt="QR code preview">
      </div>
      <a id="preview-link" href="/qrcode.png?text=hello">/qrcode.png?text=hello</a>
    </section>
  </main>

  <script>
    const form = document.querySelector("#qr-form");
    const image = document.querySelector("#preview-image");
    const link = document.querySelector("#preview-link");

    function cleanName(value) {
      return (value || "qrcode").trim().replace(/[\/\\?#]+/g, "-") || "qrcode";
    }

    function buildURL() {
      const data = new FormData(form);
      const format = data.get("format") || "png";
      const name = cleanName(data.get("name"));
      const params = new URLSearchParams();
      params.set("text", data.get("text") || "");
      params.set("size", data.get("size") || "256");
      params.set("level", data.get("level") || "m");
      params.set("margin", data.get("margin") || "4");
      return "/" + encodeURIComponent(name) + "." + format + "?" + params.toString();
    }

    function refresh() {
      const url = buildURL();
      image.src = url;
      link.href = url;
      link.textContent = url;
    }

    form.addEventListener("input", refresh);
    form.addEventListener("change", refresh);
    form.addEventListener("submit", (event) => {
      event.preventDefault();
      window.location.href = buildURL();
    });
    refresh();
  </script>
</body>
</html>
`
