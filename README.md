# stone

A Scripting Language based on Golang

## Specification 
### Types

```
comments		Comments are ignored by the interpreter
				// ...

int				The set of all signed 32-bit integers (-2147483648 to 2147483647)
				i = 123

string				A string type represents the set of string values
				s = "HelloWorld"

function 			A function type denotes the set of all functions with the same parameter and result types
				The return value is the result of the last statement
				
				add = func(a, b) { a + b } [ Anonymous function ]
				
				def max (a, b) {
                			if a > b {
                    				a
                			} else {
                    				b
                			}
				}

array				An array is a numbered sequence of elements of a single type
				a = [1, 2, 3, 4, 5]

object				See "Object Oriented Programming"
```

### Operators

```
+				A string can be added to an integer and returned as a string
				"HelloWorld" + 123 == "HelloWorld123"
				
-				Only integer types are supported
*				Only integer types are supported
/				Only integer types are supported
>				Only integer types are supported
<				Only integer types are supported

==				Returns true if the variable name points to the same object
                		If operands are the basic data type, it depends on whether the values are equal or not

=				Initialize the variable
```

### Statements

```
if				if a > b {
  					max = a
				} else {
                  			max = b
				}

while				while a > b {
  					print "HelloWorld!"
				}
```

### Clousure

```
def add(c) {
	func(x) {c + x}
}

c1 = add(1)
print(c1(2))
print("\n")
print(c1(1))

Output:
3
4
```

### Parameter Passing

```
The base data type is passed as a value
Arrays and objects are passed as references
```

### Object Oriented Programming

```
class			def Animal {
				age = 0
				def getAge {
                    			"Animal age = " + age
				}
			}
				
			def Dog extends Animal {
                		sex = "girl"
                		def getAge {
                    			"Dog age == " + age
                		}
                		def getSex {
                    			"Dog sex == " + sex
                		}
			}
				
			d = Dog.new
			d.sex = "boy"
			print(d.getAge())
			print("\n")
			print(d.getSex())

			Output:
			"Dog age == 0"
			"Dog sex == boy"
```

### Built-in functions

```
read() string				Read a word from terminal
print(any[any type]) int		Show "any" in terminal
size(s[string])	int			Return length of string
atoi(s[string]) int			Convert string to int
itoa(i[int]) string			Convert int to string
timestamp() int				Return UNIX timpstamp
```

## Requirement

* Go

## Build 

```
cp -r stone $GOPATH/src
cd $GOPATH/src/stone
go build
```

## Usage
```
./stone [source file name]
```

## References
- 两周自制脚本语言
- The Go Programming Language
