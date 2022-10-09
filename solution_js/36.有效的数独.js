/*
 * @lc app=leetcode.cn id=36 lang=javascript
 *
 * [36] 有效的数独
 */

// @lc code=start
/**
 * @param {character[][]} board
 * @return {boolean}
 */
var isValidSudoku = function(board) {
    var rowMatrix = [[],[],[],[],[],[],[],[],[],[]]  // 横坐标里的数字出现标记
    var colMatrix = [[],[],[],[],[],[],[],[],[],[]]  // 纵坐标里的数字出现标记
    var boxMatrix = [[],[],[],[],[],[],[],[],[],[]]  // 3*3 小组盒子的出现标记
    // 关于盒子的理解，一共有9个3*3的盒子
    // 纵坐标在[0,2]的为第一列，纵坐标为[3,5]的为第二列，纵坐标[6,8]的为第三列
    // 故在那一列即为纵坐标/3 ， 在这一列的第几个格子即为横坐标/3
    // 在哪一个格子即为纵坐标/3 *3 + 横坐标/3
    for(let i= 0;i< 9;i++){
        for(let j =0;j<9;j++){
            if(board[i][j]=='.'){ 
                continue // 如果格子里没有数字，则忽略
            }
            let curNum = board[i][j] - '0'   // 格子中的每一位是byte类型，我们减去字符'0'来节省空间 
                                            // 理解：ascii 中字符'0'为48，'9'为57
            if (rowMatrix[i][curNum] === 1 || colMatrix[j][curNum] === 1 || boxMatrix[parseInt(parseInt(i/3)*3+j/3) ][curNum] === 1){
                return false // 如不满足三个条件，直接返回false
            }
            // 标记三个辅助矩阵
            rowMatrix[i][curNum] = 1
			colMatrix[j][curNum] = 1
			boxMatrix[parseInt(parseInt(i/3)*3+j/3)][curNum] = 1
        }
    }
    // 三个条件都符合，返回true
    return true
};

// @lc code=end
