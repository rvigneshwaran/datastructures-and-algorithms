
search_element=90
input_list_sample  = [11,34,21,1,45,67,12,90,12,8,89,23,65,43,32]
input_list_sample = sorted(input_list_sample)
print(input_list_sample)
class BinarySearch:
    
    def __init__(self):
        print("Initializing Components")
        
# This below function is intended to do a binary search using a recursive approach
    def binary_search(self,search_list,low_index,high_index,count):
        count=count+1
        search_element_index = -1
        print("low_index :: "+str(low_index)+ " high_index :: "+str(high_index))
        if high_index >= 1 and count < 5:
            mid_index =  1 + (high_index - 1) // 2
            print("Mid Index :: "+str(mid_index)+" Mid Element :: "+str(str(search_list[mid_index])))
            if search_list[mid_index] == search_element:    
                search_element_index=mid_index
            elif search_element < search_list[mid_index]:
                print("Element should be searched on the left Half of the list")
                low_index = low_index
                high_index = mid_index - 1
                print("low_index :: "+str(low_index)+ " high_index :: "+str(high_index))
                self.binary_search(search_list,low_index,high_index,count)
            elif search_element > search_list[mid_index]:
                print("Element should be searched on the right side of the list")
                low_index = mid_index + 1
                high_index = high_index
                print("low_index :: "+str(low_index)+ " high_index :: "+str(high_index))
                self.binary_search(search_list,low_index,high_index,count)
        else:
            return search_element_index
        return search_element_index
    
binary_search_ins=BinarySearch()
total_index = len(input_list_sample) - 1
search_index = binary_search_ins.binary_search(input_list_sample,0,total_index,0)
print("The searched element :: "+str(search_element)+" is found at the index :: "+str(search_index))