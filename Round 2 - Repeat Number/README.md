# Số Lặp

Một số được gọi là **"số lặp"** nếu nó có thể được tạo thành bằng cách viết hai số nguyên dương giống nhau nối liền nhau. Ví dụ:

  * **1212** là số lặp (tạo thành từ việc nối hai số `12`)
  * **123123** là số lặp (tạo thành từ việc nối hai số `123`)
  * **11** là số lặp (tạo thành từ việc nối hai số `1`)

Trong khi đó, các số như `12`, `123122` không phải là số lặp.

Bạn được cho một số nguyên dương $N$. Nhiệm vụ của bạn là đếm xem có tất cả bao nhiêu số lặp trong đoạn $[1, N]$.

## Ví dụ

### Input

```
56
```

### Output

```
5
```

### Giải thích

Các số lặp nằm trong đoạn $[1, 56]$ là: $11, 22, 33, 44, 55$.

-----

## Định dạng Input

Input gồm một số nguyên dương $N$ duy nhất.

-----

## Ràng buộc

$1≤N≤1000000$

-----

## Định dạng Output

Output gồm một số nguyên dương là số lượng "số lặp" trong đoạn từ $[1, N]$.