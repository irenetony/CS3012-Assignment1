package lca

func CreateATree() *Node 
{
	return &Node
	{
		Left: &Node
		{
			Right: &Node
			{
				Key: 3,
				Value: 3,
			}
			Key: 1,
			Value: 1,
		}
		Right: &Node
		{
			Left: &Node
			{
				Key: 5,
				Value: 5,
			}
			Right: &Node
			{
				Key: 7,
				Value: 7,
			}

			Key: 6
			Value: 6,
		},

		Key: 4,
		Value: 4,
	}
}
func EmptyTreeTest(*testing.T)
{
	x := BinaryTree()
	expected := Node{}
	if x != expected {t.Errorf("Your program does not create a binary tree")}
}
func insertTest(*testing.T)
{
	x := CreateATree()
	x.add(4,4)
	expected := Node{Key:4, Value:4}
	if x.Value != expected.Value{t.Errorf("Your program does not insert a node correctly")}

	x.add(1,1)
	x.add(3,3)
	expected := Node
	{
		Left:  &Node
		{
			Right: &Node
			{
				Key: 3,
				Value: 3,
			},

			Key: 1, 
			Value: 1,
		},

		Key: 4,
		Value: 4,
	}
	if x.Value != expected.Value || x.Left.Value != expected.Left.Value || x.Left.Right.Value != expected.Left.Right.Value{
		t.Errorf("Your program does not insert many nodes correctly")
	}

}
func getTest(*testing.T)
{
	x := Node
	{
		Left:  &Node
		{
			Right: &Node
			{
				Key: 3,
				Value: 3,
			},

			Key: 1, 
			Value: 1,
		},

		Key: 4,
		Value: 4,
	}
	expected := 1
	value := x.get(1)

	if value.Value == expected{
		t.Errorf("Your program does not return the searched node correctly")
	}
}
func lcaTest(*testing.T)
{
	x := CreateATree()
	expected := 6
	ans := x.LCA(5,7)

	if ans != expected{t.Errorf("Your program does not return the lowest common ancestor node correctly")}

}