This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table BillingMode

Determines the billing mode for the table - on-demand or provisioned.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mode" : String,
  "ProvisionedThroughput" : ProvisionedThroughput
}

```

### YAML

```yaml

  Mode: String
  ProvisionedThroughput:
    ProvisionedThroughput

```

## Properties

`Mode`

The billing mode for the table:

- On-demand mode - `ON_DEMAND`

- Provisioned mode - `PROVISIONED`

###### Note

If you choose `PROVISIONED` mode, then you also need to specify
provisioned throughput (read and write capacity) for the table.

Valid values: `ON_DEMAND` \| `PROVISIONED`

_Required_: Yes

_Type_: String

_Allowed values_: `PROVISIONED | ON_DEMAND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisionedThroughput`

The provisioned read capacity and write capacity for the table. For more information,
see [Provisioned throughput capacity mode](../../../keyspaces/latest/devguide/readwritecapacitymode.md#ReadWriteCapacityMode.Provisioned)
in the _Amazon Keyspaces Developer Guide_.

_Required_: No

_Type_: [ProvisionedThroughput](aws-properties-cassandra-table-provisionedthroughput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoScalingSpecification

CdcSpecification

All content copied from https://docs.aws.amazon.com/.
