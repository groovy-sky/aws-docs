Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](namespace-aws-s3-s3transfer.md)
- [Progress](namespace-aws-s3-s3transfer-progress.md)

## AbstractProgressBarFormat        in package    - [Aws](package-aws.md)

AbstractYes

Defines a progress bar format.

### Table of Contents  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat-toc.md)

#### Methods  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat-toc-methods.md)

[\_\_construct()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-construct.md)
: mixed [format()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-format.md)
: string [getArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-getargs.md)
: array<string\|int, mixed> [getFormatParameters()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-getformatparameters.md)
: array<string\|int, mixed> [getFormatTemplate()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-getformattemplate.md)
: string [setArg()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-setarg.md)
: void [setArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-setargs.md)
: void To set multiple arguments at once.

### Methods  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat-methods.md)

#### \_\_construct()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-construct.md)

`
    public
                    __construct([array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$args
: array<string\|int, mixed>
= \[\]

#### format()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-format.md)

`
    public
                    format() : string`

##### Return values

string

#### getArgs()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-getargs.md)

`
    public
                    getArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getFormatParameters()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-getformatparameters.md)

`
    public
    abstract                getFormatParameters() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getFormatTemplate()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-getformattemplate.md)

`
    public
    abstract                getFormatTemplate() : string`

##### Return values

string

#### setArg()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-setarg.md)

`
    public
                    setArg(string $key, mixed $value) : void`

##### Parameters

$key
: string$value
: mixed

#### setArgs()  [header link](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-setargs.md)

To set multiple arguments at once.

`
    public
                    setArgs(array<string|int, mixed> $args) : void`

It does not override all the values, instead
it adds the arguments individually and if a value
already exists then that value will be overridden.

##### Parameters

$args
: array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-s3-s3transfer-progress-abstractprogressbarformat-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-construct.md)
  - [format()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-format.md)
  - [getArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-getargs.md)
  - [getFormatParameters()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-getformatparameters.md)
  - [getFormatTemplate()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-getformattemplate.md)
  - [setArg()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-setarg.md)
  - [setArgs()](class-aws-s3-s3transfer-progress-abstractprogressbarformat-method-setargs.md)

[Back To Top](class-aws-s3-s3transfer-progress-abstractprogressbarformat-top.md)
