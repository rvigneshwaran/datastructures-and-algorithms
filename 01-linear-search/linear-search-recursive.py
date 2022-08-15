import sys

input_list_sample  = [1, 8, 11, 12, 12, 21, 23, 32, 34, 43, 45, 65, 67, 89, 90]
print(input_list_sample)

class LinearSearchRecursive:
    
    def __init__(self):
        print("Initializing Components")
        
    def drecursive_linear_search(self,search_list,start_index,end_index,search_element):
        search_element=int(search_element)
        if end_index < start_index:
            return -1
        elif search_element == search_list[start_index]:
            return start_index
        elif search_element == search_list[end_index]:
            return end_index
        return self.drecursive_linear_search(search_list,start_index,end_index-1,search_element)
    
    def irecursive_linear_search(self,search_list,start_index,end_index,search_element):
        search_element=int(search_element)
        if start_index > len(search_list) - 1:
            return -1
        elif search_element == search_list[start_index]:
            return start_index
        elif search_element == search_list[end_index]:
            return end_index
        return self.irecursive_linear_search(search_list,start_index + 1,end_index,search_element)
    
    def print_search_results(self,result_index,search_element):
        if result_index is not None and result_index != -1:
            print("The searched element :: "+str(search_element)+" is found at the index :: "+str(result_index))
        else:
            print("The searched element :: "+str(search_element)+" is not found in the list ")
        print("+"*50)
            
    
rlinear_search_ins=LinearSearchRecursive()

# Getting total no of arguements from command line
args_length = len(sys.argv)

# Check if the command line has valid arguments
if args_length == 2:
    search_element=sys.argv[1]
    final_index = len(input_list_sample) - 1
    found_index = rlinear_search_ins.drecursive_linear_search(input_list_sample,0,final_index,search_element)
    rlinear_search_ins.print_search_results(found_index,search_element)
    found_index = rlinear_search_ins.irecursive_linear_search(input_list_sample,0,final_index,search_element)
    rlinear_search_ins.print_search_results(found_index,search_element)
else:
    print("Please enter valid arguments for searching elements in the list")
