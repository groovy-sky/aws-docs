Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)
- [Rule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.endpointv2.rule.html)

## TreeRule     extends [AbstractRule](class-aws-endpointv2-rule-abstractrule.md)   in package    - [Aws](package-aws.md)

A rule within a rule set. All rules contain a conditions property,
which can be empty, and documentation about the rule.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html#method___construct)
: mixed [evaluate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html#method_evaluate)
: mixed If a tree rule's conditions evaluate successfully, iterate over its
subordinate rules and return a result if there is one. If any of the
subsequent rules are trees, the function will recurse until it reaches
an error or an endpoint rule[getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
: array<string\|int, mixed> [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
: mixed [getRules()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html#method_getRules)
: array<string\|int, mixed>

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $definition) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>

#### evaluate()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html\#method_evaluate)

If a tree rule's conditions evaluate successfully, iterate over its
subordinate rules and return a result if there is one. If any of the
subsequent rules are trees, the function will recurse until it reaches
an error or an endpoint rule

`
    public
                    evaluate(array<string|int, mixed> $inputParameters, RulesetStandardLibrary $standardLibrary) : mixed`

##### Parameters

$inputParameters
: array<string\|int, mixed>$standardLibrary
: RulesetStandardLibrary

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

#### getRules()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html\#method_getRules)

`
    public
                    getRules() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html#method___construct)
  - [evaluate()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html#method_evaluate)
  - [getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
  - [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
  - [getRules()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html#method_getRules)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointV2.Rule.TreeRule.html#top)
