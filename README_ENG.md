<p align="center">
  <a href="README_ENG.md">
    <img src="https://img.shields.io/badge/🇬🇧_English-00D4FF?style=for-the-badge&logo=readme&logoColor=white" alt="English README">
  </a>
  <a href="README.md">
    <img src="https://img.shields.io/badge/🇺🇦_Українська-FF4D00?style=for-the-badge&logo=readme&logoColor=white" alt="Українська версія">
  </a>
</p>

<br>

# 🛡️ ADBlock-Private-DNS (ADBlock-PD)

<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/weby-homelab/adblock-pd/master/logo-dark.svg">
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/weby-homelab/adblock-pd/master/logo-light.svg">
    <img alt="ADBlock-PD Logo" src="https://raw.githubusercontent.com/weby-homelab/adblock-pd/master/logo-light.svg" width="400">
  </picture>
</p>

<p align="center">
  <em>The ultimate, heavily hardened fork of AdGuard Home for those who demand absolute privacy.</em>
</p>

<p align="center">
  <a href="https://hub.docker.com/r/webyhomelab/adblock-pd"><img src="https://img.shields.io/docker/pulls/webyhomelab/adblock-pd?style=for-the-badge&logo=docker&color=00d4ff" alt="Docker Pulls"></a>
  <a href="https://github.com/weby-homelab/adblock-pd/releases/latest"><img src="https://img.shields.io/github/v/release/weby-homelab/adblock-pd?style=for-the-badge&logo=github&color=0072ff" alt="Latest Release"></a>
  <a href="https://github.com/weby-homelab/adblock-pd/blob/master/LICENSE.txt"><img src="https://img.shields.io/badge/License-GPL_v3-blue.svg?style=for-the-badge" alt="License"></a>
</p>

---

## 📸 Interface (Bento UI / Glassmorphism)

<p align="center">
  <img src="ADBlock-PD-Light.png" width="32%" alt="Light Theme 1" />
  <img src="ADBlock-PD-Light-1.png" width="32%" alt="Light Theme 2" />
  <img src="ADBlock-PD-Light-2.png" width="32%" alt="Light Theme 3" />
</p>
<p align="center">
  <img src="ADBlock-PD-Dark.png" width="32%" alt="Dark Theme 1" />
  <img src="ADBlock-PD-Dark-1.png" width="32%" alt="Dark Theme 2" />
  <img src="ADBlock-PD-Dark-2.png" width="32%" alt="Dark Theme 3" />
</p>

## 🎯 What is this?
**ADBlock-Private-DNS (ADBlock-PD)** is a fork of the popular AdGuard Home DNS server (based on v0.107.77). Created by **Weby Homelab**, the sole purpose of this project is to **completely sever all ties with AdGuard's infrastructure (and any other external entity)**. We took a powerful filtering engine and fully "sanitized" it.

Your DNS server should belong exclusively to you. No telemetry, no hidden requests, and no downloading of third-party code without your explicit consent.

### 🏗️ Architecture

```mermaid
flowchart TD
    %% Styling
    classDef client fill:#0072ff,stroke:#00d4ff,stroke-width:2px,color:#fff,font-weight:bold,rx:10,ry:10;
    classDef server fill:#1a1a2e,stroke:#00d4ff,stroke-width:3px,color:#fff,font-weight:bold,rx:15,ry:15;
    classDef filter fill:#ff4d00,stroke:#ff8c00,stroke-width:2px,color:#fff,font-weight:bold,rx:10,ry:10;
    classDef block fill:#e84545,stroke:#903749,stroke-width:2px,color:#fff,stroke-dasharray: 5 5,rx:10,ry:10;
    classDef upstream fill:#16c79a,stroke:#11999e,stroke-width:2px,color:#fff,font-weight:bold,rx:10,ry:10;
    classDef cache fill:#f0a500,stroke:#cf7500,stroke-width:2px,color:#fff,font-weight:bold,rx:10,ry:10;
    classDef dummy fill:#555,stroke:#333,stroke-width:2px,color:#fff,stroke-dasharray: 5 5,rx:10,ry:10;

    %% Components
    subgraph Clients ["📱 Your Devices"]
        direction LR
        C1("💻 Laptops / PC"):::client
        C2("📱 Smartphones"):::client
        C3("📺 Smart TV / IoT"):::client
    end

    subgraph Docker ["🐳 Docker Environment (debian:bullseye-slim)"]
        direction TB
        
        subgraph ADBlock_PD ["🛡️ ADBlock-Private-DNS Core"]
            direction TB
            Proto["🔒 DNS Listeners<br>(DoH, DoT, DoQ, UDP/TCP 53)"]:::server
            Engine["⚡ Core DNS Engine<br>& Query Processor"]:::server
            Cache[("🗄️ DNS Cache<br>(Instant Responses)")]:::cache
            
            Proto <--> Engine
            Engine <--> Cache
        end
        
        subgraph Hardening ["🔒 Privacy & Hardening Interceptors"]
            direction LR
            H1["🔇 WHOIS Privacy<br>(Empty Stub)"]:::dummy
            H2["🛑 SafeBrowsing<br>(Redirected to 127.0.0.1)"]:::dummy
            H3["🚀 Auto-Updater<br>(Physically Purged)"]:::dummy
        end
        
        Engine -.->|Sanitized Queries| Hardening
        
        subgraph Logic ["⚙️ Advanced Filtering Logic"]
            direction TB
            Rules1["📝 Custom User Rules"]:::filter
            Rules2["🛡️ Security Blocklists<br>(Ukr Security, AdAway)"]:::filter
        end
        
        Engine ==> Logic
    end

    subgraph Outcomes ["🌐 Resolution Outcomes"]
        direction LR
        Null["🕳️ Blackhole<br>(0.0.0.0 / NXDOMAIN)"]:::block
        U1["🌍 Encrypted Upstreams<br>(Cloudflare, ControlD, etc.)"]:::upstream
    end

    subgraph Monitoring ["🏥 System Reliability"]
        HC["🩺 DNS Healthcheck<br>(Checked every 30s)"]:::upstream
    end

    %% Connections
    C1 -->|Encrypted DNS| Proto
    C2 -->|Encrypted DNS| Proto
    C3 -->|Plain DNS| Proto

    Logic =="Ads / Trackers / Phishing"==> Null
    Logic =="Clean & Safe Traffic"==> U1
    
    HC -.->|Probes port 53| Proto
    HC -.->|Auto-restarts on failure| Docker
```

## ✨ Key Hardening Features

### 🚀 Purged Updater (RCE Protection)
The original `internal/updater` module has been physically removed from the source code. The server will **never** ping `static.adtidy.org` to check for updates. This eliminates a potential Remote Code Execution (RCE) vector through compromised updates or malicious infrastructure.

### 🔒 DNS Hardening & Zero Phone-Home
- **SafeBrowsing & Parental Control:** In the original version, these features send hashes of your queries to AdGuard servers. In **ADBlock-PD**, the suffixes are nullified, and requests are forcefully redirected to `127.0.0.1`. The features are isolated.
- **WHOIS Privacy:** The built-in WHOIS client has been replaced with a dummy stub (`whois.Empty`). The IP addresses of devices in your network are never transmitted to external servers (e.g., `whois.arin.net`).

### 🔇 Zero Telemetry & Rebranding
All links in the web interface (React) that led to AdGuard trackers, analytics, or external documentation have been replaced with neutral stubs. The project has received a new adaptive SVG logo and a complete visual identity overhaul.

### 🔄 Self-Healing Architecture
The container is equipped with a built-in `HEALTHCHECK` based on the `host` utility. Every 30 seconds, the system verifies the vitality of the DNS resolver (`127.0.0.1:53`). If the service hangs, Docker will automatically restart it, guaranteeing a stable internet connection in your network.

### 🐧 Lightweight & Secure Base
The final Docker image is based on `debian:bullseye-slim`. The service runs as an unprivileged user (`UID 10001`), utilizing the `--no-permcheck` flag for a secure startup in an isolated Docker environment. The default timezone is set to `Europe/Kyiv`.

---

## 🚀 Deployment & Configuration

For a proper deployment (running Docker, completing the setup wizard, and installing SSL certificates for DoH/DoT/DoQ), please refer to our detailed guide:

📖 **[Complete Installation Guide (INSTRUCTIONS_INSTALL_ENG.md)](INSTRUCTIONS_INSTALL_ENG.md)**

---

## 🛠 Building from Source

If you want to build the project yourself, you will need Docker (for the multi-stage build process).

```bash
git clone https://github.com/weby-homelab/adblock-pd.git
cd adblock-pd
docker build -t adblock-pd:local .
```

---

## 🌿 Branches & Development Model

To maintain stability and supply-chain security, the repository adheres to a strict branching workflow:
- **`master`** — the primary release branch. Direct commits are restricted. All integrations must go through Pull Requests and pass automated CI testing. Merging into `master` triggers a build and deploy of the `latest` image to Docker Hub.
- **`feature/*`** — branches for developing new features (e.g., visual or functional upgrades).
- **`fix/*`** — branches for security patches and bug fixes.
- **`release/*`** — short-lived branches for release compilation and packaging.

Every version tag and merge commit is cryptographically GPG-signed to verify code authenticity and integrity.

---

## 📜 License & Disclaimer

This project is distributed under the **GNU General Public License v3.0 (GPL-3.0)**. 
It is provided "AS IS". The Weby Homelab team assumes no liability for any network disruptions, data loss, or other consequences resulting from the use of this software.

---

<br>
<p align="center">
  Built in Ukraine under air raid sirens &amp; blackouts ⚡<br>
  &copy; 2026 Weby Homelab
</p>
