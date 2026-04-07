This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::KeySigningKey

The `AWS::Route53::KeySigningKey` resource creates a new key-signing key (KSK) in a hosted zone. The hosted zone ID is passed as a
parameter in the KSK properties. You can specify the properties of this KSK using the `Name`, `Status`, and
`KeyManagementServiceArn ` properties of the resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53::KeySigningKey",
  "Properties" : {
      "HostedZoneId" : String,
      "KeyManagementServiceArn" : String,
      "Name" : String,
      "Status" : String
    }
}

```

### YAML

```yaml

Type: AWS::Route53::KeySigningKey
Properties:
  HostedZoneId: String
  KeyManagementServiceArn: String
  Name: String
  Status: String

```

## Properties

`HostedZoneId`

The unique string (ID) that is used to identify a hosted zone. For example: `Z00001111A1ABCaaABC11`.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Z0-9]{1,32}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyManagementServiceArn`

The Amazon resource name (ARN) for a customer managed customer master key (CMK) in AWS Key Management Service (AWS KMS ). The `KeyManagementServiceArn` must be unique for each key-signing key (KSK) in a single hosted zone. For example: `arn:aws:kms:us-east-1:111122223333:key/111a2222-a11b-1ab1-2ab2-1ab21a2b3a111`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

A string used to identify a key-signing key (KSK). `Name` can include
numbers, letters, and underscores (\_). `Name` must be unique for each
key-signing key in the same hosted zone.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_]{3,128}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Status`

A string that represents the current key-signing key (KSK) status.

Status can have one of the following values:

ACTIVE

The KSK is being used for signing.

INACTIVE

The KSK is not being used for signing.

DELETING

The KSK is in the process of being deleted.

ACTION\_NEEDED

There is a problem with the KSK that requires you to take action to
resolve. For example, the customer managed key might have been deleted,
or the permissions for the customer managed key might have been
changed.

INTERNAL\_FAILURE

There was an error during a request. Before you can continue to work with
DNSSEC signing, including actions that involve this KSK, you must correct
the problem. For example, you may need to activate or deactivate the
KSK.

_Required_: Yes

_Type_: String

_Allowed values_: `ACTIVE | INACTIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a identifier that is based on both the hosted zone ID and the KSK name properties. For example:

`{ "Ref": "Z00001111A1ABCaaABC11|KSK1" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VPC

AWS::Route53::RecordSet
