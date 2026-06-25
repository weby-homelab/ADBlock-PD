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
    <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/Weby-Homelab/ADBlock-PD/master/logo-dark.svg">
    <source media="(prefers-color-scheme: light)" srcset="https://raw.githubusercontent.com/Weby-Homelab/ADBlock-PD/master/logo-light.svg">
    <img alt="Логотип ADBlock-PD" src="https://raw.githubusercontent.com/Weby-Homelab/ADBlock-PD/master/logo-light.svg" width="400">
  </picture>
</p>

<p align="center">
  <em>Ультимативний, жорстко захищений форк популярного DNS-сервера для тих, хто вимагає абсолютної приватності.</em>
</p>

<p align="center">
  <a href="https://hub.docker.com/r/webyhomelab/adblock-pd"><img src="https://img.shields.io/docker/pulls/webyhomelab/adblock-pd?style=for-the-badge&logo=docker&color=00d4ff" alt="Завантаження з Docker"></a>
  <a href="https://github.com/Weby-Homelab/ADBlock-PD/releases/latest"><img src="https://img.shields.io/github/v/release/Weby-Homelab/ADBlock-PD?style=for-the-badge&logo=github&color=0072ff" alt="Останній реліз"></a>
  <a href="https://github.com/Weby-Homelab/ADBlock-PD/blob/master/LICENSE.txt"><img src="https://img.shields.io/badge/Ліцензія-GPL_v3-blue.svg?style=for-the-badge" alt="Ліцензія"></a>
</p>

---

## 📸 Інтерфейс (Bento UI / Glassmorphism)

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

## 🎯 Що це таке?
**ADBlock-Private-DNS (ADBlock-PD)** — це власна розробка на базі відомого DNS-сервера AdGuard Home (версії 0.107.77). Проєкт створений командою **Weby Homelab** з єдиною метою: **повністю усунути будь-які зв'язки з інфраструктурою початкових розробників та будь-якою іншою зовнішньою мережею**. Ми взяли потужний рушій фільтрації та провели його повну "стерилізацію". 

Ваш DNS-сервер повинен належати лише вам. Ніякого збору даних, жодних прихованих запитів, жодного завантаження стороннього коду без вашого відома.

### 🏗️ Архітектура (Architecture)

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
    subgraph Clients ["📱 Ваші пристрої"]
        direction LR
        C1("💻 Ноутбуки / ПК"):::client
        C2("📱 Смартфони"):::client
        C3("📺 Smart TV / IoT"):::client
    end

    subgraph Docker ["🐳 Docker Середовище (debian:bullseye-slim)"]
        direction TB
        
        subgraph ADBlock_PD ["🛡️ Ядро ADBlock-Private-DNS"]
            direction TB
            Proto["🔒 DNS Слухачі<br>(DoH, DoT, DoQ, UDP/TCP 53)"]:::server
            Engine["⚡ Головний рушій DNS<br>та обробник запитів"]:::server
            Cache[("🗄️ Кеш DNS<br>(Миттєві відповіді)")]:::cache
            
            Proto <--> Engine
            Engine <--> Cache
        end
        
        subgraph Hardening ["🔒 Блоки приватності (Hardening)"]
            direction LR
            H1["🔇 WHOIS Privacy<br>(Порожня заглушка)"]:::dummy
            H2["🛑 SafeBrowsing<br>(Переспрямовано на 127.0.0.1)"]:::dummy
            H3["🚀 Auto-Updater<br>(Фізично видалено)"]:::dummy
        end
        
        Engine -.->|Стерилізація запитів| Hardening
        
        subgraph Logic ["⚙️ Розширена логіка фільтрації"]
            direction TB
            Rules1["📝 Правила користувача"]:::filter
            Rules2["🛡️ Блок-листи безпеки<br>(Ukr Security, AdAway)"]:::filter
        end
        
        Engine ==> Logic
    end

    subgraph Outcomes ["🌐 Результат роздільної здатності (Оutcomes)"]
        direction LR
        Null["🕳️ Чорна діра<br>(0.0.0.0 / NXDOMAIN)"]:::block
        U1["🌍 Зашифровані Upstream-сервери<br>(Cloudflare, ControlD тощо)"]:::upstream
    end

    subgraph Monitoring ["🏥 Надійність системи"]
        HC["🩺 DNS Healthcheck<br>(Перевірка кожні 30 сек)"]:::upstream
    end

    %% Connections
    C1 -->|Зашифрований DNS| Proto
    C2 -->|Зашифрований DNS| Proto
    C3 -->|Стандартний DNS| Proto

    Logic =="Реклама / Трекери / Фішинг"==> Null
    Logic =="Чистий та безпечний трафік"==> U1
    
    HC -.->|Опитує порт 53| Proto
    HC -.->|Авто-рестарт при збої| Docker
```

## ✨ Ключові відмінності та посилення безпеки

### 🚀 Видалення модуля оновлень (Захист від віддаленого виконання коду)
Оригінальний модуль `internal/updater` фізично видалено з вихідного коду. Сервер **ніколи** не буде звертатися до сторонніх серверів для перевірки оновлень. Це ліквідує потенційний шлях віддаленого виконання коду через підміну файлів оновлень або злам інфраструктури розробника.

### 🔒 Захист DNS та відсутність прихованих підключень
- **Безпечний перегляд та Батьківський контроль:** В оригіналі ці функції надсилають часткові дані ваших запитів на зовнішні сервери. В **ADBlock-PD** ці адреси жорстко стерті, а запити перенаправляються на локальну адресу `127.0.0.1`. Функції повністю ізольовані.
- **Конфіденційність WHOIS:** Вбудований інструмент запитів WHOIS замінено на порожню "заглушку". IP-адреси пристроїв у вашій мережі більше не передаються зовнішнім сервірам.

### 🔇 Відсутність телеметрії та новий дизайн
Усі посилання у веб-інтерфейсі, що вели на трекери, системи аналітики або зовнішню документацію, замінено на нейтральні заглушки. Проєкт отримав новий адаптивний векторний логотип та повністю незалежний зовнішній вигляд.

### 🔄 Архітектура самовідновлення
Контейнер оснащено вбудованою перевіркою стану (`HEALTHCHECK`) на базі утиліти `host`. Система кожні 30 секунд перевіряє життєздатність DNS-служби (`127.0.0.1:53`). Якщо служба "зависає", Docker автоматично її перезапускає, гарантуючи стабільний інтернет у вашій мережі.

### 🐧 Легка та безпечна основа
Фінальний образ Docker базується на мінімалістичній операційній системі `debian:bullseye-slim`. Служба запускається від імені звичайного користувача (`UID 10001`), з доданим параметром `--no-permcheck` для безпечного запуску в ізольованому середовищі Docker. За замовчуванням встановлено київський час (`Europe/Kyiv`).

---

## 🚀 Запуск та Налаштування

Для правильного розгортання проєкту (запуск Docker, проходження майстра налаштування та встановлення SSL-сертифікатів для DoH/DoT/DoQ), будь ласка, ознайомтеся з нашим детальним посібником:

📖 **[Повний посібник зі встановлення (INSTRUCTIONS_INSTALL.md)](INSTRUCTIONS_INSTALL.md)**

---

## 🛠 Збирання з вихідного коду

Якщо ви хочете зібрати проєкт самостійно, вам знадобиться Docker (використовується багатоетапне збирання). 

```bash
git clone https://github.com/Weby-Homelab/ADBlock-PD.git
cd adblock-pd
docker build -t adblock-pd:local .
```

---

## 🌿 Гілки та модель розробки

Для забезпечення стабільності та безпеки у репозиторії діє суворий регламент роботи з гілками:
- **`master`** — основна робоча гілка. Прямі коміти в неї заборонені. Будь-які зміни інтегруються виключно через Pull Requests після обов'язкового проходження автоматичних тестів CI. Злиття в `master` автоматично оновлює образ `latest` на Docker Hub.
- **`feature/*`** — гілки для розробки нової функціональності (наприклад, візуальних змін).
- **`fix/*`** — гілки для усунення вразливостей або виправлення багів.
- **`release/*`** — гілки підготовки та публікації стабільних версій.

Кожен реліз супроводжується обов'язковим GPG-підписом тегу та комітів для гарантії ланцюжка безпеки постачання (supply-chain security).

---

## 📜 Ліцензія та Відмова від відповідальності

Цей проєкт розповсюджується під ліцензією **GNU General Public License v3.0 (GPL-3.0)**. 
Проєкт надається "ЯК Є". Команда Weby Homelab не несе відповідальності за будь-які збої в роботі мережі, втрату даних або інші наслідки використання цього програмного забезпечення.

---

<br>
<p align="center">
  Built in Ukraine under air raid sirens &amp; blackouts ⚡<br>
  &copy; 2026 Weby Homelab
</p>
