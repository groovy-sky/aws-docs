This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lightsail::InstanceSnapshot

Describes an instance snapshot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lightsail::InstanceSnapshot",
  "Properties" : {
      "InstanceName" : String,
      "InstanceSnapshotName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Lightsail::InstanceSnapshot
Properties:
  InstanceName: String
  InstanceSnapshotName: String
  Tags:
    - Tag

```

## Properties

`InstanceName`

The name the user gave the instance ( `Amazon_Linux_2023-1`).

_Required_: Yes

_Type_: String

_Pattern_: `\w[\w\-]*\w`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InstanceSnapshotName`

The name of the snapshot.

_Required_: Yes

_Type_: String

_Pattern_: `\w[\w\-]*\w`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tag keys and optional values for the resource. For more information about tags in
Lightsail, see the [Amazon Lightsail Developer Guide](https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-tags).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lightsail-instancesnapshot-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the snapshot
( `arn:aws:lightsail:us-east-2:123456789101:InstanceSnapshot/d23b5706-3322-4d83-81e5-12345EXAMPLE`).

`FromInstanceArn`

The Amazon Resource Name (ARN) of the instance from which the snapshot was created
( `arn:aws:lightsail:us-east-2:123456789101:Instance/64b8404c-ccb1-430b-8daf-12345EXAMPLE`).

`FromInstanceName`

The instance from which the snapshot was created.

`IsFromAutoSnapshot`

A Boolean value indicating whether the snapshot was created from an automatic
snapshot.

`ResourceType`

The type of resource (usually `InstanceSnapshot`).

`SizeInGb`

The size in GB of the SSD.

`State`

The state the snapshot is in.

`SupportCode`

The support code. Include this code in your email to support when you have questions about
an instance or another resource in Lightsail. This code enables our support team to look up
your Lightsail information more easily.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Location
