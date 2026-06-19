package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p := r.URL.Path

	switch {
	case p == "/" || p == "": fmt.Fprintf(w, renderDashboard())
	case p == "/buku": fmt.Fprintf(w, renderBuku())
	case p == "/jasa/alalify-tech": fmt.Fprintf(w, renderJasa())
	default: http.NotFound(w, r)
	}
}

func renderDashboard() string {
	return `<!DOCTYPE html><html lang="id"><head><meta charset="UTF-8"><title>Alif Rezky, M.Pd</title>
	<style>
		:root{--maroon:#800000; --blue:#0EA5E9; --bg:#F8FAFC;}
		body{font-family:'Inter', system-ui, sans-serif; background:var(--bg); color:#334155; margin:0;}
		.hero{padding:80px 20px; text-align:center; background:white; border-bottom:1px solid #E2E8F0;}
		h1{color:var(--maroon); font-size:3rem; font-weight:950; margin:0;}
		.grid{display:grid; grid-template-columns:repeat(auto-fit, minmax(350px, 1fr)); gap:30px; padding:40px; max-width:1100px; margin:auto;}
		.card{background:white; padding:30px; border-radius:24px; box-shadow:0 10px 15px -3px rgba(0,0,0,0.05); border:1px solid #E2E8F0;}
		.tag{display:inline-block; background:#EFF6FF; color:var(--blue); padding:4px 12px; border-radius:6px; font-weight:700; font-size:0.8rem; margin:4px;}
		.btn{padding:12px 24px; border-radius:12px; text-decoration:none; font-weight:800; display:inline-block; margin-top:20px;}
	</style></head><body>
		<div class="hero">
			<h1>Alif Rezky, M.Pd</h1>
			<p style="font-size:1.2rem; font-weight:600;">Mathematics Educator & Fullstack Developer</p>
			<a href="/buku" class="btn" style="background:var(--maroon); color:white;">Lihat Buku Cetak 📚</a>
		</div>
		<div class="grid">
			<div class="card">
				<h2>Tentang Saya</h2>
				<p>Magister Pendidikan Matematika UNM dengan passion tinggi dalam rekayasa perangkat lunak. Berpengalaman dalam membimbing santri OSN/UTBK dan membangun ekosistem digital (LMS) berbasis web modern.</p>
				<div><span class="tag">GNU/Linux</span><span class="tag">Golang</span><span class="tag">React/Next.js</span><span class="tag">Python</span></div>
			</div>
			<div class="card">
				<h2>Pengalaman Utama</h2>
				<p><strong>Master Teacher OSN</strong> @ Edumatrix (2026-Sekarang)<br>
				<strong>Online Math Tutor</strong> @ Algonova (2026-Sekarang)<br>
				<strong>Guru Matematika</strong> @ SMA IT Al Binaa (2022-2026)</p>
			</div>
			<div class="card" style="grid-column: 1 / -1;">
				<h2>Layanan Al Alify Tech</h2>
				<p>Solusi digital kustom untuk lembaga pendidikan, toko (POS), dan pengembangan aplikasi profesional.</p>
				<a href="/jasa/alalify-tech" class="btn" style="background:var(--blue); color:white;">Jelajahi Solusi Teknologi 🛠️</a>
			</div>
		</div>
	</body></html>`
}

func renderBuku() string {
	return `<!DOCTYPE html><html lang="id"><body style="font-family:sans-serif; max-width:700px; margin:auto; padding:40px;">
		<a href="/">← Kembali</a><h1 style="color:#800000;">Matematika Itu Asyik</h1>
		<p>Buku inovatif yang mengubah cara pandang santri terhadap aljabar dan geometri.</p>
		<iframe src="https://drive.google.com/file/d/17SbICWWxQOCRf_l4xOVrjAU20CBNkY0X/preview" width="100%" height="400px"></iframe>
		<a href="https://wa.me/6285256162879?text=Bismillah+pesan+buku" style="display:block; background:#25D366; padding:20px; color:white; text-align:center; border-radius:12px; font-weight:900; text-decoration:none;">PESAN VIA WHATSAPP</a>
	</body></html>`
}

func renderJasa() string {
	return `<html><body style="font-family:sans-serif; padding:40px; background:#0F172A; color:white;">
		<h1>Al Alify Tech 🛠️</h1>
		<p>Kami melayani pembuatan sistem informasi sekolah (LMS), aplikasi kasir (POS), hingga website perusahaan.</p>
		<a href="/" style="color:#38BDF8;">← Kembali ke Portofolio</a>
	</body></html>`
}
