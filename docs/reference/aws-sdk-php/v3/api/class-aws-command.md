Menu

- [Aws](namespace-aws.md)

## Command        in package    - [Aws](package-aws.md)       implements  [CommandInterface](class-aws-commandinterface.md)  Uses  [HasDataTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html)

AWS command object.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#toc-interfaces)

[CommandInterface](class-aws-commandinterface.md)A command object encapsulates the input parameters used to control the
creation of a HTTP request and processing of a HTTP response.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#toc-methods)

[\_\_clone()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method___clone)
: mixed [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method___construct)
: mixed Accepts an associative array of command options, including:[count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_count)
: int [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method_get)
: mixed [getAuthSchemes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method_getAuthSchemes)
: mixed Get auth schemes added to command as required
for endpoint resolution[getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer the command.[getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_getIterator)
: Traversable[getName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method_getName)
: string Get the name of the command[hasParam()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method_hasParam)
: bool Check if the command has a parameter by name.[offsetExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetExists)
: bool [offsetGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetGet)
: mixed\|null This method returns a reference to the variable to allow for indirect
array modification (e.g., $foo\['bar'\]\['baz'\] = 'qux').[offsetSet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetSet)
: void [offsetUnset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetUnset)
: void [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_toArray)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#methods)

#### \_\_clone()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#method___clone)

`
    public
                    __clone() : mixed`

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#method___construct)

Accepts an associative array of command options, including:

`
    public
                    __construct(string $name[, array<string|int, mixed> $args = [] ][, HandlerList $list = null ][, MetricsBuilder|null $metricsBuilder = null ]) : mixed`

- @http: (array) Associative array of transfer options.

##### Parameters

$name
: string

Name of the command

$args
: array<string\|int, mixed>
= \[\]

Arguments to pass to the command

$list
: [HandlerList](class-aws-handlerlist.md)
= null

Handler list

$metricsBuilder
: MetricsBuilder\|null
= null

#### count()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_count)

`
    public
                    count() : int`

##### Return values

int

#### get()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#method_get)

`
    public
                    get(mixed $name) : mixed`

##### Parameters

$name
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#method_get\#tags)

deprecated

#### getAuthSchemes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#method_getAuthSchemes)

Get auth schemes added to command as required
for endpoint resolution

`
    public
                    getAuthSchemes() : mixed`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#method_getAuthSchemes\#tags)

returns

array

deprecated

In favor of using the @context property bag.
Auth schemes are now accessible via the `signature_version` key
in a Command's context, if applicable.

#### getHandlerList()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#method_getHandlerList)

Get the handler list used to transfer the command.

`
    public
                    getHandlerList() : HandlerList`

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_getIterator)

`
    public
                    getIterator() : Traversable`

##### Return values

Traversable

#### getName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#method_getName)

Get the name of the command

`
    public
                    getName() : string`

##### Return values

string

#### hasParam()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html\#method_hasParam)

Check if the command has a parameter by name.

`
    public
                    hasParam(mixed $name) : bool`

##### Parameters

$name
: mixed

Name of the parameter to check

##### Return values

bool

#### offsetExists()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_offsetExists)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_offsetGet)

This method returns a reference to the variable to allow for indirect
array modification (e.g., $foo\['bar'\]\['baz'\] = 'qux').

`
    public
                &    offsetGet( $offset) : mixed|null`

##### Parameters

$offset
:

##### Return values

mixed\|null

#### offsetSet()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_offsetSet)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

#### offsetUnset()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_offsetUnset)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html\#method_toArray)

`
    public
                    toArray() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#toc-methods)
- Methods
  - [\_\_clone()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method___clone)
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method___construct)
  - [count()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_count)
  - [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method_get)
  - [getAuthSchemes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method_getAuthSchemes)
  - [getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method_getHandlerList)
  - [getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_getIterator)
  - [getName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method_getName)
  - [hasParam()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#method_hasParam)
  - [offsetExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetExists)
  - [offsetGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetGet)
  - [offsetSet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetSet)
  - [offsetUnset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_offsetUnset)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HasDataTrait.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Command.html#top)
