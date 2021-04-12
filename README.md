# go-integration

This is an example of integration test using docker-compose and github actions.

1. I created a simple client SFTP to connecto to  SFTP server
2. The test should validate the connexion between the client and the server.
3. I created a docker-compose file in which is declared the `sftp` service which contains the SFPT server  and the `tests` service in which are the integration tests
4. It is worth mentioning, test integrations using the tag `+build integration` to be executed. 
5. I created a pipe line with github actions to run the integration tests when is a PR to Master. 
