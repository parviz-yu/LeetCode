# Time complexity: O(m*n)
# Space complexity: O(m)
class Solution:
    def numUniqueEmails(self, emails: list[str]) -> int:
        unique_email = set()
        for email in emails:
            splitted = email.split("@")
            splitted[0] = splitted[0].replace(".", "")
            idx = splitted[0].find("+")
            if idx != -1:
                splitted[0] = splitted[0][:idx]

            unique_email.add("@".join(splitted))

        return len(unique_email)
