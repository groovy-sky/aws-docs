This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::VirtualMFADevice

Creates a new virtual MFA device for the AWS account. After creating
the virtual MFA, use [EnableMFADevice](https://docs.aws.amazon.com/IAM/latest/APIReference/API_EnableMFADevice.html) to attach
the MFA device to an IAM user. For more information about creating and
working with virtual MFA devices, see [Using a virtual MFA device](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_VirtualMFA.html) in
the _IAM User Guide_.

For information about the maximum number of MFA devices you can create, see [IAM and AWS STS quotas](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html) in the _IAM User Guide_.

###### Important

The seed information contained in the QR code and the Base32 string should be treated
like any other secret access information. In other words, protect the seed information
as you would your AWS access keys or your passwords. After you
provision your virtual device, you should ensure that the information is destroyed
following secure procedures.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IAM::VirtualMFADevice",
  "Properties" : {
      "Path" : String,
      "Tags" : [ Tag, ... ],
      "Users" : [ String, ... ],
      "VirtualMfaDeviceName" : String
    }
}

```

### YAML

```yaml

Type: AWS::IAM::VirtualMFADevice
Properties:
  Path: String
  Tags:
    - Tag
  Users:
    - String
  VirtualMfaDeviceName: String

```

## Properties

`Path`

The path for the virtual MFA device. For more information about paths, see [IAM\
identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) in the _IAM User Guide_.

This parameter is optional. If it is not included, it defaults to a slash (/).

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting
of either a forward slash (/) by itself or a string that must begin and end with forward slashes.
In addition, it can contain any ASCII character from the ! ( `\u0021`) through the DEL character ( `\u007F`), including
most punctuation characters, digits, and upper and lowercased letters.

_Required_: No

_Type_: String

_Pattern_: `(\u002F)|(\u002F[\u0021-\u007F]+\u002F)`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags that you want to attach to the new IAM virtual MFA device.
Each tag consists of a key name and an associated value. For more information about tagging, see [Tagging IAM resources](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the
_IAM User Guide_.

###### Note

If any one of the tags is invalid or if you exceed the allowed maximum number of tags, then the entire request
fails and the resource is not created.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iam-virtualmfadevice-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Users`

The IAM user associated with this virtual MFA device.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualMfaDeviceName`

The name of the virtual MFA device, which must be unique. Use with path to uniquely
identify a virtual MFA device.

This parameter allows (through its [regex pattern](http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric
characters with no spaces. You can also include any of the following characters: \_+=,.@-

_Required_: No

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `226`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `SerialNumber`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`SerialNumber`

Returns the serial number for the specified `AWS::IAM::VirtualMFADevice`
resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::IAM::UserToGroupAddition

Tag
