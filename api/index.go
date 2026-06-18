package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type OrderRequest struct {
	PaketNama string `json:"paketNama"`
	Amount    int64  `json:"amount"`
	Email     string `json:"email"`
}

type OrderResponse struct {
	SnapToken string `json:"snapToken"`
}

// Handler utama yang dieksekusi oleh Vercel Serverless Functions
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Clean path untuk memudahkan routing manual
	p := r.URL.Path

	// ==========================================
	// 1 & 2. ROUTING HALAMAN UTAMA (DASHBOARD & PORTFOLIO - FIX KONTRAS)
	// ==========================================
	if p == "/" || p == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Alif Rezky, M.Pd - Portofolio & Buku</title>
				<style>
					* { box-sizing: border-box; }
					body { font-family: 'Segoe UI', system-ui, sans-serif; background-color: #F8FAFC; color: #1E293B; margin: 0; padding-bottom: 80px; }
					
					.widget-bar { background: white; padding: 12px 20px; position: sticky; top: 0; z-index: 1000; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.05); display: flex; flex-wrap: wrap; justify-content: space-between; align-items: center; max-width: 1000px; margin: 0 auto; border-bottom-left-radius: 16px; border-bottom-right-radius: 16px; gap: 10px; }
					.time-text { font-size: 0.8rem; font-weight: 800; color: #64748B; }
					.time-highlight { color: #0EA5E9; font-weight: 800; }
					.hijri-highlight { color: #059669; font-weight: 800; }

					/* HERO SECTION - FIX KONTRAS TEKS */
					.hero { background: linear-gradient(135deg, #0F172A 0%, #1E3A8A 100%); padding: 80px 20px; color: white; text-align: center; }
					.hero h1 { font-weight: 950; font-size: 3rem; margin: 0; letter-spacing: -2px; color: #FFFFFF; text-shadow: 0 2px 4px rgba(0,0,0,0.3); }
					.hero p { font-size: 1.1rem; color: #93C5FD; font-weight: 600; margin: 15px 0 25px 0; }
					.hero-btn { display: inline-block; color: white; text-decoration: none; font-size: 0.8rem; font-weight: 800; padding: 12px 24px; border-radius: 12px; border: 1px solid rgba(255,255,255,0.4); margin: 5px; background: rgba(255,255,255,0.15); transition: 0.2s; }
					.hero-btn:hover { background: rgba(255,255,255,0.25); border-color: white; }

					.container { max-width: 1000px; margin: 40px auto; padding: 0 20px; display: grid; grid-template-columns: repeat(auto-fit, minmax(300px, 1fr)); gap: 30px; }
					.card { background: white; padding: 30px; border-radius: 24px; border: 1px solid #E2E8F0; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.02); }
					.section-title { font-weight: 900; border-left: 6px solid #118EEA; padding-left: 15px; margin-bottom: 20px; font-size: 1.3rem; color: #0F172A; }
					.section-title.dark { border-left-color: #0F172A; }
					
					.skill-badge { display: inline-block; background: #E0F2FE; color: #0369A1; padding: 6px 12px; border-radius: 8px; font-size: 0.75rem; font-weight: 800; margin: 4px; }
					.exp-item { background: #F8FAFC; padding: 15px; border-radius: 16px; border: 1px solid #E2E8F0; margin-bottom: 12px; }
					.exp-job { font-weight: 900; color: #0F172A; font-size: 0.95rem; }
					.exp-company { font-size: 0.85rem; color: #118EEA; font-weight: 700; }
					.exp-time { font-size: 0.75rem; color: #64748B; margin-top: 4px; }

					.book-sales-card { grid-column: 1 / -1; background: white; border: 1px solid #E2E8F0; padding: 40px 30px; border-radius: 32px; }
					.proof-badge { display: inline-flex; align-items: center; gap: 6px; background: #FEF3C7; color: #B45309; padding: 6px 14px; border-radius: 100px; border: 1px solid #F59E0B; font-size: 0.7rem; font-weight: 950; margin-bottom: 15px; }
					.book-grid { display: grid; grid-template-columns: 1fr; gap: 25px; }
					@media(min-width: 768px) { .book-grid { grid-template-columns: 200px 1fr; } }
					.book-cover { width: 100%; height: 280px; background: #0EA5E9; border-radius: 16px; display: flex; flex-direction: column; justify-content: center; align-items: center; color: white; font-weight: 900; font-style: italic; box-shadow: 0 10px 25px rgba(14,165,233,0.25); text-align: center; padding: 15px; }
					.cta-wa { display: inline-flex; align-items: center; justify-content: center; background: #25D366; color: white; padding: 16px 24px; border-radius: 16px; font-weight: 900; font-size: 0.9rem; text-decoration: none; box-shadow: 0 4px 15px rgba(37,211,102,0.3); width: 100%; text-align: center; margin-top: 15px; transition: 0.2s; }
					.cta-wa:hover { background: #22C55E; }
					
					.testi-box { background: #F8FAFC; border-left: 4px solid #0EA5E9; padding: 15px; border-radius: 8px; margin-top: 15px; font-style: italic; font-size: 0.85rem; color: #475569; }
					footer { text-align: center; margin-top: 40px; font-size: 0.75rem; color: #94A3B8; font-weight: 800; letter-spacing: 1px; }
				</style>
			</head>
			<body>

				<div class="widget-bar">
					<div class="time-text">📅 Masehi: <span class="time-highlight" id="masehi-txt">-</span></div>
					<div class="time-text">🌙 Hijriah: <span class="hijri-highlight">4 Muharram 1448 H</span></div>
					<div class="time-text">🕒 Jam: <span class="time-highlight" id="jam-txt">00.00</span></div>
				</div>

				<div class="hero">
					<h1>Alif Rezky, M.Pd</h1>
					<p>Mathematics Educator | Tech Developer | Author</p>
					<div>
						<a href="https://youtube.com/@kakalifgurumatematika" target="_blank" class="hero-btn">YouTube Channel</a>
						<a href="/jasa/alalify-tech" class="hero-btn" style="background: #0EA5E9; border-color: #0EA5E9;">Alalify Tech Services 🛠️</a>
					</div>
				</div>

				<div class="container">
					<div class="card">
						<div class="section-title">Profil & Keahlian</div>
						<p style="line-height: 1.8; color: #475569; font-size: 0.9rem; text-align: justify; margin: 0;">
							Seorang pendidik matematika profesional berlatar belakang Magister Pendidikan Matematika UNM. Berfokus tinggi pada integrasi teknologi GNU/Linux dan web modern dalam penyusunan metode ajar inklusif dan olimpiade (OSN).
						</p>
						<div style="margin-top: 20px;">
							<span class="skill-badge">GNU/Linux</span>
							<span class="skill-badge">LibreOffice</span>
							<span class="skill-badge">Next.js & Web Dev</span>
							<span class="skill-badge">Python Programming</span>
							<span class="skill-badge">Pendidikan Matematika</span>
						</div>
					</div>

					<div class="card">
						<div class="section-title dark">Pengalaman</div>
						<div class="exp-item">
							<div class="exp-job">Master Teacher OSN</div>
							<div class="exp-company">Edumatrix Indonesia</div>
							<div class="exp-time">Apr 2026 - Sekarang</div>
						</div>
						<div class="exp-item">
							<div class="exp-job">Online Math Tutor</div>
							<div class="exp-company">Algonova</div>
							<div class="exp-time">Mar 2026 - Sekarang</div>
						</div>
						<div class="exp-item">
							<div class="exp-job">Guru Matematika</div>
							<div class="exp-company">SMA IT Al Binaa Islamic Boarding School</div>
							<div class="exp-time">Sept 2022 - Jun 2026</div>
						</div>
					</div>

					<div class="book-sales-card">
						<div class="proof-badge">🔥 100+ EKSEMPLAR CETAK TERSEBAR</div>
						<div class="book-grid">
							<div class="book-cover">
								<span style="font-size: 1.4rem;">MATEMATIKA</span>
								<span style="font-size: 1.4rem; margin-bottom: 10px;">ITU ASYIK</span>
								<span style="font-size: 0.65rem; letter-spacing: 1px; font-weight: bold; opacity: 0.8;">BY ALIF REZKY</span>
							</div>
							<div>
								<h3 style="margin: 0 0 10px 0; font-size: 1.4rem; font-weight: 900; color: #0F172A;">Buku Cetak Eksklusif: "Matematika Itu Asyik"</h3>
								<p style="font-size: 0.9rem; color: #475569; line-height: 1.6; text-align: justify; margin: 0 0 15px 0;">
									Buku praktis yang menyajikan trik hitung cepat aljabar, geometri, dan logika matematika tanpa rumus buta. Dilengkapi contoh visual interaktif.
								</p>
								
								<div style="background: #F1F5F9; padding: 15px; border-radius: 12px; display: flex; align-items: center; justify-content: space-between; margin-bottom: 15px;">
									<span style="font-size: 0.85rem; font-weight: 700; color: #334155;">📖 Preview Buku Gratis (30 Halaman Pertama)</span>
									<a href="/public/preview-buku.pdf" target="_blank" style="background: #0EA5E9; color: white; text-decoration: none; padding: 8px 16px; border-radius: 8px; font-size: 0.75rem; font-weight: bold;">Unduh PDF</a>
								</div>

								<div class="testi-box">
									"Sangat menyenangkan dibaca! Metode analoginya membuat materi persamaan kuadrat yang rumit jadi gampang dipahami oleh anak didikan kami."<br>
									<span style="font-weight: bold; font-size: 0.75rem; color: #118EEA;">— Testimoni Santri Al Binaa</span>
								</div>

								<a href="https://wa.me/6285256162879?text=Bismillah+Kak+Alif,+saya+mau+pesan+buku+cetak+Matematika+Itu+Asyik+ta'." target="_blank" class="cta-wa">
									🛒 AMANKAN SALINAN CETAK VIA WHATSAPP (085256162879)
								</a>
							</div>
						</div>
					</div>
				</div>

				<footer>ALIF REZKY • POWERED BY GOLANG ENGINE VERCEL</footer>

				<script>
					function jalankanJam() {
						const sekarang = new Date();
						document.getElementById('jam-txt').innerText = sekarang.toLocaleTimeString('id-ID', {hour12: false}).replace(/:/g, '.');
						document.getElementById('masehi-txt').innerText = sekarang.toLocaleDateString('id-ID', {weekday: 'long', day: 'numeric', month: 'long', year: 'numeric'});
					}
					setInterval(jalankanJam, 1000);
					jalankanJam();
				</script>
			</body>
			</html>
		`)
		return
	}

	// ==========================================
	// 3. ROUTING JASA: ALALIFY TECH SERVICES
	// ==========================================
	if (p == "/jasa/alalify-tech") {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Alalify Tech - Software Solutions</title>
				<style>
					body { font-family: sans-serif; background-color: #0F172A; color: white; display: flex; justify-content: center; align-items: center; min-height: 100vh; margin: 0; padding: 20px; }
					.container { background: #1E293B; padding: 40px; border-radius: 28px; max-width: 500px; width: 100%%; border: 1px solid #334155; }
					h1 { color: #38BDF8; font-size: 1.8rem; margin: 0 0 10px 0; font-weight: 900; }
					p { color: #94A3B8; font-size: 0.9rem; line-height: 1.6; }
					.list-jasa { margin: 20px 0; padding-left: 20px; color: #E2E8F0; font-size: 0.88rem; line-height: 1.8; }
					.btn-wa { display: block; background: #0EA5E9; color: white; text-decoration: none; text-align: center; padding: 15px; border-radius: 12px; font-weight: bold; font-size: 0.85rem; margin-top: 25px; }
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Alalify Tech 🛠️</h1>
					<p>Devisi rekayasa digital independen berperformasi tinggi oleh Kak Alif.</p>
					<ul class="list-jasa">
						<li>🚀 Fullstack Web App Development (Next.js & Go Stack)</li>
						<li>📊 Sistem Kasir POS & Point of Sale Custom Toko</li>
						<li>🏫 Sistem Informasi & LMS Kustom Pondok Pesantren</li>
					</ul>
					<a href="https://wa.me/6285256162879?text=Bismillah+Alalify+Tech,+saya+tertarik+konsultasi+sistem+aplikasi+ta'." target="_blank" class="btn-wa">Hubungi via WhatsApp</a>
				</div>
			</body>
			</html>
		`)
		return
	}

	// ==========================================
	// 4. ENGINE PEMBACA OTOMATIS ARTIKEL MARKDOWN (FOLDER content/posts)
	// ==========================================
	if (strings.HasPrefix(p, "/posts/")) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		
		// Mengambil nama file .md dari URL slug
		slug := strings.TrimPrefix(p, "/posts/")
		namaFile := slug + ".md"

		// Membaca file mentah dari direktori project Vercel
		pathArtikel := filepath.Join("content", "posts", namaFile)
		dataMentah, err := os.ReadFile(pathArtikel)
		
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "<h2>Afdhalna, Artikel '%s' belum ditemukan di folder content/posts ta', Sobat.</h2><a href='/'>Kembali ke Home</a>", slug)
			return
		}

		// Mengubah data teks markdown mentah menjadi tampilan HTML bersih sederhana
		isiTeks := string(dataMentah)
		isiHTML := strings.ReplaceAll(isiTeks, "\n", "<br>") // Parser darurat baris baru

		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>%s - Kak Alif Math Blog</title>
				<style>
					body { font-family: sans-serif; line-height: 1.7; color: #334155; max-width: 700px; margin: 40px auto; padding: 0 20px; background: #F8FAFC; }
					.box { background: white; padding: 40px; border-radius: 20px; border: 1px solid #E2E8F0; }
					a { color: #0EA5E9; text-decoration: none; font-weight: bold; }
				</style>
			</head>
			<body>
				<p><a href="/">← Kembali ke Beranda</a></p>
				<div class="box">
					%s
				</div>
			</body>
			</html>
		`, slug, isiHTML)
		return
	}

	http.NotFound(w, r)
}
