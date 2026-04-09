# DeleteSubnetGroup

Deletes a subnet group.

###### Note

You cannot delete a subnet group if it is associated with any DAX
clusters.

## Request Syntax

```nohighlight

{
   "SubnetGroupName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[SubnetGroupName](#API_dax_DeleteSubnetGroup_RequestSyntax)**

The name of the subnet group to delete.

Type: String

Required: Yes

## Response Syntax

```nohighlight

{
   "DeletionMessage": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[DeletionMessage](#API_dax_DeleteSubnetGroup_ResponseSyntax)**

A user-specified message for this action (i.e., a reason for deleting the subnet
group).

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ServiceLinkedRoleNotFoundFault**

The specified service linked role (SLR) was not found.

HTTP Status Code: 400

**SubnetGroupInUseFault**

The specified subnet group is currently in use.

HTTP Status Code: 400

**SubnetGroupNotFoundFault**

The requested subnet group name does not refer to an existing subnet
group.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dax-2017-04-19/deletesubnetgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dax-2017-04-19/deletesubnetgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/deletesubnetgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dax-2017-04-19/deletesubnetgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/deletesubnetgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dax-2017-04-19/deletesubnetgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dax-2017-04-19/deletesubnetgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dax-2017-04-19/deletesubnetgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dax-2017-04-19/deletesubnetgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/deletesubnetgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteParameterGroup

DescribeClusters

All content copied from https://docs.aws.amazon.com/.
