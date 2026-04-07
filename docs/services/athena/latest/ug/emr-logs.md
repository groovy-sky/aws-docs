# Query Amazon EMR logs

Amazon EMR and big data applications that run on Amazon EMR produce log files. Log files are written
to the [primary node](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-master-core-task-nodes.html),
and you can also configure Amazon EMR to archive log files to Amazon S3 automatically. You can use
Amazon Athena to query these logs to identify events and trends for applications and clusters.
For more information about the types of log files in Amazon EMR and saving them to Amazon S3, see
[View log\
files](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-manage-view-web-log-files.html) in the _Amazon EMR Management Guide_.

###### Topics

- [Create and query a basic table based on Amazon EMR log files](https://docs.aws.amazon.com/athena/latest/ug/emr-create-table.html)

- [Create and query a partitioned table based on Amazon EMR logs](https://docs.aws.amazon.com/athena/latest/ug/emr-create-table-partitioned.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Example CloudTrail log queries

Query a basic table
