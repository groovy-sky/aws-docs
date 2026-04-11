This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::Subscriber

Creates a subscriber for accounts that are already enabled in Amazon Security Lake. You can
create a subscriber with access to data in the current AWS Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityLake::Subscriber",
  "Properties" : {
      "AccessTypes" : [ String, ... ],
      "DataLakeArn" : String,
      "Sources" : [ Source, ... ],
      "SubscriberDescription" : String,
      "SubscriberIdentity" : SubscriberIdentity,
      "SubscriberName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SecurityLake::Subscriber
Properties:
  AccessTypes:
    - String
  DataLakeArn: String
  Sources:
    - Source
  SubscriberDescription: String
  SubscriberIdentity:
    SubscriberIdentity
  SubscriberName: String
  Tags:
    - Tag

```

## Properties

`AccessTypes`

You can choose to notify subscribers of new objects with an Amazon Simple Queue Service
(Amazon SQS) queue or through messaging to an HTTPS endpoint provided by the
subscriber.

Subscribers can consume data by directly querying AWS Lake Formation tables in your
Amazon S3 bucket through services like Amazon Athena. This subscription
type is defined as `LAKEFORMATION`.

_Required_: Yes

_Type_: Array of String

_Allowed values_: `LAKEFORMATION | S3`

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLakeArn`

The Amazon Resource Name (ARN) used to create the data lake.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Sources`

Amazon Security Lake supports log and event collection for natively supported AWS services. For more information, see the [Amazon Security Lake User Guide](../../../security-lake/latest/userguide/source-management.md).

_Required_: Yes

_Type_: Array of [Source](aws-properties-securitylake-subscriber-source.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubscriberDescription`

The subscriber descriptions for a subscriber account. The description for a subscriber
includes `subscriberName`, `accountID`, `externalID`, and
`subscriberId`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubscriberIdentity`

The AWS identity used to access your data.

_Required_: Yes

_Type_: [SubscriberIdentity](aws-properties-securitylake-subscriber-subscriberidentity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubscriberName`

The name of your Amazon Security Lake subscriber account.

_Required_: Yes

_Type_: String

_Pattern_: `^[\\\w\s\-_:/,.@=+]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of objects, one for each tag to associate with the subscriber. For each tag, you must specify both a tag key and a tag value. A tag value cannot be null, but it can be an empty string.

_Required_: No

_Type_: Array of [Tag](aws-properties-securitylake-subscriber-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `ref` function, `ref` returns the `Subscriber` name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ResourceShareArn`

The Amazon Resource Name (ARN) of the Amazon Security Lake subscriber.

`ResourceShareName`

The ARN name of the Amazon Security Lake subscriber.

`S3BucketArn`

The Amazon Resource Name (ARN) of the S3 bucket.

`SubscriberArn`

The Amazon Resource Name (ARN) of the Security Lake subscriber.

`SubscriberRoleArn`

The Amazon Resource Name (ARN) of the role used to create the Security Lake subscriber.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Transitions

AwsLogSource

All content copied from https://docs.aws.amazon.com/.
