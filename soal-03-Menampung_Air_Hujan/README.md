# Soal 03 Menampung Air Hujan

Tulis jawabanmu di `main.go` pada folder ini.

## Deskripsi

```go
func TampungAir(tinggi []int) int
```

Diberi tinggi deretan tembok/gedung yang berdempetan (lebar tiap tembok = 1). Setelah hujan, air tertampung di cekungan antar tembok. Hitung **total satuan air** yang tertampung.

## Contoh

| Input | Output | Keterangan |
|---|---|---|
| `[0,1,0,2,1,0,1,3,2,1,2,1]` | `6` | air mengisi lembah-lembah di antara tembok |
| `[4,2,0,3,2,5]` | `9` | cekungan besar di tengah |
| `[3,3,3]` | `0` | rata, tidak ada cekungan |
| `[5]` | `0` | satu tembok tidak menampung apa pun |

## Constraint

- Panjang bisa sampai 1.000.000.

## Bonus

- Bisakah kamu selesaikan dengan memori **O(1)** tanpa array bantu?

---

> Ingat aturan main di README utama: **dilarang menggunakan AI** untuk menjawab. Boleh membuka project lama dan dokumentasi Go untuk sintaks.
