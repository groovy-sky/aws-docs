Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)
- [S3Transfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.html)
- [Progress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.s3transfer.progress.html)

## AbstractProgressBarFormat        in package    - [Aws](package-aws.md)

AbstractYes

Defines a progress bar format.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method___construct)
: mixed [format()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_format)
: string [getArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_getArgs)
: array<string\|int, mixed> [getFormatParameters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_getFormatParameters)
: array<string\|int, mixed> [getFormatTemplate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_getFormatTemplate)
: string [setArg()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_setArg)
: void [setArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_setArgs)
: void To set multiple arguments at once.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html\#method___construct)

`
    public
                    __construct([array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$args
: array<string\|int, mixed>
= \[\]

#### format()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html\#method_format)

`
    public
                    format() : string`

##### Return values

string

#### getArgs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html\#method_getArgs)

`
    public
                    getArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getFormatParameters()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html\#method_getFormatParameters)

`
    public
    abstract                getFormatParameters() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getFormatTemplate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html\#method_getFormatTemplate)

`
    public
    abstract                getFormatTemplate() : string`

##### Return values

string

#### setArg()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html\#method_setArg)

`
    public
                    setArg(string $key, mixed $value) : void`

##### Parameters

$key
: string$value
: mixed

#### setArgs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html\#method_setArgs)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method___construct)
  - [format()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_format)
  - [getArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_getArgs)
  - [getFormatParameters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_getFormatParameters)
  - [getFormatTemplate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_getFormatTemplate)
  - [setArg()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_setArg)
  - [setArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#method_setArgs)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Transfer.Progress.AbstractProgressBarFormat.html#top)
