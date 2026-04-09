# GroupConfigurationParameter

A parameter for a group configuration item. For details about group service
configuration syntax, see [Service configurations for resource\
groups](about-slg.md).

## Contents

###### Note

In the following list, the required parameters are described first.

**Name**

The name of the group configuration parameter. For the list of parameters that you can
use with each configuration item type, see [Supported resource types and\
parameters](about-slg.md#about-slg-types).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 80.

Pattern: `[a-z-]+`

Required: Yes

**Values**

The value or values to be used for the specified parameter. For the list of values you
can use with each parameter, see [Supported resource types and\
parameters](about-slg.md#about-slg-types).

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[a-zA-Z0-9:\/\._-]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/resource-groups-2017-11-27/groupconfigurationparameter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/resource-groups-2017-11-27/groupconfigurationparameter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/resource-groups-2017-11-27/groupconfigurationparameter.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GroupConfigurationItem

GroupFilter

All content copied from https://docs.aws.amazon.com/.
