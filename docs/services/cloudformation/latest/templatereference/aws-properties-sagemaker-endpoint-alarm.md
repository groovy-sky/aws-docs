This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Endpoint Alarm

An Amazon CloudWatch alarm configured to monitor metrics on an endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmName" : String
}

```

### YAML

```yaml

  AlarmName: String

```

## Properties

`AlarmName`

The name of a CloudWatch alarm in your account.

_Required_: Yes

_Type_: String

_Pattern_: `(?!\s*$).+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SageMaker::Endpoint

AutoRollbackConfig

All content copied from https://docs.aws.amazon.com/.
