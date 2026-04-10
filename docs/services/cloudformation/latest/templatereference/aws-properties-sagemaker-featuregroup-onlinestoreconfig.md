This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::FeatureGroup OnlineStoreConfig

Use this to specify the AWS Key Management Service (KMS) Key ID, or
`KMSKeyId`, for at rest data encryption. You can turn
`OnlineStore` on or off by specifying the `EnableOnlineStore` flag
at General Assembly.

The default value is `False`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableOnlineStore" : Boolean,
  "SecurityConfig" : OnlineStoreSecurityConfig,
  "StorageType" : String,
  "TtlDuration" : TtlDuration
}

```

### YAML

```yaml

  EnableOnlineStore: Boolean
  SecurityConfig:
    OnlineStoreSecurityConfig
  StorageType: String
  TtlDuration:
    TtlDuration

```

## Properties

`EnableOnlineStore`

Turn `OnlineStore` off by specifying `False` for the
`EnableOnlineStore` flag. Turn `OnlineStore` on by specifying
`True` for the `EnableOnlineStore` flag.

The default value is `False`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityConfig`

Use to specify KMS Key ID ( `KMSKeyId`) for at-rest encryption of your
`OnlineStore`.

_Required_: No

_Type_: [OnlineStoreSecurityConfig](aws-properties-sagemaker-featuregroup-onlinestoresecurityconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StorageType`

Option for different tiers of low latency storage for real-time data retrieval.

- `Standard`: A managed low latency data store for feature groups.

- `InMemory`: A managed data store for feature groups that supports very
low latency retrieval.

_Required_: No

_Type_: String

_Allowed values_: `Standard | InMemory`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TtlDuration`

Time to live duration, where the record is hard deleted after the expiration time is
reached; `ExpiresAt` = `EventTime` \+ `TtlDuration`. For
information on HardDelete, see the [DeleteRecord](../../../../reference/sagemaker/latest/apireference/api-feature-store-deleterecord.md) API in the Amazon SageMaker API Reference guide.

_Required_: No

_Type_: [TtlDuration](aws-properties-sagemaker-featuregroup-ttlduration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OfflineStoreConfig

OnlineStoreSecurityConfig

All content copied from https://docs.aws.amazon.com/.
