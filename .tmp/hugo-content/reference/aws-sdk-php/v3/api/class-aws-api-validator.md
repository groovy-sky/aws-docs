Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## Validator        in package    - [Aws](package-aws.md)

Validates a schema against a hash of input.

### Table of Contents  [header link](class-aws-api-validator-toc.md)

#### Methods  [header link](class-aws-api-validator-toc-methods.md)

[\_\_construct()](class-aws-api-validator-method-construct.md)
: mixed [validate()](class-aws-api-validator-method-validate.md)
: mixed Validates the given input against the schema.

### Methods  [header link](class-aws-api-validator-methods.md)

#### \_\_construct()  [header link](class-aws-api-validator-method-construct.md)

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

#### validate()  [header link](class-aws-api-validator-method-validate.md)

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

##### Tags  [header link](class-aws-api-validator-method-validate-tags.md)

throwsInvalidArgumentException

if the input is invalid.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-api-validator-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-api-validator-method-construct.md)
  - [validate()](class-aws-api-validator-method-validate.md)

[Back To Top](class-aws-api-validator-top.md)
