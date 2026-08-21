# Soal 04 Urutan Deploy Modul

Tulis jawabanmu di `main.go` pada folder ini.

## Deskripsi

```go
func UrutanDeploy(jumlahModul int, ketergantungan [][]int) []int
```

Sebuah sistem punya `jumlahModul` modul, diberi nomor `0` sampai `jumlahModul-1`. Beberapa modul bergantung pada modul lain: pasangan `[a, b]` berarti modul `a` baru boleh dideploy **setelah** modul `b` selesai.

Kembalikan **satu urutan deploy yang valid** — semua modul dideploy, dan setiap ketergantungan terpenuhi. Kalau mustahil (ada ketergantungan yang saling melingkar), kembalikan slice kosong `[]int{}`.

## Contoh

| Input | Output (salah satu yang valid) | Keterangan |
|---|---|---|
| `n=4, [[1,0],[2,0],[3,1],[3,2]]` | `[0,1,2,3]` | `0` dulu, lalu `1` & `2`, baru `3` |
| `n=2, [[0,1],[1,0]]` | `[]` | melingkar → mustahil |
| `n=3, []` | `[0,1,2]` | tanpa ketergantungan, urutan bebas |
| `n=5, [[1,0],[3,2]]` | `[0,2,1,3,4]` | ada modul yang berdiri sendiri (`4`) |

> Catatan: jawaban yang benar bisa **lebih dari satu**. Yang penting semua ketergantungan terpenuhi dan seluruh modul muncul tepat sekali.

## Constraint

- `jumlahModul` sampai 100.000.
- Jumlah ketergantungan sampai 200.000.
- Solusi harus efisien **O(V+E)**.

---

> Ingat aturan main di README utama: **dilarang menggunakan AI** untuk menjawab. Boleh membuka project lama dan dokumentasi Go untuk sintaks.
