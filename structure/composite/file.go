package composite

import "fmt"

//叶子节点，树的基本机构，不包含子项目，一般完成大部分实际工作

type File struct {
	name string
}

func (f *File) Search(keyword string) {
	fmt.Printf("searching for keyword %s in file %s \n", keyword, f.name)
}

func (f *File) GetName() string {
	return f.name
}

// 容器，包含叶子节点或其他容器等子项目的单位。容器不知道其子项目所属的具体类，它只通过通用的组件接口与其子项目交互。容器接收到请求后会将工作分配给自己的子项目，它可以处理中间
// 结果，然后将最终结果返回给客户端

type Folder struct {
	Component []Component
	name      string
}

//组件接口，描述了组合树中简单项目和复杂项目共有的操作

type Component interface {
	Search(string)
}

func (f *Folder) Search(keyword string) {

	fmt.Printf("searching recursively for keyword %s in folder %s \n", keyword, f.name)
	for _, composite := range f.Component {
		composite.Search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.Component = append(f.Component, c)
}

//客户端通过组件接口与所有项目交互，因此，客户端能以相同的方式与树状结构中简单活复杂项目交互

func ApplicationRun() {
	file1 := &File{"file1"}
	file2 := &File{"file2"}
	file3 := &File{"file3"}

	folder1 := &Folder{
		name: "folder1",
	}

	folder1.add(file1)
	folder2 := &Folder{
		name: "folder2",
	}

	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.Search("rose")
}
