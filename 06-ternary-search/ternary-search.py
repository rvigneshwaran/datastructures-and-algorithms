# The below is the inpit list to perform the search.
input_list_sample  = [1, 8, 11, 12, 12, 21, 23, 32, 34, 43, 45, 65, 67, 89, 90]
print(input_list_sample)

class TernarySearch:
    
    def __init__(self):
        print("Initializing Components for Search : ")
    
    def ternary_search(self,input_list,start_index,end_index,search_element):
        """ The Ternary Search is a reduce by constant and conqure algorithm where the size of the list 
        is reduced by a constant and then conqure the list.

        Args:
            input_list (List): The Input list which we are going to search.
            start_index (Number): Starting Index at which the list starts to perform ternary search operation.
            end_index (Number): Ending Index at which the list ends to perform the ternary search operation.
            search_element (Number): element to search in the list

        Returns:
            _type_: index of the search element , if the element is not present in the list it returns -1.
        """
        if end_index >= start_index:
            mid_index1 = start_index + (end_index - start_index) // 3
            mid_index2 = end_index - (end_index - start_index) // 3
            print("mid_index1 :: "+str(mid_index1)+" mid_index2 :: "+str(mid_index2))
            print("start_index :: "+str(start_index)+" end_index :: "+str(end_index))
            if search_element == input_list[mid_index1]:
                return mid_index1
            elif search_element == input_list[mid_index2]:
                return mid_index2
            elif search_element < input_list[mid_index1]:
                return self.ternary_search(input_list,start_index,mid_index1-1,search_element)
            elif search_element > input_list[mid_index2]:
                return self.ternary_search(input_list,mid_index1+1,end_index,search_element)
            else:
                return self.ternary_search(input_list,mid_index1+1,mid_index2-1,search_element)
        return -1

    def print_result(self,result_index,search_element):
        if result_index is not None:
            print(str(search_element) + " is found in the list at index :: "+str(result_index))
        else:
            print(str(search_element) + " is not found in the list")
    
ternary_search_ins = TernarySearch()
search_element_cont = [21,23,43,100,90,1]
for search_element in search_element_cont:
    start_index = 0
    end_index = len(input_list_sample) - 1
    search_index = ternary_search_ins.ternary_search(input_list_sample,start_index,end_index,search_element)
    ternary_search_ins.print_result(search_index,search_element)
    print("-----------------------------------------------------")