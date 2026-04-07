This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::EventStream

An Event Stream resource of Amazon Connect Customer Profiles.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CustomerProfiles::EventStream",
  "Properties" : {
      "DomainName" : String,
      "EventStreamName" : String,
      "Tags" : [ Tag, ... ],
      "Uri" : String
    }
}

```

### YAML

```yaml

Type: AWS::CustomerProfiles::EventStream
Properties:
  DomainName: String
  EventStreamName: String
  Tags:
    - Tag
  Uri: String

```

## Properties

`DomainName`

The unique name of the domain.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EventStreamName`

The name of the event stream.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-customerprofiles-eventstream-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Uri`

The StreamARN of the destination to deliver profile events to. For example,
arn:aws:kinesis:region:account-id:stream/stream-name.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

The timestamp of when the export was created.

`EventStreamArn`

A unique identifier for the event stream.

`State`

The operational state of destination stream for export.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

DestinationDetails
