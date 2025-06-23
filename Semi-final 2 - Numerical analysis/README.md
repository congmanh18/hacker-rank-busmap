# Đếm số dãy số thỏa mãn tổng tích

Bạn được cho hai số nguyên dương $N$ và $K$, trong đó $K$ luôn là một số chẵn. Nhiệm vụ của bạn là đếm số lượng dãy số khác nhau $A1, A2, ..., AK$ sao cho:

$A[1] * A[2] + A[3] * A[4] + ... + A[{K-1}] * A[K] = N$ và $1 ≤ A\_i ≤ N$ cho mọi $i$ từ $1$ đến $K$.

Hai dãy $A[1], A[2], ..., A[K]$ và $B[1], B[2], ..., B[K]$ được xem là khác nhau nếu tồn tại một số $i$ sao cho $A[i] ≠ B[i]$.

## Ví dụ

Với $N = 3$ và $K = 4$, chúng ta có các cách phân tích sau:

  * $1 * 1 + 1 * 2 = 3$
  * $1 * 1 + 2 * 1 = 3$
  * $1 * 2 + 1 * 1 = 3$
  * $2 * 1 + 1 * 1 = 3$

-----

## Định dạng Input

Một dòng chứa hai số nguyên $N, K$.

-----

## Ràng buộc

  * $1 ≤ N ≤ 20$
  * $2 ≤ K ≤ 10$, $K$ luôn là một số chẵn

-----

## Định dạng Output

Một dòng duy nhất chứa số lượng dãy số khác nhau thỏa mãn yêu cầu của đề bài.

-----

## Ví dụ mẫu 0

### Input

```
3 4
```

### Output

```
4
```

### Giải thích

Có 4 cách giải thích như đã nêu ở trên.