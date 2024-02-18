# medium-simple-web-api

Source code for Medium Story series which looks at different language implementations of a simple Web API.

## Directory structure

| Name               | Description                                                                                                                          |
| ------------------ | ------------------------------------------------------------------------------------------------------------------------------------ |
| database-manifests | Contains the manifest files to setup the database instance. Make sure you add your own database secret into the `secrets.yaml` file. |
| employee-portal-go | Contains source of the go service.                                                                                                   |
| employee-portal-sb | Contains the source of the java service.                                                                                             |
| k6-reports         | Contains the k6 dashboard reports and output summary.                                                                                |
| performance-tests  | Contains the load test harness code.                                                                                                 |

## Notes
If you find any improvement that can be done to these services or tests, please provide a PR with your suggestions.