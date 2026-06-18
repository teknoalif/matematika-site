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

	// 1. DASHBOARD & PORTFOLIO
	if p == "/" || p == "" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, dashboardHTML())
		return
	}

	// 2. KATALOG BUKU (Rute Baru)
	if p == "/buku" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, bukuPageHTML())
		return
	}

	// 3. JASA ALALIFY TECH
	if p == "/jasa/alalify-tech" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, jasaPageHTML())
		return
	}

	// 4. ARTIKEL PEMBUKTIAN RUMUS ABC (Spesifik)
	if p == "/posts/pembuktian-rumus-abc" {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, renderArtikelABC())
		return
	}

	// 5. ENGINE POSTS LAINNYA
	if strings.HasPrefix(p, "/posts/") {
		renderArtikel(w, p)
		return
	}

	http.NotFound(w, r)
}

// --- FUNGSI RENDER (DASHBOARD) ---
func dashboardHTML() string {
	return `<!DOCTYPE html>
	<html lang="id"><head><meta charset="UTF-8"><title>Alif Rezky, M.Pd</title>
	<style>
		body{font-family:'Segoe UI', sans-serif; background:#F8FAFC; color:#1E293B; margin:0; padding-bottom:60px;}
		.hero{background:linear-gradient(135deg, #0F172A 0%, #1E3A8A 100%); padding:60px; color:white; border-radius:24px; margin:20px; text-align:center;}
		h1{color:#800000; font-size:3.2rem; font-weight:950; margin:0;}
		.container{max-width:1200px; margin:auto; display:grid; grid-template-columns:repeat(auto-fit, minmax(300px, 1fr)); gap:20px; padding:20px;}
		.card{background:white; padding:30px; border-radius:24px; border:1px solid #E2E8F0;}
	</style></head><body>
		<div class="hero"><h1>Alif Rezky, M.Pd</h1><p>Mathematics Educator | Tech Developer</p>
		<div style="margin-top:20px;"><a href="/buku" style="color:white; padding:15px 30px; background:#800000; text-decoration:none; border-radius:12px; font-weight:900;">Katalog Buku 📚</a></div></div>
		<div class="container">
			<div class="card"><h2>Profil</h2><p>Pendidik matematika UNM, expert GNU/Linux & Web Dev.</p></div>
			<div class="card"><h2>Pengalaman</h2><ul><li>Master Teacher OSN</li><li>Guru SMA IT Al Binaa</li></ul></div>
		</div>
	</body></html>`
}

// --- FUNGSI RENDER (BUKU) ---
func bukuPageHTML() string {
	return `<!DOCTYPE html><html lang="id"><body style="font-family:sans-serif; max-width:600px; margin:auto; padding:40px;">
		<a href="/">← Kembali</a><h1 style="color:#800000;">Matematika Itu Asyik</h1>
		<iframe src="https://drive.google.com/file/d/17SbICWWxQOCRf_l4xOVrjAU20CBNkY0X/preview" width="100%" height="400px" style="border:none;"></iframe>
		<a href="https://wa.me/6285256162879?text=Bismillah+pesan+buku" style="display:block; background:#25D366; padding:20px; color:white; text-align:center; border-radius:12px; text-decoration:none; font-weight:900; margin-top:20px;">PESAN VIA WHATSAPP</a>
	</body></html>`
}

// --- FUNGSI RENDER (JASA) ---
func jasaPageHTML() string {
	return `<html><body style="background:#0F172A; color:white; padding:40px;"><h1>Alalify Tech</h1><p>Jasa Web & LMS.</p><a href="/" style="color:white;">Kembali</a></body></html>`
}

// --- FUNGSI RENDER (ARTIKEL ABC DENGAN KATEX) ---
func renderArtikelABC() string {
	return `<!DOCTYPE html><html><head>
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/katex@0.16.10/dist/katex.min.css">
	<script defer src="https://cdn.jsdelivr.net/npm/katex@0.16.10/dist/katex.min.js"></script>
	<script defer src="https://cdn.jsdelivr.net/npm/katex@0.16.10/dist/contrib/auto-render.min.js" onload="renderMathInElement(document.body)"></script>
	</head><body style="font-family:sans-serif; line-height:1.7; padding:40px; max-width:700px; margin:auto;">
		<a href="/">← Beranda</a>
		<h1>Pembuktian Rumus ABC</h1>
		<p>Rumus Kuadrat: $$x_{1,2} = \frac{-b \pm \sqrt{b^2 - 4ac}}{2a}$$</p>
		<p>Pembuktian dimulai dari bentuk umum $ax^2 + bx + c = 0$ dengan melengkapkan kuadrat sempurna...</p>
	</body></html>`
}

// --- FUNGSI RENDER (MARKDOWN GENERATOR) ---
func renderArtikel(w http.ResponseWriter, p string) {
	slug := strings.TrimPrefix(p, "/posts/")
	data, _ := os.ReadFile(filepath.Join("content", "posts", slug+".md"))
	fmt.Fprintf(w, "<html><body><div style='max-width:700px; margin:auto; padding:40px;'>%s</div></body></html>", string(data))
}
