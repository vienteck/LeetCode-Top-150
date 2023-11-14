class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        tracker = {}
        for x in nums:
            tracker[x] = True
        max = 0
        for x in nums:
            counter = 0 
            if x-1 in tracker:
                continue 
            
            while x in tracker:
                counter += 1
                x += 1 

            if counter > max:
                max = counter
        
        return max

    
longestConsecutive([100,4,200,1,3,2])