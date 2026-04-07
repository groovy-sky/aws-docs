This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Archive

Creates an archive of events with the specified settings. When you create an archive,
incoming events might not immediately start being sent to the archive. Allow a short period of
time for changes to take effect. If you do not specify a pattern to filter events sent to the
archive, all events are sent to the archive except replayed events. Replayed events are not
sent to an archive.

###### Important

If you have specified that EventBridge use a customer managed key for encrypting the source event bus, we strongly recommend you also specify a
customer managed key for any archives for the event bus as well.

For more information, see [Encrypting archives](https://docs.aws.amazon.com/eventbridge/latest/userguide/encryption-archives.html) in the _Amazon EventBridge User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Events::Archive",
  "Properties" : {
      "ArchiveName" : String,
      "Description" : String,
      "EventPattern" : Json,
      "KmsKeyIdentifier" : String,
      "RetentionDays" : Integer,
      "SourceArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Events::Archive
Properties:
  ArchiveName: String
  Description: String
  EventPattern: Json
  KmsKeyIdentifier: String
  RetentionDays: Integer
  SourceArn: String

```

## Properties

`ArchiveName`

The name for the archive to create.

_Required_: No

_Type_: String

_Pattern_: `[\.\-_A-Za-z0-9]+`

_Minimum_: `1`

_Maximum_: `48`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description for the archive.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventPattern`

An event pattern to use to filter events sent to the archive.

_Required_: No

_Type_: Json

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyIdentifier`

The identifier of the AWS KMS
customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt this archive. The identifier can be the key
Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN.

If you do not specify a customer managed key identifier, EventBridge uses an
AWS owned key to encrypt the archive.

For more information, see [Identify and view keys](https://docs.aws.amazon.com/kms/latest/developerguide/viewing-keys.html) in the _AWS Key Management Service_
_Developer Guide_.

###### Important

If you have specified that EventBridge use a customer managed key for encrypting the source event bus, we strongly recommend you also specify a
customer managed key for any archives for the event bus as well.

For more information, see [Encrypting archives](https://docs.aws.amazon.com/eventbridge/latest/userguide/encryption-archives.html) in the _Amazon EventBridge User Guide_.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RetentionDays`

The number of days to retain events for. Default value is 0. If set to 0, events are
retained indefinitely

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceArn`

The ARN of the event bus that sends events to the archive.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws([a-z]|\-)*:events:([a-z]|\d|\-)*:([0-9]{12})?:.+\/.+$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the archive name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the archive created.

## Examples

### Create an archive

The following example creates an archive for all EC2 events sent from the default
event bus that retains events in the archive for 10 days.

#### JSON

```json

{
  "SampleArchive":
    "Type" : "AWS::Events::Archive",
    "Properties" : {
        "ArchiveName" : "MyArchive",
        "Description" : "Archive for all EC2 events",
        "EventPattern" : {
              "source": [
                 "aws.ec2"
              ]
           },
        "RetentionDays" : "10",
        "SourceArn" : "arn:aws:events:us-west-2:123456789012:event-bus/default"
      }
}
```

#### YAML

```yaml

SampleArchive:
  Type: 'AWS::Events::Archive'
  Properties:
    ArchiveName: MyArchive
    Description: Archive for all EC2 events
    EventPattern:
      source:
        - "aws.ec2"
    RetentionDays: 10
    SourceArn: 'arn:aws:events:us-west-2:123456789012:event-bus/default'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Events::ApiDestination

AWS::Events::Connection
