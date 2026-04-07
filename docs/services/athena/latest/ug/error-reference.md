# Athena error catalog

Athena provides standardized error information to help you understand failed queries and
take steps after a query failure occurs. The `AthenaError` feature includes an
`ErrorCategory` field and an `ErrorType` field.
`ErrorCategory` specifies whether the cause of the failed query is due to
system error, user error, or other error. `ErrorType` provides more granular
information regarding the source of the failure. By combining the two fields, you can get a
better understanding of the circumstances surrounding and causes for the specific error that
occurred.

## Error category

The following table lists the Athena error category values and their meanings.

Error categorySource1SYSTEM2USER3OTHER

## Error type reference

The following table lists the Athena error type values and their meanings.

Error typeDescription0Query exhausted resources at this scale factor1Query exhausted resources at this scale factor2Query exhausted resources at this scale factor3Query exhausted resources at this scale factor4Query exhausted resources at this scale factor5Query exhausted resources at this scale factor6Query exhausted resources at this scale factor7Query exhausted resources at this scale factor8Query exhausted resources at this scale factor100Internal service error200Query engine had an internal error201Query engine had an internal error202Query engine had an internal error203Driver error204The metastore had an error205Query engine had an internal error206Query timed out207Query engine had an internal error208Query engine had an internal error209Failed to cancel query210Query timed out211Query engine had an internal error212Query engine had an internal error213Query engine had an internal error214Query engine had an internal error215Query engine had an internal error216Query engine had an internal error217Query engine had an internal error218Query engine had an internal error219Query engine had an internal error220Query engine had an internal error221Query engine had an internal error222Query engine had an internal error223Query engine had an internal error224Query engine had an internal error225Query engine had an internal error226Query engine had an internal error227Query engine had an internal error228Query engine had an internal error229Query engine had an internal error230Query engine had an internal error231Query engine had an internal error232Query engine had an internal error233Iceberg error234Lake Formation error235Query engine had an internal error236Query engine had an internal error237Serialization error238Failed to upload metadata to Amazon S3239General persistence error240Failed to submit query300Internal service error301Internal service error302Internal service error303Internal service error400Internal service error401Failed to write query results to Amazon S3402Failed to write query results to Amazon S31000User error1001Data error1002Data error1003DDL task failed1004Schema error1005Serialization error1006Syntax error1007Data error1008Query rejected1009Query failed1010Internal service error1011Query canceled by user1012Query engine had an internal error1013Query engine had an internal error1014Query canceled by user1100Invalid argument provided1101Invalid property provided1102Query engine had an internal error1103Invalid table property provided1104Query engine had an internal error1105Query engine had an internal error1106Invalid function argument provided1107Invalid view1108Failed to register function1109Provided Amazon S3 path not found1110Provided table or view does not exist1200Query not supported1201Provided decoder not supported1202Query type not supported1300General not found error1301General entity not found1302File not found1303Provided function or function implementation not found1304Query engine had an internal error1305Query engine had an internal error1306Amazon S3 bucket not found1307Selected engine not found1308Query engine had an internal error1400Throttling error1401Query failed due to AWS Glue throttling1402Query failed due to too many table versions in AWS Glue1403Query failed due to Amazon S3 throttling1404Query failed due to Amazon Athena throttling1405Query failed due to Amazon Athena throttling1406Query failed due to Amazon Athena throttling1500Permission error1501Amazon S3 permission error1602

Exceeded reserved capacity limit. Insufficient capacity to execute
this query.

1700Query failed due to a Lake Formation internal exception1701Query failed due to an AWS Glue internal exception9999Internal service error

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Troubleshoot issues

Code samples
