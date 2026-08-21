# Suara Terbanyak (Mayoritas Mutlak)

Tulis jawabanmu di `main.go` pada folder ini.

## Deskripsi

```go
func Mayoritas(suara []int) int
```

Diberi hasil pemilihan ketua tiap angka adalah ID kandidat yang dipilih satu pemilih. Cari kandidat yang mendapat **lebih dari setengah** total suara (mayoritas mutlak). Kalau tidak ada yang mencapai itu, kembalikan `-1`.

## Contoh

| Input | Output | Keterangan |
|---|---|---|
| `[1, 2, 1, 1, 3, 1]` | `1` | kandidat `1` dapat 4 dari 6 suara (> 3) |
| `[1, 2, 3, 1, 2]` | `-1` | tidak ada yang > 2.5 suara |
| `[7]` | `7` | satu suara, otomatis mayoritas |

## Constraint

- Jumlah suara bisa sampai 1.000.000.

## Bonus

- Bisakah kamu selesaikan dengan memori **O(1)** tanpa map atau array bantu?

---

> Ingat aturan main di README utama: **dilarang menggunakan AI** untuk menjawab. Boleh membuka project lama dan dokumentasi Go untuk sintaks.
