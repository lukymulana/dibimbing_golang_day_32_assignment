# Inventory Management API Using Vercel

## Deskripsi
API untuk manajemen inventaris yang berjalan di Vercel. API ini memungkinkan pengguna untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada item inventaris tanpa bergantung pada database eksternal, menggunakan data dummy yang disimpan dalam memori.

## Fitur
- **GET /api/items**: Mengembalikan daftar semua item dalam inventaris.
- **POST /api/items**: Menambahkan item baru ke dalam inventaris.
- **POST /api/items/update/:id**: Memperbarui item yang sudah ada berdasarkan ID.
- **GET /api/items/delete/:id**: Menghapus item berdasarkan ID.
