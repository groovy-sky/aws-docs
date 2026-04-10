This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFormation::TypeActivation

The `AWS::CloudFormation::TypeActivation` resource activates a public
third-party extension, making it available for use in stack templates.

For information about the CloudFormation registry, see [Managing\
extensions with the CloudFormation registry](../userguide/registry.md) in the
_CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFormation::TypeActivation",
  "Properties" : {
      "AutoUpdate" : Boolean,
      "ExecutionRoleArn" : String,
      "LoggingConfig" : LoggingConfig,
      "MajorVersion" : String,
      "PublicTypeArn" : String,
      "PublisherId" : String,
      "Type" : String,
      "TypeName" : String,
      "TypeNameAlias" : String,
      "VersionBump" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudFormation::TypeActivation
Properties:
  AutoUpdate: Boolean
  ExecutionRoleArn: String
  LoggingConfig:
    LoggingConfig
  MajorVersion: String
  PublicTypeArn: String
  PublisherId: String
  Type: String
  TypeName: String
  TypeNameAlias: String
  VersionBump: String

```

## Properties

`AutoUpdate`

Whether to automatically update the extension in this account and Region when a new
_minor_ version is published by the extension publisher. Major versions
released by the publisher must be manually updated.

The default is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleArn`

The name of the IAM execution role to use to activate the extension.

_Required_: No

_Type_: String

_Pattern_: `arn:.+:iam::[0-9]{12}:role/.+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingConfig`

Specifies logging configuration information for an extension.

_Required_: No

_Type_: [LoggingConfig](aws-properties-cloudformation-typeactivation-loggingconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MajorVersion`

The major version of this extension you want to activate, if multiple major versions are
available. The default is the latest major version. CloudFormation uses the latest available
_minor_ version of the major version selected.

You can specify `MajorVersion` or `VersionBump`, but not
both.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublicTypeArn`

The Amazon Resource Number (ARN) of the public extension.

Conditional: You must specify `PublicTypeArn`, or `TypeName`,
`Type`, and `PublisherId`.

_Required_: Conditional

_Type_: String

_Pattern_: `arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/.+`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublisherId`

The ID of the extension publisher.

Conditional: You must specify `PublicTypeArn`, or `TypeName`,
`Type`, and `PublisherId`.

_Required_: Conditional

_Type_: String

_Pattern_: `[0-9a-zA-Z-]{1,40}`

_Minimum_: `1`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The extension type.

Conditional: You must specify `PublicTypeArn`, or `TypeName`,
`Type`, and `PublisherId`.

_Required_: Conditional

_Type_: String

_Allowed values_: `RESOURCE | MODULE | HOOK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeName`

The name of the extension.

Conditional: You must specify `PublicTypeArn`, or `TypeName`,
`Type`, and `PublisherId`.

_Required_: Conditional

_Type_: String

_Pattern_: `[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeNameAlias`

An alias to assign to the public extension in this account and Region. If you specify an
alias for the extension, CloudFormation treats the alias as the extension type name within this
account and Region. You must use the alias to refer to the extension in your templates, API
calls, and CloudFormation console.

An extension alias must be unique within a given account and Region. You can activate the
same public resource multiple times in the same account and Region, using different type name
aliases.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}(::MODULE){0,1}`

_Minimum_: `10`

_Maximum_: `204`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionBump`

Manually updates a previously-activated type to a new major or minor version, if
available. You can also use this parameter to update the value of
`AutoUpdate`.

- `MAJOR`: CloudFormation updates the extension to the newest major version, if
one is available.

- `MINOR`: CloudFormation updates the extension to the newest minor version, if
one is available.

_Required_: No

_Type_: String

_Allowed values_: `MAJOR | MINOR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Number (ARN) of the activated
extension, in this account and Region.

`{ "Ref": "arn:aws:cloudformation:us-west-2:123456789012:type/resource/My-Example" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the activated extension in this account and
Region.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

LoggingConfig

All content copied from https://docs.aws.amazon.com/.
