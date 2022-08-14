
search_element=90
input_list_sample  = [11,34,21,1,45,67,12,90,12,8,89,23,65,43,32]
input_list_sample = sorted(input_list_sample)
print(input_list_sample)
class BinarySearch:
    
    def __init__(self):
        print("Initializing Components")
        
# This below function is intended to do a binary search using a recursive approach
    def binary_search(self,search_list,low_index,high_index):
        search_element_index = -1
        print("low_index :: "+str(low_index)+ " high_index :: "+str(high_index))
        if high_index >= 1:
            mid_index = (low_index + high_index) // 2
            print("Mid Index :: "+str(mid_index)+" Mid Element :: "+str(str(search_list[mid_index])))
            if search_list[mid_index] == search_element:
                print("Found the search element at index :: "+str(mid_index))
                return mid_index
            elif search_element < search_list[mid_index]:
                print("Element should be searched on the left Half of the list")
                high_index = mid_index - 1
                print("low_index :: "+str(low_index)+ " high_index :: "+str(high_index))
                self.binary_search(search_list,low_index,high_index)
            elif search_element > search_list[mid_index]:
                print("Element should be searched on the right side of the list")
                low_index = mid_index + 1
                print("low_index :: "+str(low_index)+ " high_index :: "+str(high_index))
                self.binary_search(search_list,low_index,high_index)
        return search_element_index
    
binary_search_ins=BinarySearch()
final_index = len(input_list_sample) - 1
found_index = binary_search_ins.binary_search(input_list_sample,0,final_index)
print("The searched element :: "+str(search_element)+" is found at the index :: "+str(found_index))