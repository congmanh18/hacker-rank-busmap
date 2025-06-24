Chúng ta được cung cấp một phương trình bậc hai có dạng:

$$
a x^2 + b x + c = 0
$$

Và được **đảm bảo có hai nghiệm phân biệt**. Nhiệm vụ là tìm giá trị:

$$
x_1^2 + x_2^2
$$

và biểu diễn nó dưới dạng **phân số tối giản** $u / v$, với $u, v$ là số nguyên dương.

---

## 🧠 Ý tưởng giải

### Bước 1: Sử dụng hệ thức Vi-ét

Với phương trình bậc hai $ax^2 + bx + c = 0$, nếu có 2 nghiệm phân biệt $x_1, x_2$, thì:

* $x_1 + x_2 = -\frac{b}{a}$
* $x_1 x_2 = \frac{c}{a}$

Ta cần tính:

$$
x_1^2 + x_2^2
$$

Sử dụng đẳng thức sau:

$$
x_1^2 + x_2^2 = (x_1 + x_2)^2 - 2x_1x_2
$$

Thay bằng hệ thức Vi-ét:

$$
x_1^2 + x_2^2 = \left(-\frac{b}{a}\right)^2 - 2 \cdot \frac{c}{a}
= \frac{b^2}{a^2} - \frac{2c}{a}
$$

### Bước 2: Quy đồng mẫu số

Để biểu diễn dưới dạng phân số:

$$
\frac{b^2}{a^2} - \frac{2c}{a} = \frac{b^2 - 2ac}{a^2}
$$

Vậy là:

$$
x_1^2 + x_2^2 = \frac{b^2 - 2ac}{a^2}
$$

### Bước 3: Rút gọn phân số $\frac{u}{v}$

* Tính tử số: $u = b^2 - 2ac$
* Tính mẫu số: $v = a^2$
* Rút gọn phân số $\frac{u}{v}$ bằng cách chia cả tử và mẫu cho **ước chung lớn nhất (GCD)**.

---

## ✅ Triển khai ví dụ mẫu

### Input:

```
1 3 2
```

### Áp dụng công thức:

* $a = 1, b = 3, c = 2$
* Tử số: $b^2 - 2ac = 9 - 4 = 5$
* Mẫu số: $a^2 = 1$

Vậy phân số là:

$$
\frac{5}{1}
\Rightarrow u = 5, v = 1
$$
---

## ✅ Tổng kết

Công thức cần nhớ:

$$
x_1^2 + x_2^2 = \frac{b^2 - 2ac}{a^2}
$$

Tính tử và mẫu, rồi rút gọn phân số là xong!
