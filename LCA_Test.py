import unittest


class MyTest(unittest.TestCase):


    def randomTest(self):
        # random TestCase
        root = Node(1)
        root.insertNode(2)
        root.insertNode(3)
        root.insertNode(4)
        root.insertNode(5)
        root.insertNode(6)
        root.insertNode(7)

        self.assertEqual(findLCA(root, 4, 5,),2)
        self.assertEqual(findLCA(root, 4, 6,),1)
        self.assertEqual(findLCA(root, 3, 4,),1)
        self.assertEqual(findLCA(root, 2, 4,),2)

    def nullTest(self):
        # null TestCase
        self.assertEqual(findLCA(root,None, None))
