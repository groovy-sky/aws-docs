Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)
- [Rule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.endpointv2.rule.html)

## AbstractRule        in package    - [Aws](package-aws.md)

AbstractYes

A rule within a rule set. All rules contain a conditions property,
which can be empty, and documentation about the rule.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html#method___construct)
: mixed [evaluate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html#method_evaluate)
: mixed [getConditions()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html#method_getConditions)
: array<string\|int, mixed> [getDocumentation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html#method_getDocumentation)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $definition) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>

#### evaluate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html\#method_evaluate)

`
    public
    abstract                evaluate(array<string|int, mixed> $inputParameters, RulesetStandardLibrary $standardLibrary) : mixed`

##### Parameters

$inputParameters
: array<string\|int, mixed>$standardLibrary
: RulesetStandardLibrary

#### getConditions()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html\#method_getConditions)

`
    public
                    getConditions() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getDocumentation()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html\#method_getDocumentation)

`
    public
                    getDocumentation() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html#method___construct)
  - [evaluate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html#method_evaluate)
  - [getConditions()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html#method_getConditions)
  - [getDocumentation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html#method_getDocumentation)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.AbstractRule.html#top)
