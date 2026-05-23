# Soal 2 — Review Website AYO (ayo.co.id)

## 1. Bug
- Saya pengguna baru, Ketika saya login menggunakan google, saya di arahkan untuk isi password dan validasi OTP, saat saya pertama kali klik `Kirim kode melalui WhatsApp` muncul error `Too Many Request`, setelah saya tunggu 1-2 menit sudah aman.
- Di header website menu `Liga Ayo` tidak bisa di buka, ter-direct ke url `https://ligaayo.com/`

---

## 2. Aksesbilitas
- Beberapa modal dan dropdown tidak bisa ditutup dengan Escape, contoh `Date` dan `Filter Waktu` di menu `Pilih Lapangan`.

---

## 3. UI/UX
- Logo `Ayo` pada footer tidak sejajar.

---

## 4. Konten
- Tidak ada halaman FAQ yang mudah ditemukan, pertanyaan seperti "Bagaimana jika venue tiba-tiba tutup?" atau "Apakah ada refund?" tidak terjawab di landing page.
- Harga venue tidak ditampilkan di listing utama, user harus masuk ke detail page baru tahu harga.

---

## 5. Keamanan
- Halaman login tidak terlihat memiliki CAPTCHA atau rate limiting yang terekspos ke user, rentan brute force.

---

# Soal 3 — Review Mobile App AYO (ayo.co.id)

## 1. Bug
- Sebagai pengguna baru, setelah berhasil login, di halaman dashboard, di samping notifikasi icon muter-muter terus, aman jika sudah memilih minimal 1 opsi.
- Di halaman `Participants` ada beberapa user yang tidak bisa di klik untuk masuk ke halaman detail user.
- Di halaman detail user, ketika geser carousel jenis olahraga, section `Community` selalu ke refresh otomatis dan tidak konsisten.
- Halaman profile di navbar bawah kanan, ketika di klik masuk ke profile orang lain yang sebelumnya sudah di lihat, harus keluar app dulu agar bisa masuk ke profile pribadi.

## 2. Aksesbilitas
- Di halaman `Participants` lalu klik salah satu user dan masuk ke halaman detail user, ketika back menggunakan navigasi android (bukan dari tombol back app diatas kiri), langsung kembali ke halaman utama, yang mana seharusnya balik ke halaman `Participants`.
- Beberapa halaman lain juga ketika `back` menggunakan navigasi dari android langsung balik ke halaman `Home`, seharusnya balik ke halaman sebelumnya.

## 3. Saran
- Di halmaan review disarankan bisa menambahkan beberapa image.