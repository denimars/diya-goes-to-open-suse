# Soal 05 Cache LRU (Least Recently Used)

Tulis jawabanmu di `main.go` pada folder ini.

## Deskripsi

Rancang sebuah cache berkapasitas tetap. Saat penuh dan ada data baru masuk, data yang **paling lama tidak diakses** harus dibuang.

```go
type LRUCache struct {
    // rancang sendiri
}

func NewLRUCache(kapasitas int) *LRUCache
func (c *LRUCache) Get(key int) (int, bool) // value + apakah ketemu
func (c *LRUCache) Put(key, value int)      // simpan / perbarui
```

## Aturan

- `Get(key)` mengembalikan nilainya dan `true` kalau ada; kalau tidak, `(0, false)`. **Mengakses sebuah key membuatnya jadi "paling baru dipakai".**
- `Put(key, value)` menyimpan atau memperbarui nilai. Kalau kapasitas penuh **dan** key-nya baru, buang key yang **paling lama tidak dipakai** lebih dulu. Memperbarui key yang sudah ada juga menjadikannya "paling baru".
- **Kedua operasi wajib rata-rata O(1).**

## Contoh (kapasitas = 2)

| Operasi | Hasil | Isi cache (baru → lama) |
|---|---|---|
| `Put(1, 10)` | | `{1}` |
| `Put(2, 20)` | | `{2, 1}` |
| `Get(1)` | `10, true` | `{1, 2}` |
| `Put(3, 30)` | buang `2` | `{3, 1}` |
| `Get(2)` | `0, false` | `{3, 1}` |
| `Put(4, 40)` | buang `1` | `{4, 3}` |
| `Get(1)` | `0, false` | `{4, 3}` |
| `Get(3)` | `30, true` | `{3, 4}` |

## Constraint

- Sampai 1.000.000 operasi.
- Kapasitas ≥ 1.
- Solusi O(n) per operasi akan dianggap **gagal** untuk kasus besar.

---

> Ingat aturan main di README utama: **dilarang menggunakan AI** untuk menjawab. Boleh membuka project lama dan dokumentasi Go untuk sintaks.