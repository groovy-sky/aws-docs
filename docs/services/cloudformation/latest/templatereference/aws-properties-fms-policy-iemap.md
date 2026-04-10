This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FMS::Policy IEMap

Specifies the AWS account IDs and AWS Organizations organizational units (OUs) to include in or exclude from the policy.
Specifying an OU is the equivalent of specifying all accounts in the OU and in any of its child OUs, including any child OUs and accounts that are added at a later time.

This is used for the policy's `IncludeMap` and `ExcludeMap`.

You can specify account IDs, OUs, or a combination:

- Specify account IDs by setting the key to `ACCOUNT`. For example, the following is a valid map:
`{“ACCOUNT” : [“accountID1”, “accountID2”]}`.

- Specify OUs by setting the key to `ORGUNIT`. For example, the following is a valid map:
`{“ORGUNIT” : [“ouid111”, “ouid112”]}`.

- Specify accounts and OUs together in a single map, separated with a comma. For example, the following is a valid map:
`{“ACCOUNT” : [“accountID1”, “accountID2”], “ORGUNIT” : [“ouid111”, “ouid112”]}`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ACCOUNT" : [ String, ... ],
  "ORGUNIT" : [ String, ... ]
}

```

### YAML

```yaml

  ACCOUNT:
    - String
  ORGUNIT:
    - String

```

## Properties

`ACCOUNT`

The account list for the map.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ORGUNIT`

The organizational unit list for the map.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IcmpTypeCode

NetworkAclCommonPolicy

All content copied from https://docs.aws.amazon.com/.
