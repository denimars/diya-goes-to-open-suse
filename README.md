# Diya Goes To openSUSE


Repo ini berisi serangkaian **soal logika** yang harus kamu pecahkan menggunakan bahasa **Go**. Yang kami nilai bukan sekadar jawaban benar, tapi **cara kamu berpikir** bagaimana kamu memilih pendekatan, menangani kasus, dan menuliskannya dengan kode yang bersih.

---

## 📋 Cara Mengerjakan

1. **Fork** repo ini ke akun GitHub-mu:
   👉 https://github.com/denimars/diya-goes-to-open-suse

2. **Clone** hasil fork-mu ke komputer lokal:
   ```bash
   git clone https://github.com/<username-kamu>/diya-goes-to-open-suse
   cd diya-goes-to-open-suse
   ```

3. **Kerjakan soal di setiap folder.** Setiap folder berisi **satu soal** (baca `README.md` di dalam folder tersebut). Tulis jawabanmu di file Go pada folder yang sama.

4. Pastikan kodemu **bisa dijalankan**:
   ```bash
   go run ./soal-01-xxx      # menjalankan satu soal
   go test ./...             # menjalankan semua test (jika disediakan)
   ```

5. **Commit & push** ke fork-mu:
   ```bash
   git add .
   git commit -m "jawaban tes logika Diya"
   git push origin main
   ```

6. **Kirim link fork-mu** ke Grup diya project batch 5

---

## ⚠️ Aturan Main

**Yang DILARANG KERAS (haram):**
- ❌ Menggunakan AI (ChatGPT, Claude, GitHub Copilot, Gemini, dan sejenisnya) untuk **menjawab / menghasilkan solusi** soal. Tes ini untuk mengukur logikamu, bukan logika mesin. Jawaban yang terdeteksi hasil AI langsung gugur.

**BOLEH:**
- ✅ Membuka **project lama** milikmu sendiri sebagai referensi.
- ✅ Membuka dokumentasi resmi Go ([go.dev](https://go.dev/doc/)), mencari **sintaks** atau fungsi standard library.
- ✅ Menggunakan editor/IDE apa pun (tanpa fitur AI completion yang menuliskan solusi untukmu).
- ✅ Bertanya jika **soalnya sendiri** kurang jelas.

> Kejujuran adalah bagian dari penilaian. Lebih baik menyerahkan solusi sederhana yang benar-benar hasil pemikiranmu, daripada solusi sempurna yang bukan buatanmu.

---

## 🗂️ Struktur Folder

Setiap soal berada dalam foldernya sendiri:

```
diya-goes-to-open-suse/
├── README.md              <- file ini
├── go.mod
├── soal-01-xxxxx/
│   ├── README.md          <- deskripsi soal
│   └── main.go            <- tulis jawabanmu di sini
├── soal-02-xxxxx/
│   ├── README.md
│   └── main.go
└── ...
```

Kerjakan **semua folder**. Kalau ada soal yang tidak sempat/tidak bisa kamu selesaikan, biarkan apa adanya dan beri catatan singkat di `main.go`-nya — itu lebih baik daripada dikosongkan tanpa keterangan.

---

## ✅ Yang Kami Nilai

- **Kebenaran** — solusi menangani contoh dan kasus tepi (termasuk yang tidak disebutkan).
- **Kejelasan** — kode mudah dibaca, penamaan variabel masuk akal.
- **Cara berpikir** — pilihan struktur data & algoritma yang tepat, bukan asal jalan.
- **Kejujuran** — hasil pemikiran sendiri (lihat Aturan Main di atas).
