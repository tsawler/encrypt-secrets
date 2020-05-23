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
and `ssl` is the Postgres SSL setting (e.g. disable for postgres, false for mysql). You can optionally specify `-key`
as well, providing a 32 character encryption key. If none is provided, one will be generated.

The secret key is also written to stdout, and to a file called `urlSignerSecret.txt`

## Example

~~~bash
./encryptSecrets -u homestead -p 'secret' -db goblender -dbtype mysql -s false -key rHbaqmfdhmdrDDPIytYhwSRzcvpOesjZ
~~~


## After Running

Once the run is done, do the following:

1. Update the secret in `/etc/supervisor/conf.d/whatever.conf`
1. Run update.sh
1. Restart the process in supervisorctl