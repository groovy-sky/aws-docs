Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## Validator        in package    - [Aws](package-aws.md)

Validates a schema against a hash of input.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html#method___construct)
: mixed [validate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html#method_validate)
: mixed Validates the given input against the schema.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html\#method___construct)

`
    public
                    __construct([array<string|int, mixed> $constraints = null ]) : mixed`

##### Parameters

$constraints
: array<string\|int, mixed>
= null

Associative array of constraints to enforce.
Accepts the following keys: "required", "min",
"max", and "pattern". If a key is not
provided, the constraint will assume false.

#### validate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html\#method_validate)

Validates the given input against the schema.

`
    public
                    validate(string $name, Shape $shape, array<string|int, mixed> $input) : mixed`

##### Parameters

$name
: string

Operation name

$shape
: [Shape](class-aws-api-shape.md)

Shape to validate

$input
: array<string\|int, mixed>

Input to validate

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html\#method_validate\#tags)

throwsInvalidArgumentException

if the input is invalid.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html#method___construct)
  - [validate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html#method_validate)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html#top)
