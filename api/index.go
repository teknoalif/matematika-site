package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	p := r.URL.Path

	// --- ROUTING HALAMAN UTAMA (DASHBOARD) ---
	if p == "/" || p == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, dashboardHTML())
		return
	}

	// --- ROUTING JASA ---
	if p == "/jasa/alalify-tech" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, jasaPageHTML())
		return
	}

	// --- ROUTING ARTIKEL MARKDOWN ---
	if strings.HasPrefix(p, "/posts/") {
		renderArtikel(w, p)
		return
	}

	http.NotFound(w, r)
}

func dashboardHTML() string {
	return `
	<!DOCTYPE html>
	<html lang="id">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Alif Rezky, M.Pd - Portofolio & Buku</title>
		<style>
			* { box-sizing: border-box; }
			body { font-family: 'Segoe UI', system-ui, sans-serif; background-color: #F8FAFC; color: #1E293B; margin: 0; padding-bottom: 60px; }
			.widget-bar { background: white; padding: 14px 40px; position: sticky; top: 0; z-index: 1000; box-shadow: 0 4px 20px rgba(0,0,0,0.05); display: flex; justify-content: space-between; align-items: center; width: 100%; border-bottom: 1px solid #E2E8F0; }
			.hero { background: linear-gradient(135deg, #0F172A 0%, #1E3A8A 100%); padding: 60px 40px; color: white; display: flex; justify-content: space-between; align-items: center; max-width: 1400px; margin: 20px auto; border-radius: 24px; }
			.hero h1 { font-weight: 950; font-size: 3.2rem; margin: 0; letter-spacing: -2px; color: #0EA5E9; }
			.hero-btn { display: inline-flex; align-items: center; color: white; text-decoration: none; font-size: 0.85rem; font-weight: 800; padding: 14px 28px; border-radius: 12px; border: 1px solid rgba(255,255,255,0.4); background: rgba(255,255,255,0.15); transition: 0.2s; }
			.container { max-width: 1400px; margin: 0 auto; padding: 0 20px; display: grid; grid-template-columns: 1fr 1fr; gap: 30px; }
			.card { background: white; padding: 35px; border-radius: 24px; border: 1px solid #E2E8F0; }
		</style>
	</head>
	<body>
		<div class="widget-bar">
			<div>📅 Masehi: <span id="masehi-txt"></span></div>
			<div>🌙 Hijriah: <span id="hijriah-txt"></span></div>
			<div>🕒 Jam: <span id="jam-txt"></span></div>
		</div>
		<div class="hero">
			<div><h1>Alif Rezky, M.Pd</h1><p>Mathematics Educator | Tech Developer | Author</p></div>
			<div>
				<a href="https://youtube.com/@kakalifgurumatematika" target="_blank" class="hero-btn">YouTube Channel</a>
				<a href="/jasa/alalify-tech" class="hero-btn" style="background:#0EA5E9;">Alalify Tech 🛠️</a>
			</div>
		</div>
		<div class="container">
			<div class="card"><h2>Profil & Keahlian</h2><p>Pendidik matematika UNM, expert GNU/Linux & Web Dev.</p></div>
			<div class="card"><h2>Pengalaman</h2><ul><li>Master Teacher OSN</li><li>Guru SMA IT Al Binaa</li></ul></div>
		</div>
		<script>
			function update() {
				const n = new Date();
				document.getElementById('jam-txt').innerText = n.toLocaleTimeString('id-ID', {hour12: false}).replace(/:/g, '.');
				document.getElementById('masehi-txt').innerText = n.toLocaleDateString('id-ID', {weekday: 'long', day: 'numeric', month: 'long', year: 'numeric'});
				// Hijriah Logic
				const d = Math.floor((n - new Date('2026-06-19'))/86400000);
				document.getElementById('hijriah-txt').innerText = (4+d) + " Muharram 1448 H";
			}
			setInterval(update, 1000); update();
		</script>
	</body>
	</html>`
}

func jasaPageHTML() string {
	return `<html><body style="background:#0F172A; color:white; padding:40px; font-family:sans-serif;">
		<h1>Alalify Tech 🛠️</h1>
		<p>Solusi Software, Web & LMS untuk Pendidikan.</p>
		<a href="/" style="color:#0EA5E9;">← Kembali ke Beranda</a>
	</body></html>`
}

func renderArtikel(w http.ResponseWriter, p string) {
	slug := strings.TrimPrefix(p, "/posts/")
	data, _ := os.ReadFile(filepath.Join("content", "posts", slug+".md"))
	fmt.Fprintf(w, "<html><body style='font-family:sans-serif; padding:40px; max-width:700px; margin:auto;'>%s</body></html>", string(data))
}
