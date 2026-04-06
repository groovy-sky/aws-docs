# Manage workgroups

In the Athena console at [https://console.aws.amazon.com/athena/](https://console.aws.amazon.com/athena/home), you can perform the following tasks:

StatementDescription[Create a workgroup](creating-workgroups.md)

Create a new workgroup.

[View the workgroup's details](https://docs.aws.amazon.com/athena/latest/ug/viewing-details-workgroups.html)View the workgroup's details, such as its name, description, data usage
limits, location of query results, expected query results bucket owner,
encryption, and control of objects written to the query results bucket. You
can also verify whether this workgroup enforces its settings, if
**Override client-side settings** is checked.[Specify a workgroup for queries](https://docs.aws.amazon.com/athena/latest/ug/specify-wkgroup-to-athena-in-which-to-run-queries.html)

Before you can run queries, you must specify to Athena which workgroup
to use. You must have permissions to the workgroup.

[Switch workgroups](https://docs.aws.amazon.com/athena/latest/ug/switching-workgroups.html)

Switch between workgroups to which you have access.

[Edit a workgroup](https://docs.aws.amazon.com/athena/latest/ug/editing-workgroups.html)Edit a workgroup and change its settings. You cannot change a workgroup's
name, but you can create a new workgroup with the same settings and a
different name.[Enable or disable a workgroup](https://docs.aws.amazon.com/athena/latest/ug/workgroups-enabled-disabled.html)

Enable or disable a workgroup. When a workgroup is disabled, its users
cannot run queries, or create new named queries. If you have access to
it, you can still view metrics, data usage limit controls, workgroup's
settings, query history, and saved queries.

[Copy a saved query between workgroups](https://docs.aws.amazon.com/athena/latest/ug/copy-a-query-between-workgroups.html)Copy a saved query between workgroups. You might want to do this if, for
example, you created a query in a preview workgroup and you want to make it
available in a nonpreview workgroup.[Delete a workgroup](https://docs.aws.amazon.com/athena/latest/ug/deleting-workgroups.html)

Delete a workgroup. If you delete a workgroup, query history, saved
queries, the workgroup's settings and per-query data limit controls are
deleted. The workgroup-wide data limit controls remain in CloudWatch, and you
can delete them individually.

The primary workgroup cannot be deleted.

[Use IAM policies to control workgroup access](workgroups-iam-policy.md)Use IAM policies to control workgroup access. For example workgroup
policies, see [Example workgroup policies](https://docs.aws.amazon.com/athena/latest/ug/example-policies-workgroup.html).[Create an Athena workgroup\
that uses IAM Identity Center authentication](https://docs.aws.amazon.com/athena/latest/ug/workgroups-identity-center.html)To use IAM Identity Center identities with Athena, you must create an IAM Identity Center enabled
workgroup. After you create the workgroup, you can use the IAM Identity Center console or
API to assign IAM Identity Center users or groups to the workgroup.[Configure minimum encryption for a workgroup](https://docs.aws.amazon.com/athena/latest/ug/workgroups-minimum-encryption.html)Enforce a minimal level of encryption in Amazon S3 for all query results from
the workgroup. Use this feature to ensure that query results are never
stored in an Amazon S3 bucket in an unencrypted state.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Override client-side settings

View the workgroup's details
