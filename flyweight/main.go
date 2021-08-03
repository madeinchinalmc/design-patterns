package main

import . "./structured"

func main(){
	treeTypeMaple := NewTreeType("枫树","红叶","无纹理")
	treeTypePolar := NewTreeType("杨树","黄叶","细微纹理")

	treeFactory := &TreeFactory{
	}
	treeFactory.TreeTypes = append(treeFactory.TreeTypes,*treeTypeMaple)
	treeFactory.TreeTypes = append(treeFactory.TreeTypes,*treeTypePolar)
	antForest := &Forest{}
	antForest.Draw()
	antForest.PlantTree(1,1,"枫树","红叶","无纹理",treeFactory)
	antForest.PlantTree(1,2,"枫树","红叶","无纹理",treeFactory)
	antForest.PlantTree(1,1,"杨树","黄叶","细微纹理",treeFactory)

	antForest.PlantTree(1,1,"松树","绿叶","大纹理",treeFactory)

	antForest.Draw()
}
