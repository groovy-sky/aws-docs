# Deleting Amazon RDS zero-ETL integrations

When you delete a zero-ETL integration, Amazon RDS removes it from the source database. Your
transactional data isn't deleted from Amazon RDS or the analytics destination, but Amazon RDS doesn't send new data to
Amazon Redshift or Amazon SageMaker.

You can only delete an integration when it has a status of `Active`,
`Failed`, `Syncing`, or `Needs attention`.

You can delete zero-ETL integrations using the AWS Management Console, the AWS CLI, or the RDS API.

###### To delete a zero-ETL integration

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. From the left navigation pane, choose **Zero-ETL integrations**.

3. Select the zero-ETL integration that you want to delete.

4. Choose **Actions**, **Delete**, and
    confirm deletion.

To delete a zero-ETL integration, use the [delete-integration](../../../cli/latest/reference/rds/delete-integration.md)
command and specify the `--integration-identifier` option.

###### Example

For Linux, macOS, or Unix:

```nohighlight

aws rds delete-integration \
    --integration-identifier ee605691-6c47-48e8-8622-83f99b1af374
```

For Windows:

```nohighlight

aws rds delete-integration ^
    --integration-identifier ee605691-6c47-48e8-8622-83f99b1af374
```

To delete a zero-ETL integration using the Amazon RDS API, use the [`DeleteIntegration`](../../../../reference/amazonrds/latest/apireference/api-deleteintegration.md) operation with the
`IntegrationIdentifier` parameter.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Modifying zero-ETL integrations

Troubleshooting zero-ETL integrations

All content copied from https://docs.aws.amazon.com/.
