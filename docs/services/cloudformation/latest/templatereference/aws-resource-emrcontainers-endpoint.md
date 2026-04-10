This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::Endpoint

This entity represents the endpoint that is managed by Amazon EMR on EKS.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMRContainers::Endpoint",
  "Properties" : {
      "ConfigurationOverrides" : ConfigurationOverrides,
      "ExecutionRoleArn" : String,
      "Name" : String,
      "ReleaseLabel" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String,
      "VirtualClusterId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EMRContainers::Endpoint
Properties:
  ConfigurationOverrides:
    ConfigurationOverrides
  ExecutionRoleArn: String
  Name: String
  ReleaseLabel: String
  Tags:
    - Tag
  Type: String
  VirtualClusterId: String

```

## Properties

`ConfigurationOverrides`

The configuration settings that are used to override existing configurations for
endpoints.

_Required_: No

_Type_: [ConfigurationOverrides](aws-properties-emrcontainers-endpoint-configurationoverrides.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecutionRoleArn`

The execution role ARN of the endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):iam::(\d{12})?:(role((\u002F)|(\u002F[\u0021-\u007F]+\u002F))[\w+=,.@-]+)$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the endpoint.

_Required_: No

_Type_: String

_Pattern_: `[0-9A-Za-z][A-Za-z0-9\-_]*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReleaseLabel`

The EMR release version to be used for the endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9._/-]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags of the endpoint.

_Required_: No

_Type_: Array of [Tag](aws-properties-emrcontainers-endpoint-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the endpoint.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VirtualClusterId`

The ID of the endpoint's virtual cluster.

_Required_: Yes

_Type_: String

_Pattern_: `[0-9a-z]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The ARN of the endpoint.

`CreatedAt`

The date and time when the endpoint was created.

`FailureReason`

The reasons why the endpoint has failed.

`Id`

The ID of the endpoint.

`SecurityGroup`

The security group configuration of the endpoint.

`ServerUrl`

The server URL of the endpoint.

`State`

The state of the endpoint.

`StateDetails`

Additional details of the endpoint state.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon EMR on EKS

Certificate

All content copied from https://docs.aws.amazon.com/.
