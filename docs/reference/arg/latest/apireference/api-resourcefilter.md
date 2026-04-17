---
title: "ResourceFilter"
---

# ResourceFilter

A filter name and value pair that is used to obtain more specific results from a list
of resources.

## Contents

###### Note

In the following list, the required parameters are described first.

**Name**

The name of the filter. Filter names are case-sensitive.

Type: String

Valid Values: `resource-type`

Required: Yes

**Values**

One or more filter values. Allowed filter values vary by resource filter name, and are
case-sensitive.

Type: Array of strings

Array Members: Minimum number of 1 item. Maximum number of 5 items.

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `AWS::[a-zA-Z0-9]+::[a-zA-Z0-9]+`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/resource-groups-2017-11-27/ResourceFilter)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/resource-groups-2017-11-27/ResourceFilter)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/resource-groups-2017-11-27/ResourceFilter)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryError

ResourceIdentifier

All content copied from https://docs.aws.amazon.com/.
