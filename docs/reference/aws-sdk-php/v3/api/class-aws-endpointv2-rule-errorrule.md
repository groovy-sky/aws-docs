Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)
- [Rule](namespace-aws-endpointv2-rule.md)

## ErrorRule     extends [AbstractRule](class-aws-endpointv2-rule-abstractrule.md)   in package    - [Aws](package-aws.md)

A rule within a rule set. All rules contain a conditions property,
which can be empty, and documentation about the rule.

### Table of Contents  [header link](class-aws-endpointv2-rule-errorrule-toc.md)

#### Methods  [header link](class-aws-endpointv2-rule-errorrule-toc-methods.md)

[\_\_construct()](class-aws-endpointv2-rule-errorrule-method-construct.md)
: mixed [evaluate()](class-aws-endpointv2-rule-errorrule-method-evaluate.md)
: null If an error rule's conditions are met, raise an
UnresolvedEndpointError containing the fully resolved error string.[getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
: array<string\|int, mixed> [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
: mixed [getError()](class-aws-endpointv2-rule-errorrule-method-geterror.md)
: array<string\|int, mixed>

### Methods  [header link](class-aws-endpointv2-rule-errorrule-methods.md)

#### \_\_construct()  [header link](class-aws-endpointv2-rule-errorrule-method-construct.md)

`
    public
                    __construct(mixed $definition) : mixed`

##### Parameters

$definition
: mixed

#### evaluate()  [header link](class-aws-endpointv2-rule-errorrule-method-evaluate.md)

If an error rule's conditions are met, raise an
UnresolvedEndpointError containing the fully resolved error string.

`
    public
                    evaluate(array<string|int, mixed> $inputParameters, RulesetStandardLibrary $standardLibrary) : null`

##### Parameters

$inputParameters
: array<string\|int, mixed>$standardLibrary
: RulesetStandardLibrary

##### Tags  [header link](class-aws-endpointv2-rule-errorrule-method-evaluate-tags.md)

throws[UnresolvedEndpointException](class-aws-exception-unresolvedendpointexception.md)

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

#### getError()  [header link](class-aws-endpointv2-rule-errorrule-method-geterror.md)

`
    public
                    getError() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-endpointv2-rule-errorrule-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-endpointv2-rule-errorrule-method-construct.md)
  - [evaluate()](class-aws-endpointv2-rule-errorrule-method-evaluate.md)
  - [getConditions()](class-aws-endpointv2-rule-abstractrule.md#method_getConditions)
  - [getDocumentation()](class-aws-endpointv2-rule-abstractrule.md#method_getDocumentation)
  - [getError()](class-aws-endpointv2-rule-errorrule-method-geterror.md)

[Back To Top](class-aws-endpointv2-rule-errorrule-top.md)
