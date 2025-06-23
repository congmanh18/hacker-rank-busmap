# Truy Vết Người Có Nguy Cơ Nhiễm Bệnh (COVID-19)

Năm 2020 đánh dấu kỷ niệm 20 năm của Thách thức, và cũng là một năm đặc biệt với đại dịch COVID-19. Mặc dù tưởng chừng Thách thức sẽ lỡ hẹn, nhưng với tinh thần "Thách thức mọi thách thức", mùa giải kỷ niệm 20 năm vẫn diễn ra. Trong trận chung kết lần thứ 20, các thí sinh có nhiệm vụ viết một thuật toán góp phần chống lại dịch COVID-19: **bài toán truy vết người có nguy cơ nhiễm bệnh**.

## Mô tả Bài toán

Có $N$ người được đánh số từ 1 đến $N$. Ta có $M$ nhóm bạn thường xuyên giao lưu với nhau. Mỗi nhóm gồm một số người thuộc $N$ người, biết rằng một người có thể thuộc nhiều nhóm bạn khác nhau. Trong một nhóm, mỗi người đều tiếp xúc gần với mọi người còn lại một cách thường xuyên.

Cho danh sách $K$ người trong số $N$ người bị phát hiện nhiễm COVID-19 (gọi là **F0**), hãy xác định những người là **F1**, **F2** và **F3** để đưa đi cách ly. Để đơn giản hóa cách thức hiển thị kết quả, chương trình của bạn chỉ cần xuất ra số lượng người là F1, F2 và F3.

Cơ chế xác định một người là F1, F2 hay F3 như sau:

  * **F0:** Người bị nhiễm bệnh.
  * **F1:** Người tiếp xúc gần với F0.
  * **F2:** Người tiếp xúc gần với F1 (và không phải F0 hoặc F1).
  * **F3:** Người tiếp xúc gần với F2 (và không phải F0, F1 hoặc F2).

-----

## Định dạng Input

  * Dòng đầu tiên: Chứa 3 số nguyên $N, M, K$ lần lượt là số người, số nhóm và số người nhiễm.
  * $M$ dòng tiếp theo: Mỗi dòng mô tả một nhóm bạn. Dòng thứ $i$ có định dạng như sau: `S[i] A[i,1] A[i,2] ... A[i,Si]`
      * Trong đó $S[i]$ là số lượng người của nhóm thứ $i$.
      * Tiếp theo là $S[i]$ số khác nhau, lần lượt là số hiệu của từng người: $A[i,1], A[i,2], \\ldots, A[i,Si]$.
  * Dòng cuối cùng: Chứa $K$ số khác nhau: $B[1], B[2], \\ldots, B[K]$ là số hiệu của những người bị nhiễm bệnh COVID-19 (tức **F0**).

-----

## Ràng buộc

  * $1 ≤ N, M, K ≤ 10^5$
  * $S[i] ≤ 10$ (Tức mỗi nhóm có tối đa 10 người)
  * $1 ≤ A[i,j], B[i] ≤ N$

-----

## Định dạng Output

Một dòng chứa 3 số nguyên $C[1], C[2], C[3]$ cách nhau bởi khoảng trắng, lần lượt là số lượng người là F1, F2, F3.

-----

## Ví dụ mẫu 0

### Input

```
5 2 1
3 3 4 5
3 1 2 3
1
```

### Output

```
2 2 0
```

### Giải thích

  * Có 5 người và 2 nhóm bạn.
  * Nhóm 1: Người 3, 4, 5.
  * Nhóm 2: Người 1, 2, 3.
  * Người 1 bị nhiễm COVID-19 nên là **F0**.
  * Người 2, 3 tiếp xúc với người 1 (trong nhóm 2) nên là **F1**.
  * Người 4, 5 tiếp xúc với người 3 (trong nhóm 1) nên là **F2**.
  * Vậy ta có 2 F1 và 2 F2 (và 0 F3).