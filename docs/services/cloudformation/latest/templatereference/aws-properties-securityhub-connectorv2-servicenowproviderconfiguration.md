This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::ConnectorV2 ServiceNowProviderConfiguration

The initial configuration settings required to establish an integration between Security Hub CSPM and ServiceNow ITSM.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InstanceName" : String,
  "SecretArn" : String
}

```

### YAML

```yaml

  InstanceName: String
  SecretArn: String

```

## Properties

`InstanceName`

The instance name of ServiceNow ITSM.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArn`

The Amazon Resource Name (ARN) of the AWS Secrets Manager secret that contains the ServiceNow credentials.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Provider

AWS::SecurityHub::DelegatedAdmin

All content copied from https://docs.aws.amazon.com/.
