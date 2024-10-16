# Time complexity: O(n)
# Space complexity: O(n)
class Solution:
    def evalRPN(self, tokens: list[str]) -> int:
        stack = []
        
        for elem in tokens:
            if elem not in "+-*/":
                stack.append(int(elem))
            else:
                operand_2 = stack.pop()
                operand_1 = stack.pop()
                match elem:
                    case "+":
                        stack.append(operand_1 + operand_2)
                    case "-":
                        stack.append(operand_1 - operand_2)
                    case "*":
                        stack.append(operand_1 * operand_2)
                    case "/":
                        stack.append(int(operand_1 / operand_2))

        return stack[-1]