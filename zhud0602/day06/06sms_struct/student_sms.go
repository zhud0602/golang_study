package main

import "fmt"

type student struct {
	id   int64
	name string
}

type sms struct {
	allStudent map[int64]student
}

//查看所有学生
func (s sms) showStudent() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", stu.id, stu.name)
	}
}

//新增学生
func (s sms) addStudent() {
	//1.根据用户输入的内容创建一个新的学生
	var (
		stuid   int64
		stuName string
	)

	fmt.Print("请输入学号:")
	fmt.Scanln(&stuid)
	fmt.Print("请输入姓名:")
	fmt.Scanln(&stuName)
	//2.把新的学生放到s.allStudent这个map中
	newStu := student{
		id:   stuid,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功！")

}

//删除学生
func (s sms) deleteStudent() {
	//1.请用户输入要删除的学生ID
	fmt.Print("请输入要删除的学生的学号：")
	var stuID int64
	fmt.Scan(&stuID)
	//2.去map中找有没有这个学生，如果没有打印个查无此人
	stuobj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	//3.有的话就删除
	fmt.Printf("你要删除的学生信息如下, 学号:%d 姓名:%s\n", stuobj.id, stuobj.name)
	var str string
	fmt.Print("是否删除? 请输入：（Y/N）")
	fmt.Scan(&str)
	if str == "N" {
		return
	}
	delete(s.allStudent, stuID)
	fmt.Println("删除成功！")
}

//修改学生
func (s sms) alterStudent() {
	//1.获取用户输入的学号
	var stuID int64
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	//2.展示该学号对应的学生信息，如果没有提示查无此人
	stuobj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息如下, 学号:%d 姓名:%s\n", stuobj.id, stuobj.name)
	var newname string
	//3.请输入修改后的学生名
	fmt.Print("请输入学生的新名字：")
	fmt.Scan(&newname)
	//4.更新学生的姓名
	stuobj.name = newname
	s.allStudent[stuID] = stuobj //更新map中的学生
	fmt.Println("修改成功")
}
