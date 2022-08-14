import sys

input_list_sample  = [1, 8, 11, 12, 12, 21, 23, 32, 34, 43, 45, 65, 67, 89, 90]
print(input_list_sample)

class BinarySearchIterative:
    
    def __init__(self):
        print("Initializing Components")
        
# This below function is intended to do a binary search using a Iterative approach
    def binary_search(self,search_list,low_index,high_index,search_element):
        search_element=int(search_element)
        while (low_index <= high_index):
            mid_index = (low_index + high_index) // 2
            print("low_index :: "+str(low_index)+" high_index :: "+str(high_index)+" mid_index :: "+str(mid_index))
            if search_list[mid_index] == search_element:
                print("Search element determined at position :: "+str(mid_index))
                return mid_index
            elif search_element < search_list[mid_index]:
                high_index = mid_index - 1
            elif search_element > search_list[mid_index]:
                low_index = mid_index + 1
        return -1
            
    
binary_search_ins=BinarySearchIterative()

# Getting total no of arguements from command line
args_length = len(sys.argv)

# Check if the command line has valid arguments
if args_length == 2:
    search_element=sys.argv[1]
    final_index = len(input_list_sample) - 1
    found_index = binary_search_ins.binary_search(input_list_sample,0,final_index,search_element)
    if found_index == -1:
        print("The searched element :: "+str(search_element)+" is not found in the list ")
    else:
        print("The searched element :: "+str(search_element)+" is found at the index :: "+str(found_index))
else:
    print("Please enter valid arguments for searching elements in the list")
