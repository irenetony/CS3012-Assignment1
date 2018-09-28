import unittest
import Generate_BinaryTree
import Lowest_Common_Ancestor
class MyTest(unittest.TestCase):

    # random TestCase
    root = Generate_BinaryTree.Node(5)
    root.insertNode(2)
    root.insertNode(3)
    root.insertNode(4)
    root.insertNode(1)
    root.insertNode(6)
    root.insertNode(7)

    def randomTest(self):

        # root.Generate_BinaryTree.PrintTree()
        self.assertEqual(Lowest_Common_Ancestor.findLCA(root, 4, 5,),2)
        self.assertEqual(Lowest_Common_Ancestor.findLCA((root, 4, 6,),1)
        self.assertEqual(Lowest_Common_Ancestor.findLCA(root, 3, 4,),1)
        self.assertEqual(Lowest_Common_Ancestor.findLCA((root, 2, 4,),2)

    # def nullTest(self):
    #     # null TestCase
    #     self.assertEqual(findLCA(root,None, None))
if __name__ == '__main__':
    unittest.main()
