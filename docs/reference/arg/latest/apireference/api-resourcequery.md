# ResourceQuery

The query you can use to define a resource group or a search for resources. A
`ResourceQuery` specifies both a query `Type` and a
`Query` string as JSON string objects. See the examples section for
example JSON strings. For more information about creating a resource group with a
resource query, see [Build queries and groups in\
AWS Resource Groups](../../../../services/arg/latest/userguide/gettingstarted-query.md) in the _AWS Resource Groups User Guide_

When you combine all of the elements together into a single string, any double quotes
that are embedded inside another double quote pair must be escaped by preceding the
embedded double quote with a backslash character (\\). For example, a complete
`ResourceQuery` parameter must be formatted like the following CLI
parameter example:

`--resource-query
                '{"Type":"TAG_FILTERS_1_0","Query":"{\"ResourceTypeFilters\":[\"AWS::AllSupported\"],\"TagFilters\":[{\"Key\":\"Stage\",\"Values\":[\"Test\"]}]}"}'`

In the preceding example, all of the double quote characters in the value part of the
`Query` element must be escaped because the value itself is surrounded by
double quotes. For more information, see [Quoting\
strings](../../../../services/cli/latest/userguide/cli-usage-parameters-quoting-strings.md) in the _AWS Command Line Interface User Guide_.

For the complete list of resource types that you can use in the array value for
`ResourceTypeFilters`, see [Resources\
you can use with AWS Resource Groups and Tag Editor](../../../../services/arg/latest/userguide/supported-resources.md) in the
_AWS Resource Groups User Guide_. For example:

`"ResourceTypeFilters":["AWS::S3::Bucket", "AWS::EC2::Instance"]`

## Contents

###### Note

In the following list, the required parameters are described first.

**Query**

The query that defines a group or a search. The contents depends on the value of the
`Type` element.

- `ResourceTypeFilters` – Applies to all
`ResourceQuery` objects of either `Type`. This element
contains one of the following two items:

- The value `AWS::AllSupported`. This causes the
ResourceQuery to match resources of any resource type that also match
the query.

- A list (a JSON array) of resource type identifiers that limit the
query to only resources of the specified types. For the complete list of
resource types that you can use in the array value for
`ResourceTypeFilters`, see [Resources you can use with AWS Resource Groups and Tag\
Editor](../../../../services/arg/latest/userguide/supported-resources.md) in the _AWS Resource Groups User Guide_.

Example: `"ResourceTypeFilters": ["AWS::AllSupported"]` or
`"ResourceTypeFilters": ["AWS::EC2::Instance",
                        "AWS::S3::Bucket"]`

- `TagFilters` – applicable only if `Type` =
`TAG_FILTERS_1_0`. The `Query` contains a JSON string
that represents a collection of simple tag filters. The JSON string uses a
syntax similar to the `
                             GetResources
                          ` operation, but uses only the `
                              ResourceTypeFilters
                          ` and `
                             TagFilters
                          ` fields. If you specify more than one tag key,
only resources that match all tag keys, and at least one value of each specified
tag key, are returned in your query. If you specify more than one value for a
tag key, a resource matches the filter if it has a tag key value that matches
_any_ of the specified values.

For example, consider the following sample query for resources that have two
tags, `Stage` and `Version`, with two values each:

`[{"Stage":["Test","Deploy"]},{"Version":["1","2"]}]`

The results of this resource query could include the following.

- An Amazon EC2 instance that has the following two tags:
`{"Stage":"Deploy"}`, and
`{"Version":"2"}`

- An S3 bucket that has the following two tags:
`{"Stage":"Test"}`, and
`{"Version":"1"}`

The resource query results would _not_ include the
following items in the results, however.

- An Amazon EC2 instance that has only the following tag:
`{"Stage":"Deploy"}`.

The instance does not have **all** of the
tag keys specified in the filter, so it is excluded from the
results.

- An RDS database that has the following two tags:
`{"Stage":"Archived"}` and
`{"Version":"4"}`

The database has all of the tag keys, but none of those keys has an
associated value that matches at least one of the specified values in
the filter.

Example: `"TagFilters": [ { "Key": "Stage", "Values": [ "Gamma", "Beta" ]
                        }`

- `StackIdentifier` – applicable only if `Type` =
`CLOUDFORMATION_STACK_1_0`. The value of this parameter is the
Amazon Resource Name (ARN) of the CloudFormation stack whose resources you want included
in the group.

Type: String

Length Constraints: Maximum length of 4096.

Pattern: `[\s\S]*`

Required: Yes

**Type**

The type of the query to perform. This can have one of two values:

- _`CLOUDFORMATION_STACK_1_0:`_ Specifies that you
want the group to contain the members of an CloudFormation stack. The `Query`
contains a `StackIdentifier` element with an Amazon resource name (ARN) for a CloudFormation
stack.

- _`TAG_FILTERS_1_0:`_ Specifies that you want the
group to include resource that have tags that match the query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `^\w+$`

Valid Values: `TAG_FILTERS_1_0 | CLOUDFORMATION_STACK_1_0`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/resource-groups-2017-11-27/resourcequery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/resource-groups-2017-11-27/resourcequery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/resource-groups-2017-11-27/resourcequery.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceIdentifier

ResourceStatus

All content copied from https://docs.aws.amazon.com/.
