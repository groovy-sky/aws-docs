This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::Group

Creates a new group.

For information about the number of groups you can create, see [Limitations\
on IAM Entities](../../../iam/latest/userguide/limitationsonentities.md) in the _IAM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::Group",
  "Properties" : {
      "GroupName" : String,
      "ManagedPolicyArns" : [ String, ... ],
      "Path" : String,
      "Policies" : [ Policy, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IAM::Group
Properties:
  GroupName: String
  ManagedPolicyArns:
    - String
  Path: String
  Policies:
    - Policy

```

## Properties

`GroupName`

The name of the group to create. Do not include the path in this value.

The group name must be unique within the account. Group names are not distinguished by
case. For example, you cannot create groups named both "ADMINS" and "admins". If you don't
specify a name, CloudFormation generates a unique physical ID and uses that ID for
the group name.

###### Important

If you specify a name, you cannot perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must
replace the resource, specify a new name.

If you specify a name, you must specify the `CAPABILITY_NAMED_IAM` value to
acknowledge your template's capabilities. For more information, see [Acknowledging IAM Resources in CloudFormation\
Templates](../userguide/using-iam-template.md#using-iam-capabilities).

###### Important

Naming an IAM resource can cause an unrecoverable error if you reuse
the same template in multiple Regions. To prevent this, we recommend using
`Fn::Join` and `AWS::Region` to create a Region-specific name,
as in the following example: `{"Fn::Join": ["", [{"Ref": "AWS::Region"}, {"Ref":
               "MyResourceName"}]]}`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManagedPolicyArns`

The Amazon Resource Name (ARN) of the IAM policy you want to attach.

For more information about ARNs, see [Amazon Resource Names (ARNs)](../../../../general/latest/gr/aws-arns-and-namespaces.md) in the _AWS General Reference_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The path to the group. For more information about paths, see [IAM\
identifiers](../../../iam/latest/userguide/using-identifiers.md) in the _IAM User Guide_.

This parameter is optional. If it is not included, it defaults to a slash (/).

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting
of either a forward slash (/) by itself or a string that must begin and end with forward slashes.
In addition, it can contain any ASCII character from the ! ( `\u0021`) through the DEL character ( `\u007F`), including
most punctuation characters, digits, and upper and lowercased letters.

_Required_: No

_Type_: String

_Pattern_: `(\u002F)|(\u002F[\u0021-\u007E]+\u002F)`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Policies`

Adds or updates an inline policy document that is embedded in the specified IAM group. To view AWS::IAM::Group snippets, see [Declaring\
an IAM Group Resource](../userguide/quickref-iam.md#scenario-iam-group).

###### Important

The name of each inline policy for a role, user, or group must be unique. If you
don't choose unique names, updates to the IAM identity will fail.

For information about limits on the number of inline policies that you can embed in a
group, see [Limitations on IAM\
Entities](../../../iam/latest/userguide/limitationsonentities.md) in the _IAM User Guide_.

_Required_: No

_Type_: Array of [Policy](aws-properties-iam-group-policy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `GroupName`. For example:
`mystack-mygroup-1DZETITOWEKVO`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) for the specified `AWS::IAM::Group`
resource. For example:
`arn:aws:iam::123456789012:group/mystack-mygroup-1DZETITOWEKVO`.

## See also

- To view `AWS::IAM::Group` template example snippets, see [Declaring an IAM Group Resource](../userguide/quickref-iam.md#scenario-iam-group).

- [CreateGroup](../../../../reference/iam/latest/apireference/api-creategroup.md) in the _AWS Identity and Access Management API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IAM::AccessKey

Policy

All content copied from https://docs.aws.amazon.com/.
