# stone
![Build Status](https://img.shields.io/teamcity/codebetter/bt428.svg)

A Scripting Language

## Specification 
### Types

```
comments		Comments are ignored by the interpreter
				// ...

int				The set of all signed 32-bit integers (-2147483648 to 2147483647)
				i = 123

string			A string type represents the set of string values
				s = "HelloWorld"

function 		A function type denotes the set of all functions with the same parameter and result types
				The return value is the result of the last statement
				
				add = func(a, b) { a + b } [ Anonymous function ]
				
				def max (a, b) {
                	if a > b {
                    	a
                	} else {
                    	b
                	}
				}

array			An array is a numbered sequence of elements of a single type
				a = [1, 2, 3, 4, 5]

object			See "OOP"
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
=				Initialize the variable
```

### Statements

```
if				if a > b {
  					max = a
				} else {
                  	max = b
				}

while			while a > b {
  					print "HelloWorld!"
				}
```

### OOP

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
				d.getSex()
				
				Output
				"Dog sex == boy"
```




## Build 
```

```

## Usage
```

```