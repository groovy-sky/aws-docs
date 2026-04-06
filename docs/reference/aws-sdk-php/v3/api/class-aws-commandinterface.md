Menu

- [Aws](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.html)

## CommandInterface    extends  ArrayAccess, Countable, IteratorAggregate   in    - [Aws](https://docs.aws.amazon.com/aws-sdk-php/v3/api/package-Aws.html)

A command object encapsulates the input parameters used to control the
creation of a HTTP request and processing of a HTTP response.

Using the toArray() method will return the input parameters of the command
as an associative array.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html\#toc-methods)

[getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html#method_getHandlerList)
: [HandlerList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html)Get the handler list used to transfer the command.[getName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html#method_getName)
: string Get the name of the command[hasParam()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html#method_hasParam)
: bool Check if the command has a parameter by name.[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html#method_toArray)
: array<string\|int, mixed> Converts the command parameters to an array

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html\#methods)

#### getHandlerList()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html\#method_getHandlerList)

Get the handler list used to transfer the command.

`
    public
                    getHandlerList() : HandlerList`

##### Return values

[HandlerList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html)

#### getName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html\#method_getName)

Get the name of the command

`
    public
                    getName() : string`

##### Return values

string

#### hasParam()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html\#method_hasParam)

Check if the command has a parameter by name.

`
    public
                    hasParam(string $name) : bool`

##### Parameters

$name
: string

Name of the parameter to check

##### Return values

bool

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html\#method_toArray)

Converts the command parameters to an array

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html#toc-methods)
- Methods
  - [getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html#method_getHandlerList)
  - [getName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html#method_getName)
  - [hasParam()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html#method_hasParam)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html#top)
