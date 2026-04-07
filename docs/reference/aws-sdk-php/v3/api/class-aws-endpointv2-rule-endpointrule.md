Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)
- [Rule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.endpointv2.rule.html)

## EndpointRule     extends [AbstractRule](class-aws-endpointv2-rule-abstractrule.md)   in package    - [Aws](package-aws.md)

A rule within a rule set. All rules contain a conditions property,
which can be empty, and documentation about the rule.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html#method___construct)
: mixed [evaluate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html#method_evaluate)
: [RulesetEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.RulesetEndpoint.html)If all the rule's conditions are met, return the resolved
endpoint object.[getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
: array<string\|int, mixed> [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
: mixed [getEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html#method_getEndpoint)
: array<string\|int, mixed>

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $definition) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>

#### evaluate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html\#method_evaluate)

If all the rule's conditions are met, return the resolved
endpoint object.

`
    public
                    evaluate(array<string|int, mixed> $inputParameters, RulesetStandardLibrary $standardLibrary) : RulesetEndpoint`

##### Parameters

$inputParameters
: array<string\|int, mixed>$standardLibrary
: RulesetStandardLibrary

##### Return values

[RulesetEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Ruleset.RulesetEndpoint.html)
—

\| null

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

#### getEndpoint()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html\#method_getEndpoint)

`
    public
                    getEndpoint() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html#method___construct)
  - [evaluate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html#method_evaluate)
  - [getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
  - [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
  - [getEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html#method_getEndpoint)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.EndpointRule.html#top)
