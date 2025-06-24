# Tìm dãy con có tích tạo ra nhiều số 0 cuối cùng nhất

Thomas mới học về phép nhân và rất thích những số tròn chục. Cậu ấy thấy thú vị khi nhân các số với nhau có thể tạo ra một số có nhiều số 0 ở cuối. Hôm nay, Thomas được thầy cho một dãy số nguyên không âm có độ dài $N$. Thomas nhận thấy rằng khi nhân các số liền nhau trong dãy, cậu ấy có thể nhận được một số có rất nhiều số 0 ở cuối.

Hãy giúp Thomas chọn ra một dãy con khác rỗng gồm các số liên tiếp nhau của dãy số đã cho, sao cho tích của các số trong dãy con đó sẽ tạo được một số có nhiều số 0 ở cuối nhất có thể. Lưu ý rằng số `0` là một số đặc biệt và được coi là không có số 0 nào ở cuối.

Cụ thể, cho dãy số $A[1], A[2],... A[N]$, hãy giúp Thomas tìm một dãy con từ $L$ đến $R$ ($1 ≤ L ≤ R ≤ N$) của dãy số sao cho $A[L] * A[L+1] * ... * A[R-1] * A[R] = X$. Mục tiêu là tìm $X$ sao cho nó có nhiều số 0 ở cuối nhất, tức là $X = Y * 10^K$ với $Y$ là một số nguyên dương và $K$ có giá trị lớn nhất có thể. Nếu $X = 0$, coi như giá trị $K = 0$.

Vì có thể có nhiều dãy con thỏa mãn yêu cầu đề bài, bạn chỉ cần xuất ra số lượng số 0 nhiều nhất của kết quả phép nhân dãy số tối ưu tìm được, tức là giá trị $K$.

## Định dạng Input

  * Dòng đầu tiên chứa số nguyên $N$.
  * Dòng tiếp theo chứa $N$ số nguyên không âm $A[1], A[2], ..., A[N]$.

-----

## Ràng buộc

  * $1 ≤ N ≤ 100000$
  * $0 ≤ A[i] ≤ 10^9$

-----

## Định dạng Output

Một số duy nhất là số lượng số 0 ở cuối nhiều nhất của kết quả phép nhân dãy con gồm các số liền kề của dãy $A$.

-----

## Ví dụ mẫu 0

### Input

```
5
1 3 4 5 30
```

### Output

```
2
```

### Giải thích

Khi nhân các số ở dãy con $[4, 5, 30]$ ta có $4 * 5 * 30 = 600$. Như vậy, số có nhiều chữ số 0 nhất mà ta tạo được là $600$, có $2$ chữ số 0 ở cuối ($600 = 6 * 10^2$).

-----