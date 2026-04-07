Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)
- [Ruleset](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.endpointv2.ruleset.html)

## Ruleset        in package    - [Aws](package-aws.md)

A collection of rules, parameter definitions and a class of helper functions
used to resolve either an endpoint or an error.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html\#toc)

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html\#toc-properties)

[$standardLibrary](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#property_standardLibrary)
: RulesetStandardLibrary

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#method___construct)
: mixed [evaluate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#method_evaluate)
: mixed Evaluate the ruleset against the input parameters.[getParameters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#method_getParameters)
: array<string\|int, mixed> [getRules()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#method_getRules)
: array<string\|int, mixed> [getVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#method_getVersion)
: mixed

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html\#properties)

#### $standardLibrary  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html\#property_standardLibrary)

`
    public
        RulesetStandardLibrary
    $standardLibrary
    `

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $ruleset, array<string|int, mixed> $partitions) : mixed`

##### Parameters

$ruleset
: array<string\|int, mixed>$partitions
: array<string\|int, mixed>

#### evaluate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html\#method_evaluate)

Evaluate the ruleset against the input parameters.

`
    public
                    evaluate(array<string|int, mixed> $inputParameters) : mixed`

Return the first rule the parameters match against.

##### Parameters

$inputParameters
: array<string\|int, mixed>

#### getParameters()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html\#method_getParameters)

`
    public
                    getParameters() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getRules()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html\#method_getRules)

`
    public
                    getRules() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getVersion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html\#method_getVersion)

`
    public
                    getVersion() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#toc-methods)
- Properties
  - [$standardLibrary](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#property_standardLibrary)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#method___construct)
  - [evaluate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#method_evaluate)
  - [getParameters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#method_getParameters)
  - [getRules()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#method_getRules)
  - [getVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#method_getVersion)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.Ruleset.html#top)
