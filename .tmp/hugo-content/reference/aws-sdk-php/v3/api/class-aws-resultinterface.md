Menu

- [Aws](namespace-aws.md)

## ResultInterface    extends  ArrayAccess, IteratorAggregate, Countable   in    - [Aws](package-aws.md)

Represents an AWS result object that is returned from executing an operation.

### Table of Contents  [header link](class-aws-resultinterface-toc.md)

#### Methods  [header link](class-aws-resultinterface-toc-methods.md)

[\_\_toString()](class-aws-resultinterface-method-tostring.md)
: string Provides debug information about the result object[get()](class-aws-resultinterface-method-get.md)
: mixed\|null Get a specific key value from the result model.[hasKey()](class-aws-resultinterface-method-haskey.md)
: bool Check if the model contains a key by name[search()](class-aws-resultinterface-method-search.md)
: mixed Returns the result of executing a JMESPath expression on the contents
of the Result model.[toArray()](class-aws-resultinterface-method-toarray.md)
: array<string\|int, mixed> Convert the result to an array.

### Methods  [header link](class-aws-resultinterface-methods.md)

#### \_\_toString()  [header link](class-aws-resultinterface-method-tostring.md)

Provides debug information about the result object

`
    public
                    __toString() : string`

##### Return values

string

#### get()  [header link](class-aws-resultinterface-method-get.md)

Get a specific key value from the result model.

`
    public
                    get(string $key) : mixed|null`

##### Parameters

$key
: string

Key to retrieve.

##### Return values

mixed\|null
—

Value of the key or NULL if not found.

#### hasKey()  [header link](class-aws-resultinterface-method-haskey.md)

Check if the model contains a key by name

`
    public
                    hasKey(string $name) : bool`

##### Parameters

$name
: string

Name of the key to retrieve

##### Return values

bool

#### search()  [header link](class-aws-resultinterface-method-search.md)

Returns the result of executing a JMESPath expression on the contents
of the Result model.

`
    public
                    search(string $expression) : mixed`

$result = $client->execute($command);
$jpResult = $result->search('foo.\*.bar\[?baz > `10`\]');

##### Parameters

$expression
: string

JMESPath expression to execute

##### Tags  [header link](class-aws-resultinterface-method-search-tags.md)

link[JMESPath documentation](http://jmespath.readthedocs.org/en/latest)

##### Return values

mixed
—

Returns the result of the JMESPath expression.

#### toArray()  [header link](class-aws-resultinterface-method-toarray.md)

Convert the result to an array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-resultinterface-toc-constants.md)
  - [Methods](class-aws-resultinterface-toc-methods.md)
- Methods
  - [\_\_toString()](class-aws-resultinterface-method-tostring.md)
  - [get()](class-aws-resultinterface-method-get.md)
  - [hasKey()](class-aws-resultinterface-method-haskey.md)
  - [search()](class-aws-resultinterface-method-search.md)
  - [toArray()](class-aws-resultinterface-method-toarray.md)

[Back To Top](class-aws-resultinterface-top.md)
