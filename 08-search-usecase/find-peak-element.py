import math

input_list_contends = [10, 2, 4, 61, 13, 5, 81]
print(input_list_contends)

class FindPeakElement:
    
    def findPeakElementbyApp1(self):
        peak_element_list = []
        if(input_list_contends is not None and len(input_list_contends) > 0):
            for index,element in enumerate(input_list_contends):
                if index != 0 and index != len(input_list_contends) - 1:
                    if input_list_contends[index] > input_list_contends[index - 1] and input_list_contends[index] > input_list_contends[index + 1]:
                        peak_element_list.append(element)
                elif index == 0 and input_list_contends[index] > input_list_contends[index + 1]:
                    peak_element_list.append(element)
                elif index == len(input_list_contends) - 1 and input_list_contends[index] > input_list_contends[index - 1]:
                    peak_element_list.append(element)
        return peak_element_list
    

    def print_execution_result(self,fixed_peak_element,approach_identifier):
        if fixed_peak_element is not None and len(fixed_peak_element) > 0:
            print("Peak Element is found using "+str(approach_identifier)+" approach :: \n "+str(fixed_peak_element))
        else:
            print("No Peak Elementin the list")

# find the element by using Linear Search    
find_peakelement_ins = FindPeakElement()
peak_element_list = find_peakelement_ins.findPeakElementbyApp1()
find_peakelement_ins.print_execution_result(peak_element_list,"Linear Search")