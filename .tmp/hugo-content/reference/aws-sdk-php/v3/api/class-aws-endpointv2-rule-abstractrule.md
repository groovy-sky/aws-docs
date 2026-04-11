Menu

- [Aws](namespace-aws.md)
- [EndpointV2](namespace-aws-endpointv2.md)
- [Rule](namespace-aws-endpointv2-rule.md)

## AbstractRule        in package    - [Aws](package-aws.md)

AbstractYes

A rule within a rule set. All rules contain a conditions property,
which can be empty, and documentation about the rule.

### Table of Contents  [header link](class-aws-endpointv2-rule-abstractrule-toc.md)

#### Methods  [header link](class-aws-endpointv2-rule-abstractrule-toc-methods.md)

[\_\_construct()](class-aws-endpointv2-rule-abstractrule-method-construct.md)
: mixed [evaluate()](class-aws-endpointv2-rule-abstractrule-method-evaluate.md)
: mixed [getConditions()](class-aws-endpointv2-rule-abstractrule-method-getconditions.md)
: array<string\|int, mixed> [getDocumentation()](class-aws-endpointv2-rule-abstractrule-method-getdocumentation.md)
: mixed

### Methods  [header link](class-aws-endpointv2-rule-abstractrule-methods.md)

#### \_\_construct()  [header link](class-aws-endpointv2-rule-abstractrule-method-construct.md)

`
    public
                    __construct(array<string|int, mixed> $definition) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>

#### evaluate()  [header link](class-aws-endpointv2-rule-abstractrule-method-evaluate.md)

`
    public
    abstract                evaluate(array<string|int, mixed> $inputParameters, RulesetStandardLibrary $standardLibrary) : mixed`

##### Parameters

$inputParameters
: array<string\|int, mixed>$standardLibrary
: RulesetStandardLibrary

#### getConditions()  [header link](class-aws-endpointv2-rule-abstractrule-method-getconditions.md)

`
    public
                    getConditions() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getDocumentation()  [header link](class-aws-endpointv2-rule-abstractrule-method-getdocumentation.md)

`
    public
                    getDocumentation() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-endpointv2-rule-abstractrule-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-endpointv2-rule-abstractrule-method-construct.md)
  - [evaluate()](class-aws-endpointv2-rule-abstractrule-method-evaluate.md)
  - [getConditions()](class-aws-endpointv2-rule-abstractrule-method-getconditions.md)
  - [getDocumentation()](class-aws-endpointv2-rule-abstractrule-method-getdocumentation.md)

[Back To Top](class-aws-endpointv2-rule-abstractrule-top.md)
