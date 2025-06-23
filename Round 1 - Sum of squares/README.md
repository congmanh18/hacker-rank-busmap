# Tổng nghiệm bình phương

Bạn được cung cấp một phương trình bậc hai có dạng: `a * x^2 + b * x + c = 0`. Phương trình này **đảm bảo có hai nghiệm phân biệt** $x₁$ và $x₂$. Giá trị $x₁² + x₂²$ có thể được biểu diễn dưới dạng một phân số tối giản $u/v$, trong đó $u$ và $v$ là hai số nguyên dương. Nhiệm vụ của bạn là tìm hai giá trị $u$ và $v$.


## Định dạng Input

Input bao gồm một dòng duy nhất chứa ba số nguyên $a, b, c$.
Dữ liệu input luôn đảm bảo phương trình trên có hai nghiệm phân biệt.

### Ràng buộc
```
-10 ≤ a, b, c ≤ 10
```
### Định dạng Output

Output gồm hai số nguyên $u$ và $v$ cách nhau bởi một dấu cách.

-----

## Ví dụ mẫu 0

### Input

```
1 3 2
```

### Output

```
5 1
```

### Giải thích

Phương trình $x^2 + 3x + 2 = 0$ có hai nghiệm là $x = -1$ và $x = -2$.
Tổng bình phương của hai nghiệm là $(-1)^2 + (-2)^2 = 1 + 4 = 5$.
Giá trị này có thể được viết dưới dạng phân số tối giản là $5/1$.