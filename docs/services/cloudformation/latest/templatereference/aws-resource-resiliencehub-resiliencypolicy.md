This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResilienceHub::ResiliencyPolicy

Defines a resiliency policy.

###### Note

AWS Resilience Hub allows you to provide a value of zero for `rtoInSecs`
and `rpoInSecs` of your resiliency policy. But, while assessing your application,
the lowest possible assessment result is near zero. Hence, if you provide value zero for
`rtoInSecs` and `rpoInSecs`, the estimated workload RTO and
estimated workload RPO result will be near zero and the **Compliance**
**status** for your application will be set to **Policy**
**breached**.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ResilienceHub::ResiliencyPolicy",
  "Properties" : {
      "DataLocationConstraint" : String,
      "Policy" : PolicyMap,
      "PolicyDescription" : String,
      "PolicyName" : String,
      "Tags" : {Key: Value, ...},
      "Tier" : String
    }
}

```

### YAML

```yaml

Type: AWS::ResilienceHub::ResiliencyPolicy
Properties:
  DataLocationConstraint: String
  Policy:
    PolicyMap
  PolicyDescription: String
  PolicyName: String
  Tags:
    Key: Value
  Tier: String

```

## Properties

`DataLocationConstraint`

Specifies a high-level geographical location constraint for where your resilience policy
data can be stored.

_Required_: No

_Type_: String

_Allowed values_: `AnyLocation | SameContinent | SameCountry`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Policy`

The resiliency policy.

_Required_: Yes

_Type_: [PolicyMap](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-resiliencehub-resiliencypolicy-policymap.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyDescription`

Description of the resiliency policy.

_Required_: No

_Type_: String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyName`

The name of the policy

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9_\-]{1,59}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags assigned to the resource. A tag is a label that you assign to an AWS resource.
Each tag consists of a key/value pair.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,128}`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tier`

The tier for this resiliency policy, ranging from the highest severity
( `MissionCritical`) to lowest ( `NonCritical`).

_Required_: Yes

_Type_: String

_Allowed values_: `MissionCritical | Critical | Important | CoreServices | NonCritical`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

The returned Amazon Resource Name (ARN) for the resiliency policy.

### Fn::GetAtt

Amazon Resource Name (ARN) for the resiliency policy.

`PolicyArn`

Amazon Resource Name (ARN) of the resiliency policy.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResourceMapping

FailurePolicy
