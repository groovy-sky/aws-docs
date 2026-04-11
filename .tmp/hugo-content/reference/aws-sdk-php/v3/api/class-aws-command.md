Menu

- [Aws](namespace-aws.md)

## Command        in package    - [Aws](package-aws.md)       implements  [CommandInterface](class-aws-commandinterface.md)  Uses  [HasDataTrait](class-aws-hasdatatrait.md)

AWS command object.

### Table of Contents  [header link](class-aws-command-toc.md)

#### Interfaces  [header link](class-aws-command-toc-interfaces.md)

[CommandInterface](class-aws-commandinterface.md)A command object encapsulates the input parameters used to control the
creation of a HTTP request and processing of a HTTP response.

#### Methods  [header link](class-aws-command-toc-methods.md)

[\_\_clone()](class-aws-command-method-clone.md)
: mixed [\_\_construct()](class-aws-command-method-construct.md)
: mixed Accepts an associative array of command options, including:[count()](class-aws-hasdatatrait-method-count.md)
: int [get()](class-aws-command-method-get.md)
: mixed [getAuthSchemes()](class-aws-command-method-getauthschemes.md)
: mixed Get auth schemes added to command as required
for endpoint resolution[getHandlerList()](class-aws-command-method-gethandlerlist.md)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer the command.[getIterator()](class-aws-hasdatatrait-method-getiterator.md)
: Traversable[getName()](class-aws-command-method-getname.md)
: string Get the name of the command[hasParam()](class-aws-command-method-hasparam.md)
: bool Check if the command has a parameter by name.[offsetExists()](class-aws-hasdatatrait-method-offsetexists.md)
: bool [offsetGet()](class-aws-hasdatatrait-method-offsetget.md)
: mixed\|null This method returns a reference to the variable to allow for indirect
array modification (e.g., $foo\['bar'\]\['baz'\] = 'qux').[offsetSet()](class-aws-hasdatatrait-method-offsetset.md)
: void [offsetUnset()](class-aws-hasdatatrait-method-offsetunset.md)
: void [toArray()](class-aws-hasdatatrait-method-toarray.md)
: mixed

### Methods  [header link](class-aws-command-methods.md)

#### \_\_clone()  [header link](class-aws-command-method-clone.md)

`
    public
                    __clone() : mixed`

#### \_\_construct()  [header link](class-aws-command-method-construct.md)

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

#### count()  [header link](class-aws-hasdatatrait-method-count.md)

`
    public
                    count() : int`

##### Return values

int

#### get()  [header link](class-aws-command-method-get.md)

`
    public
                    get(mixed $name) : mixed`

##### Parameters

$name
: mixed

##### Tags  [header link](class-aws-command-method-get-tags.md)

deprecated

#### getAuthSchemes()  [header link](class-aws-command-method-getauthschemes.md)

Get auth schemes added to command as required
for endpoint resolution

`
    public
                    getAuthSchemes() : mixed`

##### Tags  [header link](class-aws-command-method-getauthschemes-tags.md)

returns

array

deprecated

In favor of using the @context property bag.
Auth schemes are now accessible via the `signature_version` key
in a Command's context, if applicable.

#### getHandlerList()  [header link](class-aws-command-method-gethandlerlist.md)

Get the handler list used to transfer the command.

`
    public
                    getHandlerList() : HandlerList`

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-hasdatatrait-method-getiterator.md)

`
    public
                    getIterator() : Traversable`

##### Return values

Traversable

#### getName()  [header link](class-aws-command-method-getname.md)

Get the name of the command

`
    public
                    getName() : string`

##### Return values

string

#### hasParam()  [header link](class-aws-command-method-hasparam.md)

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

#### offsetExists()  [header link](class-aws-hasdatatrait-method-offsetexists.md)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](class-aws-hasdatatrait-method-offsetget.md)

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

#### offsetSet()  [header link](class-aws-hasdatatrait-method-offsetset.md)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

#### offsetUnset()  [header link](class-aws-hasdatatrait-method-offsetunset.md)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

#### toArray()  [header link](class-aws-hasdatatrait-method-toarray.md)

`
    public
                    toArray() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-command-toc-methods.md)
- Methods
  - [\_\_clone()](class-aws-command-method-clone.md)
  - [\_\_construct()](class-aws-command-method-construct.md)
  - [count()](class-aws-hasdatatrait-method-count.md)
  - [get()](class-aws-command-method-get.md)
  - [getAuthSchemes()](class-aws-command-method-getauthschemes.md)
  - [getHandlerList()](class-aws-command-method-gethandlerlist.md)
  - [getIterator()](class-aws-hasdatatrait-method-getiterator.md)
  - [getName()](class-aws-command-method-getname.md)
  - [hasParam()](class-aws-command-method-hasparam.md)
  - [offsetExists()](class-aws-hasdatatrait-method-offsetexists.md)
  - [offsetGet()](class-aws-hasdatatrait-method-offsetget.md)
  - [offsetSet()](class-aws-hasdatatrait-method-offsetset.md)
  - [offsetUnset()](class-aws-hasdatatrait-method-offsetunset.md)
  - [toArray()](class-aws-hasdatatrait-method-toarray.md)

[Back To Top](class-aws-command-top.md)
