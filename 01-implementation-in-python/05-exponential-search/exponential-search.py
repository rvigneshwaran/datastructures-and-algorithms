import math 

input_contends_list  = [10, 12, 13, 16, 18, 19, 20, 21, 22, 23, 24, 33, 35, 42, 47]
print(input_contends_list)

class ExponentialSearch:
    
    def __init__(self):
        print("Initializing Components for Search : ")
        
    def binary_search(self,input_list,start_index,end_index,search_element):
        """ The below function is inetnded to perform the binary search on the input list and return the index of th
        search elememt , if the element is found in the input list the function returns the index of the element 
        else it returns -1 to denote the element is not present in the input list.

        Args:
            input_list (List): The Input list provided to seach the element 
            start_index (Number): initial index of the list to search
            end_index (Number): boundary of element to search
            search_element (Number): Element to search inside the list

        Returns:
            _type_: _description_
        """
        if end_index >= start_index:
            mid_index = (start_index + end_index) // 2
            if input_list[mid_index] == search_element:
                return mid_index
            elif search_element < input_list[mid_index]:
                end_index = mid_index - 1
                return self.binary_search(input_list,start_index,end_index,search_element)
            elif search_element > input_list[mid_index]:
                start_index = mid_index + 1
                return self.binary_search(input_list,start_index,end_index,search_element)
        return -1
            
    def exponential_search(self,input_list,search_element):
        # If the seqarch element is present in the first or last index of the List.
        if search_element == input_list[0]:
            return 0
        elif search_element == input_list[len(input_list) - 1]:
            return len(input_list) - 1
        # Determine the range at which binary search should be done.
        range_index=1
        # Jumping in to exponential positions - now exponential positions are taken in the power of 2.
        while range_index < len(input_list) and input_list[range_index] < search_element:
            range_index = range_index * 2
            print("Range :: "+str(range_index))
        minimum_value = min(range_index,len(input_list) - 1)
        print("Searching Between --> Start Index :: "+str(range_index/2)+" End Index :: "+str(minimum_value))
        return self.binary_search(input_list,range_index/2,minimum_value,search_element)

    def print_result(self,result_index,search_element):
        if result_index is not None:
            print(str(search_element) + " is found in the list at index :: "+str(result_index))
        else:
            print(str(search_element) + " is not found in the list")
    
exponential_search_ins = ExponentialSearch()
search_element_cont = [18,23,100,10,47,42,12]
start_index = 0
end_index = len(input_contends_list) - 1
for search_element in search_element_cont:
    search_index = exponential_search_ins.exponential_search(input_contends_list,search_element)
    exponential_search_ins.print_result(search_index,search_element)
    print("-------------------------------------")