# ASCII Art Web Generator

A high-performance full-stack web application built in Go that transforms user input text into stylized ASCII art banners in real-time. This project evolves a series of command-line tools into a unified, cloud-native web service optimized for minimal runtime file I/O and low-latency serverless edge distribution.

Live Demo: [ascii-art-web-generator.vercel.app](https://ascii-art-web-generator.vercel.app/)

---

## 🛠️ System Architecture & Engineering Choices

### 1. Memory-Resident Asset Pipeline (`go:embed`)
Instead of invoking expensive operating system disk reads (`os.ReadFile`) on every single user request to parse banner files (`standard.txt`, `shadow.txt`, `thinkertoy.txt`), this application leverages Go's native `embed` package. Banners and HTML templates are baked directly into the compiled application binary at compile time. This reduces runtime file descriptor overhead to zero and ensures lightning-fast execution in ephemeral environments.

### 2. High-Performance Text Layout Scanner
The core rendering engine processes strings through an explicit buffer management system. To handle custom color highlighting, the engine uses a state-tracking algorithm that scans the input text, flags matching sub-strings, and dynamically wraps targeted characters in inline HTML `rgba`/hex wrappers without breaking the structural padding of the vertical ASCII canvas grids.

### 3. Serverless Optimization & Resource Isolation
Designed specifically to run inside Vercel's Edge network, the repository decouples static UI elements from dynamic execution pipelines:
* **Edge CDN:** CSS layouts and brand assets live in a dedicated `static/` hierarchy served instantly via edge routing tables.
* **Serverless Lambda:** The root Go HTTP handler runs within a dedicated, isolated serverless binary path, intercepting incoming payload streams cleanly.

---

## 🚀 Tech Stack

* **Backend Engine:** Go (Golang) 1.20+
* **Frontend Web UI:** Semantic HTML5, CSS3 Grid / Flexbox Architecture
* **Testing:** Table-Driven Go Unit Tests (`testing`)
* **Deployment/Hosting:** Vercel Serverless Function Engine

---

## 💻 Local Development Setup

To run this project locally, ensure you have Go installed on your system (Linux, WSL, macOS, or Windows).

1. Clone the repository to your machine:
   ```bash
   git clone [https://github.com/dsunday05034-crypto/ASCII-Art-Web-Generator.git](https://github.com/dsunday05034-crypto/ASCII-Art-Web-Generator.git)
   cd ASCII-Art-Web-Generator

```

2. Start the local multi-plexed development server:
```bash
go run main.go

```


3. Open your preferred web browser and navigate to:
```text
http://localhost:8080

```



---

## 🧪 Automated Testing Suite

The project includes an idiomatic table-driven test suite to verify the structural string matching and HTML coloring components of the engine.

To run the automated tests locally, move into the package core directory and execute the test runner:

```bash
cd api
go test -v

```
---

## 👤 Author

* **dsunday05034-crypto** - [GitHub Profile](https://github.com/dsunday05034-crypto)