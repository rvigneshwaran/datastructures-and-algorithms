
input_contends_list  = [10, 12, 13, 16, 18, 19, 20, 21, 22, 23, 24, 33, 35, 42, 47]
print(input_contends_list)

class InterpolationSearch:
    
    def __init__(self):
        print("Initializing Components for Search : ")
        
    def interpolation_search(self,input_list,start_index,end_index,search_element):
        found_index = None
        if start_index < end_index and search_element >= input_list[start_index] and search_element <= input_list[end_index]:
            print("start_index :: "+str(start_index)+" end_index :: "+str(end_index))
            position = start_index + ((end_index - start_index) // (input_list[end_index] - input_list[start_index]) * (search_element - input_list[start_index]))
            print("The Value of the derived position is :: "+str(position))
            if input_list[position] == search_element:
                return position
            elif search_element < input_list[position]:
                end_index = position - 1
                print("start_index :: "+str(start_index)+" end_index :: "+str(end_index))
                return self.interpolation_search(input_list,start_index,end_index,search_element)
            elif search_element > input_list[position]:
                start_index = position + 1
                print("start_index :: "+str(start_index)+" end_index :: "+str(end_index))
                return self.interpolation_search(input_list,start_index,end_index,search_element)
        return found_index

    def print_result(self,result_index,search_element):
        if result_index is not None:
            print(str(search_element) + " is found in the list at index :: "+str(result_index))
        else:
            print(str(search_element) + " is not found in the list")
    
interpolation_search_ins = InterpolationSearch()
search_element_cont = [18,23,100,10,47]
start_index = 0
end_index = len(input_contends_list) - 1
for search_element in search_element_cont:
    search_index = interpolation_search_ins.interpolation_search(input_contends_list,start_index,end_index,search_element)
    interpolation_search_ins.print_result(search_index,search_element)
    print("-------------------------------------------------------------------------------------------------------------------------")