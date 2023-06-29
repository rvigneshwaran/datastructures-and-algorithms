import math

input_list_contends = [10, 2, 4, 61, 13, 5, 81]
print(input_list_contends)

class FindFixedPoint:
    
    def findFixedPointByApp2(self,start_index,end_index):
        if(input_list_contends is not None and len(input_list_contends) > 0):
            if end_index > start_index:
                mid_element = start_index + (end_index - start_index) // 2
                if input_list_contends[mid_element] == mid_element:
                    return mid_element
                elif mid_element + 1 <= input_list_contends[end_index]:
                    return self.findFixedPointByApp2(mid_element+1,end_index)
                elif mid_element -1  <= input_list_contends[start_index]:
                    return self.findFixedPointByApp2(start_index,mid_element - 1)
                else:
                    return -1
            return -1
        
    
    def findFixedPointbyApp3(self):
        if input_list_contends is not None and len(input_list_contends) > 0:
            list_length = len(input_list_contends)
            block_size = int(math.sqrt(list_length))
            start_index = 0
            end_endex = block_size
        return -1
    
    def findFixedPointbyApp1(self):
        if(input_list_contends is not None and len(input_list_contends) > 0):
            for index,element in enumerate(input_list_contends):
                if index == element:
                    print("Found Fixed Element in index :: "+str(index))
                    return index
            return -1

    def print_execution_result(self,fixed_point_index,approach_identifier):
        if fixed_point_index is not None and str(fixed_point_index) != -1:
            print("Elements of the Fixed Point using "+str(approach_identifier)+" approach :: \n "+str(fixed_point_index))
        else:
            print("No Fixed Points in the list")

# find the element by using Linear Search    
find_fixedpoint_ins = FindFixedPoint()
fixed_point_list = find_fixedpoint_ins.findFixedPointbyApp1()
find_fixedpoint_ins.print_execution_result(fixed_point_list,"Linear Search")

# find the element by using Binary Search
fixed_point_list_01 = find_fixedpoint_ins.findFixedPointByApp2(0,len(input_list_contends)-1)
find_fixedpoint_ins.print_execution_result(fixed_point_list_01,"Binary Search")