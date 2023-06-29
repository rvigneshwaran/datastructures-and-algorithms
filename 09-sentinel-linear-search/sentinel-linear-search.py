import sys

input_list_sample = [10, 12, 13, 16, 18, 19, 20, 21, 22, 23, 24, 33, 35, 42, 47]
target_value = 22

"""
    The sentinel_linear_search function takes an input_list_sample arr and a target value target as input. 
    It sets the last element of the input_list_sample as the sentinel value temporarily, allowing us to 
    eliminate the need for an explicit boundary check within the loop.

    The function then iterates through the input_list_sample using a while loop, checking each element 
    until the target value is found. If the target value is found, it returns the index of 
    the element. Finally, the program checks the return value and prints the appropriate 
    message based on whether the target value was found or not.  
"""

class SentinelLinearSearch:
    
    def sentinel_linear_search(self,input_list_sample, target):
        # Set the last element as the sentinel
        last_element = input_list_sample[-1]
        input_list_sample[-1] = target

        index = 0
        while input_list_sample[index] != target:
            index += 1

        # Restore the original last element
        input_list_sample[-1] = last_element

        if index < len(input_list_sample) - 1 or input_list_sample[-1] == target:
            return index
        else:
            return -1


# Instace of the class Sentinel Search Instance
sentinel_search_ins = SentinelLinearSearch()

result = sentinel_search_ins.sentinel_linear_search(input_list_sample, target_value)

if result != -1:
    print(f"Element {target_value} found at index {result}")
else:
    print(f"Element {target_value} not found in the array.")
