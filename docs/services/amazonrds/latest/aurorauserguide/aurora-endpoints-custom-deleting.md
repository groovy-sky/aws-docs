# Deleting a custom endpoint

Delete a custom endpoint using the AWS Management Console, AWS CLI, or the Amazon RDS API.

To delete a custom endpoint with the AWS Management Console, go to the cluster detail page,
select the appropriate custom endpoint, and select the **Delete**
action.

![Delete custom endpoint page.](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/AuroraDeleteCustomEndpoint.png)

To delete a custom endpoint with the AWS CLI, run the [delete-db-cluster-endpoint](../../../cli/latest/reference/rds/delete-db-cluster-endpoint.md) command.

The following command deletes a custom endpoint. You don't need to specify the
associated cluster, but you must specify the region.

For Linux, macOS, or Unix:

```nohighlight

aws rds delete-db-cluster-endpoint --db-cluster-endpoint-identifier custom-end-point-id \
  --region region_name

```

For Windows:

```nohighlight

aws rds delete-db-cluster-endpoint --db-cluster-endpoint-identifier custom-end-point-id ^
  --region region_name

```

To delete a custom endpoint with the RDS API, run the [DeleteDBClusterEndpoint](../../../../reference/amazonrds/latest/apireference/api-deletedbclusterendpoint.md)
operation.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Editing a custom
endpoint

AWS CLI examples for custom
endpoints

All content copied from https://docs.aws.amazon.com/.
