This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VerifiedPermissions::PolicyStore

Creates a policy store. A policy store is a container for policy resources. You can
create a separate policy store for each of your applications.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VerifiedPermissions::PolicyStore",
  "Properties" : {
      "DeletionProtection" : DeletionProtection,
      "Description" : String,
      "EncryptionSettings" : EncryptionSettings,
      "Schema" : SchemaDefinition,
      "Tags" : [ Tag, ... ],
      "ValidationSettings" : ValidationSettings
    }
}

```

### YAML

```yaml

Type: AWS::VerifiedPermissions::PolicyStore
Properties:
  DeletionProtection:
    DeletionProtection
  Description: String
  EncryptionSettings:
    EncryptionSettings
  Schema:
    SchemaDefinition
  Tags:
    - Tag
  ValidationSettings:
    ValidationSettings

```

## Properties

`DeletionProtection`

Specifies whether the policy store can be deleted. If enabled, the policy store can't
be deleted.

The default state is `DISABLED`.

_Required_: No

_Type_: [DeletionProtection](aws-properties-verifiedpermissions-policystore-deletionprotection.md)

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Descriptive text that you can provide to help with identification
of the current policy store.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionSettings`

A structure that contains the encryption configuration for the policy store and child resources.

This data type is used as a request parameter in the [CreatePolicyStore](../../../../reference/verifiedpermissions/latest/apireference/api-createpolicystore.md) operation.

_Required_: No

_Type_: [EncryptionSettings](aws-properties-verifiedpermissions-policystore-encryptionsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schema`

Creates or updates the policy schema in a policy store. Cedar can use the schema to
validate any Cedar policies and policy templates submitted to the policy store. Any
changes to the schema validate only policies and templates submitted after the schema
change. Existing policies and templates are not re-evaluated against the changed schema.
If you later update a policy, then it is evaluated against the new schema at that
time.

_Required_: No

_Type_: [SchemaDefinition](aws-properties-verifiedpermissions-policystore-schemadefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The list of key-value pairs to associate with the policy store.

_Required_: No

_Type_: Array of [Tag](aws-properties-verifiedpermissions-policystore-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidationSettings`

Specifies the validation setting for this policy store.

Currently, the only valid and required value is `Mode`.

###### Important

We recommend that you turn on `STRICT` mode only after you define a
schema. If a schema doesn't exist, then `STRICT` mode causes any policy
to fail validation, and Verified Permissions rejects the policy. You can turn off validation by
using the [UpdatePolicyStore](../../../../reference/verifiedpermissions/latest/apireference/api-updatepolicystore.md). Then, when you have a schema defined, use [UpdatePolicyStore](../../../../reference/verifiedpermissions/latest/apireference/api-updatepolicystore.md) again to turn validation back on.

_Required_: Yes

_Type_: [ValidationSettings](aws-properties-verifiedpermissions-policystore-validationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique id of the new or updated policy store. For
example:

`{ "Ref": "PSEXAMPLEabcdefg111111" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The [Amazon Resource Name (ARN)](../../../../general/latest/gr/aws-arns-and-namespaces.md) of the new or updated policy store.

`PolicyStoreId`

The unique ID of the new or updated policy store.

## Examples

### Creating a policy store with a schema and verification enabled

The following example creates a policy store that is configured with a schema
and policy validation against that schema turned on.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "AWS CloudFormation sample template for creating a policy store for Verified Permissions.",
    "Resources": {
        "MyPolicyStore": {
            "Type": "AWS::VerifiedPermissions::PolicyStore",
            "Properties": {
                "Schema": {
                    "CedarJson": "{\"PhotoApp\":{\"commonTypes\":{\"PersonType\":{\"type\":\"Record\",\"attributes\":{\"age\":{\"type\":\"Long\"},\"name\":{\"type\":\"String\"}}},\"ContextType\":{\"type\":\"Record\",\"attributes\":{\"ip\":{\"type\":\"Extension\",\"name\":\"ipaddr\",\"required\":false},\"authenticated\":{\"type\":\"Boolean\",\"required\":true}}}},\"entityTypes\":{\"User\":{\"shape\":{\"type\":\"Record\",\"attributes\":{\"userId\":{\"type\":\"String\"},\"personInformation\":{\"type\":\"PersonType\"}}},\"memberOfTypes\":[\"UserGroup\"]},\"UserGroup\":{\"shape\":{\"type\":\"Record\",\"attributes\":{}}},\"Photo\":{\"shape\":{\"type\":\"Record\",\"attributes\":{\"account\":{\"type\":\"Entity\",\"name\":\"Account\",\"required\":true},\"private\":{\"type\":\"Boolean\",\"required\":true}}},\"memberOfTypes\":[\"Album\",\"Account\"]},\"Album\":{\"shape\":{\"type\":\"Record\",\"attributes\":{}}},\"Account\":{\"shape\":{\"type\":\"Record\",\"attributes\":{}}}},\"actions\":{\"viewPhoto\":{\"appliesTo\":{\"principalTypes\":[\"User\",\"UserGroup\"],\"resourceTypes\":[\"Photo\"],\"context\":{\"type\":\"ContextType\"}}},\"createPhoto\":{\"appliesTo\":{\"principalTypes\":[\"User\",\"UserGroup\"],\"resourceTypes\":[\"Photo\"],\"context\":{\"type\":\"ContextType\"}}},\"listPhotos\":{\"appliesTo\":{\"principalTypes\":[\"User\",\"UserGroup\"],\"resourceTypes\":[\"Photo\"],\"context\":{\"type\":\"ContextType\"}}}}}}"
                },
                "ValidationSettings": {
                    "Mode": "STRICT"
                }
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: >-
  Description": "AWS CloudFormation sample template for creating a policy store for Verified Permissions."
Resources:
  MyPolicyStore:
    Type: AWS::VerifiedPermissions::PolicyStore
    Properties:
      Schema:
        CedarJson: '{"PhotoApp":{"commonTypes":{"PersonType":{"type":"Record","attributes":{"age":{"type":"Long"},"name":{"type":"String"}}},"ContextType":{"type":"Record","attributes":{"ip":{"type":"Extension","name":"ipaddr","required":false},"authenticated":{"type":"Boolean","required":true}}}},"entityTypes":{"User":{"shape":{"type":"Record","attributes":{"userId":{"type":"String"},"personInformation":{"type":"PersonType"}}},"memberOfTypes":["UserGroup"]},"UserGroup":{"shape":{"type":"Record","attributes":{}}},"Photo":{"shape":{"type":"Record","attributes":{"account":{"type":"Entity","name":"Account","required":true},"private":{"type":"Boolean","required":true}}},"memberOfTypes":["Album","Account"]},"Album":{"shape":{"type":"Record","attributes":{}}},"Account":{"shape":{"type":"Record","attributes":{}}}},"actions":{"viewPhoto":{"appliesTo":{"principalTypes":["User","UserGroup"],"resourceTypes":["Photo"],"context":{"type":"ContextType"}}},"createPhoto":{"appliesTo":{"principalTypes":["User","UserGroup"],"resourceTypes":["Photo"],"context":{"type":"ContextType"}}},"listPhotos":{"appliesTo":{"principalTypes":["User","UserGroup"],"resourceTypes":["Photo"],"context":{"type":"ContextType"}}}}}}'
      ValidationSettings:
        Mode: STRICT
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TemplateLinkedPolicyDefinition

DeletionProtection

All content copied from https://docs.aws.amazon.com/.
