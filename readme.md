# Encrypt Secrets

A one shot program to encrypt secrets for existing goblender apps


## Usage

1: Generate the binary:

~~~
env GOOS=linux GOARCH=amd64  go build -o encryptSecrets *.go
~~~

2: Backup the database

3: Copy the file to the **root directory** of the application.

4: Run with flags:

~~~bash
./encryptSecrects -u username -p password -db databaseName -dbtype postgres -s ssl 
~~~

where `username` and `password` are the db credentials, dbtype is the database type, -db is the datbase name, 
and `ssl` is the Postgres SSL setting (e.g. disable for postgres, false for mysql).

## Example

~~~bash
./encryptSecrets -u root -p marlow11 -dbName goblender -dbtype mysql -ssl false
~~~