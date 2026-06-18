package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Handler utama yang dieksekusi oleh Vercel Serverless Functions
func Handler(w http.ResponseWriter, r *http.Request) {
	// Aturan CORS agar web tetap fleksibel
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	p := r.URL.Path

	// 1. ROUTING HALAMAN UTAMA (DASHBOARD)
	if p == "/" || p == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, dashboardHTML())
		return
	}

	// 2. ROUTING KATALOG BUKU
	if p == "/buku" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, bukuPageHTML())
		return
	}

	// 3. ROUTING JASA ALALIFY TECH
	if p == "/jasa/alalify-tech" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, jasaPageHTML())
		return
	}

	// 4. ENGINE ARTIKEL MARKDOWN
	if strings.HasPrefix(p, "/posts/") {
		renderArtikel(w, p)
		return
	}

	http.NotFound(w, r)
}

// ==========================================
// KUMPULAN FUNGSI HTML (LOGIKA TAMPILAN)
// ==========================================

func dashboardHTML() string {
	return `
	<!DOCTYPE html>
	<html lang="id">
	<head>
		<meta charset="UTF-8">
		<title>Alif Rezky, M.Pd - Portofolio & Buku</title>
		<style>
			* { box-sizing: border-box; }
			body { font-family: 'Segoe UI', sans-serif; background-color: #F8FAFC; color: #1E293B; margin: 0; padding-bottom: 60px; }
			.widget-bar { background: white; padding: 14px 40px; position: sticky; top: 0; z-index: 1000; box-shadow: 0 4px 20px rgba(0,0,0,0.05); display: flex; justify-content: space-between; align-items: center; width: 100%; border-bottom: 1px solid #E2E8F0; }
			.hero { background: linear-gradient(135deg, #0F172A 0%, #1E3A8A 100%); padding: 60px 40px; color: white; display: flex; justify-content: space-between; align-items: center; max-width: 1400px; margin: 20px auto; border-radius: 24px; }
			.hero h1 { font-weight: 950; font-size: 3.2rem; color: #800000; background: none; filter: drop-shadow(0 2px 4px rgba(0,0,0,0.3)); }
			.hero-btn { display: inline-flex; align-items: center; color: white; text-decoration: none; font-size: 0.85rem; font-weight: 800; padding: 14px 28px; border-radius: 12px; border: 1px solid rgba(255,255,255,0.4); background: #800000; transition: 0.2s; }
			.container { max-width: 1400px; margin: 0 auto; padding: 0 20px; display: grid; grid-template-columns: 1fr 1fr; gap: 30px; }
			.card { background: white; padding: 35px; border-radius: 24px; border: 1px solid #E2E8F0; }
		</style>
	</head>
	<body>
		<div class="widget-bar">
			<div id="masehi-txt">Loading...</div>
			<div id="hijriah-txt">Loading...</div>
			<div id="jam-txt">Loading...</div>
		</div>
		<div class="hero">
			<div><h1>Alif Rezky, M.Pd</h1><p>Mathematics Educator | Tech Developer | Author</p></div>
			<div style="display:flex; gap:10px;">
				<a href="/buku" class="hero-btn">Katalog Buku 📚</a>
				<a href="/jasa/alalify-tech" class="hero-btn" style="background:#0EA5E9;">Alalify Tech 🛠️</a>
			</div>
		</div>
		<div class="container">
			<div class="card"><h2>Profil & Keahlian</h2><p>Pendidik matematika profesional, lulusan UNM, aktif mengembangkan solusi teknologi inklusif.</p></div>
			<div class="card"><h2>Pengalaman</h2><p>Master Teacher OSN, Algonova, SMA IT Al Binaa.</p></div>
		</div>
		<script>
			function update() {
				const now = new Date();
				document.getElementById('jam-txt').innerText = now.toLocaleTimeString('id-ID', {hour12: false}).replace(/:/g, '.');
				document.getElementById('masehi-txt').innerText = now.toLocaleDateString('id-ID', {weekday: 'long', day: 'numeric', month: 'long', year: 'numeric'});
				const tglPatokan = new Date('2026-06-19').setHours(0,0,0,0);
				const selisih = Math.floor((now.setHours(0,0,0,0) - tglPatokan) / 86400000);
				document.getElementById('hijriah-txt').innerText = (4 + selisih) + " Muharram 1448 H";
			}
			setInterval(update, 1000); update();
		</script>
	</body>
	</html>`
}

func bukuPageHTML() string {
	return `<!DOCTYPE html><html><body><div style="max-width:800px; margin:auto; padding:40px;">
	<a href="/">← Kembali</a><h1 style="color:#800000;">Matematika Itu Asyik</h1>
	<iframe src="https://drive.google.com/file/d/17SbICWWxQOCRf_l4xOVrjAU20CBNkY0X/preview" width="100%" height="400px"></iframe>
	<a href="https://wa.me/6285256162879?text=Bismillah+pesan+buku" style="display:block; padding:20px; background:#25D366; color:white; text-align:center; border-radius:12px; text-decoration:none; font-weight:900;">PESAN VIA WHATSAPP</a>
	</div></body></html>`
}

func jasaPageHTML() string {
	return `<html><body style="background:#0F172A; color:white; padding:40px;"><h1>Alalify Tech</h1><p>Jasa Web & LMS.</p><a href="/">Kembali</a></body></html>`
}

func renderArtikel(w http.ResponseWriter, p string) {
	slug := strings.TrimPrefix(p, "/posts/")
	data, _ := os.ReadFile(filepath.Join("content", "posts", slug+".md"))
	fmt.Fprintf(w, "<html><body><div style='max-width:700px; margin:auto;'>%s</div></body></html>", string(data))
}
