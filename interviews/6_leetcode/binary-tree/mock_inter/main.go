//type Node struct {
//    Val string
//    Children1 *Node
//    Children2 *Node
//}



e = {
	value: "a",
	children: [
	  {
		value: "b",
		children: [
		  {
			value: "d",
			children: [
			  {
				value: "h",
				children: [
				  {
					value: "i",
					children: [],
				  },
				],
			  },
			],
		  },
		  {
			value: "e",
			children: [],
		  },
		],
	  },
	  {
		value: "c",
		children: [
		  {
			value: "f",
			children: [],
		  },
		  {
			value: "g",
			children: [],
		  },
		],
	  },
	],
  };
  
  type Node struct {
	  value string
	  children *[]Node
  }
  
  //func preOrder(tree Node, res *string[]) {
  //    *res = append(*res, tree.value)
  //    for i := range tree.children {
  //        preOrder(tree.children[i], &res)
  //    }
  //}
  
  //func flatten(tree Node) string[] {
	//  res := make(string[], 0)
  //    preOrder(tree, &res)
  //}
  type Node struct {
	  value string
	  children *[]Node
  }
  
  func preOrder(tree Node, res *string[]) {
	  *res = append(*res, tree.value)
	  for i := range tree.children {
		  preOrder(tree.children[i])
	  }
  }
   Node
	  res := make(string[], 0)
	  preOrder()
  
  
  func flatten(tree Node) string[] {
	  res := make(string[], 0)
	  // стек или очередь
	  queue := Queue{}
	  res = append(res, tree.value)
	  queue.Push(tree)
	  for queue.Size() != 0 {
		  root = queue.Pop()
		  for _, child := range root.children {
			  queue.Push(child)
		  }
		  
	  }
	  
	  return res
  }
  
  
  
  