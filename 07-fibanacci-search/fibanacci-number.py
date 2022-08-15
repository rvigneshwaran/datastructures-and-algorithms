import sys

class FibanacciNumber:
    
    def calculateNFibanocciNumber(self,input_number):
        """ The below function willhelp to determine the value of the nth Fibanacci number,
        the nth fibanacci number is determined by using the recursive function.
        reducing by constant values and determining the sum of values
        follows the convention f(n) = f(n-1) + f(n-2)
        
        Args:
            input_number (number): Number for which the fibanacci is calculated

        Returns:
            _type_: nth fibanacci number
        """
        if input_number >= 0:
            if input_number == 0:
                return 0
            elif input_number == 1:
                return 1
            else:
                return (self.calculateNFibanocciNumber(input_number - 1) + self.calculateNFibanocciNumber(input_number - 2))

fibanocci_instance = FibanacciNumber()
# Getting total no of arguements from command line
args_length = len(sys.argv)

# Check if the command line has valid arguments
if args_length == 2:
    nthNumber = int(sys.argv[1])
    if nthNumber > 0:
        nthNumberResult = fibanocci_instance.calculateNFibanocciNumber(nthNumber)
        print("The "+str(nthNumber)+"th Fibanacci number is :: "+str(nthNumberResult))
    else:
        print("Please enter valid positive Integer")
else:
    print("Please enter valid arguments for deriving the nth Fibanacci Number")