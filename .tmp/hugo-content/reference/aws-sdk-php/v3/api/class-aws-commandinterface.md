Menu

- [Aws](namespace-aws.md)

## CommandInterface    extends  ArrayAccess, Countable, IteratorAggregate   in    - [Aws](package-aws.md)

A command object encapsulates the input parameters used to control the
creation of a HTTP request and processing of a HTTP response.

Using the toArray() method will return the input parameters of the command
as an associative array.

### Table of Contents  [header link](class-aws-commandinterface-toc.md)

#### Methods  [header link](class-aws-commandinterface-toc-methods.md)

[getHandlerList()](class-aws-commandinterface-method-gethandlerlist.md)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer the command.[getName()](class-aws-commandinterface-method-getname.md)
: string Get the name of the command[hasParam()](class-aws-commandinterface-method-hasparam.md)
: bool Check if the command has a parameter by name.[toArray()](class-aws-commandinterface-method-toarray.md)
: array<string\|int, mixed> Converts the command parameters to an array

### Methods  [header link](class-aws-commandinterface-methods.md)

#### getHandlerList()  [header link](class-aws-commandinterface-method-gethandlerlist.md)

Get the handler list used to transfer the command.

`
    public
                    getHandlerList() : HandlerList`

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getName()  [header link](class-aws-commandinterface-method-getname.md)

Get the name of the command

`
    public
                    getName() : string`

##### Return values

string

#### hasParam()  [header link](class-aws-commandinterface-method-hasparam.md)

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

#### toArray()  [header link](class-aws-commandinterface-method-toarray.md)

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
  - [Constants](class-aws-commandinterface-toc-constants.md)
  - [Methods](class-aws-commandinterface-toc-methods.md)
- Methods
  - [getHandlerList()](class-aws-commandinterface-method-gethandlerlist.md)
  - [getName()](class-aws-commandinterface-method-getname.md)
  - [hasParam()](class-aws-commandinterface-method-hasparam.md)
  - [toArray()](class-aws-commandinterface-method-toarray.md)

[Back To Top](class-aws-commandinterface-top.md)
