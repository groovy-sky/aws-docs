# Deleting records

The following procedure explains how to delete records using the Route 53 console. For information about how to
delete records using the Route 53 API, see
[ChangeResourceRecordSets](https://docs.aws.amazon.com/Route53/latest/APIReference/API_ChangeResourceRecordSets.html) in the _Amazon Route 53 API Reference_.

###### Note

Your changes to records take time to propagate to the Route 53 DNS servers.
Currently, the only way to verify that changes have propagated is to use the
[GetChange](https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetChange.html) API action. Changes generally propagate to all
Route 53 name servers within 60 seconds.

###### To delete records

1. Sign in to the AWS Management Console and open the Route 53 console at
    [https://console.aws.amazon.com/route53/](https://console.aws.amazon.com/route53).

2. On the Hosted Zones page, choose the row for the hosted zone that contains records that you want to delete.

3. In the list of records, select the record that you want to delete.

To select multiple, consecutive records, choose the first row, hold the **Shift** key,
    and choose the last row. To select multiple, nonconsecutive records, choose the first row, hold the
    **Ctrl** key, and choose additional rows.

You can't delete the records that have a value of **NS** or **SOA**
    for **Type**.

4. Choose **Delete**.

5. Choose **Delete** to close the dialog box.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Editing records

Listing records
