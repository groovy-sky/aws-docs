Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)
- [Rule](namespace-aws-endpointv2-rule.md)

## EndpointRule     extends [AbstractRule](class-aws-endpointv2-rule-abstractrule.md)   in package    - [Aws](package-aws.md)

A rule within a rule set. All rules contain a conditions property,
which can be empty, and documentation about the rule.

### Table of Contents  [header link](class-aws-endpointv2-rule-endpointrule-toc.md)

#### Methods  [header link](class-aws-endpointv2-rule-endpointrule-toc-methods.md)

[\_\_construct()](class-aws-endpointv2-rule-endpointrule-method-construct.md)
: mixed [evaluate()](class-aws-endpointv2-rule-endpointrule-method-evaluate.md)
: [RulesetEndpoint](class-aws-endpointv2-ruleset-rulesetendpoint.md)If all the rule's conditions are met, return the resolved
endpoint object.[getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
: array<string\|int, mixed> [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
: mixed [getEndpoint()](class-aws-endpointv2-rule-endpointrule-method-getendpoint.md)
: array<string\|int, mixed>

### Methods  [header link](class-aws-endpointv2-rule-endpointrule-methods.md)

#### \_\_construct()  [header link](class-aws-endpointv2-rule-endpointrule-method-construct.md)

`
    public
                    __construct(array<string|int, mixed> $definition) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>

#### evaluate()  [header link](class-aws-endpointv2-rule-endpointrule-method-evaluate.md)

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

[RulesetEndpoint](class-aws-endpointv2-ruleset-rulesetendpoint.md)
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

#### getEndpoint()  [header link](class-aws-endpointv2-rule-endpointrule-method-getendpoint.md)

`
    public
                    getEndpoint() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-endpointv2-rule-endpointrule-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-endpointv2-rule-endpointrule-method-construct.md)
  - [evaluate()](class-aws-endpointv2-rule-endpointrule-method-evaluate.md)
  - [getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
  - [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
  - [getEndpoint()](class-aws-endpointv2-rule-endpointrule-method-getendpoint.md)

[Back To Top](class-aws-endpointv2-rule-endpointrule-top.md)
