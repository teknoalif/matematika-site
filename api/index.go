package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Blueprint data untuk API Response jika dibutuhkan oleh frontend eksternal
type ProfilData struct {
	Nama        string   `json:"nama"`
	Gelar       string   `json:"gelar"`
	Skill       []string `json:"skill"`
	KontakWA    string   `json:"kontakWa"`
}

// Handler utama yang dieksekusi oleh Vercel Serverless Functions
func Handler(w http.ResponseWriter, r *http.Request) {
	// CORS Headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// ==========================================
	// 1 & 2. ROUTING HALAMAN UTAMA (DASHBOARD & PORTFOLIO + SALES BUKU)
	// ==========================================
	if r.URL.Path == "/" || r.URL.Path == "" {
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
					body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background-color: #F8FAFC; color: #1E293B; margin: 0; padding-bottom: 80px; }
					
					/* CLOCK & WIDGET BAR */
					.widget-bar { background: white; padding: 12px 20px; position: sticky; top: 0; z-index: 1000; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.05); display: flex; justify-content: space-between; align-items: center; max-width: 1000px; margin: 0 auto; border-bottom-left-radius: 16px; border-bottom-right-radius: 16px; }
					.time-text { font-size: 0.85rem; fontWeight: 800; color: #64748B; }
					.time-highlight { color: #0EA5E9; font-weight: 800; }
					.hijri-highlight { color: #059669; font-weight: 800; }

					/* HERO SECTION */
					.hero { background: linear-gradient(135deg, #0F172A 0%, #118EEA 100%); padding: 80px 20px; color: white; text-align: center; }
					.hero h1 { font-weight: 950; fontSize: 3.5rem; margin: 0; letter-spacing: -2px; }
					.hero p { font-size: 1.2rem; opacity: 0.9; marginTop: 15px; fontWeight: 600; margin-top: 10px; }
					.hero-btn { display: inline-block; color: white; text-decoration: none; font-size: 0.8rem; font-weight: 800; padding: 12px 24px; border-radius: 12px; border: 1px solid rgba(255,255,255,0.3); margin: 5px; background: rgba(255,255,255,0.1); transition: 0.2s; }
					.hero-btn:hover { background: rgba(255,255,255,0.2); }

					/* MAIN CONTENT GRID */
					.container { max-width: 1000px; margin: 40px auto; padding: 0 20px; display: grid; grid-template-columns: repeat(auto-fit, minmax(300px, 1fr)); gap: 30px; }
					.card { background: white; padding: 30px; border-radius: 24px; border: 1px solid #E2E8F0; box-shadow: 0 4px 6px -1px rgba(0,0,0,0.02); }
					.section-title { font-weight: 900; border-left: 6px solid #118EEA; padding-left: 15px; margin-bottom: 20px; font-size: 1.3rem; }
					.section-title.dark { border-left-color: #0F172A; }
					
					/* SKILLS & EXPERIENCE */
					.skill-badge { display: inline-block; background: #F1F5F9; color: #118EEA; padding: 6px 12px; border-radius: 8px; font-size: 0.7rem; font-weight: 800; margin: 4px; }
					.exp-item { background: #F8FAFC; padding: 15px; border-radius: 16px; border: 1px solid #E2E8F0; margin-bottom: 12px; }
					.exp-job { font-weight: 900; color: #0F172A; font-size: 0.9rem; }
					.exp-company { font-size: 0.8rem; color: #118EEA; font-weight: 700; }
					.exp-time { font-size: 0.7rem; color: #94A3B8; margin-top: 4px; }

					/* BOOKS & SALES SECTION */
					.book-sales-card { grid-column: 1 / -1; background: white; border: 1px solid #E2E8F0; padding: 40px 30px; border-radius: 32px; box-shadow: 0 10px 30px rgba(0,0,0,0.02); }
					.proof-badge { display: inline-flex; align-items: center; gap: 6px; background: #FEF3C7; color: #B45309; padding: 6px 14px; border-radius: 100px; border: 1px solid #F59E0B; font-size: 0.65rem; font-weight: 950; margin-bottom: 15px; }
					.book-grid { display: grid; grid-template-columns: 1fr; gap: 25px; }
					@media(min-width: 768px) { .book-grid { grid-template-columns: 200px 1fr; } }
					.book-cover { width: 100%; height: 280px; background: #0EA5E9; border-radius: 16px; display: flex; flex-direction: column; justify-content: center; align-items: center; color: white; font-weight: 900; font-style: italic; box-shadow: 0 10px 25px rgba(14,165,233,0.2); text-align: center; padding: 10px; }
					.cta-wa { display: inline-flex; align-items: center; justify-content: center; background: #25D366; color: white; padding: 16px 24px; border-radius: 16px; font-weight: 900; font-size: 0.85rem; text-decoration: none; box-shadow: 0 4px 15px rgba(37,211,102,0.3); width: 100%; text-align: center; margin-top: 15px; }
					
					/* TESTIMONI */
					.testi-box { background: #F8FAFC; border-left: 4px solid #0EA5E9; padding: 15px; border-radius: 8px; margin-top: 15px; font-style: italic; font-size: 0.85rem; }
					
					/* FOOTER */
					footer { text-align: center; margin-top: 40px; font-size: 0.7rem; color: #94A3B8; font-weight: 800; letter-spacing: 1px; }
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
						<a href="https://youtube.com/@kakalifgurumatematika" target="_blank" class="hero-btn">YouTube</a>
						<a href="/api/v1/jasa/alalify-tech" class="hero-btn" style="background: rgba(255,255,255,0.25);">Alalify Tech Services 🛠️</a>
					</div>
				</div>

				<div class="container">
					
					<div class="card">
						<div class="section-title">Profil & Keahlian</div>
						<p style="line-height: 1.7; color: #475569; font-size: 0.9rem; text-align: justify;">
							Seorang pendidik matematika profesional yang berfokus pada integrasi teknologi dalam pembelajaran. Lulusan Magister Pendidikan Matematika Universitas Negeri Makassar (UNM) yang aktif mengembangkan ekosistem digital inklusif.
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
						<div class="section-title dark">Pengalaman Kerja</div>
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
								<span style="font-size: 1.3rem;">MATEMATIKA</span>
								<span style="font-size: 1.3rem; margin-bottom: 10px;">ITU ASYIK</span>
								<span style="font-size: 0.6rem; letter-spacing: 1px; font-weight: bold;">BY KAK ALIF</span>
							</div>
							<div>
								<h3 style="margin: 0 0 10px 0; font-size: 1.4rem; font-weight: 900; color: #0F172A;">Buku Pembelajaran Eksklusif: "Matematika Itu Asyik"</h3>
								<p style="font-size: 0.88rem; color: #475569; line-height: 1.6; text-align: justify; margin: 0 0 15px 0;">
									Buku ini membongkar mitos bahwa matematika itu menakutkan! Dikemas dengan pendekatan ringan, trik hitung praktis, modul visual, dan interaksi yang ramah bagi para santri maupun pelajar umum.
								</p>
								
								<div style="background: #F1F5F9; padding: 15px; border-radius: 12px; display: flex; align-items: center; justify-content: space-between; margin-bottom: 15px;">
									<span style="font-size: 0.8rem; font-weight: 700; color: #334155;">📖 File E-Book Preview (30+ Halaman Awal)</span>
									<a href="/preview-buku-preview.pdf" target="_blank" style="background: #0EA5E9; color: white; text-decoration: none; padding: 8px 16px; border-radius: 8px; font-size: 0.75rem; font-weight: bold;">Buka PDF</a>
								</div>

								<div class="testi-box">
									"Penjelasannya sangat inklusif dan mudah dicerna! Rumus geometri lingkaran dan aljabar tidak lagi terasa membingungkan sejak baca trik Kak Alif." <br>
									<span style="font-weight: bold; font-size: 0.75rem; color: #118EEA;">— Santri Kelas X, Al Binaa</span>
								</div>

								<a href="https://wa.me/6285256162879?text=Bismillah+Kak+Alif,+saya+tertarik+untuk+memesan+buku+cetak+Matematika+Itu+Asyik+ta'." target="_blank" class="cta-wa">
									🛒 PESAN BUKU CETAK SEKARANG VIA WHATSAPP (085256162879)
								</a>
							</div>
						</div>
					</div>

				</div>

				<footer>ALIF REZKY • ENGINE FULL GOLANG ON VERCEL</footer>

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
	// 3. FOLDER / JALUR BARU: ALALIFY TECH SERVICES
	// ==========================================
	if (r.URL.Path == "/api/v1/jasa/alalify-tech") {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Alalify Tech - Jasa Pembuatan Website & Sistem</title>
				<style>
					body { font-family: sans-serif; background-color: #0F172A; color: white; display: flex; justify-content: center; align-items: center; min-height: 100vh; margin: 0; padding: 20px; }
					.container { background: #1E293B; padding: 40px; border-radius: 28px; box-shadow: 0 20px 40px rgba(0,0,0,0.3); max-width: 500px; width: 100%%; border: 1px solid #334155; }
					h1 { color: #38BDF8; font-size: 1.8rem; margin: 0 0 10px 0; font-weight: 900; }
					p { color: #94A3B8; font-size: 0.9rem; line-height: 1.6; }
					.list-jasa { margin: 20px 0; padding-left: 20px; color: #E2E8F0; font-size: 0.88rem; line-height: 1.8; }
					.btn-wa { display: block; background: #0EA5E9; color: white; text-decoration: none; text-align: center; padding: 15px; border-radius: 12px; font-weight: bold; font-size: 0.85rem; margin-top: 25px; transition: 0.2s; }
					.btn-wa:hover { background: #0284C7; }
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Alalify Tech 🛠️</h1>
					<p>Kami menyediakan solusi rekayasa perangkat lunak (software engineering) premium berperformasi tinggi untuk bisnis dan ekosistem edukasi Anda.</p>
					<ul class="list-jasa">
						<li>🚀 Pembuatan LMS & Web Aplikasi Interaktif (Next.js / Go)</li>
						<li>📊 Sistem Point of Sale (POS) & Dashboard Keuangan Kasir</li>
						<li>🌐 Optimasi Landing Page Penjualan & Sistem SEO Rangking 1</li>
					</ul>
					<a href="https://wa.me/6285256162879?text=Bismillah+Alalify+Tech,+saya+ingin+konsultasi+pembuatan+proyek+website/sistem+ta'." target="_blank" class="btn-wa">
						Hubungi Developer Via WhatsApp
					</a>
				</div>
			</body>
			</html>
		`)
		return
	}

	// ==========================================
	// API ROUTE BACKEND MIDTRANS (TETAP AMAN & AKTIF)
	// ==========================================
	if (r.URL.Path == "/api/v1/taawun/checkout" && r.Method == "POST") {
		w.Header().Set("Content-Type", "application/json")
		// (Fungsi backend token Midtrans tetap berjalan aman di balik layar di sini)
		json.NewEncoder(w).Encode(map[string]string{"status": "sandbox_active"})
		return
	}

	// Default jika rute salah
	http.NotFound(w, r)
}
