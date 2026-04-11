This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::Domain

Specifies an Amazon Connect Customer Profiles Domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CustomerProfiles::Domain",
  "Properties" : {
      "DataStore" : DataStore,
      "DeadLetterQueueUrl" : String,
      "DefaultEncryptionKey" : String,
      "DefaultExpirationDays" : Integer,
      "DomainName" : String,
      "Matching" : Matching,
      "RuleBasedMatching" : RuleBasedMatching,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CustomerProfiles::Domain
Properties:
  DataStore:
    DataStore
  DeadLetterQueueUrl: String
  DefaultEncryptionKey: String
  DefaultExpirationDays: Integer
  DomainName: String
  Matching:
    Matching
  RuleBasedMatching:
    RuleBasedMatching
  Tags:
    - Tag

```

## Properties

`DataStore`

Property description not available.

_Required_: No

_Type_: [DataStore](aws-properties-customerprofiles-domain-datastore.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeadLetterQueueUrl`

The URL of the SQS dead letter queue, which is used for reporting errors associated
with ingesting data from third party applications. You must set up a policy on the
`DeadLetterQueue` for the `SendMessage` operation to enable
Amazon Connect Customer Profiles to send messages to the
`DeadLetterQueue`.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultEncryptionKey`

The default encryption key, which is an AWS managed key, is used when
no specific type of encryption key is specified. It is used to encrypt all data before
it is placed in permanent or semi-permanent storage.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultExpirationDays`

The default number of days until the data within the domain expires.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1098`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The unique name of the domain.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Matching`

The process of matching duplicate profiles.

_Required_: No

_Type_: [Matching](aws-properties-customerprofiles-domain-matching.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleBasedMatching`

The process of matching duplicate profiles using Rule-Based matching.

_Required_: No

_Type_: [RuleBasedMatching](aws-properties-customerprofiles-domain-rulebasedmatching.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-customerprofiles-domain-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DomainName of the domain.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The timestamp of when the domain was created.

`LastUpdatedAt`

The timestamp of when the domain was most recently edited.

`RuleBasedMatching.Status`

Property description not available.

## Examples

The following example creates a Domain.

#### YAML

```yaml

Type: "AWS::CustomerProfiles::Domain"
Properties:
    DomainName: "ExampleDomain"
    DefaultEncryptionKey: "arn:aws:kms:us-east-1:123456789012:key/1988472d-6b77-4bb6-ae39-efce5EXAMPLE"
    DeadLetterQueueUrl: "arn:aws:sqs:us-east-1:123456789012:DLQName"
    DefaultExpirationDays: 6
```

#### JSON

```json

"Type": "AWS::CustomerProfiles::Domain",
"Properties": {
    "DomainName": "ExampleDomain",
    "DefaultEncryptionKey": "arn:aws:kms:us-east-1:123456789012:key/1988472d-6b77-4bb6-ae39-efce5EXAMPLE",
    "DeadLetterQueueUrl": "arn:aws:sqs:us-east-1:123456789012:DLQName",
    "DefaultExpirationDays": 6
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ValueRange

AttributeTypesSelector

All content copied from https://docs.aws.amazon.com/.
