# ValidDBInstanceModificationsMessage

Information about valid modifications that you can make to your DB instance.
Contains the result of a successful call to the
`DescribeValidDBInstanceModifications` action.
You can use this information when you call
`ModifyDBInstance`.

## Contents

###### Note

In the following list, the required parameters are described first.

**AdditionalStorage**

The valid additional storage options for the DB instance.

Type: [ValidAdditionalStorageOptions](api-validadditionalstorageoptions.md) object

Required: No

**Storage.ValidStorageOptions.N**

Valid storage options for your DB instance.

Type: Array of [ValidStorageOptions](api-validstorageoptions.md) objects

Required: No

**SupportsDedicatedLogVolume**

Indicates whether a DB instance supports using a dedicated log volume (DLV).

Type: Boolean

Required: No

**ValidProcessorFeatures.AvailableProcessorFeature.N**

Valid processor features for your DB instance.

Type: Array of [AvailableProcessorFeature](api-availableprocessorfeature.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/validdbinstancemodificationsmessage.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/validdbinstancemodificationsmessage.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/validdbinstancemodificationsmessage.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ValidAdditionalStorageOptions

ValidStorageOptions

All content copied from https://docs.aws.amazon.com/.
