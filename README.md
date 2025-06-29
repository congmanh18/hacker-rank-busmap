# hacker-rank-busmap
### 1. **Sliding Window (Cửa sổ trượt)**

**👉 Bài toán:** Tìm độ dài chuỗi con dài nhất có tổng ≤ `k`

```go
// Cho mảng nums và số k, tìm độ dài đoạn con liên tiếp có tổng không vượt quá k
func longestSubarraySumLEK(nums []int, k int) int {
    left, sum, maxLen := 0, 0, 0
    for right := 0; right < len(nums); right++ {
        sum += nums[right]
        for sum > k {
            sum -= nums[left]
            left++
        }
        maxLen = max(maxLen, right-left+1)
    }
    return maxLen
}
```

---

### 2. **Two Pointers (Hai con trỏ)**

**👉 Bài toán:** Kiểm tra mảng có cặp phần tử có tổng bằng `target` (đã sắp xếp)

```go
func hasPairWithSum(arr []int, target int) bool {
    left, right := 0, len(arr)-1
    for left < right {
        sum := arr[left] + arr[right]
        if sum == target {
            return true
        } else if sum < target {
            left++
        } else {
            right--
        }
    }
    return false
}
```

---

### 3. **Fast and Slow Pointers (Con trỏ nhanh chậm)**

**👉 Bài toán:** Phát hiện vòng lặp trong danh sách liên kết

```go
func hasCycle(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```

---

### 4. **Merge Intervals (Gộp khoảng)**

**👉 Bài toán:** Gộp các khoảng chồng lấn

```go
func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    res := [][]int{intervals[0]}
    for _, interval := range intervals[1:] {
        last := res[len(res)-1]
        if interval[0] <= last[1] {
            last[1] = max(last[1], interval[1])
        } else {
            res = append(res, interval)
        }
    }
    return res
}
```

---

### 5. **Cyclic Sort (Sắp xếp vòng lặp)**

**👉 Bài toán:** Tìm số thiếu trong mảng \[1..n]

```go
func missingNumber(nums []int) int {
    i := 0
    for i < len(nums) {
        if nums[i] < len(nums) && nums[i] != nums[nums[i]] {
            nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
        } else {
            i++
        }
    }
    for i := 0; i < len(nums); i++ {
        if nums[i] != i {
            return i
        }
    }
    return len(nums)
}
```

---

### 6. **In-Place Reversal of Linked List**

**👉 Bài toán:** Đảo ngược danh sách liên kết

```go
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head
    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }
    return prev
}
```

---

### 7. **Tree BFS (Duyệt cây theo chiều rộng)**

**👉 Bài toán:** Trả về các node theo từng tầng

```go
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    queue := []*TreeNode{root}
    res := [][]int{}
    for len(queue) > 0 {
        level := []int{}
        size := len(queue)
        for i := 0; i < size; i++ {
            node := queue[0]
            queue = queue[1:]
            level = append(level, node.Val)
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        res = append(res, level)
    }
    return res
}
```

---

### 8. **Tree DFS (Duyệt cây theo chiều sâu)**

**👉 Bài toán:** Duyệt theo tiền tố

```go
func preorderTraversal(root *TreeNode) []int {
    var res []int
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        res = append(res, node.Val)
        dfs(node.Left)
        dfs(node.Right)
    }
    dfs(root)
    return res
}
```

---

### 9. **Two Heaps (Hai heap)**

**👉 Bài toán:** Tìm median liên tục

```go
// Cần sử dụng container/heap cho Go
// MinHeap lưu nửa lớn hơn, MaxHeap lưu nửa nhỏ hơn
```

---

### 10. **Subsets (Tập con)**

**👉 Bài toán:** Sinh tất cả tập con

```go
func subsets(nums []int) [][]int {
    res := [][]int{{}}
    for _, num := range nums {
        n := len(res)
        for i := 0; i < n; i++ {
            subset := append([]int{}, res[i]...)
            subset = append(subset, num)
            res = append(res, subset)
        }
    }
    return res
}
```

---

### 11. **Modified Binary Search**

**👉 Bài toán:** Tìm phần tử đầu tiên ≥ target

```go
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] >= target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return left
}
```

---

### 12. **Top K Elements (K phần tử lớn nhất)**

**👉 Bài toán:** Tìm k phần tử lớn nhất

```go
// Dùng heap: Push từng phần tử, nếu size > k thì pop phần tử nhỏ nhất
```

---

### 13. **K-Way Merge (Hợp nhất K dãy)**

**👉 Bài toán:** Merge K mảng đã sắp xếp

```go
// Dùng min heap để theo dõi phần tử đầu mỗi mảng
```

---

### 14. **Topological Sort (Sắp xếp topo)**

**👉 Bài toán:** Sắp xếp thứ tự thực hiện công việc (dạng DAG)

```go
// Dùng BFS (Kahn's algorithm) với in-degree
```