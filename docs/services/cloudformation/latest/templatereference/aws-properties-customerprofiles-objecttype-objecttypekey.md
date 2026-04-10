This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::ObjectType ObjectTypeKey

An object that defines the Key element of a ProfileObject. A Key is a special element
that can be used to search for a customer profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldNames" : [ String, ... ],
  "StandardIdentifiers" : [ String, ... ]
}

```

### YAML

```yaml

  FieldNames:
    - String
  StandardIdentifiers:
    - String

```

## Properties

`FieldNames`

The reference for the key name of the fields map.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StandardIdentifiers`

The types of keys that a ProfileObject can have. Each ProfileObject can have only 1
UNIQUE key but multiple PROFILE keys. PROFILE means that this key can be used to tie an
object to a PROFILE. UNIQUE means that it can be used to uniquely identify an object. If
a key a is marked as SECONDARY, it will be used to search for profiles after all other
PROFILE keys have been searched. A LOOKUP\_ONLY key is only used to match a profile but
is not persisted to be used for searching of the profile. A NEW\_ONLY key is only used if
the profile does not already exist before the object is ingested, otherwise it is only
used for matching objects to profiles.

_Required_: No

_Type_: Array of String

_Allowed values_: `PROFILE | UNIQUE | SECONDARY | LOOKUP_ONLY | NEW_ONLY | ASSET | CASE | ORDER | AIR_PREFERENCE | AIR_BOOKING | AIR_SEGMENT | HOTEL_PREFERENCE | HOTEL_STAY_REVENUE | HOTEL_RESERVATION | LOYALTY | LOYALTY_TRANSACTION | LOYALTY_PROMOTION | WEB_ANALYTICS | DEVICE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObjectTypeField

Tag

All content copied from https://docs.aws.amazon.com/.
