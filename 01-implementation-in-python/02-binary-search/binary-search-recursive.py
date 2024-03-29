
search_element=8
input_list_sample  = [1, 8, 11, 12, 12, 21, 23, 32, 34, 43, 45, 65, 67, 89, 90]
class BinarySearchRecursive:
    
    def __init__(self):
        print("Initializing Components")
        
# This below function is intended to do a binary search using a recursive approach
    def binary_search(self,search_list,low_index,high_index):
        # Setting default search index to -1 to denote the element is not present in the list
        search_index=-1
        print("low_index :: "+str(low_index)+ " high_index :: "+str(high_index))
        if high_index >= low_index:
            mid_index = (low_index + high_index) // 2
            print("Mid Index :: "+str(mid_index)+" Mid Element :: "+str(str(search_list[mid_index])))
            if search_list[mid_index] == search_element:
                search_index = mid_index
            elif search_element < search_list[mid_index]:
                print("Element should be searched on the left Half of the list")
                high_index = mid_index - 1
                print("low_index :: "+str(low_index)+ " high_index :: "+str(high_index))
                return self.binary_search(search_list,low_index,high_index)
            elif search_element > search_list[mid_index]:
                print("Element should be searched on the right side of the list")
                low_index = mid_index + 1
                print("low_index :: "+str(low_index)+ " high_index :: "+str(high_index))
                return self.binary_search(search_list,low_index,high_index)
        return search_index
            
    
binary_search_ins=BinarySearchRecursive()
final_index = len(input_list_sample) - 1
found_index = binary_search_ins.binary_search(input_list_sample,0,final_index)
if found_index == -1:
    print("The searched element :: "+str(search_element)+" is not found in the list ")
else:
    print("The searched element :: "+str(search_element)+" is found at the index :: "+str(found_index))
