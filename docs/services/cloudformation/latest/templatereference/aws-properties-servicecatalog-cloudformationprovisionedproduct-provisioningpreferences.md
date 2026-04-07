This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::CloudFormationProvisionedProduct ProvisioningPreferences

The user-defined preferences that will be applied when updating a provisioned
product. Not all preferences are applicable to all provisioned product type

One or more AWS accounts that will have access to the provisioned product.

Applicable only to a `CFN_STACKSET` provisioned product type.

The AWS accounts specified should be within the list of accounts in the
`STACKSET` constraint. To get the list of accounts in the
`STACKSET` constraint, use the `DescribeProvisioningParameters`
operation.

If no values are specified, the default value is all accounts from the
`STACKSET` constraint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StackSetAccounts" : [ String, ... ],
  "StackSetFailureToleranceCount" : Integer,
  "StackSetFailureTolerancePercentage" : Integer,
  "StackSetMaxConcurrencyCount" : Integer,
  "StackSetMaxConcurrencyPercentage" : Integer,
  "StackSetOperationType" : String,
  "StackSetRegions" : [ String, ... ]
}

```

### YAML

```yaml

  StackSetAccounts:
    - String
  StackSetFailureToleranceCount: Integer
  StackSetFailureTolerancePercentage: Integer
  StackSetMaxConcurrencyCount: Integer
  StackSetMaxConcurrencyPercentage: Integer
  StackSetOperationType: String
  StackSetRegions:
    - String

```

## Properties

`StackSetAccounts`

One or more AWS accounts where the provisioned product will be available.

Applicable only to a `CFN_STACKSET` provisioned product type.

The specified accounts should be within the list of accounts from the `STACKSET` constraint. To get the list of accounts in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.

If no values are specified, the default value is all acounts from the `STACKSET` constraint.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackSetFailureToleranceCount`

The number of accounts, per Region, for which this operation can fail before AWS Service Catalog stops the operation in that Region. If the operation is stopped in a Region, AWS Service Catalog doesn't attempt the operation in any subsequent Regions.

Applicable only to a `CFN_STACKSET` provisioned product type.

Conditional: You must specify either `StackSetFailureToleranceCount` or `StackSetFailureTolerancePercentage`, but not both.

The default value is `0` if no value is specified.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackSetFailureTolerancePercentage`

The percentage of accounts, per Region, for which this stack operation can fail before AWS Service Catalog stops the operation in that Region. If the operation is stopped in a Region, AWS Service Catalog doesn't attempt the operation in any subsequent Regions.

When calculating the number of accounts based on the specified percentage, AWS Service Catalog rounds down to the next whole number.

Applicable only to a `CFN_STACKSET` provisioned product type.

Conditional: You must specify either `StackSetFailureToleranceCount` or `StackSetFailureTolerancePercentage`, but not both.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackSetMaxConcurrencyCount`

The maximum number of accounts in which to perform this operation at one time. This is dependent on the value of `StackSetFailureToleranceCount`. `StackSetMaxConcurrentCount` is at most one more than the `StackSetFailureToleranceCount`.

Note that this setting lets you specify the maximum for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.

Applicable only to a `CFN_STACKSET` provisioned product type.

Conditional: You must specify either `StackSetMaxConcurrentCount` or `StackSetMaxConcurrentPercentage`, but not both.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackSetMaxConcurrencyPercentage`

The maximum percentage of accounts in which to perform this operation at one time.

When calculating the number of accounts based on the specified percentage, AWS Service Catalog rounds down to the next whole number. This is true except in cases where rounding down would result is zero. In this case, AWS Service Catalog sets the number as `1` instead.

Note that this setting lets you specify the maximum for operations. For large deployments, under certain circumstances the actual number of accounts acted upon concurrently may be lower due to service throttling.

Applicable only to a `CFN_STACKSET` provisioned product type.

Conditional: You must specify either `StackSetMaxConcurrentCount` or `StackSetMaxConcurrentPercentage`, but not both.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackSetOperationType`

Determines what action AWS Service Catalog performs to a stack set or a stack instance represented by the provisioned product. The default value is `UPDATE` if nothing is specified.

Applicable only to a `CFN_STACKSET` provisioned product type.

CREATE

Creates a new stack instance in the stack set represented by the provisioned product. In this case, only new stack instances are created based on accounts and Regions; if new ProductId or ProvisioningArtifactID are passed, they will be ignored.

UPDATE

Updates the stack set represented by the provisioned product and also its stack instances.

DELETE

Deletes a stack instance in the stack set represented by the provisioned product.

_Required_: No

_Type_: String

_Allowed values_: `CREATE | UPDATE | DELETE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StackSetRegions`

One or more AWS Regions where the provisioned product will be available.

Applicable only to a `CFN_STACKSET` provisioned product type.

The specified Regions should be within the list of Regions from the `STACKSET` constraint. To get the list of Regions in the `STACKSET` constraint, use the `DescribeProvisioningParameters` operation.

If no values are specified, the default value is all Regions from the `STACKSET` constraint.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ProvisioningParameter

Tag
