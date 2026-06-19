package handler

import (
	"fmt"
	"net/http"
)

// Handler ini adalah entrypoint standar Vercel untuk Go
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	p := r.URL.Path
	switch p {
	case "/buku": 
		fmt.Fprintf(w, renderBuku())
	case "/jasa/alalify-tech": 
		fmt.Fprintf(w, renderJasa())
	default: 
		fmt.Fprintf(w, dashboardHTML())
	}
}

func dashboardHTML() string {
	return `<!DOCTYPE html><html lang="id" dir="ltr"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Alif Rezky, M.Pd - Portofolio</title>
	<style>
		body{font-family:system-ui, sans-serif; background:#F1F5F9; margin:0; padding-bottom:50px;}
		.hero{background:#0F172A; padding:60px 20px; color:white; text-align:center;}
		.main-wrapper{max-width:900px; margin:-40px auto 0; padding:0 20px;}
		.card{background:white; padding:30px; border-radius:16px; margin-bottom:20px; box-shadow:0 4px 6px rgba(0,0,0,0.03);}
		h2{color:#800000; font-weight:900; border-left:6px solid #800000; padding-left:15px;}
		.btn-yt{background:#0EA5E9; color:white; padding:12px 24px; border-radius:8px; text-decoration:none; font-weight:800; display:inline-block; margin:5px;}
		.btn-nav{background:transparent; color:white; padding:8px 16px; border:1px solid #334155; border-radius:6px; text-decoration:none; font-size:0.9rem; margin:0 5px;}
		.video-wrapper{position:relative; width:100%; padding-bottom:56.25%; height:0; border-radius:12px; overflow:hidden; background:#000;}
		.video-wrapper iframe{position:absolute; top:0; left:0; width:100%; height:100%;}
	</style></head><body>
		<div class="hero">
			<h1>Alif Rezky, M.Pd</h1>
			<p>Mathematics Educator | Tech Developer | Author</p>
			<a href="https://youtube.com/@kakalifgurumatematika" class="btn-yt">YouTube Channel 📺</a>
			<div style="margin-top:20px;">
				<a href="/buku" class="btn-nav">Katalog Buku 📚</a>
				<a href="/jasa/alalify-tech" class="btn-nav">Jasa Al Alify Tech 🛠️</a>
			</div>
		</div>
		<div class="main-wrapper">
			<div class="card">
				<h2>Profil (ID / AR)</h2>
				<p><b>ID:</b> Muslim & Fullstack Developer, berfokus pada edukasi Islam, Matematika, dan Teknologi.</p>
				<p dir="rtl" style="text-align:right;"><b>AR:</b> أنا مسلم ومطور برمجيات متكامل، وأسعى بإذن الله إلى تعليم المجتمع عن الإسلام والرياضيات والتكنولوجيا.</p>
			</div>
			<div class="card">
				<h2>Riwayat Pengalaman / الخبرات</h2>
				<ul>
					<li><b>Math Tutor (Online)</b>, Algonova (Mar 2026-Sekarang)</li>
					<li><b>Guru Matematika</b>, SMA IT Al Binaa (Sept 2022-Jun 2026)</li>
					<li><b>Master Teacher</b>, Brain Academy by Ruangguru (Okt-Des 2023)</li>
					<li><b>Asisten Dosen</b>, UNM (Okt 2018-Sept 2022)</li>
				</ul>
			</div>
			<div class="card">
				<h2>Video Pembelajaran</h2>
				<div class="video-wrapper">
					<iframe src="https://www.youtube.com/embed/_3pqgVhtDBg" allowfullscreen></iframe>
				</div>
			</div>
		</div>
	</body></html>`
}

func renderBuku() string {
	return `<!DOCTYPE html><html lang="id"><body style="font-family:system-ui, sans-serif; max-width:700px; margin:auto; padding:40px 20px;">
		<a href="/" style="color:#800000; font-weight:bold; text-decoration:none;">← Kembali ke Beranda</a>
		<h1 style="color:#800000; margin-top:20px;">Matematika Itu Asyik</h1>
		<div style="border:2px solid #E2E8F0; border-radius:12px; overflow:hidden; margin-bottom:20px;">
			<iframe src="https://drive.google.com/file/d/17SbICWWxQOCRf_l4xOVrjAU20CBNkY0X/preview" width="100%" height="500px" style="border:none;"></iframe>
		</div>
		<a href="https://wa.me/6285256162879?text=Bismillah%2C%20saya%20ingin%20memesan%20buku%20Matematika%20Itu%20Asyik" style="display:block; background:#25D366; padding:20px; color:white; text-align:center; border-radius:12px; text-decoration:none; font-weight:900;">PESAN VIA WHATSAPP 📱</a>
	</body></html>`
}

func renderJasa() string {
	return `<html><body style="background:#0F172A; color:white; padding:40px; font-family:sans-serif;">
		<h1>Al Alify Tech 🛠️</h1>
		<p>Solusi Software, Web & LMS untuk Pendidikan.</p>
		<a href="/" style="color:#0EA5E9;">← Kembali ke Beranda</a>
	</body></html>`
}
