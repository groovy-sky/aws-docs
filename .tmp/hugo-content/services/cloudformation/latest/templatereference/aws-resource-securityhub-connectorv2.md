This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::ConnectorV2

Grants permission to create a connectorV2 based on input parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::ConnectorV2",
  "Properties" : {
      "Description" : String,
      "KmsKeyArn" : String,
      "Name" : String,
      "Provider" : Provider,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::ConnectorV2
Properties:
  Description: String
  KmsKeyArn: String
  Name: String
  Provider:
    Provider
  Tags:
    Key: Value

```

## Properties

`Description`

The description of the connectorV2.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

The Amazon Resource Name (ARN) of KMS key used to encrypt secrets for the connectorV2.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The unique name of the connectorV2.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Provider`

The third-party provider detail for a service configuration.

_Required_: Yes

_Type_: [Provider](aws-properties-securityhub-connectorv2-provider.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to add to the connectorV2 when you create.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]{1,128}$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ConnectorV2Arn` for the `ConnectorV2Arn` resource created: `arn:aws:securityhub:region:123456789012:connectorv2/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConnectorArn`

The ARN of the V2 connector.

`ConnectorId`

The unique identifier of the V2 connector.

`ConnectorStatus`

The status of the V2 connector.

`CreatedAt`

The timestamp when the V2 connector was created.

`LastCheckedAt`

The most recent timestamp when the V2 connector was checked on health status.

`LastUpdatedAt`

The most recent timestamp when the V2 connector was updated.

`Message`

The message of the V2 connector when connector status is FAILED\_TO\_CONNECT.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecurityHubPolicy

JiraCloudProviderConfiguration

All content copied from https://docs.aws.amazon.com/.
