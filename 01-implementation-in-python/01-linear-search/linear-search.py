
input_list_sample1  = [11,34,21,1,45,67,1,90,12,8,89,23,65,43,32]

class LinearSearch:
    
    def __init__(self):
        print("Initializing Components for Search : ")
        
    def linear_search(self,input_list,search_element):
        found_index = None
        if input_list is not None and search_element is not None:
            for index,input_element in enumerate(input_list):
                if input_element == search_element:
                    found_index = index
                    break
        return found_index

    def print_result(self,result_index,search_element):
        if result_index is not None:
            print(str(search_element) + " is found in the list at index :: "+str(result_index))
        else:
            print(str(search_element) + " is not found in the list")
    
linear_search_ins = LinearSearch()
search_element_cont = [21,1,32]
for search_element in search_element_cont:
    search_index = linear_search_ins.linear_search(input_list_sample1,search_element)
    linear_search_ins.print_result(search_index,search_element)