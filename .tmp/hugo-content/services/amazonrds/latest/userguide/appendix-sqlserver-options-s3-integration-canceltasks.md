# Canceling a task

To cancel S3 integration tasks, use the `msdb.dbo.rds_cancel_task` stored procedure with the `task_id`
parameter. Delete and list tasks that are in progress can't be cancelled. The following example shows a request to cancel a
task.

```SQL

exec msdb.dbo.rds_cancel_task @task_id = 1234;
```

To get an overview of all tasks and their task IDs, use the `rds_fn_task_status` function as described in [Monitoring the status of a file transfer task](appendix-sqlserver-options-s3-integration-using-monitortasks.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Monitoring file transfers

Disabling S3 integration

All content copied from https://docs.aws.amazon.com/.
