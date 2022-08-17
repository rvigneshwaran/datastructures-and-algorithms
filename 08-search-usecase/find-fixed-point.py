input_list_contends = [10, 2, 4, 61, 13, 5, 81]
print(input_list_contends)

class FindFixedPoint:
    
    def findFixedPointbyApp1(self):
        fixed_point_list = []
        if(input_list_contends is not None and len(input_list_contends) > 0):
            for index,element in enumerate(input_list_contends):
                if index == element:
                    print("Found Fixed Element in index :: "+str(index))
                    fixed_point_list.append(element)
        return fixed_point_list

    def print_execution_result(self,fixed_point_list):
        if fixed_point_list is not None and len(fixed_point_list) > 0:
            print("Elements of the Fixed Point in the list are :: \n "+str(fixed_point_list))
        else:
            print("No Fixed Points in the list")
    
find_fixedpoint_ins = FindFixedPoint()
fixed_point_list = find_fixedpoint_ins.findFixedPointbyApp1()
find_fixedpoint_ins.print_execution_result(fixed_point_list)