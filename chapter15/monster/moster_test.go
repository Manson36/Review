package monster

import "testing"

//测试用例，测试Store方法
func TestStore(t *testing.T){

	//先创建一个Monster实例
	monster := Monster{
		Name:"红孩儿",
		Age:222,
		Skill:"三味真火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store错误，希望为%v，实际为%v",true,res)
	}
	t.Logf("monster.Store()测试成功")
}

func TestReStore(t *testing.T){

	//先创建一个Monster的实例
	var monster =&Monster{}
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore错误，希望为%v，实际为%v",true,res)
	}
	//进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore错误，希望为%v，实际为%v","红孩儿",monster.Name)
	}
	t.Logf("monster.ReStore()测试成功")
}