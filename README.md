<div align="center">

<img src="https://user-images.githubusercontent.com/130087473/235356807-32b80288-7808-4f66-a6f2-fcbe7ab34b72.png" alt="WebPalm Banner" width="600">

<h2> 🌐 Advanced Web Traversal & Data Extraction 🚀 </h2>

[![GitHub release](https://img.shields.io/github/v/release/XORbit01/webpalm?color=blue&label=release)]()
[![GitHub license](https://img.shields.io/github/license/XORbit01/webpalm?color=green)]()
[![GitHub issues](https://img.shields.io/github/issues/XORbit01/webpalm?color=red)]()
[![GitHub stars](https://img.shields.io/github/stars/XORbit01/webpalm?color=yellow)]()
[![GitHub forks](https://img.shields.io/github/forks/XORbit01/webpalm?color=orange)]()
[![GitHub watchers](https://img.shields.io/github/watchers/XORbit01/webpalm?color=blue)]()

🔍 **Crawl websites efficiently, extract structured data, and visualize connections.** 🕵️‍♂️

<img src="https://github.com/XORbit01/webpalm/assets/130087473/6c601672-f278-431d-854b-0a9876a2fafd" alt="WebPalm Preview" width="600">

</div>

---

## 🗺️ Table of Contents
- [`📦 Installation`](#-installation)
- [`⚡ Features`](#-features)
- [`🚀 Usage`](#-usage)
- [`📌 Examples`](#-examples)
- [`📜 Regex Patterns`](#-regex-patterns)
- [`🤝 Contributing`](#-contributing)

---

## ⚡ Features

🌳 **Structured Web-Tree Generation** - Easily visualize website structure  
🕵️ **Regex-Based Data Extraction** - Extract key information efficiently  
⚡ **High-Speed Multi-threading** - Optimized for performance  
📂 **Multiple Export Formats** - Save results in JSON, XML, or TXT  
🎨 **Colorized Output & Robust Error Handling** - Enhanced readability

---

## 📦 Installation

### 📥 Download Binary

### 📥 Compile from Source
```sh
git clone https://github.com/XORbit01/webpalm.git
cd webpalm
go build -o webpalm && ./webpalm
```  
👉 [Download Latest Release](https://github.com/XORbit01/webpalm/releases/latest)

### 📥 Install via Go
```sh
go install github.com/XORbit01/webpalm/v2@latest
```

---

## 🚀 Usage

```sh
webpalm -h
```

### ⚙️ Common Flags
```yaml
🌎 -i, --include     # Include only specific domains (e.g., google.com, facebook.com)
🔗 -u, --url         # Target website  
📏 -l, --level       # Depth of traversal  
❌ -x, --exclude     # Exclude status codes (e.g., 404, 500)  
💾 -o, --output      # Save results (JSON, XML, TXT)  
🚀 -w, --worker      # Multi-threading workers  
🔍 --regexes         # Extract data using regex  
```

---

## 📌 Examples

### 🌲 Generate a Website Map
```sh
webpalm -u https://example.com -l2
```

### 💬 Extract Comments from Pages
```sh
webpalm -u https://example.com -l1 --regexes comments="\<\!--.*?-->" -o results.json
```

### 🚀 Crawl with Multi-threading
```sh
webpalm -u https://example.com -l3 -w 50
```

### 💾 Export Results
```sh
webpalm -u https://example.com -l2 -o output.xml
```

---

## 📜 Regex Patterns

| 🔍 Purpose   | 📜 Regex Pattern |
|-----------|--------------|
| 📧 Emails    | `[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+` |
| 💬 Comments  | `\<\!--.*?-->` |
| 🔑 Tokens    | `[a-zA-Z0-9]{32}` |
| 🔐 Passwords | `\bpassword\b.{0,10}` |

📌 *Escape special characters if needed.*

---

## 🤝 Contributing
💡 Pull requests are welcome! Open an issue before major changes.  
📢 **Discord:** `xorbit.`

---


