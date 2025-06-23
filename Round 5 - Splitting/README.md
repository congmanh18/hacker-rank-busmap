# Đếm cách chia dãy thành hai phần có tổng bằng nhau

Bạn được cung cấp một dãy số nguyên $a$ gồm $N$ số. Nhiệm vụ của bạn là đếm số cách chia dãy $a$ thành hai phần:

  * **Phần thứ nhất:** Gồm các phần tử liên tục từ vị trí $1$ đến vị trí $i$.
  * **Phần thứ hai:** Gồm các phần tử còn lại, từ vị trí $i+1$ đến $N$.

Sao cho tổng các phần tử của mỗi phần là bằng nhau.


## Định dạng Input

  * Dòng đầu tiên: Một số nguyên dương $N$, là số lượng phần tử trong dãy $a$.
  * Dòng thứ hai: $N$ số nguyên, số nguyên thứ $i$ là giá trị thứ $i$ của dãy $a$.

-----

## Ràng buộc

  * $1 ≤ N ≤ 100$
  * $-2 ≤ a_{i} ≤ 2$

-----

## Định dạng Output

Một số nguyên không âm duy nhất, tương ứng là số cách chia dãy thành hai phần sao cho tổng các phần tử trong mỗi phần bằng nhau.

-----

## Ví dụ mẫu 0

### Input

```
5
0 0 1 -1 0
```

### Output

```
3
```

### Giải thích

Các cách chia dãy thành hai phần có tổng bằng nhau là:

  * **Cách 1:** $[0]$ và $[0, 1, -1, 0]$
  * **Cách 2:** $[0, 0]$ và $[1, -1, 0]$
  * **Cách 3:** $[0, 0, 1, -1]$ và $[0]$