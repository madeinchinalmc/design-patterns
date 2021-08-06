package flyweight

import "fmt"

// 享元类包含一个树的部分状态。这些成员变量保存的数值对于特定树而言是唯一
//的。例如，你在这里找不到树的坐标。但这里有很多树木之间所共有的纹理和颜
//色。由于这些数据的体积通常非常大，所以如果让每棵树都其进行保存的话将耗
//费大量内存。因此，我们可将纹理、颜色和其他重复数据导出到一个单独的对象
//中，然后让众多的单个树对象去引用它。

type TreeType struct {
	Name string
	Color string
	Texture string
}

func NewTreeType(name,color,texture string)*TreeType{
	return &TreeType{
		name,
		color,
		texture,
	}
}

func (t *TreeType)Draw(x,y int){
	fmt.Printf("name:%v,color:%v,texture:%v, x:%v y:%v \n",t.Name,t.Color,t.Texture,x,y)
}

// 享元工厂决定是否复用已有享元或者创建一个新的对象。

type TreeFactory struct {
	TreeTypes []TreeType
}

func (tf *TreeFactory) GetTreeType(name,color,texture string)TreeType{
	for _, t := range tf.TreeTypes{
		if t.Color == color && t.Name == name && t.Texture == texture{
			return t
		}
	}
	t := TreeType{
		name,color,texture,
	}
	fmt.Printf("new tree type is name:%v,color:%v,texture:%v \n",t.Name,t.Color,t.Texture)
	tf.TreeTypes = append(tf.TreeTypes,t)
	return t
}
// 情景对象包含树状态的外在部分。程序中可以创建数十亿个此类对象，因为它们 // 体积很小:仅有两个整型坐标和一个引用成员变量。

type Tree struct {
	X int
	Y int
	TreeType TreeType
}

func NewTree(x,y int , treeType TreeType)Tree{
	return Tree{
		x,y,treeType,
	}
}

func (t *Tree)Draw(){
	t.TreeType.Draw(t.X,t.Y)
}

type Forest struct {
	Trees []Tree
}

func (f *Forest) PlantTree(x,y int,name,color,texture string,treeFactory *TreeFactory){
	treeType := treeFactory.GetTreeType(name,color,texture)
	tree := NewTree(x,y,treeType)
	f.Trees = append(f.Trees,tree)
}

func  (f *Forest)Draw() {
	for _,tree := range f.Trees{
		tree.Draw()
	}
}