Menu

- [Aws](namespace-aws.md)

## ResultInterface    extends  ArrayAccess, IteratorAggregate, Countable   in    - [Aws](package-aws.md)

Represents an AWS result object that is returned from executing an operation.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html\#toc-methods)

[\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#method___toString)
: string Provides debug information about the result object[get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#method_get)
: mixed\|null Get a specific key value from the result model.[hasKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#method_hasKey)
: bool Check if the model contains a key by name[search()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#method_search)
: mixed Returns the result of executing a JMESPath expression on the contents
of the Result model.[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#method_toArray)
: array<string\|int, mixed> Convert the result to an array.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html\#methods)

#### \_\_toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html\#method___toString)

Provides debug information about the result object

`
    public
                    __toString() : string`

##### Return values

string

#### get()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html\#method_get)

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

#### hasKey()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html\#method_hasKey)

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

#### search()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html\#method_search)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html\#method_search\#tags)

link[JMESPath documentation](http://jmespath.readthedocs.org/en/latest)

##### Return values

mixed
—

Returns the result of the JMESPath expression.

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html\#method_toArray)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#toc-methods)
- Methods
  - [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#method___toString)
  - [get()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#method_get)
  - [hasKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#method_hasKey)
  - [search()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#method_search)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html#top)
