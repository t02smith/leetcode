
from typing import List

class Solution1:
    def minCostClimbingStairs(self, cost: List[int]):
        minCost = [0,0]
        
        for i in range(2, len(cost)+1):
            minCost.append(min(minCost[i-1]+cost[i-1], minCost[i-2]+cost[i-2]))
            
        return minCost[len(cost)]
    
class Solution2:
    def minCostClimbingStairs(self, cost: List[int]):
        for i in range(2, len(cost)):
            cost[i] += min(cost[i-1], cost[i-2])
            
        return min(cost[len(cost)-1], cost[len(cost)-2])
        
s = Solution2()

assert s.minCostClimbingStairs([1,100,1,1,1,100,1,1,100,1]) == 6
assert s.minCostClimbingStairs([10,15,20]) == 15
    
print("all tests passed")