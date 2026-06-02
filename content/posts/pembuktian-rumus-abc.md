---
title: "Pembuktian Rumus ABC (Rumus Kuadrat) dari Bentuk Umum"
date: 2026-06-02T08:00:00+07:00
draft: false
math: true
summary: "Bagaimana asal-usul Rumus ABC yang sering kita gunakan? Mari kita buktikan bersama menggunakan metode melengkapkan kuadrat sempurna."
categories: ["materi"]
tags: ["matematika", "aljabar", "persamaan-kuadrat"]
---

Rumus Kuadrat atau yang lebih akrab kita sebut sebagai **Rumus ABC** digunakan untuk mencari akar-akar dari persamaan kuadrat $ax^2 + bx + c = 0$. Rumusnya adalah:

$$x_{1,2} = \frac{-b \pm \sqrt{b^2 - 4ac}}{2a}$$

Namun, dari manakah rumus tersebut berasal? Mari kita buktikan rumus ini langkah demi langkah dengan metode **Melengkapkan Kuadrat Sempurna**.

---

### Langkah-Langkah Pembuktian

**1. Tulis Bentuk Umum Persamaan Kuadrat**
$$ax^2 + bx + c = 0$$

**2. Pindahkan Konstanta $c$ ke Ruas Kanan**
$$ax^2 + bx = -c$$

**3. Bagi Kedua Ruas dengan Koefisien $a$**
Agar koefisien dari $x^2$ menjadi 1, kita bagi seluruh suku dengan $a$:
$$x^2 + \frac{b}{a}x = -\frac{c}{a}$$

**4. Lengkapkan Kuadrat Sempurna**
Tambahkan kedua ruas dengan kuadrat dari setengah koefisien $x$, yaitu $\left(\frac{b}{2a}\right)^2$:
$$x^2 + \frac{b}{a}x + \left(\frac{b}{2a}\right)^2 = -\frac{c}{a} + \left(\frac{b}{2a}\right)^2$$

**5. Ubah Ruas Kiri Menjadi Bentuk Kuadrat Sempurna**
Ubah ruas kiri menjadi bentuk $(x + p)^2$ dan jabarkan kuadrat di ruas kanan:
$$\left(x + \frac{b}{2a}\right)^2 = -\frac{c}{a} + \frac{b^2}{4a^2}$$

**6. Samakan Penyebut di Ruas Kanan**
Samakan penyebut ruas kanan menjadi $4a^2$:
$$\left(x + \frac{b}{2a}\right)^2 = \frac{-4ac + b^2}{4a^2}$$
$$\left(x + \frac{b}{2a}\right)^2 = \frac{b^2 - 4ac}{4a^2}$$

**7. Akarkan Kedua Ruas**
Untuk menghilangkan pangkat kuadrat di ruas kiri, kita tarik akar di ruas kanan (jangan lupa tambahkan tanda $\pm$):
$$x + \frac{b}{2a} = \pm \sqrt{\frac{b^2 - 4ac}{4a^2}}$$

Karena $\sqrt{4a^2} = 2a$, maka bentuknya menjadi:
$$x + \frac{b}{2a} = \pm \frac{\sqrt{b^2 - 4ac}}{2a}$$

**8. Pindahkan $\frac{b}{2a}$ ke Ruas Kanan**
$$x = -\frac{b}{2a} \pm \frac{\sqrt{b^2 - 4ac}}{2a}$$

Karena penyebutnya sudah sama-sama $2a$, kita bisa satukan rumusnya menjadi:
$$x_{1,2} = \frac{-b \pm \sqrt{b^2 - 4ac}}{2a}$$

**Q.E.D.** (Selesai Dibuktikan).

---

{{< banner-buku >}}
