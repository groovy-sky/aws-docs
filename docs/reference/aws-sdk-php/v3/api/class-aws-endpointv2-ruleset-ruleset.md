Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)
- [Ruleset](namespace-aws-endpointv2-ruleset.md)

## Ruleset        in package    - [Aws](package-aws.md)

A collection of rules, parameter definitions and a class of helper functions
used to resolve either an endpoint or an error.

### Table of Contents  [header link](class-aws-endpointv2-ruleset-ruleset-toc.md)

#### Properties  [header link](class-aws-endpointv2-ruleset-ruleset-toc-properties.md)

[$standardLibrary](class-aws-endpointv2-ruleset-ruleset-property-standardlibrary.md)
: RulesetStandardLibrary

#### Methods  [header link](class-aws-endpointv2-ruleset-ruleset-toc-methods.md)

[\_\_construct()](class-aws-endpointv2-ruleset-ruleset-method-construct.md)
: mixed [evaluate()](class-aws-endpointv2-ruleset-ruleset-method-evaluate.md)
: mixed Evaluate the ruleset against the input parameters.[getParameters()](class-aws-endpointv2-ruleset-ruleset-method-getparameters.md)
: array<string\|int, mixed> [getRules()](class-aws-endpointv2-ruleset-ruleset-method-getrules.md)
: array<string\|int, mixed> [getVersion()](class-aws-endpointv2-ruleset-ruleset-method-getversion.md)
: mixed

### Properties  [header link](class-aws-endpointv2-ruleset-ruleset-properties.md)

#### $standardLibrary  [header link](class-aws-endpointv2-ruleset-ruleset-property-standardlibrary.md)

`
    public
        RulesetStandardLibrary
    $standardLibrary
    `

### Methods  [header link](class-aws-endpointv2-ruleset-ruleset-methods.md)

#### \_\_construct()  [header link](class-aws-endpointv2-ruleset-ruleset-method-construct.md)

`
    public
                    __construct(array<string|int, mixed> $ruleset, array<string|int, mixed> $partitions) : mixed`

##### Parameters

$ruleset
: array<string\|int, mixed>$partitions
: array<string\|int, mixed>

#### evaluate()  [header link](class-aws-endpointv2-ruleset-ruleset-method-evaluate.md)

Evaluate the ruleset against the input parameters.

`
    public
                    evaluate(array<string|int, mixed> $inputParameters) : mixed`

Return the first rule the parameters match against.

##### Parameters

$inputParameters
: array<string\|int, mixed>

#### getParameters()  [header link](class-aws-endpointv2-ruleset-ruleset-method-getparameters.md)

`
    public
                    getParameters() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getRules()  [header link](class-aws-endpointv2-ruleset-ruleset-method-getrules.md)

`
    public
                    getRules() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getVersion()  [header link](class-aws-endpointv2-ruleset-ruleset-method-getversion.md)

`
    public
                    getVersion() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Properties](class-aws-endpointv2-ruleset-ruleset-toc-properties.md)
  - [Methods](class-aws-endpointv2-ruleset-ruleset-toc-methods.md)
- Properties
  - [$standardLibrary](class-aws-endpointv2-ruleset-ruleset-property-standardlibrary.md)
- Methods
  - [\_\_construct()](class-aws-endpointv2-ruleset-ruleset-method-construct.md)
  - [evaluate()](class-aws-endpointv2-ruleset-ruleset-method-evaluate.md)
  - [getParameters()](class-aws-endpointv2-ruleset-ruleset-method-getparameters.md)
  - [getRules()](class-aws-endpointv2-ruleset-ruleset-method-getrules.md)
  - [getVersion()](class-aws-endpointv2-ruleset-ruleset-method-getversion.md)

[Back To Top](class-aws-endpointv2-ruleset-ruleset-top.md)
