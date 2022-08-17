
input_list_contends = [1, 2, 4, 6, 3, 7, 8]
print(input_list_contends)

class MissingNumber:
        
    def findMissingByApproach1(self,input_list,nElement):
        missing_element = None
        if input_list is not None and len(input_list) > 0 and nElement is not None and nElement > 0:
            sum_element = sum(input_list)
            print("The Sum of N Numbers :: "+str(sum_element))
            summation = (nElement * (nElement+1)) / 2
            print("The sum of First N Elements :: "+str(summation))
            missing_element = summation - sum_element
        return missing_element

nthElement=8 
missing_number_instance = MissingNumber()
missing_element = missing_number_instance.findMissingByApproach1(input_list_contends,nthElement)
print("The Missing element is :: "+str(missing_element))