import math
from re import search

from numpy import size 

# The below is the inpit list to perform the search.
input_list_sample  = [1, 8, 11, 12, 12, 21, 23, 32, 34, 43, 45, 65, 67, 89, 90]
print(input_list_sample)

class JumpSearch:
    
    def __init__(self):
        print("Initializing Components for Search : ")
        
    def linear_search(self,input_list,search_element,start_index):
        if input_list is not None and search_element is not None:
            for index,indv_element in enumerate(input_list):
                if indv_element == search_element:
                    return index + start_index
        return -1
        
    def jump_search(self,input_list,search_element):
        list_length=len(input_list)
        block_size=int(math.sqrt(list_length))
        print("The Block size of the whole list is :: "+str(block_size))
        start_index=0
        end_index=block_size
        if search_element >= input_list[start_index] and search_element <= input_list[list_length-1]:
            while input_list[end_index] <= search_element and end_index < list_length:
                start_index=end_index
                end_index=end_index+block_size
                # restricting the end_index to be the final element in the list
                if end_index > list_length - 1:
                    end_index = list_length
            print("Start Index :: "+str(start_index)+" End Index :: "+str(end_index))
            return self.linear_search(input_list[start_index:end_index],search_element,start_index)
        else:
            return -1

    def print_result(self,result_index,search_element):
        if result_index is not None:
            print(str(search_element) + " is found in the list at index :: "+str(result_index))
        else:
            print(str(search_element) + " is not found in the list")
    
jump_search_ins = JumpSearch()
search_element_cont = [21,23,43,100,90]
for search_element in search_element_cont:
    search_index = jump_search_ins.jump_search(input_list_sample,search_element)
    jump_search_ins.print_result(search_index,search_element)