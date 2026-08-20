# Soal Logika — Test Kandidat Diya (Golang)

> Pecahkan dengan menulis kode Go. Ceritakan dulu pendekatanmu sebelum ngoding, baru implementasikan.

---

## Periode Transaksi Terbaik

```go
func BestPeriod(transaksi []int) int
```

Diberi daftar transaksi harian sebuah unit usaha (positif = pemasukan, negatif = pengeluaran). Cari **total terbesar** dari satu periode berurutan (hari-hari yang bersambung). Minimal ambil satu hari.

### Contoh

| Input | Output | Keterangan |
|---|---|---|
| `[-2, 1, -3, 4, -1, 2, 1, -5, 4]` | `6` | hari ke-4 s/d ke-7: `4 + -1 + 2 + 1` |
| `[-3, -1, -2]` | `-1` | kalau semua rugi, ambil hari yang paling kecil ruginya |
| `[5]` | `5` | satu hari saja |

### Constraint

- Panjang bisa sampai 100.000 hari.
- Boleh didiskusikan efisiensinya.

---

<!-- ===================================================================== -->
<!-- BAGIAN DI BAWAH INI UNTUK PENILAI — hapus sebelum diberikan ke kandidat -->
<!-- ===================================================================== -->

## Kunci Jawaban & Rubrik (untuk penilai)

### Brute-force O(n²)

Masih diterima untuk junior. Coba semua pasangan awal–akhir:

```go
func BestPeriod(t []int) int {
    best := t[0]
    for i := 0; i < len(t); i++ {
        sum := 0
        for j := i; j < len(t); j++ {
            sum += t[j]
            if sum > best {
                best = sum
            }
        }
    }
    return best
}
```

### Solusi Optimal O(n) — Kadane's Algorithm

Yang dicari dari kandidat mid ke atas:

```go
func BestPeriod(t []int) int {
    best, cur := t[0], t[0]
    for i := 1; i < len(t); i++ {
        // lanjutkan periode sebelumnya, atau mulai baru dari hari ini
        if cur+t[i] > t[i] {
            cur = cur + t[i]
        } else {
            cur = t[i]
        }
        if cur > best {
            best = cur
        }
    }
    return best
}
```

Inti idenya satu kalimat: *di tiap hari, putuskan lanjut periode berjalan atau mulai baru mana yang lebih besar.*

### Penilaian

- ⭐ Menghasilkan solusi benar (walau O(n²)) + sadar kasus semua-negatif → paham dasar.
- ⭐⭐ Sampai ke Kadane, atau O(n²) yang bersih plus tahu cara memperbaikinya.
- ⭐⭐⭐ Kadane benar, jelaskan kenapa `best` diinisialisasi `t[0]` (bukan `0`), dan bisa cerita variasi (misal: kalau diminta *indeks* periodenya, bukan cuma totalnya).

### Jebakan Umum

Inisialisasi `best := 0`. Ini lolos contoh pertama tapi **salah** untuk `[-3, -1, -2]` (mengembalikan `0`, bukan `-1`). Kandidat yang ketahuan jebakan ini lalu memperbaikinya sendiri = sinyal bagus.