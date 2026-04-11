# Step 6: Query the data in the DynamoDB table

In this step, you will use HiveQL to query the Features table in DynamoDB. Try the
following Hive queries:

1. All of the feature types ( `feature_class`) in alphabetical
    order:

```nohighlight

SELECT DISTINCT feature_class
FROM ddb_features
ORDER BY feature_class;
```

2. All of the lakes that begin with the letter "M":

```nohighlight

SELECT feature_name, state_alpha
FROM ddb_features
WHERE feature_class = 'Lake'
AND feature_name LIKE 'M%'
ORDER BY feature_name;
```

3. States with at least three features higher than a mile (5,280
    feet):

```nohighlight

SELECT state_alpha, feature_class, COUNT(*)
FROM ddb_features
WHERE elev_in_ft > 5280
GROUP by state_alpha, feature_class
HAVING COUNT(*) >= 3
ORDER BY state_alpha, feature_class;
```

###### Next step

[Step 7: (Optional) clean up](emrfordynamodb-tutorial-cleanup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Step 5: Copy data to DynamoDB

Step 7: (Optional) clean up

All content copied from https://docs.aws.amazon.com/.
