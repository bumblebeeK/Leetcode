/*
 * @lc app=leetcode.cn id=729 lang=rust
 *
 * [729] 我的日程安排表 I
 */

// @lc code=start
#[derive(Default)]
struct MyCalendar {
    scope: std::collections::BTreeMap<i32, i32>,
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyCalendar {

    fn new() -> Self {
        Default::default()
    }
    
    fn book(&mut self, start: i32, end: i32) -> bool {
        for (&k,_) in self.scope.range(start..).take(1) {
            if k < end {
                return false
            }
        }
        for (_, &v) in self.scope.range(..start).rev().take(1) {
            if v > start {
                return false
            }
        }
        self.scope.insert(start, end);
        true 
    }
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * let obj = MyCalendar::new();
 * let ret_1: bool = obj.book(start, end);
 */
// @lc code=end

