This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KMS::Key

The `AWS::KMS::Key` resource specifies an [KMS key](../../../kms/latest/developerguide/concepts.md#kms_keys) in AWS Key Management Service. You can use this resource to create symmetric encryption KMS keys,
asymmetric KMS keys for encryption or signing, and symmetric HMAC KMS keys. You can use
`AWS::KMS::Key` to create [multi-Region\
primary keys](../../../kms/latest/developerguide/multi-region-keys-overview.md#mrk-primary-key) of all supported types. To replicate a multi-Region key, use the
`AWS::KMS::ReplicaKey` resource.

###### Important

If you change the value of the `KeySpec`, `KeyUsage`,
`Origin`, or `MultiRegion` properties of an existing KMS key, the
update request fails, regardless of the value of the [`UpdateReplacePolicy` attribute](../userguide/aws-attribute-updatereplacepolicy.md). This prevents you from accidentally
deleting a KMS key by changing any of its immutable property values.

###### Note

AWS KMS replaced the term _customer master key_
_(CMK)_ with _AWS KMS key_ and _KMS key_.
The concept has not changed. To prevent breaking changes, AWS KMS is keeping
some variations of this term.

You can use symmetric encryption KMS keys to encrypt and decrypt small amounts of data,
but they are more commonly used to generate data keys and data key pairs. You can also use a
symmetric encryption KMS key to encrypt data stored in AWS services that are [integrated with AWS KMS](https://aws.amazon.com/kms/features). For more information, see [Symmetric encryption KMS keys](../../../kms/latest/developerguide/concepts.md#symmetric-cmks) in the
_AWS Key Management Service Developer Guide_.

You can use asymmetric KMS keys to encrypt and decrypt data or sign messages and verify
signatures. To create an asymmetric key, you must specify an asymmetric `KeySpec`
value and a `KeyUsage` value. For details, see [Asymmetric keys in AWS KMS](../../../kms/latest/developerguide/symmetric-asymmetric.md) in the
_AWS Key Management Service Developer Guide_.

You can use HMAC KMS keys (which are also symmetric keys) to generate and verify
hash-based message authentication codes. To create an HMAC key, you must specify an HMAC
`KeySpec` value and a `KeyUsage` value of
`GENERATE_VERIFY_MAC`. For details, see [HMAC keys in AWS KMS](../../../kms/latest/developerguide/hmac.md) in the
_AWS Key Management Service Developer Guide_.

You can also create symmetric encryption, asymmetric, and HMAC multi-Region primary keys.
To create a multi-Region primary key, set the `MultiRegion` property to
`true`. For information about multi-Region keys, see [Multi-Region keys in AWS KMS](../../../kms/latest/developerguide/multi-region-keys-overview.md) in the _AWS Key Management Service Developer_
_Guide_.

You cannot use the `AWS::KMS::Key` resource to specify a KMS key with [imported key\
material](../../../kms/latest/developerguide/importing-keys.md) or a KMS key in a [custom key\
store](../../../kms/latest/developerguide/custom-key-store-overview.md).

**Regions**

AWS KMS CloudFormation resources are available in all Regions in which AWS KMS and CloudFormation are supported.
You can use the `AWS::KMS::Key` resource to create and manage all KMS key types that are supported in a Region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KMS::Key",
  "Properties" : {
      "BypassPolicyLockoutSafetyCheck" : Boolean,
      "Description" : String,
      "Enabled" : Boolean,
      "EnableKeyRotation" : Boolean,
      "KeyPolicy" : Json,
      "KeySpec" : String,
      "KeyUsage" : String,
      "MultiRegion" : Boolean,
      "Origin" : String,
      "PendingWindowInDays" : Integer,
      "RotationPeriodInDays" : Integer,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::KMS::Key
Properties:
  BypassPolicyLockoutSafetyCheck: Boolean
  Description: String
  Enabled: Boolean
  EnableKeyRotation: Boolean
  KeyPolicy: Json
  KeySpec: String
  KeyUsage: String
  MultiRegion: Boolean
  Origin: String
  PendingWindowInDays: Integer
  RotationPeriodInDays: Integer
  Tags:
    - Tag

```

## Properties

`BypassPolicyLockoutSafetyCheck`

Skips ("bypasses") the key policy lockout safety check. The default value is false.

###### Important

Setting this value to true increases the risk that the KMS key becomes unmanageable. Do not set this value to true indiscriminately.

For more information, see
[Default key policy](../../../kms/latest/developerguide/key-policy-default.md#prevent-unmanageable-key)
in the _AWS Key Management Service Developer Guide_.

Use this parameter only when you intend to prevent the principal that is making the request from making a subsequent
[PutKeyPolicy](../../../../reference/kms/latest/apireference/api-putkeypolicy.md) request on the KMS key.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the KMS key. Use a description that helps you to distinguish this KMS key from
others in the account, such as its intended use.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Specifies whether the KMS key is enabled. Disabled KMS keys cannot be used in cryptographic
operations.

When `Enabled` is `true`, the _key state_ of the
KMS key is `Enabled`. When `Enabled` is `false`, the key state of
the KMS key is `Disabled`. The default value is `true`.

The actual key state of the KMS key might be affected by actions taken outside of
CloudFormation, such as running the [EnableKey](../../../../reference/kms/latest/apireference/api-enablekey.md), [DisableKey](../../../../reference/kms/latest/apireference/api-disablekey.md),
or [ScheduleKeyDeletion](../../../../reference/kms/latest/apireference/api-schedulekeydeletion.md) operations.

For information about the key states of a KMS key, see [Key state: Effect on your KMS key](../../../kms/latest/developerguide/key-state.md) in the
_AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableKeyRotation`

Enables automatic rotation of the key material for the specified KMS key. By default,
automatic key rotation is not enabled.

AWS KMS supports automatic rotation only for symmetric encryption KMS keys ( `KeySpec` = `SYMMETRIC_DEFAULT`). For asymmetric KMS keys, HMAC KMS
keys, and KMS keys with Origin `EXTERNAL`, omit the `EnableKeyRotation`
property or set it to `false`.

To enable automatic key rotation of the key material for a multi-Region KMS key, set `EnableKeyRotation` to `true` on the primary key (created by using `AWS::KMS::Key`).
AWS KMS copies the rotation status to all replica keys. For details, see [Rotating multi-Region keys](../../../kms/latest/developerguide/multi-region-keys-manage.md#multi-region-rotate) in the _AWS Key Management Service Developer Guide_.

When you enable automatic rotation, AWS KMS automatically creates new key
material for the KMS key one year after the enable date and every year
thereafter. AWS KMS retains all key material until you delete the KMS key. For
detailed information about automatic key rotation, see [Rotating KMS keys](../../../kms/latest/developerguide/rotate-keys.md) in the
_AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyPolicy`

The key policy to attach to the KMS key.

If you provide a key policy, it must meet the following criteria:

- The key policy must allow the caller to make a subsequent [PutKeyPolicy](../../../../reference/kms/latest/apireference/api-putkeypolicy.md) request on the
KMS key. This reduces the risk that the KMS key becomes unmanageable. For more information, see
[Default key policy](../../../kms/latest/developerguide/key-policies.md#key-policy-default-allow-root-enable-iam) in the
_AWS Key Management Service Developer Guide_. (To omit this condition, set `BypassPolicyLockoutSafetyCheck` to true.)

- Each statement in the key policy must contain one or more principals. The principals
in the key policy must exist and be visible to AWS KMS. When you create a new
AWS principal (for example, an IAM user or role), you might need to
enforce a delay before including the new principal in a key policy because the new
principal might not be immediately visible to AWS KMS. For more information,
see [Changes that I make are not always immediately visible](../../../iam/latest/userguide/troubleshoot-general.md#troubleshoot_general_eventual-consistency) in the _AWS Identity and Access Management User Guide_.

If you do not provide a key policy, AWS KMS attaches a default key policy to the KMS key. For more information, see [Default key policy](../../../kms/latest/developerguide/key-policies.md#key-policy-default) in the _AWS Key Management Service Developer_
_Guide_.

A key policy document can include only the following characters:

- Printable ASCII characters

- Printable characters in the Basic Latin and Latin-1 Supplement character set

- The tab ( `\u0009`), line feed ( `\u000A`), and carriage return ( `\u000D`) special characters

_Minimum_: `1`

_Maximum_: `32768`

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeySpec`

Specifies the type of KMS key to create. The default value, `SYMMETRIC_DEFAULT`,
creates a KMS key with a 256-bit symmetric key for encryption and decryption. In China Regions, `SYMMETRIC_DEFAULT`
creates a 128-bit symmetric key that uses SM4 encryption. You can't change the
`KeySpec` value after the KMS key is created. For help choosing a
key spec for your KMS key, see [Choosing a KMS key type](../../../kms/latest/developerguide/symm-asymm-choose.md) in the _AWS Key Management Service Developer_
_Guide_.

The `KeySpec` property determines
the type of key material in the KMS key and the algorithms that the KMS key supports. To further restrict the algorithms that can
be used with the KMS key, use a condition key in its key policy or IAM policy. For more
information, see [AWS KMS condition keys](../../../kms/latest/developerguide/policy-conditions.md#conditions-kms) in the _AWS Key Management Service Developer_
_Guide_.

###### Important

If you change the value of the `KeySpec` property on an existing KMS key, the update request fails, regardless of the value of the
[`UpdateReplacePolicy` attribute](../userguide/aws-attribute-updatereplacepolicy.md).
This prevents you from accidentally deleting a KMS key by changing an immutable property value.

###### Note

[AWS\
services that are integrated with AWS KMS](https://aws.amazon.com/kms/features) use symmetric encryption KMS keys to
protect your data. These services do not support encryption with asymmetric KMS keys. For help determining
whether a KMS key is asymmetric, see [Identifying asymmetric\
KMS keys](../../../kms/latest/developerguide/find-symm-asymm.md) in the _AWS Key Management Service Developer Guide_.

AWS KMS supports the following key specs for KMS keys:

- Symmetric encryption key (default)

- `SYMMETRIC_DEFAULT` (AES-256-GCM)

- HMAC keys (symmetric)

- `HMAC_224`

- `HMAC_256`

- `HMAC_384`

- `HMAC_512`

- Asymmetric RSA key pairs (encryption and decryption _or_ signing and verification)

- `RSA_2048`

- `RSA_3072`

- `RSA_4096`

- Asymmetric NIST-recommended elliptic curve key pairs (signing and verification _or_ deriving shared secrets)

- `ECC_NIST_P256` (secp256r1)

- `ECC_NIST_P384` (secp384r1)

- `ECC_NIST_P521` (secp521r1)

- `ECC_NIST_EDWARDS25519` (ed25519) - signing and verification only

- **Note:** For ECC\_NIST\_EDWARDS25519 KMS keys, the
ED25519\_SHA\_512 signing algorithm requires [`MessageType:RAW`](../../../../reference/kms/latest/apireference/api-sign.md#KMS-Sign-request-MessageType), while ED25519\_PH\_SHA\_512 requires [`MessageType:DIGEST`](../../../../reference/kms/latest/apireference/api-sign.md#KMS-Sign-request-MessageType). These message types cannot be used interchangeably.

- Other asymmetric elliptic curve key pairs (signing and verification)

- `ECC_SECG_P256K1` (secp256k1), commonly used for
cryptocurrencies.

- Asymmetric ML-DSA key pairs (signing and verification)

- `ML_DSA_44`

- `ML_DSA_65`

- `ML_DSA_87`

- SM2 key pairs (encryption and decryption _or_ signing and verification _or_ deriving shared secrets)

- `SM2` (China Regions only)

_Required_: No

_Type_: String

_Allowed values_: `SYMMETRIC_DEFAULT | RSA_2048 | RSA_3072 | RSA_4096 | ECC_NIST_P256 | ECC_NIST_P384 | ECC_NIST_P521 | ECC_SECG_P256K1 | HMAC_224 | HMAC_256 | HMAC_384 | HMAC_512 | SM2 | ML_DSA_44 | ML_DSA_65 | ML_DSA_87 | ECC_NIST_EDWARDS25519`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KeyUsage`

Determines the [cryptographic\
operations](../../../kms/latest/developerguide/concepts.md#cryptographic-operations) for which you can use the KMS key. The default value is
`ENCRYPT_DECRYPT`. This property is required for asymmetric KMS keys and HMAC KMS keys. You can't
change the `KeyUsage` value after the KMS key is created.

###### Important

If you change the value of the `KeyUsage` property on an existing KMS key,
the update request fails, regardless of the value of the [`UpdateReplacePolicy` attribute](../userguide/aws-attribute-updatereplacepolicy.md). This prevents you from accidentally
deleting a KMS key by changing an immutable property value.

Select only one valid value.

- For symmetric encryption KMS keys, omit the parameter or specify
`ENCRYPT_DECRYPT`.

- For HMAC KMS keys (symmetric), specify `GENERATE_VERIFY_MAC`.

- For asymmetric KMS keys with RSA key pairs, specify `ENCRYPT_DECRYPT` or
`SIGN_VERIFY`.

- For asymmetric KMS keys with NIST-recommended elliptic curve key pairs, specify
`SIGN_VERIFY` or `KEY_AGREEMENT`.

- For asymmetric KMS keys with `ECC_SECG_P256K1` key pairs, specify
`SIGN_VERIFY`.

- For asymmetric KMS keys with ML-DSA key pairs, specify
`SIGN_VERIFY`.

- For asymmetric KMS keys with SM2 key pairs (China Regions only), specify
`ENCRYPT_DECRYPT`, `SIGN_VERIFY`, or `KEY_AGREEMENT`.

_Required_: No

_Type_: String

_Allowed values_: `ENCRYPT_DECRYPT | SIGN_VERIFY | GENERATE_VERIFY_MAC | KEY_AGREEMENT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MultiRegion`

Creates a multi-Region primary key that you can replicate in other AWS Regions. You can't change the
`MultiRegion` value after the KMS key is created.

For a list of AWS Regions in which multi-Region keys are supported, see [Multi-Region keys in AWS KMS](../../../kms/latest/developerguide/multi-region-keys-overview.md) in the _AWS Key Management Service Developer Guide_.

###### Important

If you change the value of the `MultiRegion` property on an existing KMS key,
the update request fails, regardless of the value of the [`UpdateReplacePolicy` attribute](../userguide/aws-attribute-updatereplacepolicy.md). This prevents you from accidentally
deleting a KMS key by changing an immutable property value.

For a multi-Region key, set to this property to `true`. For a single-Region
key, omit this property or set it to `false`. The default value is
`false`.

_Multi-Region keys_ are an AWS KMS feature that lets you
create multiple interoperable KMS keys in different AWS Regions. Because these
KMS keys have the same key ID, key material, and other metadata, you can use them to encrypt data
in one AWS Region and decrypt it in a different AWS Region
without making a cross-Region call or exposing the plaintext data. For more information, see
[Multi-Region keys](../../../kms/latest/developerguide/multi-region-keys-overview.md) in the _AWS Key Management Service Developer_
_Guide_.

You can create a symmetric encryption, HMAC, or asymmetric multi-Region KMS key, and you can
create a multi-Region key with imported key material. However, you cannot create a
multi-Region key in a custom key store.

To create a replica of this primary key in a different AWS Region ,
create an [AWS::KMS::ReplicaKey](../userguide/aws-resource-kms-replicakey.md) resource in a CloudFormation stack in the replica Region.
Specify the key ARN of this primary key.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Origin`

The source of the key material for the KMS key. You cannot change the origin after you
create the KMS key. The default is `AWS_KMS`, which means that AWS KMS
creates the key material.

To [create a KMS key with no key\
material](../../../kms/latest/developerguide/importing-keys-create-cmk.md) (for imported key material), set this value to `EXTERNAL`. For
more information about importing key material into AWS KMS, see [Importing Key\
Material](../../../kms/latest/developerguide/importing-keys.md) in the _AWS Key Management Service Developer Guide_.

You can ignore `ENABLED` when Origin is `EXTERNAL`. When a KMS key
with Origin `EXTERNAL` is created, the key state is `PENDING_IMPORT` and
`ENABLED` is `false`. After you import the key material,
`ENABLED` updated to `true`. The KMS key can then be used for
Cryptographic Operations.

###### Note

- CloudFormation doesn't support creating an `Origin` parameter of the
`AWS_CLOUDHSM` or `EXTERNAL_KEY_STORE` values.

- `EXTERNAL` is not supported for ML-DSA keys.

_Required_: No

_Type_: String

_Allowed values_: `AWS_KMS | EXTERNAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PendingWindowInDays`

Specifies the number of days in the waiting period before AWS KMS deletes a
KMS key that has been removed from a CloudFormation stack. Enter a value between 7 and 30 days.
The default value is 30 days.

When you remove a KMS key from a CloudFormation stack, AWS KMS schedules the
KMS key for deletion and starts the mandatory waiting period. The `PendingWindowInDays`
property determines the length of waiting period. During the waiting period, the key state of
KMS key is `Pending Deletion` or `Pending Replica Deletion`, which prevents
the KMS key from being used in cryptographic operations. When the waiting period expires, AWS KMS permanently deletes the KMS key.

AWS KMS will not delete a [multi-Region primary\
key](../../../kms/latest/developerguide/multi-region-keys-overview.md) that has replica keys. If you remove a multi-Region primary key from a
CloudFormation stack, its key state changes to `PendingReplicaDeletion` so it
cannot be replicated or used in cryptographic operations. This state can persist indefinitely.
When the last of its replica keys is deleted, the key state of the primary key changes to
`PendingDeletion` and the waiting period specified by
`PendingWindowInDays` begins. When this waiting period expires, AWS KMS deletes the primary key. For details, see [Deleting multi-Region\
keys](../../../kms/latest/developerguide/multi-region-keys-delete.md) in the _AWS Key Management Service Developer Guide_.

You cannot use a CloudFormation template to cancel deletion of the KMS key after you remove it
from the stack, regardless of the waiting period. If you specify a KMS key in your template, even
one with the same name, CloudFormation creates a new KMS key. To cancel deletion of a KMS key, use the
AWS KMS console or the [CancelKeyDeletion](../../../../reference/kms/latest/apireference/api-cancelkeydeletion.md)
operation.

For information about the `Pending Deletion` and `Pending Replica
        Deletion` key states, see [Key state: Effect on your KMS key](../../../kms/latest/developerguide/key-state.md) in the
_AWS Key Management Service Developer Guide_. For more information about
deleting KMS keys, see the [ScheduleKeyDeletion](../../../../reference/kms/latest/apireference/api-schedulekeydeletion.md)
operation in the _AWS Key Management Service API Reference_ and [Deleting KMS\
keys](../../../kms/latest/developerguide/deleting-keys.md) in the _AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `7`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RotationPeriodInDays`

Specifies a custom period of time between each rotation date. If no
value is specified, the default value is 365 days.

The rotation period defines the number of days after you enable automatic key rotation
that AWS KMS will rotate your key material, and the number of days between each automatic
rotation thereafter.

You can use the [`kms:RotationPeriodInDays`](../../../kms/latest/developerguide/conditions-kms.md#conditions-kms-rotation-period-in-days) condition key to further constrain the
values that principals can specify in the `RotationPeriodInDays` parameter.

For more information about rotating KMS keys and automatic rotation, see [Rotating keys](../../../kms/latest/developerguide/rotate-keys.md) in the _AWS Key Management Service Developer Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `90`

_Maximum_: `2560`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Assigns one or more tags to the replica key.

###### Note

Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see
[ABAC for\
AWS KMS](../../../kms/latest/developerguide/abac.md) in the _AWS Key Management Service Developer_
_Guide_.

For information about tags in AWS KMS, see [Tagging keys](../../../kms/latest/developerguide/tagging-keys.md) in the
_AWS Key Management Service Developer Guide_. For information about tags
in CloudFormation, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-kms-key-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the key ID, such as
`1234abcd-12ab-34cd-56ef-1234567890ab`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the KMS key, such as
`arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab`.

For information about the key ARN of a KMS key, see [Key ARN](../../../kms/latest/developerguide/concepts.md#key-id-key-ARN) in the
_AWS Key Management Service Developer Guide_.

`KeyId`

The key ID of the KMS key, such as
`1234abcd-12ab-34cd-56ef-1234567890ab`.

For information about the key ID of a KMS key, see [Key ID](../../../kms/latest/developerguide/concepts.md#key-id-key-id) in the
_AWS Key Management Service Developer Guide_.

## Examples

- [Create a symmetric encryption KMS key](#aws-resource-kms-key--examples--Create_a_symmetric_encryption_KMS_key)

- [Create a symmetric encryption KMS key with a resource tag](#aws-resource-kms-key--examples--Create_a_symmetric_encryption_KMS_key_with_a_resource_tag)

- [Create an asymmetric KMS key](#aws-resource-kms-key--examples--Create_an_asymmetric_KMS_key)

- [Create an HMAC KMS key](#aws-resource-kms-key--examples--Create_an_HMAC_KMS_key)

- [Create a multi-Region primary key](#aws-resource-kms-key--examples--Create_a_multi-Region_primary_key)

### Create a symmetric encryption KMS key

The following example creates a symmetric encryption KMS key. The key policy for the
KMS key allows `Alice` to manage the key and allows `Bob` to view
the KMS key and use it in cryptographic operations. It also allows the AWS account (root) full access to the key. This prevents you from losing
control of the key if both `Alice` and `Bob` are deleted from the
account.

#### JSON

```json

"myKey": {
        "Type": "AWS::KMS::Key",
        "Properties": {
            "Description": "An example symmetric encryption KMS key",
            "EnableKeyRotation": true,
            "PendingWindowInDays": 20,
            "KeyPolicy": {
                "Version": "2012-10-17",
                "Id": "key-default-1",
                "Statement": [
                    {
                        "Sid": "Enable IAM User Permissions",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:root"
                        },
                        "Action": "kms:*",
                        "Resource": "*"
                    },
                    {
                        "Sid": "Allow administration of the key",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:user/Alice"
                        },
                        "Action": [
                            "kms:Create*",
                            "kms:Describe*",
                            "kms:Enable*",
                            "kms:List*",
                            "kms:Put*",
                            "kms:Update*",
                            "kms:Revoke*",
                            "kms:Disable*",
                            "kms:Get*",
                            "kms:Delete*",
                            "kms:ScheduleKeyDeletion",
                            "kms:CancelKeyDeletion"
                        ],
                        "Resource": "*"
                    },
                    {
                        "Sid": "Allow use of the key",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:user/Bob"
                        },
                        "Action": [
                            "kms:DescribeKey",
                            "kms:Encrypt",
                            "kms:Decrypt",
                            "kms:ReEncrypt*",
                            "kms:GenerateDataKey",
                            "kms:GenerateDataKeyWithoutPlaintext"
                        ],
                        "Resource": "*"
                    }
                ]
            }
        }
    }
```

#### YAML

```yaml

myKey:
  Type: 'AWS::KMS::Key'
  Properties:
    Description: An example symmetric encryption KMS key
    EnableKeyRotation: true
    PendingWindowInDays: 20
    KeyPolicy:
      Version: 2012-10-17
      Id: key-default-1
      Statement:
        - Sid: Enable IAM User Permissions
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:root'
          Action: 'kms:*'
          Resource: '*'
        - Sid: Allow administration of the key
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:user/Alice'
          Action:
            - 'kms:Create*'
            - 'kms:Describe*'
            - 'kms:Enable*'
            - 'kms:List*'
            - 'kms:Put*'
            - 'kms:Update*'
            - 'kms:Revoke*'
            - 'kms:Disable*'
            - 'kms:Get*'
            - 'kms:Delete*'
            - 'kms:ScheduleKeyDeletion'
            - 'kms:CancelKeyDeletion'
          Resource: '*'
        - Sid: Allow use of the key
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:user/Bob'
          Action:
            - 'kms:DescribeKey'
            - 'kms:Encrypt'
            - 'kms:Decrypt'
            - 'kms:ReEncrypt*'
            - 'kms:GenerateDataKey'
            - 'kms:GenerateDataKeyWithoutPlaintext'
          Resource: '*'
```

### Create a symmetric encryption KMS key with a resource tag

The following example creates a symmetric encryption KMS key with one resource
tag.

###### Note

Tagging or untagging a KMS key can allow or deny permission to the KMS key. For details, see
[ABAC for\
AWS KMS](../../../kms/latest/developerguide/abac.md) in the _AWS Key Management Service Developer_
_Guide_.

#### JSON

```json

"myKeyWithTag": {
        "Type": "AWS::KMS::Key",
        "Properties": {
            "KeyPolicy": {
                "Version": "2012-10-17",
                "Id": "key-default-1",
                "Statement": [
                    {
                        "Sid": "Enable IAM User Permissions",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": {
                                "Fn::Join": [
                                    "",
                                    [
                                        "arn:aws:iam::",
                                        {
                                            "Ref": "AWS::AccountId"
                                        },
                                        ":root"
                                    ]
                                ]
                            }
                        },
                        "Action": "kms:*",
                        "Resource": "*"
                    }
                ]
            },
            "Tags": [
                {
                    "Key": {
                        "Ref": "Key"
                    },
                    "Value": {
                        "Ref": "Value"
                    }
                }
            ]
        },
        "Parameters": {
            "Key": {
                "Type": "String"
            },
            "Value": {
                "Type": "String"
            }
        }
    }
```

#### YAML

```yaml

myKeyWithTag:
  Type: 'AWS::KMS::Key'
  Properties:
    KeyPolicy:
      Version: 2012-10-17
      Id: key-default-1
      Statement:
        - Sid: Enable IAM User Permissions
          Effect: Allow
          Principal:
            AWS: !Join
              - ''
              - - 'arn:aws:iam::'
                - !Ref 'AWS::AccountId'
                - ':root'
          Action: 'kms:*'
          Resource: '*'
    Tags:
      - Key: !Ref Key
        Value: !Ref Value
  Parameters:
    Key:
      Type: String
    Value:
      Type: String
```

### Create an asymmetric KMS key

The following example creates an RSA asymmetric KMS key for signing and verification. For
an asymmetric KMS key, you must specify `KeySpec` and `KeyUsage`
properties. The `EnableKeyRotation` property must be omitted or set to
`false`.

#### JSON

```json

"RSASigningKey": {
        "Type": "AWS::KMS::Key",
        "Properties": {
            "Description": "RSA-3072 asymmetric KMS key for signing and verification",
            "KeySpec": "RSA_3072",
            "KeyUsage": "SIGN_VERIFY",
            "KeyPolicy": {
                "Version": "2012-10-17",
                "Id": "key-default-1",
                "Statement": [
                    {
                        "Sid": "Enable IAM User Permissions",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:root"
                        },
                        "Action": "kms:*",
                        "Resource": "*"
                    },
                    {
                        "Sid": "Allow administration of the key",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:role/Admin"
                        },
                        "Action": [
                            "kms:Create*",
                            "kms:Describe*",
                            "kms:Enable*",
                            "kms:List*",
                            "kms:Put*",
                            "kms:Update*",
                            "kms:Revoke*",
                            "kms:Disable*",
                            "kms:Get*",
                            "kms:Delete*",
                            "kms:ScheduleKeyDeletion",
                            "kms:CancelKeyDeletion"
                        ],
                        "Resource": "*"
                    },
                    {
                        "Sid": "Allow use of the key",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:role/Developer"
                        },
                        "Action": [
                            "kms:Sign",
                            "kms:Verify",
                            "kms:DescribeKey"
                        ],
                        "Resource": "*"
                    }
                ]
            }
        }
    }
```

#### YAML

```yaml

RSASigningKey:
  Type: 'AWS::KMS::Key'
  Properties:
    Description: RSA-3072 asymmetric KMS key for signing and verification
    KeySpec: RSA_3072
    KeyUsage: SIGN_VERIFY
    KeyPolicy:
      Version: 2012-10-17
      Id: key-default-1
      Statement:
        - Sid: Enable IAM User Permissions
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:root'
          Action: 'kms:*'
          Resource: '*'
        - Sid: Allow administration of the key
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:role/Admin'
          Action:
            - 'kms:Create*'
            - 'kms:Describe*'
            - 'kms:Enable*'
            - 'kms:List*'
            - 'kms:Put*'
            - 'kms:Update*'
            - 'kms:Revoke*'
            - 'kms:Disable*'
            - 'kms:Get*'
            - 'kms:Delete*'
            - 'kms:ScheduleKeyDeletion'
            - 'kms:CancelKeyDeletion'
          Resource: '*'
        - Sid: Allow use of the key
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:role/Developer'
          Action:
            - 'kms:Sign'
            - 'kms:Verify'
            - 'kms:DescribeKey'
          Resource: '*'
```

### Create an HMAC KMS key

The following example creates an HMAC KMS key. For
an HMAC KMS key, you must specify an HMAC `KeySpec` and the GENERATE\_VERIFY\_MAC value for the `KeyUsage`
property. Omit the `EnableKeyRotation` property or it set to
`false`.

#### JSON

```json

{
    "HMACExampleKey": {
        "Type": "AWS::KMS::Key",
        "Properties": {
            "Description": "HMAC_384 key for tokens",
            "KeySpec": "HMAC_384",
            "KeyUsage": "GENERATE_VERIFY_MAC",
            "KeyPolicy": {
                "Version": "2012-10-17",
                "Id": "key-default-1",
                "Statement": [
                    {
                        "Sid": "Enable IAM User Permissions",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:root"
                        },
                        "Action": "kms:*",
                        "Resource": "*"
                    },
                    {
                        "Sid": "Allow administration of the key",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:role/Admin"
                        },
                        "Action": [
                            "kms:Create*",
                            "kms:Describe*",
                            "kms:Enable*",
                            "kms:List*",
                            "kms:Put*",
                            "kms:Update*",
                            "kms:Revoke*",
                            "kms:Disable*",
                            "kms:Get*",
                            "kms:Delete*",
                            "kms:ScheduleKeyDeletion",
                            "kms:CancelKeyDeletion"
                        ],
                        "Resource": "*"
                    },
                    {
                        "Sid": "Allow use of the key",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:role/Developer"
                        },
                        "Action": [
                            "kms:GenerateMac",
                            "kms:VerifyMac",
                            "kms:DescribeKey"
                        ],
                        "Resource": "*"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

HMACExampleKey:
  Type: 'AWS::KMS::Key'
  Properties:
    Description: HMAC_384 key for tokens
    KeySpec: HMAC_384
    KeyUsage: GENERATE_VERIFY_MAC
    KeyPolicy:
      Version: 2012-10-17
      Id: key-default-1
      Statement:
        - Sid: Enable IAM User Permissions
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:root'
          Action: 'kms:*'
          Resource: '*'
        - Sid: Allow administration of the key
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:role/Admin'
          Action:
            - 'kms:Create*'
            - 'kms:Describe*'
            - 'kms:Enable*'
            - 'kms:List*'
            - 'kms:Put*'
            - 'kms:Update*'
            - 'kms:Revoke*'
            - 'kms:Disable*'
            - 'kms:Get*'
            - 'kms:Delete*'
            - 'kms:ScheduleKeyDeletion'
            - 'kms:CancelKeyDeletion'
          Resource: '*'
        - Sid: Allow use of the key
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:role/Developer'
          Action:
            - 'kms:GenerateMac'
            - 'kms:VerifyMac'
            - 'kms:DescribeKey'
          Resource: '*'
```

### Create a multi-Region primary key

The following example creates a multi-Region primary key. This example key is a
symmetric encryption KMS key, but you can create multi-Region versions of asymmetric KMS
keys and HMAC KMS keys.

_Multi-Region keys_ are an AWS KMS feature that lets you create
multiple interoperable KMS keys in different AWS Regions. Because these KMS keys have the same key
ID, key material, and other metadata, you can use them to encrypt data in one AWS Region
and decrypt it in a different AWS Region without making a cross-Region call or exposing
the plaintext data. For more information, see [Multi-Region\
keys](../../../kms/latest/developerguide/multi-region-keys-overview.md) in the _AWS Key Management Service Developer_
_Guide_.

To replicate this primary key into a different AWS Region, use the [AWS::KMS::ReplicaKey](../userguide/aws-resource-kms-replica-key.md) CloudFormation resource.

#### JSON

```json

"myPrimaryKey": {
        "Type": "AWS::KMS::Key",
        "Properties": {
            "Description": "An example multi-Region primary key",
            "MultiRegion": true,
            "EnableKeyRotation": true,
            "PendingWindowInDays": 10,
            "KeyPolicy": {
                "Version": "2012-10-17",
                "Id": "key-default-1",
                "Statement": [
                    {
                        "Sid": "Enable IAM User Permissions",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:root"
                        },
                        "Action": "kms:*",
                        "Resource": "*"
                    },
                    {
                        "Sid": "Allow administration of the key",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:user/Alice"
                        },
                        "Action": [
                            "kms:ReplicateKey",
                            "kms:Create*",
                            "kms:Describe*",
                            "kms:Enable*",
                            "kms:List*",
                            "kms:Put*",
                            "kms:Update*",
                            "kms:Revoke*",
                            "kms:Disable*",
                            "kms:Get*",
                            "kms:Delete*",
                            "kms:ScheduleKeyDeletion",
                            "kms:CancelKeyDeletion"
                        ],
                        "Resource": "*"
                    },
                    {
                        "Sid": "Allow use of the key",
                        "Effect": "Allow",
                        "Principal": {
                            "AWS": "arn:aws:iam::111122223333:user/Bob"
                        },
                        "Action": [
                            "kms:DescribeKey",
                            "kms:Encrypt",
                            "kms:Decrypt",
                            "kms:ReEncrypt*",
                            "kms:GenerateDataKey",
                            "kms:GenerateDataKeyWithoutPlaintext"
                        ],
                        "Resource": "*"
                    }
                ]
            }
        }
    }
```

#### YAML

```yaml

myPrimaryKey:
  Type: 'AWS::KMS::Key'
  Properties:
    Description: An example multi-Region primary key
    MultiRegion: true
    EnableKeyRotation: true
    PendingWindowInDays: 10
    KeyPolicy:
      Version: 2012-10-17
      Id: key-default-1
      Statement:
        - Sid: Enable IAM User Permissions
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:root'
          Action: 'kms:*'
          Resource: '*'
        - Sid: Allow administration of the key
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:user/Alice'
          Action:
            - 'kms:ReplicateKey'
            - 'kms:Create*'
            - 'kms:Describe*'
            - 'kms:Enable*'
            - 'kms:List*'
            - 'kms:Put*'
            - 'kms:Update*'
            - 'kms:Revoke*'
            - 'kms:Disable*'
            - 'kms:Get*'
            - 'kms:Delete*'
            - 'kms:ScheduleKeyDeletion'
            - 'kms:CancelKeyDeletion'
          Resource: '*'
        - Sid: Allow use of the key
          Effect: Allow
          Principal:
            AWS: 'arn:aws:iam::111122223333:user/Bob'
          Action:
            - 'kms:DescribeKey'
            - 'kms:Encrypt'
            - 'kms:Decrypt'
            - 'kms:ReEncrypt*'
            - 'kms:GenerateDataKey'
            - 'kms:GenerateDataKeyWithoutPlaintext'
          Resource: '*'
```

## See also

- [AWS KMS keys](../../../kms/latest/developerguide/concepts.md#kms_keys) in the _AWS Key Management Service_
_Developer Guide_.

- [CreateKey](../../../../reference/kms/latest/apireference/api-createkey.md) in the _AWS Key Management Service API_
_Reference_.

- [Creating\
keys](../../../kms/latest/developerguide/create-keys.md) in the _AWS Key Management Service Developer_
_Guide_.

- [Creating\
asymmetric KMS keys](../../../kms/latest/developerguide/asymm-create-key.md) in the _AWS Key Management Service Developer_
_Guide_.

- [Creating\
HMAC KMS keys](../../../kms/latest/developerguide/hmac-create-key.md) in the _AWS Key Management Service Developer_
_Guide_.

- [Creating multi-Region primary keys](../../../kms/latest/developerguide/create-primary-keys.md) in the _AWS Key Management Service_
_Developer Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::KMS::Alias

Tag

All content copied from https://docs.aws.amazon.com/.
