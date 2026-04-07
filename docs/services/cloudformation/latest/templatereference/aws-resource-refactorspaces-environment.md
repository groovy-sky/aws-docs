This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RefactorSpaces::Environment

###### Note

AWS Migration Hub is no longer open to new customers as of November 7, 2025. For capabilities similar to AWS Migration Hub, explore [AWS Migration Hub](https://aws.amazon.com/transform).

Creates an AWS Migration Hub Refactor Spaces environment. The caller owns the environment resource, and all
Refactor Spaces applications, services, and routes created within the environment. They are referred
to as the _environment owner_. The environment owner has cross-account
visibility and control of Refactor Spaces resources that are added to the environment by other
accounts that the environment is shared with.

When creating an environment with a [CreateEnvironment:NetworkFabricType](https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/APIReference/API_CreateEnvironment.html#migrationhubrefactorspaces-CreateEnvironment-request-NetworkFabricType) of `TRANSIT_GATEWAY`, Refactor Spaces
provisions a transit gateway to enable services in VPCs to communicate directly across
accounts. If [CreateEnvironment:NetworkFabricType](https://docs.aws.amazon.com/migrationhub-refactor-spaces/latest/APIReference/API_CreateEnvironment.html#migrationhubrefactorspaces-CreateEnvironment-request-NetworkFabricType) is `NONE`, Refactor Spaces does not create
a transit gateway and you must use your network infrastructure to route traffic to services
with private URL endpoints.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RefactorSpaces::Environment",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "NetworkFabricType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RefactorSpaces::Environment
Properties:
  Description: String
  Name: String
  NetworkFabricType: String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the environment.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-_\s\.\!\*\#\@\']+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the environment.

_Required_: No

_Type_: String

_Pattern_: `^(?!env-)[a-zA-Z0-9]+[a-zA-Z0-9-_ ]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkFabricType`

The network fabric type of the environment.

_Required_: No

_Type_: String

_Allowed values_: `TRANSIT_GATEWAY | NONE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the environment.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-refactorspaces-environment-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the environment, for example,
`env-1234654123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the environment.

`EnvironmentIdentifier`

The unique identifier of the environment.

`TransitGatewayId`

The ID of the AWS Transit Gateway set up by the environment.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
