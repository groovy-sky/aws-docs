Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)
- [Rule](namespace-aws-endpointv2-rule.md)

## TreeRule     extends [AbstractRule](class-aws-endpointv2-rule-abstractrule.md)   in package    - [Aws](package-aws.md)

A rule within a rule set. All rules contain a conditions property,
which can be empty, and documentation about the rule.

### Table of Contents  [header link](class-aws-endpointv2-rule-treerule-toc.md)

#### Methods  [header link](class-aws-endpointv2-rule-treerule-toc-methods.md)

[\_\_construct()](class-aws-endpointv2-rule-treerule-method-construct.md)
: mixed [evaluate()](class-aws-endpointv2-rule-treerule-method-evaluate.md)
: mixed If a tree rule's conditions evaluate successfully, iterate over its
subordinate rules and return a result if there is one. If any of the
subsequent rules are trees, the function will recurse until it reaches
an error or an endpoint rule[getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
: array<string\|int, mixed> [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
: mixed [getRules()](class-aws-endpointv2-rule-treerule-method-getrules.md)
: array<string\|int, mixed>

### Methods  [header link](class-aws-endpointv2-rule-treerule-methods.md)

#### \_\_construct()  [header link](class-aws-endpointv2-rule-treerule-method-construct.md)

`
    public
                    __construct(array<string|int, mixed> $definition) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>

#### evaluate()  [header link](class-aws-endpointv2-rule-treerule-method-evaluate.md)

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

#### getRules()  [header link](class-aws-endpointv2-rule-treerule-method-getrules.md)

`
    public
                    getRules() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-endpointv2-rule-treerule-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-endpointv2-rule-treerule-method-construct.md)
  - [evaluate()](class-aws-endpointv2-rule-treerule-method-evaluate.md)
  - [getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
  - [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
  - [getRules()](class-aws-endpointv2-rule-treerule-method-getrules.md)

[Back To Top](class-aws-endpointv2-rule-treerule-top.md)
