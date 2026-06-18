package main

import (
	"bytes"         // Untuk mengolah data request sebelum dikirim ke internet
	"encoding/json" // Untuk mengubah objek data menjadi format JSON (dan sebaliknya)
	"fmt"           // Untuk menampilkan log teks di terminal server
	"log"           // Untuk mencatat log jika terjadi masalah fatal pada server
	"net/http"      // Core library utama Go untuk menjalankan HTTP Web Server
	"os"            // Untuk mendeteksi Port otomatis dari server hosting cloud
	"strings"       // Untuk manipulasi string jika diperlukan
	"time"          // Untuk membuat stempel waktu unik pada Order ID
)

// Blueprint data yang dikirim oleh Next.js (Frontend)
type OrderRequest struct {
	PaketNama string `json:"paketNama"` // Nama paket taawun atau buku
	Amount    int64  `json:"amount"`    // Nominal uang (wajib angka bulat)
	Email     string `json:"email"`     // Email donatur/pembeli
}

// Blueprint data respon balik berisi Token Snap untuk Next.js
type OrderResponse struct {
	SnapToken string `json:"snapToken"`
}

// Blueprint data profil Kak Alif untuk portofolio digital
type ProfilResponse struct {
	Nama        string   `json:"nama"`
	Gelar       string   `json:"gelar"`
	Keahlian    []string `json:"keahlian"`
	BukuPopuler []string `json:"bukuPopuler"`
	KontakWA    string   `json:"kontakWa"`
}

// Fungsi internal untuk menembak API Midtrans Sandbox secara langsung
func panggilMidtransAPI(req OrderRequest) (string, error) {
	// Kunci Rahasia Midtrans Sandbox Kak Alif
	serverKey := "SB-Mid-Server-XXXXXXXXXXXX"

	// Membuat ID Transaksi unik berbasis waktu agar tidak bentrok di sistem Midtrans
	orderID := fmt.Sprintf("TAAWUN-%d", time.Now().UnixNano())

	// Menyusun struktur data sesuai dokumentasi API Midtrans Snap
	payload := map[string]interface{}{
		"transaction_details": map[string]interface{}{
			"order_id":     orderID,
			"gross_amount": req.Amount,
		},
		"customer_details": map[string]interface{}{
			"email": req.Email,
		},
	}

	// Mengubah objek data menjadi baris teks JSON
	bytesPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Menyiapkan tembakan HTTP POST ke server gerbang Midtrans Sandbox
	urlMidtrans := "https://app.sandbox.midtrans.com/snap/v1/transactions"
	httpReq, err := http.NewRequest("POST", urlMidtrans, bytes.NewBuffer(bytesPayload))
	if err != nil {
		return "", err
	}

	// Mengatur autentikasi keamanan menggunakan metode Basic Auth bawaan Midtrans
	httpReq.SetBasicAuth(serverKey, "")
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept", "application/json")

	// Eksekusi pengiriman data ke Midtrans
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Membaca dan membongkar kotak respon dari Midtrans
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	// Mengambil string token Snap dari dalam data respon
	token, ada := result["token"].(string)
	if !ada {
		return "", fmt.Errorf("gagal mendapatkan token dari midtrans, periksa server key")
	}

	return token, nil
}

// Handler 1: Mengurus pembuatan transaksi otomatis Midtrans Snap
func handleCheckout(w http.ResponseWriter, r *http.Request) {
	// Aturan CORS (Cross-Origin Resource Sharing) agar Next.js di Vercel bisa mengakses API ini
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

	// Jika mendeteksi preflight request dari browser, langsung setujui aman
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Validasi wajib menggunakan metode POST
	if r.Method != "POST" {
		http.Error(w, "Metode HTTP wajib POST!", http.StatusMethodNotAllowed)
		return
	}

	var req OrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Format JSON kiriman rusak", http.StatusBadRequest)
		return
	}

	// Jalankan fungsi tembak API Midtrans
	snapToken, err := panggilMidtransAPI(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Bungkus token ke format JSON dan kirim balik ke Next.js browser donatur
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(OrderResponse{SnapToken: snapToken})
}

// Handler 2: Menyajikan data portofolio & profil Kak Alif riil dari kakalif.jamia.id
func handleProfil(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	dataProfil := ProfilResponse{
		Nama:     "Alif Rezky (Daeng Lewa / Abu Uwais)",
		Gelar:    "M.Pd. (Magister Pendidikan Matematika - Universitas Negeri Makassar)",
		Keahlian: []string{"Professional Mathematics Educator (OSN & UTBK)", "Web Developer (Next.js, React, Supabase, Go)", "Published Book Author"},
		BukuPopuler: []string{
			"Matematika Itu Asyik (100+ Eksemplar Tersebar)",
			"Belajar Python dari Nol",
			"TULIMATIKA (Modul Inklusif Bahasa Isyarat)",
			"Langkah Kecil, Karya Besar",
		},
		KontakWA: "6285256162879",
	}

	json.NewEncoder(w).Encode(dataProfil)
}

// Fungsi utama yang mengendalikan siklus hidup server
func main() {
	// Pendaftaran jalur URL Endpoint API Go
	http.HandleFunc("/api/v1/taawun/checkout", handleCheckout)
	http.HandleFunc("/api/v1/profil", handleProfil)

	// Mengambil port dinamis dari lingkungan cloud hosting (seperti Render / Railway)
	// Jika berjalan di laptop lokal, otomatis menggunakan fallback port 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Alhamdulillah, Server Backend Go sukses aktif di port :%s...\n", port)
	
	// Menyalakan server HTTP agar bersiap siaga melayani data secara non-stop
	log.Fatal(http.ListenAndServe(":"+port, nil))
} 
