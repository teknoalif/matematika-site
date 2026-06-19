package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	p := r.URL.Path
	switch {
	case p == "/" || p == "": fmt.Fprintf(w, dashboardHTML())
	case p == "/buku": fmt.Fprintf(w, renderBuku())
	case p == "/jasa/alalify-tech": fmt.Fprintf(w, renderJasa())
	default: http.NotFound(w, r)
	}
}

func dashboardHTML() string {
	return `<!DOCTYPE html><html lang="id" dir="ltr"><head><meta charset="UTF-8"><title>Alif Rezky, M.Pd</title>
	<style>
		body{font-family:system-ui, sans-serif; background:#F1F5F9; margin:0; padding-bottom:50px;}
		.hero{background:#0F172A; padding:60px 20px; color:white; text-align:center;}
		.main-wrapper{max-width:900px; margin:-40px auto 0; padding:0 20px;}
		.card{background:white; padding:30px; border-radius:16px; margin-bottom:20px; box-shadow:0 4px 6px rgba(0,0,0,0.03);}
		h2{color:#800000; font-weight:900; border-left:6px solid #800000; padding-left:15px;}
		.btn-yt{background:#0EA5E9; color:white; padding:12px 24px; border-radius:8px; text-decoration:none; font-weight:800; display:inline-block; margin:5px;}
	</style></head><body>
		<div class="hero">
			<h1>Alif Rezky, M.Pd</h1>
			<p>Mathematics Educator | Tech Developer | Author</p>
			<a href="https://youtube.com/@kakalifgurumatematika" class="btn-yt">YouTube Channel 📺</a>
		</div>
		<div class="main-wrapper">
			<div class="card">
				<h2>Profil (ID / AR)</h2>
				<p><b>ID:</b> Muslim & Fullstack Developer yang bertujuan memberikan edukasi tentang Islam, Matematika, dan Teknologi.</p>
				<p dir="rtl" style="text-align:right;"><b>AR:</b> أنا مسلم ومطور برمجيات متكامل، وأسعى بإذن الله إلى تعليم المجتمع عن الإسلام والرياضيات والتكنولوجيا.</p>
			</div>
			<div class="card">
				<h2>Riwayat Pengalaman / الخبرات</h2>
				<ul>
					<li><b>Math Tutor (Online)</b>, Algonova (Mar 2026-Sekarang / مارس ٢٠٢٦م - الآن)</li>
					<li><b>Guru Matematika</b>, SMA IT Al Binaa (Sept 2022-Jun 2026 / سبتمبر ٢٠٢٢م - يونيو ٢٠٢٦م)</li>
					<li><b>Master Teacher</b>, Brain Academy by Ruangguru (Okt-Des 2023 / أكتوبر - ديسمبر ٢٠٢٣م)</li>
					<li><b>Asisten Dosen</b>, UNM (Okt 2018-Sept 2022 / أكتوبر ٢٠١٨م - سبتمبر ٢٠٢٢م)</li>
				</ul>
			</div>
			<div class="card">
				<h2>Bahasa & Skill</h2>
				<p><b>Bahasa:</b> Indonesia, Inggris, Makassar, Bisindo | العربية (مستوى أساسي)</p>
				<p><b>Teknologi:</b> Python, JS, CSS, PHP, ActionScript 2.0</p>
			</div>
			<div class="card">
				<h2>Video Pembelajaran</h2>
				<div style="position:relative; padding-bottom:56.25%; height:0; overflow:hidden;">
					<iframe style="position:absolute; top:0; left:0; width:100%; height:100%; border-radius:12px;" src="https://www.youtube.com/embed/_3pqgVhtDBg"></iframe>
				</div>
			</div>
		</div>
	</body></html>`
}

func renderBuku() string {
	return `<html><body style="font-family:sans-serif; max-width:600px; margin:auto; padding:40px;">
		<a href="/">← Kembali</a><h1 style="color:#800000;">Matematika Itu Asyik</h1>
		<iframe src="https://drive.google.com/file/d/17SbICWWxQOCRf_l4xOVrjAU20CBNkY0X/preview" width="100%" height="400px"></iframe>
	</body></html>`
}

func renderJasa() string {
	return `<html><body style="background:#0F172A; color:white; padding:40px; font-family:sans-serif;">
		<h1>Al Alify Tech 🛠️</h1>
		<p>Pengembangan Web, LMS, dan Aplikasi POS Kustom.</p>
		<a href="/" style="color:#0EA5E9;">← Kembali ke Beranda</a>
	</body></html>`
}
