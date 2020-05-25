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
./encryptSecrects -u username -p password -db databaseName -dbtype postgres -s ssl -port port
~~~

The flags:

~~~bash
% ./encryptSecrets -help                                            
Usage of ./encryptSecrets:
  -db string
        Database name
  -dbtype string
        Database type (postgres or mysql
  -key string
        Secret key (32 chars)
  -p string
        DB Password
  -port string
        Database type (postgres or mysql
  -s string
        SSL Settings (default "disable")
  -u string
        DB Username
~~~

`-key` is an optional 32 character encryption key. If none is provided, one will be generated.

The secret key is also written to stdout, and to a file called `urlSignerSecret.txt`

## Example

~~~bash
./encryptSecrets -u homestead -p 'secret' -db goblender -dbtype mysql -port 3306 -s false -key rHbaqmfdhmdrDDPIytYhwSRzcvpOesjZ
~~~


## After Running

Once the run is complete, do the following:

1. Update the secret in `/etc/supervisor/conf.d/whatever.conf`
1. Run update.sh
1. If necessary, add the key to supervisor configuration
1. Restart the process in supervisorctl