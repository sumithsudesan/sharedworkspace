Integrating Grafana with InfluxDB to visualize the data(metrics).


Configuring Grafana to Connect to InfluxDB
1. Access Grafana: Get the Grafana service NodePort or use port-forwarding to access Grafana's web UI (http://<Grafana-IP>:<NodePort>).
2. Add InfluxDB Data Source:
   - Log in to Grafana with the username and password
   - Navigate to Configuration (gear icon) > Data Sources.
   - Click Add data source and Choose InfluxDB.
   - Configure the following:
      - influxdb URL, token, DB 

3. RBAC for grafana
   - Grafana needs the appropriate permissions to read data from InfluxDB. This is typically managed through InfluxDB’s own authentication and authorization mechanisms.