AWS Audit Manager will no longer be open to new customers starting
April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers
can continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](audit-manager-availability-change.md).

# Confirming the status of evidence finder

After you submit your request to enable evidence finder, it takes up to 10 minutes to
enable the feature and create an event data store. As soon as the event data store is
created, all new evidence is ingested into the event data store moving forward.

When evidence finder is enabled and the event data store is created, we backfill the
newly created event data store with up to two years’ worth of your past evidence. This
process happens automatically and takes up to seven days to complete.

Follow the steps on this page to check and understand the status of your request to
enable evidence finder.

## Prerequisites

Make sure that you followed the steps to enable evidence finder. For instructions,
see [Enabling evidence finder](evidence-finder-settings-enable.md).

## Procedure

You can check the current status of evidence finder using the Audit Manager console, the
AWS CLI, or the Audit Manager API.

Audit Manager console

###### To see the current status of evidence finder on the Audit Manager console

1. Open the AWS Audit Manager console at [https://console.aws.amazon.com/auditmanager/home](https://console.aws.amazon.com/auditmanager/home).

2. In the left navigation pane, choose
    **Settings**.

3. Under **Enable evidence finder –**
**optional**, review the current status.

Each status is defined as follows:

StatusDescription

**Evidence finder isn't**
**enabled**

You haven't successfully enabled evidence
finder yet.

**You have requested to enable**
**evidence finder**

Your request is pending the event data store
being created.

**Evidence finder is**
**enabled**

The event data store was created. You can
now use evidence finder.

Depending how much evidence you have, it
takes up to seven days to backfill the new event
data store with your past evidence data. A blue
information panel indicates that the data backfill
is in progress. Feel free to start exploring
evidence finder in the meantime. However, keep in
mind that not all data is available until the
backfill is complete.

**You have requested to disable**
**evidence finder**

Your request is pending the event data store
being deleted.

**Evidence finder has been**
**disabled**

Evidence finder has been permanently
disabled and the event data store is deleted.

AWS CLI

###### To see the current status of evidence finder in the AWS CLI

Run the [get-settings](https://docs.aws.amazon.com/cli/latest/reference/auditmanager/get-settings.html) command with the `--attribute`
parameter set to `EVIDENCE_FINDER_ENABLEMENT`.

```

aws auditmanager get-settings --attribute EVIDENCE_FINDER_ENABLEMENT
```

This returns the following information:

###### enablementStatus

This attribute shows the current status of evidence finder.

- `ENABLE_IN_PROGRESS` – You requested to
enable evidence finder. An event data store is currently being
created to support evidence finder queries.

- `ENABLED` – An event data store was created
and evidence finder is enabled. We recommend waiting seven days
until the event data store is backfilled with your past evidence
data. You can use evidence finder in the meantime, but not all
data is available until the backfill is complete.

- `DISABLE_IN_PROGRESS` – You requested to
disable evidence finder, and your request is pending the event
data store being deleted.

- `DISABLED` – You permanently disabled
evidence finder and the event data store is deleted. You can't
re-enable evidence finder after this point.

###### backfillStatus

This attribute shows the current status of the evidence data
backfill.

- `NOT_STARTED` – The backfill hasn’t started
yet.

- `IN_PROGRESS` – The backfill is in progress.
This takes up to seven days to complete, depending on the amount
of evidence data.

- `COMPLETED` – The backfill is complete. All
of your past evidence is now queryable.

Audit Manager API

###### To see the current status of evidence finder using the API

Call the [GetSettings](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_GetSettings.html) operation with the `attribute`
parameter set to `EVIDENCE_FINDER_ENABLEMENT`. This
returns the following information:

###### enablementStatus

This attribute shows the current status of evidence finder.

- `ENABLE_IN_PROGRESS` \- You requested to enable
evidence finder. An event data store is currently being created
to support evidence finder queries.

- `ENABLED` \- An event data store was created and
evidence finder is enabled. We recommend waiting seven days
until the event data store is backfilled with your past evidence
data. You can use evidence finder in the meantime, but not all
data is available until the backfill is complete.

- `DISABLE_IN_PROGRESS` \- You requested to disable
evidence finder, and your request is pending the deletion of the
event data store.

- `DISABLED` \- You permanently disabled evidence
finder and the event data store is deleted. You can't re-enable
evidence finder after this point.

###### backfillStatus

This attribute shows the current status of the evidence data
backfill.

- `NOT_STARTED` means that the backfill hasn’t
started yet.

- `IN_PROGRESS` means that the backfill is in
progress. This takes up to seven days to complete, depending on
the amount of evidence data.

- `COMPLETED` means that the backfill is complete.
All of your past evidence is now queryable.

For more information, see [evidenceFinderEnablement](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_EvidenceFinderEnablement.html) in the _Audit Manager API Reference_.

## Next steps

After evidence finder is successfully enabled, you can start using the feature. We
recommend waiting seven days until the event data store is backfilled with your past
evidence data. You can use evidence finder in the meantime, but not all data might
be available until the backfill is complete.

To get started with evidence finder, see [Searching for evidence in evidence finder](search-for-evidence-in-evidence-finder.md).

## Additional resources

- [Troubleshooting evidence finder issues](evidence-finder-issues.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Enabling evidence finder

Disabling evidence finder
