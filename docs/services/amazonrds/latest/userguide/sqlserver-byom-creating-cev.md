---
title: "Creating and managing BYOM engine versions for RDS for SQL Server"
---

# Creating and managing BYOM engine versions for RDS for SQL Server

Creating a BYOM engine version requires a one-time setup (License Mobility verification, obtaining RTM media, and uploading it to S3). After setup, you can create engine versions using the Amazon RDS console or the AWS CLI.

## Before you begin

### Step 1: Complete Microsoft License Mobility verification

License Mobility through Software Assurance is required for BYOM. You must have active Software Assurance coverage on your SQL Server licenses before deploying them on AWS.

To complete the verification process:

1. **Review eligibility** Visit the [AWS License Mobility page](https://aws.amazon.com/windows/resources/licensemobility) to review requirements and eligibility criteria.

2. **Submit the verification form** Download and complete the [Microsoft License Mobility verification form](https://www.microsoft.com/en-us/licensing/licensing-programs/software-assurance-license-mobility.aspx). Submit the form to Microsoft through your licensing partner or directly.

3. **Wait for approval** Microsoft sends you a confirmation email when your verification is approved.

You can create BYOM instances only after you receive License Mobility approval from Microsoft. Keep records of your approval for compliance audits.

###### Important

AWS does not validate or enforce Microsoft licensing compliance. You are responsible for maintaining compliance with your Microsoft licensing agreements.

### Step 2: Obtain SQL Server installation media

Download the Release to Manufacturing (RTM) file for the SQL Server major version you want to use. The RTM file is the base installation media for a SQL Server major version and edition (for example, SQL Server 2019 Enterprise Edition).

You can obtain the RTM file from one of the following Microsoft sources:

- **Visual Studio subscription** If you have an active Visual Studio subscription, download the RTM file from the [subscriber downloads page](https://visualstudio.microsoft.com/subscriptions).

- **Microsoft 365 admin center** If you purchased products through volume licensing, download the RTM file from the [Microsoft 365 admin center Downloads page](https://admin.microsoft.com/).

###### Important

Download only the English language Core-based ISO file. Do not use the Server+CAL-based ISO file.

### Step 3: Upload installation media to Amazon S3

Upload the SQL Server RTM file to an Amazon S3 bucket in the same AWS Region and account where you plan to create the BYOM engine version. To upload the RTM file using the AWS CLI:

```bash

aws s3 cp SQLServer2022-x64-ENU-Enterprise.iso s3://my-sqlserver-media/ISOs/
```

###### Note

If you do not provide a Cumulative Update (CU) file during BYOM engine creation, Amazon RDS automatically downloads the required CU from Microsoft when you create the engine version.

## Creating a BYOM engine version

Use `describe-db-engine-versions` to verify which engine versions are eligible for creating a BYOM engine version. For supported versions, review [Bring Your Own Media (BYOM) for RDS for SQL Server](sqlserver-byom.md).

01. Sign in to the [AWS Management Console](https://console.aws.amazon.com/rds) and open the Amazon RDS console.

02. In the navigation pane, choose **Custom engine versions**.

03. From the dropdown, select **RDS for SQL Server**.

04. Choose **Create custom engine version**.

05. For **Engine type**, choose **SQL Server**.

06. For **Database management type**, choose **Amazon RDS**.

07. For **Edition**, choose one of the following:

- SQL Server Enterprise Edition

- SQL Server Standard Edition

08. For **Engine version**, choose the SQL Server version that matches your installation files (for example, SQL Server 2022 `16.00.4175.1.v1`).

09. For **Installation media**, enter the S3 URI that contains your SQL Server RTM media, or choose **Browse S3** to locate and select the file.

10. Choose **Create custom engine version**.

The console displays the new BYOM engine version with an initial status of **Pending validation**. The status progresses automatically. For more information, see [BYOM engine version lifecycle states](#sqlserver-byom-creating-cev.lifecycle).

Use the `create-custom-db-engine-version` command to create a BYOM engine version from your installation media. Amazon RDS automatically downloads the required cumulative update from Microsoft.

```bash

aws rds create-custom-db-engine-version \
    --engine sqlserver-ee \
    --engine-version 16.00.4175.1.v1 \
    --database-installation-files-s3-bucket-name my-sqlserver-media \
    --database-installation-files-s3-prefix "ISOs/" \
    --database-installation-files "SQLServer2022-x64-ENU-Enterprise.iso"
```

**Response**

```json

{
    "Engine": "sqlserver-ee",
    "MajorEngineVersion": "16.00",
    "EngineVersion": "16.00.4175.1.v1",
    "DatabaseInstallationFiles": [
        "SQLServer2022-x64-ENU-Enterprise.iso"
    ],
    "DBEngineDescription": "Microsoft SQL Server Enterprise Edition",
    "DBEngineVersionArn": "arn:aws:rds:us-east-1:123456789012:cev:sqlserver-ee/16.00.4175.1.v1/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111",
    ....
    "Status": "pending-validation",
    .....
}
```

**CLI parameters**

ParameterRequiredDescription`--engine`YesThe database engine type: `sqlserver-ee`, `sqlserver-se`.`--engine-version`YesThe target engine version (for example, `16.00.4175.1.v1`).`--database-installation-files-s3-bucket-name`YesThe name of the S3 bucket that contains your installation media. The bucket must be in the same AWS Region as the BYOM engine version.`--database-installation-files-s3-prefix`NoThe S3 key prefix (folder path) where the installation media files are stored.`--database-installation-files`YesThe file names of the installation media.

###### Note

When using the AWS CLI, the `--database-installation-files-s3-bucket-name`, `--database-installation-files-s3-prefix`, and `--database-installation-files` parameters are required only for the first BYOM engine version you create for a major version and edition. For subsequent minor versions within the same major version, only `--engine` and `--engine-version` are required.

## Describing BYOM engine versions

After you create a BYOM engine version, use `describe-db-engine-versions` to check its status and details. The output returns two records for the same engine version: one for License Included ( `DBEngineMediaType: None`) and one for BYOM ( `DBEngineMediaType: Customer Provided`).

**Example:**

```bash

aws rds describe-db-engine-versions \
    --engine sqlserver-ee \
    --engine-version 16.00.4175.1.v1 \
    --include-all \
    --output table \
    --query "DBEngineVersions[].{Engine: Engine, Version: EngineVersion, Status: Status, MediaType: DBEngineMediaType}"
```

**Response:**

```

-----------------------------------------------------------------------
|                      DescribeDBEngineVersions                       |
+--------------+---------------------+------------+-------------------+
|    Engine    |      MediaType      |  Status    |      Version      |
+--------------+---------------------+------------+-------------------+
|  sqlserver-ee|  None               |  available |  16.00.4175.1.v1  |
|  sqlserver-ee|  Customer Provided  |  available |  16.00.4175.1.v1  |
+--------------+---------------------+------------+-------------------+
```

## BYOM engine version lifecycle states

StateDescription`pending-validation`RDS accepted the request and queued it for processing.`validating`RDS is validating RTM media file.`creating`RDS is installing SQL Server from RTM media file.`available`The BYOM engine version is ready for instance creation.`incompatible_installation_media`The installation files are invalid or corrupted. This is a terminal state. Delete the BYOM engine version and retry with correct files.

## Modifying a BYOM engine version

You can modify a BYOM engine version using the AWS Management Console or the AWS CLI. You can modify the BYOM engine version description or its availability status. Your BYOM engine version has one of the following status values:

- **available** – You can use this BYOM engine version to create a new DB instance or upgrade a DB instance. This is the default status for a newly created BYOM engine version.

- **inactive** – You can't create or upgrade a DB instance with this BYOM engine version. You can't restore a DB snapshot to create a new DB instance with this BYOM engine version.

You can change the BYOM engine version status from `available` to `inactive` or from `inactive` to `available`. You might change the status to `inactive` to prevent the accidental use of a BYOM engine version or to make a discontinued BYOM engine version eligible for use again.

```bash

aws rds modify-custom-db-engine-version \
    --engine sqlserver-ee \
    --engine-version 16.00.4175.1.v1 \
    --status inactive
```

**Response**

```json

{
    "Engine": "sqlserver-ee",
    "MajorEngineVersion": "16.00",
    "EngineVersion": "16.00.4175.1.v1",
    ....
    "Status": "inactive",
    ....
}
```

###### Note

Existing instances running on a BYOM engine version that you set to inactive continue to operate normally. Only new instance creation and restore operations are affected.

## Deleting a BYOM engine version

Use `delete-custom-db-engine-version` to permanently remove a BYOM engine version. After deletion, the engine version returns to a single LI-only entry.

**Prerequisites for deletion**

Before you can delete a BYOM engine version, ensure the following:

- No DB instances are using the BYOM version.

- No DB snapshots reference the BYOM version.

If any of the RDS resources exist, the API fails the deletion request.

**Deleting a BYOM engine version (CLI)**

```bash

aws rds delete-custom-db-engine-version \
    --engine sqlserver-ee \
    --engine-version 16.00.4175.1.v1
```

**Response**

```json

{
    "Engine": "sqlserver-ee",
    "MajorEngineVersion": "16.00",
    "EngineVersion": "16.00.4175.1.v1",
    .....
    "Status": "deleting",
    ....
}
```

## Considerations

- You provide the RTM ISO file once per major version and edition. All minor versions within the same major version and edition will reuse the same installation media.

- When providing your SQL Server RTM for Bring Your Own Media, you must only use the Core-based RTM ISO and not the Server + CAL based RTM ISO file.

## Next steps

After your BYOM engine version reaches `available` status, you can launch an RDS for SQL Server instance using it. See [Creating a BYOM DB instance for RDS for SQL Server](sqlserver-byom-creating-instance.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Bring Your Own Media (BYOM)

Creating a BYOM DB instance

All content copied from https://docs.aws.amazon.com/.
