📂 Excel & CSV Merger – Versi Golang 🚀

🔄 Aplikasi ini memungkinkan pengguna menggabungkan banyak file Excel (.xlsx) dan CSV menjadi satu dataset besar dalam format CSV.

📌 Fitur Utama

✅ Mudah digunakan – GUI sederhana & intuitif
✅ Dukungan format – Bisa gabungkan .xlsx & .csv
✅ Otomatis menyimpan hasil dalam satu file CSV
✅ Menggunakan Go + Fyne – Performa tinggi

📸 Screenshot

<img width="612" alt="image" src="https://github.com/user-attachments/assets/242444f8-c115-4e90-a6bd-94e2177e29d2" />
<img width="677" alt="image" src="https://github.com/user-attachments/assets/671a9add-594e-41fd-a63d-69f20403f39a" />

💻 Teknologi yang Digunakan
	•	Golang
	•	Fyne (GUI)
	•	Excelize (Untuk membaca Excel)
	•	Encoding/CSV (Untuk membaca & menulis CSV)

📌 Instalasi & Cara Menjalankan

1️⃣ Clone Repository

git clone https://github.com/santoso-py/go-excelcsv-merger.git
cd excel-csv-merger-go

2️⃣ Install Dependencies

go mod tidy

3️⃣ Jalankan Aplikasi

go run main.go

🔥 Cara Penggunaan

1️⃣ Pilih folder berisi file Excel/CSV yang ingin digabungkan
2️⃣ Pilih folder tujuan untuk menyimpan file hasil
3️⃣ Klik tombol “Merge Files”
4️⃣ File hasil akan tersimpan dalam format combined.csv di folder tujuan 🎉

📜 Lisensi

📄 MIT License – Bebas digunakan, dimodifikasi, dan didistribusikan dengan mencantumkan kredit.

💡 Kontribusi & Feedback

🎯 Masukan dan kontribusi sangat dihargai! Feel free to open an issue or pull request. 🚀

🔗 GitHub Repository
📌 https://github.com/santoso-py/go-excelcsv-merger

🔥 Penjelasan Singkat:

	•	Aplikasi ini dibuat menggunakan Golang dengan GUI berbasis Fyne.
 
	•	Excelize digunakan untuk membaca file .xlsx.
 
	•	encoding/csv digunakan untuk membaca dan menulis file .csv.
 
	•	Fitur utama: Pilih folder input dan output, lalu gabungkan semua file dalam satu klik.
 
	•	Hasil akhirnya: File combined.csv tersimpan di folder tujuan.

🔥 Kelebihan dibanding versi Python:
	•	Eksekusi lebih cepat, karena Go lebih optimal dalam pemrosesan data.
 
	•	Standalone binary (tidak butuh interpreter seperti Python).
 
	•	Lebih ringan, karena tidak ada dependency berat seperti Pandas.
