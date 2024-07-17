Overview
Collect metrics from different application in cluster.
Allow the client to save metrics data

HTTP service : '/v1/metrics'

Input JSON:

type Metric struct {
    Service string  `json:"service"`
    Metric  string  `json:"metric"`
    Value   float64 `json:"value"`
}

TODO: Quering from influxdb. Need to add support for - '/v1/query' GET

InfluxDB

The INFLUXDB_TOKEN is an authentication token used by InfluxDB to secure access to its API and resources. InfluxDB uses tokens to control permissions for reading and writing data to the database. When you set up InfluxDB, you create a token that is used to authenticate requests to the database.

How to Get an InfluxDB Token

Initial Setup:
    When you first set up InfluxDB 2.x, it guides you through an onboarding process where you create an initial user, organization, and bucket. During this process, an authentication token is generated.

Generate a New Token via InfluxDB UI:
    Log in to the InfluxDB web interface.
    Go to the Data section in the sidebar.
    Click on Tokens.
    Click on Generate Token and choose the appropriate permissions (e.g., read/write for your bucket).

Generate a New Token via InfluxDB CLI:
    Use the InfluxDB CLI (influx) to generate a new token:
        influx auth create --org <your-org> --read-buckets --write-buckets

    influx auth create --org <your-org> --read-buckets --write-buckets
    Replace <your-org> with your organization name. Adjust permissions as needed.