This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::RemediationConfiguration SsmControls

AWS Systems Manager (SSM) specific remediation controls.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConcurrentExecutionRatePercentage" : Integer,
  "ErrorPercentage" : Integer
}

```

### YAML

```yaml

  ConcurrentExecutionRatePercentage: Integer
  ErrorPercentage: Integer

```

## Properties

`ConcurrentExecutionRatePercentage`

The maximum percentage of remediation actions allowed to run in parallel on the non-compliant resources for that specific rule. You can specify a percentage, such as 10%. The default value is 10.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ErrorPercentage`

The percentage of errors that are allowed before SSM stops running automations on non-compliant resources for that specific rule.
You can specify a percentage of errors, for example 10%. If you do not specifiy a percentage, the default is 50%.
For example, if you set the ErrorPercentage to 40% for 10 non-compliant resources, then SSM stops running the automations when the fifth error is received.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RemediationParameterValue

AWS::Config::StoredQuery

All content copied from https://docs.aws.amazon.com/.
