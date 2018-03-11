# openssl

##### 秘密鍵の作成

```
$ openssl genrsa -out private.key 2048
```

##### サーバ証明書要求の作成

```
$ openssl req -new -key private.key -out server.csr
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [XX]:JP
State or Province Name (full name) []:Tokyo
Locality Name (eg, city) [Default City]:xx
Organization Name (eg, company) [Default Company Ltd]:xx
Organizational Unit Name (eg, section) []:xx
Common Name (eg, your name or your server's hostname) []:xxxxxx
Email Address []:

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []:
An optional company name []:
```

### 証明書（通称オレオレ証明書）の作成

```
$ openssl x509 -req -in server.csr -signkey private.key -out server.crt
```
