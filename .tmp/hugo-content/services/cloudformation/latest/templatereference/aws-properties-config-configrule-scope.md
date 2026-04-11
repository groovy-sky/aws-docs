This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConfigRule Scope

Defines which resources trigger an evaluation for an AWS Config
rule. The scope can include one or more resource types, a
combination of a tag key and value, or a combination of one resource
type and one resource ID. Specify a scope to constrain which
resources trigger an evaluation for a rule. Otherwise, evaluations
for the rule are triggered when any resource in your recording group
changes in configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComplianceResourceId" : String,
  "ComplianceResourceTypes" : [ String, ... ],
  "TagKey" : String,
  "TagValue" : String
}

```

### YAML

```yaml

  ComplianceResourceId: String
  ComplianceResourceTypes:
    - String
  TagKey: String
  TagValue: String

```

## Properties

`ComplianceResourceId`

The ID of the only AWS resource that you want to trigger an
evaluation for the rule. If you specify a resource ID, you must
specify one resource type for
`ComplianceResourceTypes`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `768`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceResourceTypes`

The resource types of only those AWS resources that you want to
trigger an evaluation for the rule. You can only specify one type if
you also specify a resource ID for
`ComplianceResourceId`.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagKey`

The tag key that is applied to only those AWS resources that
you want to trigger an evaluation for the rule.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagValue`

The tag value applied to only those AWS resources that you want
to trigger an evaluation for the rule. If you specify a value for
`TagValue`, you must also specify a value for
`TagKey`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Multiple Resource Types with Tag-Based Scope](#aws-properties-config-configrule-scope--examples--Multiple_Resource_Types_with_Tag-Based_Scope)

- [Single Resource Specific Scope](#aws-properties-config-configrule-scope--examples--Single_Resource_Specific_Scope)

### Multiple Resource Types with Tag-Based Scope

This example configures AWS Config to evaluate both Amazon EC2 instances and volumes that are tagged with " `Environment` = `Production`". This is useful when you want to monitor compliance for multiple resource types that share specific tags.

#### YAML

```yaml

Scope:
  ComplianceResourceTypes:
    - "AWS::EC2::Instance"
    - "AWS::EC2::Volume"
  TagKey: "Environment"
  TagValue: "Production"
```

#### JSON

```json

{
  "Scope": {
    "ComplianceResourceTypes": [
      "AWS::EC2::Instance",
      "AWS::EC2::Volume"
    ],
    "TagKey": "Environment",
    "TagValue": "Production"
  }
}
```

### Single Resource Specific Scope

This example shows how to target a specific Amazon EC2 instance for evaluation using its resource ID. When using `ComplianceResourceId`, you must specify exactly one resource type in `ComplianceResourceTypes`.

#### YAML

```yaml

Scope:
  ComplianceResourceId: "i-1234567890abcdef0"
  ComplianceResourceTypes:
    - "AWS::EC2::Instance"
```

#### JSON

```json

{
  "Scope": {
    "ComplianceResourceId": "i-1234567890abcdef0",
    "ComplianceResourceTypes": [
      "AWS::EC2::Instance"
    ]
  }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationModeConfiguration

Source

All content copied from https://docs.aws.amazon.com/.
