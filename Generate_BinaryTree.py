# Create a binary Tree

class Node:
    def __init__(self, data):
        self.left = None
        self.right = None
        self.data = data

    def insertNode(self, data):
        if self.data:
            if data < self.data:
                if self.left is None:
                    self.left = Node(data)
                else:
                    self.left.insertNode(data)
            elif data > self.data:
                if self.right is None:
                    self.right = Node(data)
                else:
                    self.right.insertNode(data)
        else:
            self.data = data
