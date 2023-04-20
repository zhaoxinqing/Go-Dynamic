# README

## Git

> git branch -m master master  
> git fetch origin  
> git branch -u origin/master master  
> git remote set-head origin -a

## 日志

有错误，立刻打印，然后回抛

---

## MySQL & ElasticSearch

- MySQL：擅长事务类型操作，可以确保数据的安全和一致性；
- ElasticSearch：擅长海量数据的搜索、分析、计算；

```json
// 查询
GET collectibles_period_per_minutes_1440/_search
{
  "query": {
    "match_all": {}
  },
  "sort": [
    {
      "created_at": {
        "order": "desc"
      }
    }
  ],
  "size":30
}
```

## 工作

- 事事有汇报


## Go 面向对象编程：方法、函数

在 go 语言中，结构体就像是类的一种简化形式。go 的方法是作用在接收者（receiver）（接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针）上的一个函数，接收者是某种类型的变量，因此方法是一种特殊类型的变量。  
一个类型加上它的方法等价于面向对象的一个类。（类方法）类型和作用在它上面定义的方法。

```go
type Person struct {
    Name string
    Age  int
}

func NewPerson(name string, age int) *Person {
    return &Person{
        Name: name,
        Age:  age,
    }
}

func (p *Person) AddAge() {
    p.Age += 1
}
 
func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s, I am %d years old", p.Name, p.Age)
}


func main() {
    person := NewPerson("Tom", 18)
    person.AddAge()
    person.SayHello()
}

// 通过指针定义了AddAge方法，然后定义了一个NewPerson方法，该方法返回一个指向Person对象的指针。在main函数中，我们通过NewPerson创建了一个Person对象，然后调用AddAge方法增加了Person的年龄，并输出了Person的信息。
```

## 重载，重载的作用

重载，就是函数或者方法有相同的名称，但是参数列表不相同的情形，这样的同名不同参数的函数或者方法之间，互相称之为重载函数或者重载方法。通俗的理解就是省了给method重新命名了，差不多的都用一个。重载的最直接作用是方便了程序员可以根据不同的参数个数，顺序，类型，自动匹配方法，减少写多个函数名或方法名的重复步骤。

- 重载例子：

```java
public class OverloadDemo {
 
    //1. test()方法第一次重载，没有参数
    void test() {
        System.out.println("No parameters");
    }
 
    //2. test()方法第二次重载，一个整型参数
    void test(int a) {
        System.out.println("a: " + a);
    }
 
    //3. test()方法第三次重载，两个整型参数
    void test(int a, int b) {
        System.out.println("a and b: " + a + " " + b);
    }
 
    //4. test()方法第四次重载，一个双精度型参数
    double test(double a) {
        System.out.println("double a: " + a);
        return a * a;//返回a*a的值
    }
}
```

- 输出结果：

```java
public class Overload {
   public static void main(String args[]){
        OverloadDemo ob=new OverloadDemo();
        double result;
        ob.test();            //定义的是test()方法
        ob.test(10);             //定义的是test(int a)方法
        ob.test(10,20);            //定义的是test(int a,int b)方法
        result=ob.test(123.23);       //定义的是test(double a)方法
        System.out.println("result of ob.test(123.23): "+result);   //输出result的值
   }
}
 
// 运行结果：    No parameters
// a: 10
// a and b: 10 20
// double a: 123.23
// result of ob.test(123.23): 15185.6329
```

## 深拷贝、浅拷贝

在有指针的情况下，浅拷贝只是增加了一个指针指向已经存在的内存，而深拷贝就是增加一个指针并且申请一个新的内存，使这个增加的指针指向这个新的内存，采用深拷贝的情况下，释放内存的时候就不会出现在浅拷贝时重复释放同一内存的错误。

## 格式化打印 - TestPrintf

```go
func TestPrintf() {
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "walala", Age: 18}
	fmt.Printf("person:%v\n", person)  // person:{walala 18}
	fmt.Printf("person:%+v\n", person) // person:{Name:walala Age:18}

	personPit := &Person{Name: "walala02", Age: 30}
	fmt.Printf("personPit:%v\n", personPit)  // personPit:&{walala02 30}
	fmt.Printf("personPit:%+v\n", personPit) // personPit:&{Name:walala02 Age:30}

	fmt.Printf("personPit:%v\n", &personPit)  // personPit:0xc00014ab58 (& 拿到的是指针)
	fmt.Printf("personPit:%+v\n", &personPit) // personPit:0xc00014ab58

	fmt.Printf("personPit:%v\n", *personPit)  // personPit:{walala02 30} (* 指针指向的值)
	fmt.Printf("personPit:%+v\n", *personPit) // personPit:{Name:walala02 Age:30}
}
```
