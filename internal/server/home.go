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
    html, body { margin: 0; min-height: 100%; background: #14130f; }
    * { box-sizing: border-box; }
    .app {
      --bg: #14130f;
      --panel: #1c1a16;
      --panel2: #1a1814;
      --line: #2c2924;
      --fg: #f2efe8;
      --muted: #8d877b;
      --dim: #5f5a51;
      --num: #4d4941;
      --accent: #ff5c1a;
      --accent-ink: #14130f;
      --sel: #221f1a;
      --chip-bg: #f2efe8;
      --chip-fg: #14130f;
      min-height: 100vh;
      background: var(--bg);
      color: var(--fg);
      font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Microsoft YaHei", Arial, sans-serif;
      -webkit-font-smoothing: antialiased;
    }
    .app[data-theme="light"] {
      --bg: #faf8f4;
      --panel: #ffffff;
      --panel2: #f3f0ea;
      --line: #e3ded4;
      --fg: #17160f;
      --muted: #6f6a60;
      --dim: #a09a8e;
      --num: #bdb7aa;
      --accent: #d2450b;
      --accent-ink: #ffffff;
      --sel: #efebe3;
      --chip-bg: #17160f;
      --chip-fg: #faf8f4;
    }
    button, input, textarea, select { font: inherit; color: inherit; }
    button { cursor: pointer; }
    button:disabled { cursor: not-allowed; opacity: .42; }
    a { color: var(--accent); text-decoration: none; }
    a:hover { color: var(--fg); }
    input, textarea, select {
      width: 100%;
      border: 1px solid var(--line);
      border-radius: 11px;
      background: var(--panel);
      outline: none;
      padding: 14px 16px;
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace;
      font-size: 14px;
    }
    input:focus, textarea:focus, select:focus { border-color: var(--accent); }
    textarea {
      min-height: 140px;
      line-height: 1.65;
      resize: vertical;
    }
    input[type="range"] {
      accent-color: var(--accent);
      padding: 0;
    }
    .mono {
      font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, "Liberation Mono", monospace;
    }
    .brand {
      display: flex;
      align-items: center;
      gap: 12px;
      min-width: 0;
      padding: 4px 12px 24px;
    }
    .mark {
      width: 20px;
      height: 20px;
      border-radius: 5px;
      background: var(--accent);
      flex: none;
    }
    .brand-name {
      font-size: 15.5px;
      font-weight: 650;
      white-space: nowrap;
    }
    .brand-code {
      color: var(--dim);
      font-size: 10.5px;
      letter-spacing: .18em;
      text-transform: uppercase;
    }
    .top-actions {
      display: flex;
      align-items: center;
      gap: 8px;
      margin-top: auto;
      padding: 18px 12px 3px;
    }
    .seg {
      display: flex;
      border: 1px solid var(--line);
      border-radius: 9px;
      overflow: hidden;
    }
    .seg button,
    .theme {
      border: 0;
      background: transparent;
      color: var(--muted);
      padding: 7px 13px;
      font-size: 11px;
      letter-spacing: .08em;
    }
    .seg button[aria-pressed="true"] {
      background: var(--sel);
      color: var(--fg);
    }
    .theme {
      display: flex;
      align-items: center;
      gap: 8px;
      border: 1px solid var(--line);
      border-radius: 9px;
    }
    .theme:hover, .seg button:hover { color: var(--fg); }
    .dot {
      width: 11px;
      height: 11px;
      border-radius: 999px;
      border: 1.5px solid var(--muted);
      display: block;
    }
    .app[data-theme="light"] .dot { background: var(--fg); }
    .frame {
      display: grid;
      grid-template-columns: 212px minmax(0, 1fr) 396px;
      align-content: start;
      min-height: 100vh;
    }
    .rail {
      display: flex;
      flex-direction: column;
      gap: 2px;
      padding: 16px 12px;
      border-right: 1px solid var(--line);
    }
    .rail-list {
      display: flex;
      flex-direction: column;
      gap: 2px;
    }
    .eyebrow {
      color: var(--muted);
      font-size: 10.5px;
      letter-spacing: .14em;
      text-transform: uppercase;
    }
    .railcap {
      color: var(--dim);
      font-size: 10.5px;
      text-transform: uppercase;
      padding: 6px 12px 12px;
      letter-spacing: .16em;
    }
    .rail-list button {
      border: 0;
      border-radius: 9px;
      background: transparent;
      color: var(--muted);
      display: flex;
      align-items: center;
      gap: 12px;
      padding: 11px 12px;
      text-align: left;
      white-space: nowrap;
    }
    .rail-list button[aria-current="page"] {
      background: var(--sel);
      color: var(--fg);
    }
    .rail-list button span {
      color: var(--num);
      font-size: 10px;
    }
    .rail-list button b {
      font-size: 13.5px;
      font-weight: 500;
    }
    .rail-list button[aria-current="page"] span { color: var(--accent); }
    main {
      max-width: 640px;
      padding: 34px 40px 60px;
      display: flex;
      flex-direction: column;
      gap: 32px;
    }
    h1 {
      margin: 0;
      font-size: 26px;
      font-weight: 600;
      letter-spacing: 0;
      line-height: 1.18;
    }
    .subhead {
      margin-top: 6px;
      color: var(--muted);
      font-size: 13.5px;
      line-height: 1.6;
    }
    form {
      display: flex;
      flex-direction: column;
      gap: 22px;
    }
    .mode-panel {
      display: none;
      flex-direction: column;
      gap: 22px;
    }
    .mode-panel[data-active="true"] {
      display: flex;
    }
    .field-grid {
      display: grid;
      grid-template-columns: repeat(2, minmax(0, 1fr));
      gap: 22px;
    }
    .span-2 {
      grid-column: 1 / -1;
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
      gap: 18px;
      align-items: baseline;
    }
    .count {
      color: var(--dim);
      font-size: 11px;
    }
    .encoded {
      border-top: 1px solid var(--line);
      padding-top: 18px;
      display: flex;
      flex-direction: column;
      gap: 8px;
    }
    .encoded-output {
      color: var(--dim);
      font-size: 11.5px;
      line-height: 1.7;
      overflow-wrap: anywhere;
      white-space: pre-wrap;
    }
    .previewcol {
      position: sticky;
      top: 0;
      align-self: start;
      border-left: 1px solid var(--line);
      padding: 30px 30px 34px;
      display: flex;
      flex-direction: column;
      gap: 18px;
      max-height: 100vh;
      overflow-y: auto;
    }
    .preview-head {
      display: flex;
      align-items: baseline;
      justify-content: space-between;
      gap: 16px;
    }
    .meta {
      color: var(--dim);
      font-size: 11px;
      white-space: nowrap;
    }
    .qr-frame {
      border-radius: 16px;
      overflow: hidden;
      background: #ffffff;
      padding: 14px;
      display: grid;
      place-items: center;
      aspect-ratio: 1;
      border: 1px solid var(--line);
      flex: none;
      max-width: 340px;
      width: 100%;
      align-self: center;
    }
    .qr-frame img {
      width: 100%;
      height: 100%;
      object-fit: contain;
      image-rendering: pixelated;
    }
    .download, .mbar-download {
      width: 100%;
      padding: 14px 0;
      border: 0;
      border-radius: 11px;
      background: var(--accent);
      color: var(--accent-ink);
      font-size: 14px;
      font-weight: 650;
    }
    .secondary-row {
      display: flex;
      gap: 9px;
    }
    .secondary {
      flex: 1;
      padding: 12px 0;
      border: 1px solid var(--line);
      border-radius: 11px;
      background: transparent;
      color: var(--muted);
      font-size: 13px;
    }
    .secondary:hover {
      border-color: var(--accent);
      color: var(--fg);
    }
    .controls {
      border-top: 1px solid var(--line);
      padding-top: 20px;
      display: flex;
      flex-direction: column;
      gap: 22px;
    }
    .control {
      display: flex;
      flex-direction: column;
      gap: 11px;
    }
    .control-head {
      display: flex;
      justify-content: space-between;
      align-items: baseline;
      gap: 12px;
    }
    .value {
      font-size: 12px;
    }
    .chip-row {
      display: flex;
      gap: 6px;
    }
    .chip {
      flex: 1;
      padding: 11px 0;
      border-radius: 9px;
      border: 1px solid var(--line);
      background: transparent;
      color: var(--muted);
      font-size: 12px;
    }
    .chip[aria-pressed="true"] {
      border-color: var(--chip-bg);
      background: var(--chip-bg);
      color: var(--chip-fg);
    }
    .palette-row {
      display: flex;
      flex-wrap: wrap;
      gap: 9px;
    }
    .swatch {
      width: 46px;
      height: 46px;
      border-radius: 11px;
      border: 1.5px solid var(--line);
      background: var(--swatch-bg);
      padding: 0;
      display: grid;
      place-items: center;
    }
    .swatch[aria-pressed="true"] {
      border-color: var(--accent);
    }
    .swatch span {
      width: 20px;
      height: 20px;
      border-radius: 5px;
      background: var(--swatch-fg);
      display: block;
    }
    .switch {
      display: flex;
      align-items: center;
      gap: 12px;
      border: 0;
      background: transparent;
      color: var(--muted);
      padding: 0;
      text-align: left;
    }
    .switch-track {
      width: 38px;
      height: 22px;
      border-radius: 999px;
      background: var(--line);
      display: flex;
      align-items: center;
      padding: 3px;
      transition: background .18s;
    }
    .switch[aria-pressed="true"] .switch-track {
      background: var(--accent);
      justify-content: flex-end;
    }
    .switch-knob {
      width: 16px;
      height: 16px;
      border-radius: 999px;
      background: #fff;
      display: block;
    }
    .filebox {
      display: flex;
      align-items: center;
      gap: 8px;
      border: 1px solid var(--line);
      border-radius: 11px;
      background: var(--panel);
      padding: 0 14px;
    }
    .filebox:focus-within { border-color: var(--accent); }
    .filebox input {
      flex: 1;
      min-width: 0;
      border: 0;
      background: transparent;
      padding: 13px 0;
    }
    .filebox span {
      color: var(--dim);
      font-size: 13px;
    }
    .level-hint, .toast {
      color: var(--dim);
      font-size: 11.5px;
      line-height: 1.5;
      min-height: 17px;
    }
    .url-box {
      border-top: 1px solid var(--line);
      padding-top: 18px;
      display: flex;
      flex-direction: column;
      gap: 9px;
    }
    .preview-link {
      color: var(--accent);
      font-size: 11.5px;
      line-height: 1.7;
      overflow-wrap: anywhere;
    }
    .sheetgrip, .scrim, .mbar { display: none; }
    @media (max-width: 1080px) {
      .frame { grid-template-columns: 1fr; }
      .rail {
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 10px;
        overflow-x: auto;
        scrollbar-width: none;
        border-right: 0;
        border-bottom: 1px solid var(--line);
        position: sticky;
        top: 0;
        z-index: 30;
        background: var(--bg);
        padding: 10px;
      }
      .brand {
        flex: none;
        padding: 0 6px;
      }
      .top-actions {
        flex: none;
        order: 3;
        margin-top: 0;
        padding: 0;
        justify-content: flex-end;
      }
      .rail-list {
        flex: 1 0 auto;
        flex-direction: row;
        align-items: center;
        overflow-x: auto;
        scrollbar-width: none;
      }
      .rail::-webkit-scrollbar { display: none; }
      .rail-list::-webkit-scrollbar { display: none; }
      .railcap { display: none; }
      .rail-list button {
        flex: none;
        padding: 12px 14px;
        align-self: center;
      }
      main {
        max-width: none;
        padding: 26px 20px 132px;
      }
      .previewcol {
        position: fixed;
        left: 0;
        right: 0;
        bottom: 0;
        top: auto;
        z-index: 50;
        border-left: 0;
        border-top: 1px solid var(--line);
        border-radius: 20px 20px 0 0;
        background: var(--bg);
        max-height: 88vh;
        overflow-y: auto;
        padding: 12px 20px calc(22px + env(safe-area-inset-bottom));
        box-shadow: 0 -24px 60px rgba(0,0,0,.45);
        transform: translateY(101%);
        transition: transform .28s cubic-bezier(.4,0,.2,1);
      }
      .previewcol[data-open="1"] { transform: translateY(0); }
      .sheetgrip {
        display: block;
        align-self: center;
        width: 44px;
        height: 5px;
        border-radius: 999px;
        background: var(--line);
        border: 0;
        padding: 0;
        margin-bottom: 4px;
      }
      .scrim {
        display: block;
        position: fixed;
        inset: 0;
        z-index: 40;
        background: rgba(0,0,0,.55);
        transition: opacity .28s;
      }
      .scrim[data-open="0"] {
        opacity: 0;
        pointer-events: none;
      }
      .mbar {
        position: fixed;
        left: 0;
        right: 0;
        bottom: 0;
        z-index: 45;
        display: flex;
        align-items: center;
        gap: 14px;
        padding: 10px 16px calc(10px + env(safe-area-inset-bottom));
        background: var(--bg);
        border-top: 1px solid var(--line);
      }
      .mbar-open {
        flex: 1;
        min-width: 0;
        display: flex;
        align-items: center;
        gap: 12px;
        border: 0;
        background: transparent;
        padding: 0;
        text-align: left;
      }
      .mini {
        width: 46px;
        height: 46px;
        border-radius: 9px;
        background: #ffffff;
        border: 1px solid var(--line);
        padding: 4px;
        display: grid;
        place-items: center;
        flex: none;
      }
      .mini img {
        width: 100%;
        height: 100%;
        object-fit: contain;
      }
      .mbar-title {
        display: flex;
        flex-direction: column;
        gap: 3px;
        min-width: 0;
      }
      .mbar-title strong {
        color: var(--fg);
        font-size: 13px;
      }
      .mbar-title span {
        color: var(--dim);
        font-size: 10.5px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
      .mbar-download {
        width: auto;
        flex: none;
        padding: 14px 20px;
      }
    }
    @media (max-width: 640px) {
      .rail {
        display: grid;
        grid-template-areas:
          "brand actions"
          "tabs tabs";
        grid-template-columns: minmax(0, 1fr) auto;
        align-items: center;
        gap: 14px 10px;
        overflow: hidden;
        padding: 16px;
      }
      .brand {
        grid-area: brand;
        flex: none;
        padding: 0;
      }
      .brand-code { display: none; }
      .brand-name { font-size: 15px; }
      .top-actions {
        grid-area: actions;
        flex: none;
        order: 0;
        margin-left: 0;
        justify-content: flex-end;
      }
      .seg button, .theme {
        padding: 7px 11px;
      }
      .rail-list {
        grid-area: tabs;
        flex: none;
        display: flex;
        align-items: center;
        justify-content: flex-start;
        width: 100%;
        gap: 8px;
        overflow-x: auto;
        padding-bottom: 2px;
      }
      .rail-list button {
        flex: 0 0 56px;
        justify-content: center;
        gap: 0;
        padding: 0 6px;
        height: 40px;
        min-width: 56px;
        align-self: center;
      }
      .rail-list button span { display: none; }
      .rail-list button b { font-size: 13px; }
      main { padding: 24px 16px 132px; }
      h1 { font-size: 24px; }
      .field-grid { grid-template-columns: 1fr; }
      .span-2 { grid-column: auto; }
      .secondary-row, .chip-row { gap: 7px; }
      .mbar-download { padding: 13px 16px; }
    }
  </style>
</head>
<body>
  <div class="app" id="app" data-theme="dark">
    <div class="frame">
      <nav class="rail" aria-label="Code type">
        <div class="brand">
          <div class="mark" aria-hidden="true"></div>
          <div class="brand-name" data-i18n="brand">码上生成</div>
          <div class="brand-code mono">qrtool</div>
        </div>
        <div class="rail-list">
          <div class="railcap mono" data-i18n="kind">内容类型</div>
          <button type="button" data-kind="url"><span class="mono">01</span><b data-i18n="typeUrl">网址</b></button>
          <button type="button" data-kind="text"><span class="mono">02</span><b data-i18n="typeText">纯文本</b></button>
          <button type="button" data-kind="wifi" aria-current="page"><span class="mono">03</span><b data-i18n="typeWifi">WiFi</b></button>
          <button type="button" data-kind="card"><span class="mono">04</span><b data-i18n="typeCard">名片</b></button>
          <button type="button" data-kind="mail"><span class="mono">05</span><b data-i18n="typeMail">邮件</b></button>
        </div>
        <div class="top-actions">
          <div class="seg" aria-label="Language">
            <button class="mono" type="button" data-lang="zh" aria-pressed="true">中</button>
            <button class="mono" type="button" data-lang="en" aria-pressed="false">EN</button>
          </div>
          <button class="theme mono" id="theme" type="button" title="Toggle theme">
            <span class="dot" aria-hidden="true"></span>
            <span data-i18n="dark">暗色</span>
          </button>
        </div>
      </nav>

      <main>
        <section>
          <h1 id="heading">生成 WiFi 连接码</h1>
          <div class="subhead" id="subheading">客人扫码即可入网，无需口头告知密码。</div>
        </section>

        <form id="qr-form">
          <div class="mode-panel" data-panel="url">
            <label>
              <span class="eyebrow mono" data-i18n="linkLabel">链接地址</span>
              <input id="url" type="url" value="https://example.com/hello" placeholder="https://">
            </label>
          </div>

          <div class="mode-panel" data-panel="text">
            <label>
              <div class="label-row">
                <span class="eyebrow mono" data-i18n="textLabel">文本内容</span>
                <span id="count" class="count mono">5 / 2953</span>
              </div>
              <textarea id="text" maxlength="2953">hello</textarea>
            </label>
          </div>

          <div class="mode-panel" data-panel="wifi" data-active="true">
            <label>
              <span class="eyebrow mono" data-i18n="ssid">网络名称 SSID</span>
              <input id="ssid" value="Cafe_Guest" placeholder="MyNetwork">
            </label>
            <div class="control">
              <span class="eyebrow mono" data-i18n="enc">加密方式</span>
              <div class="chip-row" id="enc-row">
                <button class="chip mono" type="button" data-enc="WPA" aria-pressed="true">WPA</button>
                <button class="chip mono" type="button" data-enc="WEP" aria-pressed="false">WEP</button>
                <button class="chip mono" type="button" data-enc="nopass" aria-pressed="false" data-i18n="nopass">无密码</button>
              </div>
            </div>
            <label id="password-field">
              <span class="eyebrow mono" data-i18n="password">密码</span>
              <input id="password" value="coffee2026">
            </label>
            <button class="switch" id="hidden" type="button" aria-pressed="false">
              <span class="switch-track"><span class="switch-knob"></span></span>
              <span data-i18n="hidden">隐藏网络（不广播 SSID）</span>
            </button>
          </div>

          <div class="mode-panel" data-panel="card">
            <div class="field-grid">
              <label>
                <span class="eyebrow mono" data-i18n="cardName">姓名</span>
                <input id="card-name" value="张伟">
              </label>
              <label>
                <span class="eyebrow mono" data-i18n="cardOrg">公司</span>
                <input id="card-org" value="码上生成">
              </label>
              <label>
                <span class="eyebrow mono" data-i18n="cardTitle">职位</span>
                <input id="card-title" value="产品设计">
              </label>
              <label>
                <span class="eyebrow mono" data-i18n="cardTel">电话</span>
                <input id="card-tel" value="+86 138 0000 0000">
              </label>
              <label class="span-2">
                <span class="eyebrow mono" data-i18n="cardEmail">邮箱</span>
                <input id="card-email" type="email" value="wei@example.com">
              </label>
              <label class="span-2">
                <span class="eyebrow mono" data-i18n="cardSite">网站</span>
                <input id="card-site" value="example.com">
              </label>
            </div>
          </div>

          <div class="mode-panel" data-panel="mail">
            <label>
              <span class="eyebrow mono" data-i18n="mailTo">收件人</span>
              <input id="mail-to" type="email" value="hi@example.com">
            </label>
            <label>
              <span class="eyebrow mono" data-i18n="mailSubject">主题</span>
              <input id="mail-subject" value="你好">
            </label>
            <label>
              <span class="eyebrow mono" data-i18n="mailBody">正文</span>
              <textarea id="mail-body"></textarea>
            </label>
          </div>

        </form>

        <section class="encoded" aria-label="Encoded payload">
          <div class="eyebrow mono" data-i18n="encoded">编码内容</div>
          <div id="payload" class="encoded-output mono">WIFI:T:WPA;S:Cafe_Guest;P:coffee2026;;</div>
        </section>
      </main>

      <aside class="previewcol" id="sheet" data-open="0" aria-label="Preview">
        <button class="sheetgrip" id="close-sheet" type="button" aria-label="Close preview"></button>

        <div class="preview-head">
          <div class="eyebrow mono" data-i18n="preview">实时预览</div>
          <div id="meta" class="meta mono">256 x 256 · PNG</div>
        </div>

        <div class="qr-frame">
          <img id="preview-image" alt="QR code preview">
        </div>

        <div>
          <button id="download" class="download" type="button">下载 PNG</button>
          <div class="secondary-row" style="margin-top:9px">
            <button id="copy-image" class="secondary" type="button" data-i18n="copyImg">复制图片</button>
            <button id="copy-link" class="secondary" type="button" data-i18n="copyLink">复制链接</button>
          </div>
          <div id="toast" class="toast" role="status" aria-live="polite"></div>
        </div>

        <div class="controls">
          <label class="control">
            <div class="control-head">
              <span class="eyebrow mono" data-i18n="size">导出尺寸</span>
              <span class="value mono"><span id="size-value">256</span> px</span>
            </div>
            <input id="size" type="range" min="128" max="1024" step="32" value="256">
          </label>

          <div class="control">
            <div class="eyebrow mono" data-i18n="color">配色</div>
            <div class="palette-row" id="palette-row">
              <button class="swatch" type="button" data-palette="classic" data-fg="14130f" data-bg="ffffff" aria-pressed="true" title="Classic" style="--swatch-fg:#14130f;--swatch-bg:#ffffff"><span></span></button>
              <button class="swatch" type="button" data-palette="ember" data-fg="d2450b" data-bg="fff3ec" aria-pressed="false" title="Ember" style="--swatch-fg:#d2450b;--swatch-bg:#fff3ec"><span></span></button>
              <button class="swatch" type="button" data-palette="pine" data-fg="1f4d3d" data-bg="f0f5f1" aria-pressed="false" title="Pine" style="--swatch-fg:#1f4d3d;--swatch-bg:#f0f5f1"><span></span></button>
              <button class="swatch" type="button" data-palette="indigo" data-fg="24326b" data-bg="eef1f8" aria-pressed="false" title="Indigo" style="--swatch-fg:#24326b;--swatch-bg:#eef1f8"><span></span></button>
              <button class="swatch" type="button" data-palette="inverse" data-fg="f2efe8" data-bg="14130f" aria-pressed="false" title="Inverse" style="--swatch-fg:#f2efe8;--swatch-bg:#14130f"><span></span></button>
            </div>
          </div>

          <div class="control">
            <div class="eyebrow mono" data-i18n="format">格式</div>
            <div class="chip-row" id="format-row">
              <button class="chip mono" type="button" data-format="png" aria-pressed="true">PNG</button>
              <button class="chip mono" type="button" data-format="svg" aria-pressed="false">SVG</button>
            </div>
          </div>

          <div class="control">
            <div class="eyebrow mono" data-i18n="level">容错等级</div>
            <div class="chip-row" id="level-row">
              <button class="chip mono" type="button" data-level="l" aria-pressed="false">L</button>
              <button class="chip mono" type="button" data-level="m" aria-pressed="true">M</button>
              <button class="chip mono" type="button" data-level="q" aria-pressed="false">Q</button>
              <button class="chip mono" type="button" data-level="h" aria-pressed="false">H</button>
            </div>
            <div id="level-hint" class="level-hint">约 15% 容错，日常使用的平衡点</div>
          </div>

          <label class="control">
            <div class="control-head">
              <span class="eyebrow mono" data-i18n="margin">留白边距</span>
              <span class="value mono"><span id="margin-value">4</span> <span data-i18n="modules">模块</span></span>
            </div>
            <input id="margin" type="range" min="0" max="8" step="1" value="4">
          </label>

          <label class="control">
            <div class="eyebrow mono" data-i18n="filename">文件名</div>
            <div class="filebox">
              <input id="filename" value="qrcode" autocomplete="off">
              <span id="ext" class="mono">.png</span>
            </div>
          </label>
        </div>

        <div class="url-box">
          <div class="eyebrow mono">URL</div>
          <a id="preview-link" class="preview-link mono" href="/qrcode.png?text=hello">/qrcode.png?text=hello</a>
        </div>
      </aside>

      <div class="scrim" id="scrim" data-open="0"></div>

      <div class="mbar">
        <button class="mbar-open" id="open-sheet" type="button">
          <span class="mini"><img id="mini-image" alt=""></span>
          <span class="mbar-title">
            <strong data-i18n="preview">实时预览</strong>
            <span id="mobile-meta" class="mono">256 x 256 · PNG · M</span>
          </span>
        </button>
        <button id="mobile-download" class="mbar-download" type="button" data-i18n="download">下载</button>
      </div>
    </div>
  </div>

  <script>
    const I18N = {
      zh: {
        brand: "码上生成", kind: "内容类型", dark: "暗色", light: "亮色",
        typeUrl: "网址", typeText: "纯文本", typeWifi: "WiFi", typeCard: "名片", typeMail: "邮件",
        openQR: "打开二维码",
        headingUrl: "生成一个网址二维码",
        headingText: "生成一段文本二维码",
        headingWifi: "生成 WiFi 连接码",
        headingCard: "生成电子名片 vCard",
        headingMail: "生成写信二维码",
        subheadingUrl: "扫码直接打开链接，适合海报、名片和活动物料。",
        subheadingText: "任意文字原样编码，扫码后以文本形式显示。",
        subheadingWifi: "客人扫码即可入网，无需口头告知密码。",
        subheadingCard: "扫码一键存入通讯录，字段留空则不写入。",
        subheadingMail: "扫码打开邮件应用，收件人与主题已填好。",
        linkLabel: "链接地址", ssid: "网络名称 SSID", enc: "加密方式", nopass: "无密码",
        password: "密码", hidden: "隐藏网络（不广播 SSID）",
        cardName: "姓名", cardOrg: "公司", cardTitle: "职位", cardTel: "电话", cardEmail: "邮箱", cardSite: "网站",
        mailTo: "收件人", mailSubject: "主题", mailBody: "正文",
        textLabel: "文本内容", encoded: "编码内容", preview: "实时预览",
        download: "下载", copyImg: "复制图片", copyLink: "复制链接",
        size: "导出尺寸", format: "格式", color: "配色", level: "容错等级", margin: "留白边距",
        filename: "文件名", modules: "模块",
        copiedLink: "链接已复制。", copiedImage: "图片已复制。", clipboardUnavailable: "当前环境不可用剪贴板。",
        hints: {
          l: "约 7% 容错，图案最简洁",
          m: "约 15% 容错，日常使用的平衡点",
          q: "约 25% 容错，适合印刷与贴纸",
          h: "约 30% 容错，可覆盖中心 Logo"
        }
      },
      en: {
        brand: "QR Studio", kind: "Content type", dark: "Dark", light: "Light",
        typeUrl: "URL", typeText: "Text", typeWifi: "Wi-Fi", typeCard: "vCard", typeMail: "Email",
        openQR: "Open QR Code",
        headingUrl: "Make a link QR code",
        headingText: "Make a text QR code",
        headingWifi: "Make a Wi-Fi join code",
        headingCard: "Make a digital business card",
        headingMail: "Make a compose-email code",
        subheadingUrl: "Scans straight to the page - posters, cards, event signage.",
        subheadingText: "Any characters, encoded verbatim and shown as text.",
        subheadingWifi: "Guests scan to connect - no reading passwords aloud.",
        subheadingCard: "Saves to contacts in one tap. Blank fields are omitted.",
        subheadingMail: "Opens the mail app with recipient and subject prefilled.",
        linkLabel: "Destination URL", ssid: "Network name (SSID)", enc: "Security", nopass: "None",
        password: "Password", hidden: "Hidden network (SSID not broadcast)",
        cardName: "Name", cardOrg: "Company", cardTitle: "Role", cardTel: "Phone", cardEmail: "Email", cardSite: "Website",
        mailTo: "To", mailSubject: "Subject", mailBody: "Body",
        textLabel: "Plain text", encoded: "Encoded payload", preview: "Live preview",
        download: "Download", copyImg: "Copy image", copyLink: "Copy link",
        size: "Export size", format: "Format", color: "Colour", level: "Error correction", margin: "Quiet zone",
        filename: "File name", modules: "modules",
        copiedLink: "Link copied.", copiedImage: "Image copied.", clipboardUnavailable: "Clipboard is unavailable.",
        hints: {
          l: "~7% recovery - sparsest pattern",
          m: "~15% recovery - the everyday balance",
          q: "~25% recovery - good for print and stickers",
          h: "~30% recovery - survives a center logo"
        }
      }
    };

    const app = document.querySelector("#app");
    const heading = document.querySelector("#heading");
    const subheading = document.querySelector("#subheading");
    const text = document.querySelector("#text");
    const url = document.querySelector("#url");
    const ssid = document.querySelector("#ssid");
    const password = document.querySelector("#password");
    const passwordField = document.querySelector("#password-field");
    const hidden = document.querySelector("#hidden");
    const cardName = document.querySelector("#card-name");
    const cardOrg = document.querySelector("#card-org");
    const cardTitle = document.querySelector("#card-title");
    const cardTel = document.querySelector("#card-tel");
    const cardEmail = document.querySelector("#card-email");
    const cardSite = document.querySelector("#card-site");
    const mailTo = document.querySelector("#mail-to");
    const mailSubject = document.querySelector("#mail-subject");
    const mailBody = document.querySelector("#mail-body");
    const size = document.querySelector("#size");
    const margin = document.querySelector("#margin");
    const filename = document.querySelector("#filename");
    const image = document.querySelector("#preview-image");
    const miniImage = document.querySelector("#mini-image");
    const link = document.querySelector("#preview-link");
    const meta = document.querySelector("#meta");
    const mobileMeta = document.querySelector("#mobile-meta");
    const payload = document.querySelector("#payload");
    const count = document.querySelector("#count");
    const ext = document.querySelector("#ext");
    const levelHint = document.querySelector("#level-hint");
    const toast = document.querySelector("#toast");
    const sheet = document.querySelector("#sheet");
    const scrim = document.querySelector("#scrim");
    const state = {
      lang: "zh",
      theme: "dark",
      kind: "wifi",
      format: "png",
      level: "m",
      enc: "WPA",
      hidden: false,
      fg: "14130f",
      bg: "ffffff"
    };

    function t(key) {
      return I18N[state.lang][key];
    }

    function cleanName(value) {
      return (value || "qrcode").trim().replace(/[\/\\?#]+/g, "-") || "qrcode";
    }

    function escapeWiFi(value) {
      return value.replace(/([\\;,":])/g, "\\$1");
    }

    function vcardLine(name, value) {
      return value ? name + ":" + value : "";
    }

    function payloadValue() {
      if (state.kind === "url") {
        return url.value || "";
      }
      if (state.kind === "wifi") {
        const security = state.enc === "nopass" ? "nopass" : state.enc;
        return "WIFI:T:" + security + ";S:" + escapeWiFi(ssid.value || "") + ";" +
          (state.enc === "nopass" ? "" : "P:" + escapeWiFi(password.value || "") + ";") +
          (state.hidden ? "H:true;" : "") + ";";
      }
      if (state.kind === "card") {
        return [
          "BEGIN:VCARD",
          "VERSION:3.0",
          vcardLine("N", cardName.value),
          vcardLine("FN", cardName.value),
          vcardLine("ORG", cardOrg.value),
          vcardLine("TITLE", cardTitle.value),
          vcardLine("TEL", cardTel.value),
          vcardLine("EMAIL", cardEmail.value),
          vcardLine("URL", cardSite.value),
          "END:VCARD"
        ].filter(Boolean).join("\n");
      }
      if (state.kind === "mail") {
        const params = new URLSearchParams();
        if (mailSubject.value) params.set("subject", mailSubject.value);
        if (mailBody.value) params.set("body", mailBody.value);
        const query = params.toString();
        return "mailto:" + (mailTo.value || "") + (query ? "?" + query : "");
      }
      return text.value || "";
    }

    function buildURL() {
      const params = new URLSearchParams();
      params.set("text", payloadValue());
      params.set("size", size.value || "256");
      params.set("level", state.level);
      params.set("margin", margin.value || "4");
      params.set("fg", state.fg);
      params.set("bg", state.bg);
      return "/" + encodeURIComponent(cleanName(filename.value)) + "." + state.format + "?" + params.toString();
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

    function setPressed(selector, attr, value) {
      document.querySelectorAll(selector).forEach((button) => {
        button.setAttribute("aria-pressed", button.getAttribute(attr) === value ? "true" : "false");
      });
    }

    function setSheet(open) {
      sheet.dataset.open = open ? "1" : "0";
      scrim.dataset.open = open ? "1" : "0";
    }

    function downloadCurrent() {
      const a = document.createElement("a");
      a.href = buildURL();
      a.download = "";
      document.body.append(a);
      a.click();
      a.remove();
    }

    function refresh() {
      const url = buildURL();
      const upper = state.format.toUpperCase();
      const info = size.value + " x " + size.value + " · " + upper;
      const payloadText = payloadValue();

      image.src = url;
      miniImage.src = url;
      link.href = url;
      link.textContent = url;
      meta.textContent = info;
      mobileMeta.textContent = info + " · " + state.level.toUpperCase();
      payload.textContent = payloadText || " ";
      count.textContent = (text.value || "").length + " / 2953";
      ext.textContent = "." + state.format;
      document.querySelector("#size-value").textContent = size.value;
      document.querySelector("#margin-value").textContent = margin.value;
      document.querySelector("#download").textContent = t("download") + " " + upper;
      document.querySelector("#mobile-download").textContent = t("download");
      levelHint.textContent = I18N[state.lang].hints[state.level];
    }

    function setKind(kind) {
      state.kind = kind;
      document.querySelectorAll("[data-kind]").forEach((button) => {
        if (button.dataset.kind === kind) {
          button.setAttribute("aria-current", "page");
        } else {
          button.removeAttribute("aria-current");
        }
      });
      document.querySelectorAll("[data-panel]").forEach((panel) => {
        panel.dataset.active = panel.dataset.panel === kind ? "true" : "false";
      });
      heading.textContent = I18N[state.lang]["heading" + kind[0].toUpperCase() + kind.slice(1)];
      subheading.textContent = I18N[state.lang]["subheading" + kind[0].toUpperCase() + kind.slice(1)];
      refresh();
    }

    function applyLang(lang) {
      state.lang = lang;
      document.documentElement.lang = lang;
      document.querySelectorAll("[data-i18n]").forEach((node) => {
        const value = I18N[lang][node.dataset.i18n];
        if (value) node.textContent = value;
      });
      document.querySelectorAll("[data-lang]").forEach((button) => {
        button.setAttribute("aria-pressed", button.dataset.lang === lang ? "true" : "false");
      });
      setKind(state.kind);
      refresh();
    }

    function applyTheme(theme) {
      state.theme = theme;
      app.dataset.theme = theme;
      document.querySelector("#theme [data-i18n]").dataset.i18n = theme === "dark" ? "dark" : "light";
      document.querySelector("#theme [data-i18n]").textContent = I18N[state.lang][theme === "dark" ? "dark" : "light"];
    }

    document.querySelectorAll("input, textarea").forEach((input) => {
      input.addEventListener("input", refresh);
    });
    size.addEventListener("input", refresh);
    margin.addEventListener("input", refresh);
    document.querySelector("#qr-form").addEventListener("submit", (event) => {
      event.preventDefault();
      window.location.href = buildURL();
    });
    document.querySelector("#download").addEventListener("click", downloadCurrent);
    document.querySelector("#mobile-download").addEventListener("click", downloadCurrent);
    document.querySelector("#open-sheet").addEventListener("click", () => setSheet(true));
    document.querySelector("#close-sheet").addEventListener("click", () => setSheet(false));
    scrim.addEventListener("click", () => setSheet(false));

    document.querySelectorAll("[data-format]").forEach((button) => {
      button.addEventListener("click", () => {
        state.format = button.dataset.format;
        setPressed("[data-format]", "data-format", state.format);
        refresh();
      });
    });

    document.querySelectorAll("[data-enc]").forEach((button) => {
      button.addEventListener("click", () => {
        state.enc = button.dataset.enc;
        setPressed("[data-enc]", "data-enc", state.enc);
        passwordField.style.display = state.enc === "nopass" ? "none" : "";
        refresh();
      });
    });

    document.querySelectorAll("[data-palette]").forEach((button) => {
      button.addEventListener("click", () => {
        state.fg = button.dataset.fg;
        state.bg = button.dataset.bg;
        setPressed("[data-palette]", "data-palette", button.dataset.palette);
        document.querySelector(".qr-frame").style.background = "#" + state.bg;
        document.querySelector(".mini").style.background = "#" + state.bg;
        refresh();
      });
    });

    hidden.addEventListener("click", () => {
      state.hidden = !state.hidden;
      hidden.setAttribute("aria-pressed", state.hidden ? "true" : "false");
      refresh();
    });

    document.querySelectorAll("[data-kind]").forEach((button) => {
      button.addEventListener("click", () => setKind(button.dataset.kind));
    });

    document.querySelectorAll("[data-level]").forEach((button) => {
      button.addEventListener("click", () => {
        state.level = button.dataset.level;
        setPressed("[data-level]", "data-level", state.level);
        refresh();
      });
    });

    document.querySelectorAll("[data-lang]").forEach((button) => {
      button.addEventListener("click", () => applyLang(button.dataset.lang));
    });

    document.querySelector("#theme").addEventListener("click", () => {
      applyTheme(state.theme === "dark" ? "light" : "dark");
    });

    document.querySelector("#copy-link").addEventListener("click", async () => {
      try {
        await navigator.clipboard.writeText(absoluteURL(buildURL()));
        showToast(t("copiedLink"));
      } catch {
        showToast(t("clipboardUnavailable"));
      }
    });

    document.querySelector("#copy-image").addEventListener("click", async () => {
      try {
        const response = await fetch(buildURL());
        const blob = await response.blob();
        await navigator.clipboard.write([new ClipboardItem({ [blob.type]: blob })]);
        showToast(t("copiedImage"));
      } catch {
        showToast(t("clipboardUnavailable"));
      }
    });

    applyLang("zh");
    applyTheme("dark");
  </script>
</body>
</html>
`
