Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)
- [Rule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.endpointv2.rule.html)

## ErrorRule     extends [AbstractRule](class-aws-endpointv2-rule-abstractrule.md)   in package    - [Aws](package-aws.md)

A rule within a rule set. All rules contain a conditions property,
which can be empty, and documentation about the rule.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html#method___construct)
: mixed [evaluate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html#method_evaluate)
: null If an error rule's conditions are met, raise an
UnresolvedEndpointError containing the fully resolved error string.[getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
: array<string\|int, mixed> [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
: mixed [getError()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html#method_getError)
: array<string\|int, mixed>

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html\#method___construct)

`
    public
                    __construct(mixed $definition) : mixed`

##### Parameters

$definition
: mixed

#### evaluate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html\#method_evaluate)

If an error rule's conditions are met, raise an
UnresolvedEndpointError containing the fully resolved error string.

`
    public
                    evaluate(array<string|int, mixed> $inputParameters, RulesetStandardLibrary $standardLibrary) : null`

##### Parameters

$inputParameters
: array<string\|int, mixed>$standardLibrary
: RulesetStandardLibrary

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html\#method_evaluate\#tags)

throws[UnresolvedEndpointException](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.UnresolvedEndpointException.html)

##### Return values

null

#### getConditions()  [header link](class-aws-endpointv2-rule-abstractrule.md\#method_getConditions)

`
    public
                    getConditions() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getDocumentation()  [header link](class-aws-endpointv2-rule-abstractrule.md\#method_getDocumentation)

`
    public
                    getDocumentation() : mixed`

#### getError()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html\#method_getError)

`
    public
                    getError() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html#method___construct)
  - [evaluate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html#method_evaluate)
  - [getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
  - [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
  - [getError()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html#method_getError)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.ErrorRule.html#top)
