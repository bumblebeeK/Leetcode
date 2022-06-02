/*
 * @lc app=leetcode.cn id=450 lang=rust
 *
 * [450] 删除二叉搜索树中的节点
 */

// @lc code=start
// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn delete_node(root: Option<Rc<RefCell<TreeNode>>>, key: i32) -> Option<Rc<RefCell<TreeNode>>> {
        Solution::dfs(&root,key)
    }
    fn dfs(node: &Option<Rc<RefCell<TreeNode>>>, key: i32) -> Option<Rc<RefCell<TreeNode>>> {
        if let Some(n) = node {
            let val = n.borrow().val;
            match val.cmp(&key) {
                std::cmp::Ordering::Greater =>{
                    let l = Solution::dfs(&n.borrow().left,key);
                    n.borrow_mut().left = l;
                }
                std::cmp::Ordering::Less => {
                    let r = Solution::dfs(&n.borrow().right,key);
                    n.borrow_mut().right = r;
                }
                std::cmp::Ordering::Equal => {
                   if n.borrow().left.is_none() {
                       return n.borrow().right.clone()
                   }
                   if n.borrow().right.is_none() {
                       return n.borrow().left.clone()
                   }
                   let next = Solution::search(&n.borrow().right);
                   if let Some(val) = next {
                       let r = Solution::dfs(&n.borrow().right,val);
                       n.borrow_mut().val = val;
                       n.borrow_mut().right = r;
                   }
                }
            }
        }
        node.clone()
    }
    fn search(node: &Option<Rc<RefCell<TreeNode>>>) -> Option<i32> {
        if let Some(n) = node {
            if n.borrow().left.is_some() {
                Solution::search(&n.borrow().left)
            } else {
                Some(n.borrow().val)
            }
        } else {
            None
        }
    }
}
// @lc code=end

