# Using Amazon Q Developer with Amazon EMR Studio

This page describes how to set up and activate Amazon Q Developer for Amazon EMR Studio. Once
activated, Amazon Q can make code recommendations automatically as you write your ETL
code.

###### Note

Amazon Q supports Python, which can be used to code ETL scripts for Spark jobs in Amazon EMR Studio.

Use the following procedure to set up Amazon EMR Studio to work with Amazon Q.

1. Set up [Amazon EMR Studio Notebook](../../../emr/latest/managementguide/emr-managed-notebooks-create.md).

2. Attach the following policy to the IAM user role for Amazon EMR Studio Notebook.

###### Note

The `codewhisperer` prefix is a legacy name from a service that merged
with Amazon Q Developer. For more information, see
[Amazon Q Developer rename - Summary of changes](service-rename.md).

JSON

```json

{
       "Version":"2012-10-17",
       "Statement": [
           {
               "Sid": "AmazonQDeveloperPermissions",
               "Effect": "Allow",
               "Action": [
                   "codewhisperer:GenerateRecommendations"
               ],
               "Resource": "*"
           }
       ]
}

```

3. Open the [Amazon EMR console](https://console.aws.amazon.com/emr).

4. Under Amazon EMR Studio, choose **Workspaces (Notebooks).**

5. Select your desired Workspace and choose **Quick launch**.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JupyterLab

AWS Glue Studio

All content copied from https://docs.aws.amazon.com/.
