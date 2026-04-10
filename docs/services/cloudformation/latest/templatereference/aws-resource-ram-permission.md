This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RAM::Permission

Creates a customer managed permission for a specified resource type that you can attach to resource shares.
It is created in the AWS Region in which you call the operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RAM::Permission",
  "Properties" : {
      "Name" : String,
      "PolicyTemplate" : Json,
      "ResourceType" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RAM::Permission
Properties:
  Name: String
  PolicyTemplate: Json
  ResourceType: String
  Tags:
    - Tag

```

## Properties

`Name`

Specifies the name of the customer managed permission. The name must be unique within the
AWS Region.

_Required_: Yes

_Type_: String

_Pattern_: `[\w.-]*`

_Minimum_: `1`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyTemplate`

A string in JSON format string that contains the following elements of a
resource-based policy:

- **Effect**: must be set to
`ALLOW`.

- **Action**: specifies the actions that are
allowed by this customer managed permission. The list must contain only actions that are supported by
the specified resource type. For a list of all actions supported by each
resource type, see [Actions, resources, and condition keys for AWS services](../../../service-authorization/latest/reference/reference-policies-actions-resources-contextkeys.md) in the
_AWS Identity and Access Management User Guide_.

- **Condition**: (optional) specifies conditional
parameters that must evaluate to true when a user attempts an action for that
action to be allowed. For more information about the Condition element, see
[IAM\
policies: Condition element](../../../iam/latest/userguide/reference-policies-elements-condition.md) in the _AWS Identity and Access Management User_
_Guide_.

This template can't include either the `Resource` or
`Principal` elements. Those are both filled in by AWS RAM when it instantiates
the resource-based policy on each resource shared using this managed permission. The
`Resource` comes from the ARN of the specific resource that you are sharing.
The `Principal` comes from the list of identities added to the resource
share.

_Required_: Yes

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceType`

Specifies the name of the resource type that this customer managed permission applies
to.

The format is `<service-code>:<resource-type>` and is not case sensitive. For example, to specify an Amazon EC2 Subnet, you can
use the string `ec2:subnet`. To see the list of valid values for this
parameter, query the [ListResourceTypes](../../../../reference/ram/latest/apireference/api-listresourcetypes.md)
operation.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Specifies a list of one or more tag key and value pairs to attach to the
permission.

_Required_: No

_Type_: Array of [Tag](aws-properties-ram-permission-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the permission.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the new permission.

`IsResourceTypeDefault`

Specifies whether this permission is the default for new resource shares that include
resources of the associated resource type.

`PermissionType`

The type of managed permission. This can be one of the following values:

- **AWS\_MANAGED\_PERMISSION** – AWS
created and manages this managed permission. You can associate it with your
resource shares, but you can't modify it.

- **CUSTOMER\_MANAGED\_PERMISSION** – You, or another
principal in your account created this managed permission. You can associate it
with your resource shares and create new versions that have different
permissions.

`Version`

The version number for this version of the permission.

## See also

- [CreatePermission](../../../../reference/ram/latest/apireference/api-createpermission.md) in the _AWS Resource Access Manager API Reference_

- [AWS Resource Access Manager User Guide](../../../ram/latest/userguide.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Resource Access Manager

Tag

All content copied from https://docs.aws.amazon.com/.
