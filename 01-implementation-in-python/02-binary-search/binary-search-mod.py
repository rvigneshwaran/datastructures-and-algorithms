input_list_sample  = [1, 8, 11, 12, 12, 21, 23, 32, 34, 43, 45, 65, 67, 89, 90]
print("Input List :: \n "+str(input_list_sample))

class BinarySearchMod:
    
    def __init__(self):
        print("Initializing Components")
        
    def binary_search_mod(self,search_list,low_index,high_index,search_element):
        """ The below function is intended to perform binary search with less number of comparisions 
        using the iterative approach , basically a modified version of the binary search in Iterative approach.

        Args:
            search_list (_type_): _description_
            low_index (_type_): _description_
            high_index (_type_): _description_
            search_element (_type_): _description_

        Returns:
            _type_: _description_
        """
        while (high_index - low_index) > 1 :
            mid_index =  low_index + (high_index - low_index) // 2
            print("low_index :: "+str(low_index)+" high_index :: "+str(high_index)+" mid_index :: "+str(mid_index))
            if search_element > search_list[mid_index]:
                low_index = mid_index
            else:
                high_index = mid_index
        if search_list[low_index] == search_element:
            return low_index
        elif search_list[high_index] == search_element:
            return high_index
        else:
            return -1
        
    def binary_search(self,search_list,low_index,high_index,search_element):
        while (low_index <= high_index):
            # For Bigger Values of Integres this might cause overflow , So this should be slitly modified mid = 
            mid_index = (low_index + high_index) // 2
            print("low_index :: "+str(low_index)+" high_index :: "+str(high_index)+" mid_index :: "+str(mid_index))
            if search_list[mid_index] == search_element:
                #print("Search element determined at position :: "+str(mid_index))
                return mid_index
            elif search_element < search_list[mid_index]:
                high_index = mid_index - 1
            elif search_element > search_list[mid_index]:
                low_index = mid_index + 1
        return -1
    
    def print_search_results(self,result_index,search_element):
        if result_index is not None and result_index != -1:
            print("The searched element :: "+str(search_element)+" is found at the index :: "+str(result_index))
        else:
            print("The searched element :: "+str(search_element)+" is not found in the list ")
            
binary_search_mod=BinarySearchMod()

# possible list of elements to search from the list
search_element_cont = [21,34,45,100,90,1]

if search_element_cont is not None and len(search_element_cont) > 0:
    for search_element in search_element_cont:
        start_index = 0
        final_index = len(input_list_sample) - 1
        print("-"*50)
        found_index = binary_search_mod.binary_search(input_list_sample,0,final_index,search_element)
        binary_search_mod.print_search_results(found_index,search_element)
        print("*"*25)
        mod_found_index = binary_search_mod.binary_search_mod(input_list_sample,0,final_index,search_element)
        binary_search_mod.print_search_results(mod_found_index,search_element)
else:
    print("There are no elements to search, Kindly provide search elements")
