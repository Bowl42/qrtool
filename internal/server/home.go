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
      color-scheme: dark;
      --bg: #111217;
      --panel: #191d26;
      --panel-soft: #151923;
      --line: #303747;
      --line-strong: #3b4457;
      --text: #eef2ff;
      --muted: #aab4c8;
      --dim: #6f7b91;
      --accent: #75a7ff;
      --accent-ink: #07111f;
      --accent-hover: #8bb6ff;
      --white: #ffffff;
    }
    * { box-sizing: border-box; }
    html, body { margin: 0; min-height: 100%; }
    body {
      background: var(--bg);
      color: var(--text);
      font: 15px/1.5 system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
      -webkit-font-smoothing: antialiased;
    }
    button, input, select, textarea { color: inherit; font: inherit; }
    button { cursor: pointer; }
    a { color: var(--accent); text-decoration: none; }
    a:hover { color: var(--text); }
    .topbar {
      min-height: 61px;
      border-bottom: 1px solid var(--line);
      display: flex;
      align-items: center;
      justify-content: space-between;
      gap: 20px;
      padding: 16px 26px;
    }
    .brand {
      display: flex;
      align-items: center;
      gap: 12px;
      min-width: 0;
    }
    .mark {
      width: 20px;
      height: 20px;
      border-radius: 5px;
      background: var(--accent);
      box-shadow: 0 0 0 4px rgba(117, 167, 255, .08);
      flex: none;
    }
    .brand-name {
      font-size: 15px;
      font-weight: 700;
    }
    .brand-code {
      color: var(--dim);
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
      font-size: 10px;
      letter-spacing: .16em;
      text-transform: uppercase;
    }
    .shell {
      display: grid;
      grid-template-columns: 188px minmax(0, 1fr) 396px;
      min-height: calc(100vh - 61px);
    }
    .rail {
      border-right: 1px solid var(--line);
      padding: 16px 12px;
    }
    .rail-title,
    .eyebrow {
      color: var(--dim);
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
      font-size: 10.5px;
      letter-spacing: .14em;
      text-transform: uppercase;
    }
    .rail-title { padding: 6px 12px 12px; }
    .rail-button {
      width: 100%;
      border: 0;
      border-radius: 9px;
      background: var(--panel-soft);
      color: var(--text);
      display: flex;
      align-items: center;
      gap: 12px;
      padding: 11px 12px;
      text-align: left;
    }
    .rail-button span {
      color: var(--dim);
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
      font-size: 10px;
    }
    main {
      max-width: 680px;
      padding: 34px 40px 64px;
      display: flex;
      flex-direction: column;
      gap: 30px;
    }
    h1 {
      margin: 0;
      font-size: 28px;
      font-weight: 700;
      letter-spacing: 0;
      line-height: 1.15;
    }
    .subhead {
      margin-top: 7px;
      color: var(--muted);
      font-size: 14px;
      max-width: 560px;
    }
    form {
      display: flex;
      flex-direction: column;
      gap: 22px;
    }
    label {
      display: flex;
      flex-direction: column;
      gap: 9px;
      min-width: 0;
    }
    .label-row {
      display: flex;
      justify-content: space-between;
      gap: 14px;
      align-items: baseline;
    }
    .count {
      color: var(--dim);
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
      font-size: 11px;
    }
    input, select, textarea {
      width: 100%;
      border: 1px solid var(--line);
      border-radius: 11px;
      background: var(--panel);
      outline: none;
      padding: 14px 16px;
    }
    input:focus, select:focus, textarea:focus {
      border-color: var(--accent);
    }
    textarea {
      min-height: 142px;
      line-height: 1.65;
      resize: vertical;
    }
    .field-grid {
      display: grid;
      grid-template-columns: repeat(2, minmax(0, 1fr));
      gap: 18px 16px;
    }
    .submit {
      min-height: 48px;
      border: 0;
      border-radius: 11px;
      background: var(--accent);
      color: var(--accent-ink);
      font-weight: 750;
    }
    .submit:hover { background: var(--accent-hover); }
    .encoded {
      border-top: 1px solid var(--line);
      padding-top: 18px;
      display: grid;
      gap: 8px;
    }
    .encoded-output {
      color: var(--dim);
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
      font-size: 11.5px;
      line-height: 1.7;
      overflow-wrap: anywhere;
      white-space: pre-wrap;
    }
    .preview {
      position: sticky;
      top: 0;
      align-self: start;
      min-height: calc(100vh - 61px);
      border-left: 1px solid var(--line);
      padding: 30px;
      display: flex;
      flex-direction: column;
      gap: 18px;
    }
    .preview-head {
      display: flex;
      justify-content: space-between;
      align-items: baseline;
      gap: 16px;
    }
    .meta {
      color: var(--dim);
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
      font-size: 11px;
      white-space: nowrap;
    }
    .qr-frame {
      width: 100%;
      max-width: 340px;
      aspect-ratio: 1;
      align-self: center;
      border: 1px solid var(--line);
      border-radius: 16px;
      background: var(--white);
      padding: 14px;
      display: grid;
      place-items: center;
      overflow: hidden;
    }
    .qr-frame img {
      width: 100%;
      height: 100%;
      object-fit: contain;
      image-rendering: pixelated;
    }
    .actions {
      display: grid;
      gap: 9px;
    }
    .download {
      width: 100%;
      border: 0;
      border-radius: 11px;
      background: var(--accent);
      color: var(--accent-ink);
      padding: 14px 0;
      font-weight: 750;
    }
    .ghost-row {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 9px;
    }
    .ghost {
      border: 1px solid var(--line);
      border-radius: 11px;
      background: transparent;
      color: var(--muted);
      padding: 12px 10px;
      font-size: 13px;
    }
    .ghost:hover {
      border-color: var(--accent);
      color: var(--text);
    }
    .link-box {
      border-top: 1px solid var(--line);
      padding-top: 18px;
      display: grid;
      gap: 9px;
    }
    .preview-link {
      color: var(--accent);
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace;
      font-size: 12px;
      line-height: 1.6;
      overflow-wrap: anywhere;
    }
    .toast {
      min-height: 18px;
      color: var(--muted);
      font-size: 12px;
    }
    .mobile-preview { display: none; }
    @media (max-width: 1080px) {
      .shell {
        grid-template-columns: 1fr;
      }
      .rail {
        position: sticky;
        top: 0;
        z-index: 2;
        border-right: 0;
        border-bottom: 1px solid var(--line);
        background: var(--bg);
        padding: 10px;
      }
      .rail-title { display: none; }
      .rail-button {
        width: auto;
        padding: 12px 14px;
      }
      main {
        max-width: none;
        padding: 26px 20px 112px;
      }
      .preview {
        position: static;
        min-height: 0;
        border-left: 0;
        border-top: 1px solid var(--line);
      }
    }
    @media (max-width: 720px) {
      .topbar {
        padding: 14px 18px;
      }
      .brand-code { display: none; }
      main {
        padding: 24px 16px 30px;
      }
      h1 {
        font-size: 25px;
      }
      .field-grid {
        grid-template-columns: 1fr;
      }
      .preview {
        padding: 22px 16px 28px;
      }
      .qr-frame {
        max-width: 312px;
      }
      .ghost-row {
        grid-template-columns: 1fr;
      }
    }
  </style>
</head>
<body>
  <header class="topbar">
    <div class="brand">
      <div class="mark" aria-hidden="true"></div>
      <div class="brand-name">QR Generator</div>
      <div class="brand-code">qrtool</div>
    </div>
  </header>

  <div class="shell">
    <nav class="rail" aria-label="Code type">
      <div class="rail-title">kind</div>
      <button class="rail-button" type="button" aria-current="page">
        <span>01</span>
        Text QR Code
      </button>
    </nav>

    <main>
      <section>
        <h1>Generate QR code images from text.</h1>
        <div class="subhead">Create direct PNG or SVG URLs that can be opened, cached, saved, or embedded.</div>
      </section>

      <form id="qr-form">
        <label>
          <div class="label-row">
            <span class="eyebrow">text</span>
            <span id="count" class="count">5 / 2953</span>
          </div>
          <textarea name="text" required autofocus maxlength="2953">hello</textarea>
        </label>

        <div class="field-grid">
          <label>
            <span class="eyebrow">file name</span>
            <input name="name" value="qrcode" autocomplete="off">
          </label>
          <label>
            <span class="eyebrow">format</span>
            <select name="format">
              <option value="png">PNG</option>
              <option value="svg">SVG</option>
            </select>
          </label>
          <label>
            <span class="eyebrow">size</span>
            <input name="size" type="number" min="64" max="2048" step="1" value="256">
          </label>
          <label>
            <span class="eyebrow">level</span>
            <select name="level">
              <option value="l">L</option>
              <option value="m" selected>M</option>
              <option value="q">Q</option>
              <option value="h">H</option>
            </select>
          </label>
          <label>
            <span class="eyebrow">margin</span>
            <input name="margin" type="number" min="0" max="20" step="1" value="4">
          </label>
        </div>

        <button class="submit" type="submit">Open QR Code</button>
      </form>

      <section class="encoded" aria-label="Encoded payload">
        <div class="eyebrow">encoded</div>
        <div id="payload" class="encoded-output">hello</div>
      </section>
    </main>

    <aside class="preview" aria-label="Preview">
      <div class="preview-head">
        <div class="eyebrow">preview</div>
        <div id="meta" class="meta">256 x 256 · png</div>
      </div>

      <div class="qr-frame">
        <img id="preview-image" alt="QR code preview">
      </div>

      <div class="actions">
        <button id="download" class="download" type="button">Download PNG</button>
        <div class="ghost-row">
          <button id="copy-image" class="ghost" type="button">Copy Image</button>
          <button id="copy-link" class="ghost" type="button">Copy Link</button>
        </div>
        <div id="toast" class="toast" role="status" aria-live="polite"></div>
      </div>

      <div class="link-box">
        <div class="eyebrow">url</div>
        <a id="preview-link" class="preview-link" href="/qrcode.png?text=hello">/qrcode.png?text=hello</a>
      </div>
    </aside>
  </div>

  <script>
    const form = document.querySelector("#qr-form");
    const text = form.elements.text;
    const image = document.querySelector("#preview-image");
    const link = document.querySelector("#preview-link");
    const meta = document.querySelector("#meta");
    const payload = document.querySelector("#payload");
    const count = document.querySelector("#count");
    const toast = document.querySelector("#toast");
    const download = document.querySelector("#download");
    const copyLink = document.querySelector("#copy-link");
    const copyImage = document.querySelector("#copy-image");

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

    function absoluteURL(path) {
      return new URL(path, window.location.href).toString();
    }

    function showToast(message) {
      toast.textContent = message;
      window.clearTimeout(showToast.timer);
      showToast.timer = window.setTimeout(() => {
        toast.textContent = "";
      }, 1800);
    }

    function refresh() {
      const data = new FormData(form);
      const url = buildURL();
      const format = data.get("format") || "png";
      const size = data.get("size") || "256";
      const value = data.get("text") || "";

      image.src = url;
      link.href = url;
      link.textContent = url;
      meta.textContent = size + " x " + size + " · " + format;
      payload.textContent = value || " ";
      count.textContent = value.length + " / 2953";
      download.textContent = "Download " + format.toUpperCase();
    }

    form.addEventListener("input", refresh);
    form.addEventListener("change", refresh);
    form.addEventListener("submit", (event) => {
      event.preventDefault();
      window.location.href = buildURL();
    });

    download.addEventListener("click", () => {
      const a = document.createElement("a");
      a.href = buildURL();
      a.download = "";
      document.body.append(a);
      a.click();
      a.remove();
    });

    copyLink.addEventListener("click", async () => {
      try {
        await navigator.clipboard.writeText(absoluteURL(buildURL()));
        showToast("Link copied.");
      } catch {
        showToast("Clipboard is unavailable.");
      }
    });

    copyImage.addEventListener("click", async () => {
      try {
        const response = await fetch(buildURL());
        const blob = await response.blob();
        await navigator.clipboard.write([
          new ClipboardItem({ [blob.type]: blob }),
        ]);
        showToast("Image copied.");
      } catch {
        showToast("Clipboard is unavailable.");
      }
    });

    refresh();
  </script>
</body>
</html>
`
