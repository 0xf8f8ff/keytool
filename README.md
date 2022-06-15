# KEYTOOL

A simple key generator. 

Compile: 

```
go install
```

Run:

```
keytool <filename>
```

to generate two keys: a master key and a passphrase. The master key will be encrypted with the passphrase.

The encrypted master key will be written to the provided file. If no file name in the command-line, will write the master key to the `key32` file in the working directory.

Will print the base 64 encoded passphrase to the standard output.