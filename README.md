# ❌ No-as-a-Service-Go

<p align="center">
  <img src="https://raw.githubusercontent.com/hotheadhacker/no-as-a-service/main/assets/imgs/naas-with-no-logo-bunny.png" width="800" alt="No-as-a-Service Banner" width="70%"/>
</p>


Ever needed a graceful way to say “no”?  
This tiny API returns random, generic, creative, and sometimes hilarious rejection reasons — perfectly suited for any scenario: personal, professional, student life, dev life, or just because.

Built for humans, excuses, and humor.

---

## 🚀 API Usage

**Base URL**
```
https://localhost:8080/no
```

**Method:** `GET`  
**Rate Limit:** `120 requests per minute per IP`

### 🔄 Example Request
```http
GET /no
```

### ✅ Example Response
```json
{
  "reason": "This feels like something Future Me would yell at Present Me for agreeing to."
}
```

Use it in apps, bots, landing pages, Slack integrations, rejection letters, or wherever you need a polite (or witty) no.

---

## 🛠️ Self-Hosting

Want to run it yourself? It’s lightweight and simple.

### 1. Clone this repository
```bash
git clone https://github.com/itsnuyen/no-as-a-service-go.git
cd no-as-a-service-go
```

### 3. Start the server
```bash
go run main.go
```

The API will be live at:
```
http://localhost:8080/no
```

You can also change the port using an environment variable:
```bash
SERVER_PORT=8099 go run main.go
```

---

## 📁 Project Structure

```
no-as-service/
├── Dockerfile          # Minimal Docker File
├── main.go             # Standard Go Http Service
├── reasons.json        # 1000+ universal rejection reasons
└── README.md
```

---

> Want to use no-as-a-service in your own project? Check out the usage section in this README and start returning **"no"** like a pro.
---

## 👤 Author

Created with creative stubbornness by [hotheadhacker](https://github.com/hotheadhacker)

Ported to go by [itsnuyen](https://github.com/itsnuyen)

---

## 📄 License

MIT — do whatever, just don’t say yes when you should say no.
