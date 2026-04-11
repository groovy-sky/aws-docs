This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityAgent::Application

The `AWS::SecurityAgent::Application` resource specifies a Security Agent application. An application provides the top-level configuration for Security Agent, including identity and access management settings and encryption options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityAgent::Application",
  "Properties" : {
      "DefaultKmsKeyId" : String,
      "IdCConfiguration" : IdCConfiguration,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SecurityAgent::Application
Properties:
  DefaultKmsKeyId: String
  IdCConfiguration:
    IdCConfiguration
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`DefaultKmsKeyId`

The identifier of the default Amazon Web Services KMS key to use for encrypting data in the application.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdCConfiguration`

The IAM Identity Center configuration for the application.

_Required_: No

_Type_: [IdCConfiguration](aws-properties-securityagent-application-idcconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of the IAM role to associate with the application.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to associate with the application.

_Required_: No

_Type_: Array of [Tag](aws-properties-securityagent-application-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application ID. For example:

`{ "Ref": "MyApplication" }`

For the application `MyApplication`, `Ref` returns the unique identifier of the Security Agent application.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ApplicationId`

The unique identifier of the Security Agent application. For example: `app-0123456789abcdef0`.

`ApplicationName`

The name of the Security Agent application.

`Domain`

The domain associated with the Security Agent application.

`IdCConfiguration.IdCApplicationArn`

The Amazon Resource Name (ARN) of the IAM Identity Center application associated with the Security Agent application.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfig

IdCConfiguration

All content copied from https://docs.aws.amazon.com/.
