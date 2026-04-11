This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::SequenceStore

Creates a sequence store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Omics::SequenceStore",
  "Properties" : {
      "AccessLogLocation" : String,
      "Description" : String,
      "ETagAlgorithmFamily" : String,
      "FallbackLocation" : String,
      "Name" : String,
      "PropagatedSetLevelTags" : [ String, ... ],
      "S3AccessPolicy" : Json,
      "SseConfig" : SseConfig,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Omics::SequenceStore
Properties:
  AccessLogLocation: String
  Description: String
  ETagAlgorithmFamily: String
  FallbackLocation: String
  Name: String
  PropagatedSetLevelTags:
    - String
  S3AccessPolicy: Json
  SseConfig:
    SseConfig
  Tags:
    Key: Value

```

## Properties

`AccessLogLocation`

Location of the access logs.

_Required_: No

_Type_: String

_Pattern_: `^$|^s3://([a-z0-9][a-z0-9-.]{1,61}[a-z0-9])/?((.{1,800})/)?$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the store.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ETagAlgorithmFamily`

The algorithm family of the ETag.

_Required_: No

_Type_: String

_Allowed values_: `MD5up | SHA256up | SHA512up`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FallbackLocation`

An S3 location that is used to store files that have failed a direct upload.

_Required_: No

_Type_: String

_Pattern_: `^$|^s3://([a-z0-9][a-z0-9-.]{1,61}[a-z0-9])/?((.{1,1024})/)?$`

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the store.

_Required_: Yes

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropagatedSetLevelTags`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `128 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3AccessPolicy`

Property description not available.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SseConfig`

Server-side encryption (SSE) settings for the store.

_Required_: No

_Type_: [SseConfig](aws-properties-omics-sequencestore-sseconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags for the store.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the details of this resource. For example:

`{ "Ref": "SequenceStore.CreationTime" }` `Ref` returns the timestamp for when the sequence store was created.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The store's ARN.

`CreationTime`

When the store was created.

`S3AccessPointArn`

This is ARN of the access point associated with the S3 bucket storing read sets.

`S3Uri`

The S3 URI of the sequence store.

`SequenceStoreId`

The store's ID.

`Status`

Status of the sequence store.

`StatusMessage`

The status message of the sequence store.

`UpdateTime`

The last-updated time of the Sequence Store.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Omics::RunGroup

SseConfig

All content copied from https://docs.aws.amazon.com/.
