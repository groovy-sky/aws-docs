This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EventSchemas::Discoverer

Use the `AWS::EventSchemas::Discoverer` resource to specify a
_discoverer_ that is associated with an event bus. A discoverer
allows the Amazon EventBridge Schema Registry to automatically generate schemas based on
events on an event bus.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EventSchemas::Discoverer",
  "Properties" : {
      "CrossAccount" : Boolean,
      "Description" : String,
      "SourceArn" : String,
      "Tags" : [ TagsEntry, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EventSchemas::Discoverer
Properties:
  CrossAccount: Boolean
  Description: String
  SourceArn: String
  Tags:
    - TagsEntry

```

## Properties

`CrossAccount`

Allows for the discovery of the event schemas that are sent to the event bus from another account.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the discoverer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceArn`

The ARN of the event bus.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags associated with the resource.

_Required_: No

_Type_: Array of [TagsEntry](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eventschemas-discoverer-tagsentry.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic
function, `Ref` returns the ARN of the discoverer.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DiscovererArn`

The ARN of the discoverer.

`DiscovererId`

The ID of the discoverer.

`State`

The state of the discoverer.

## Examples

### Generate Schemas for Events on the Default Event Bus

#### YAML

```yaml

Resources:
  MyDiscoverer:
    Type: AWS::EventSchemas::Discoverer
    Properties:
      SourceArn: 'arn:aws:events:us-west-2:012345678910:event-bus/default'
      Description: 'discover all custom schemas'

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon EventBridge Schemas

TagsEntry
