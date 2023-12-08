# 查看cert证书，cert证书是明文存储的，因为它仅仅包含public key, the identity information, and the signature这些允许所有人查看的信息 
- cd grpc-go/examples/data/x509
- openssl x509 -in ca_cert.pem -noout -text
- 输出：
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            d4:2f:c9:be:ed:80:b0:e4
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = US, ST = CA, L = SVL, O = gRPC, CN = test-server_ca
        Validity
            Not Before: Mar 18 21:44:56 2022 GMT
            Not After : Mar 15 21:44:56 2032 GMT
        Subject: C = US, ST = CA, L = SVL, O = gRPC, CN = test-server_ca
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (4096 bit)
                Modulus:
                    00:d1:a6:84:14:10:e5:fd:e7:e1:48:60:24:bb:17:
                    1e:1d:c2:13:77:a3:3b:1d:fa:a1:bf:87:9e:7c:bd:
                    19:bf:ce:aa:e0:38:14:c2:a3:4d:36:92:ae:12:88:
                    41:aa:6a:94:22:ce:0d:a8:b9:d1:72:92:44:12:5c:
                    82:cb:ec:a9:69:3d:51:d2:f6:df:37:fb:0b:30:0f:
                    9a:0e:42:9b:3a:56:5b:08:29:28:f4:ee:2b:95:be:
                    80:b0:35:fa:e3:0a:79:c0:5d:1f:98:2a:4a:cc:10:
                    dd:30:f5:7c:c7:e7:26:ea:bd:98:d7:c2:1a:c3:24:
                    a6:56:aa:4b:95:40:4c:d1:4f:56:7a:92:ee:eb:37:
                    39:72:84:21:39:81:69:03:f8:c0:ea:46:3f:53:e0:
                    a9:7e:ca:d3:c6:99:11:e0:32:40:96:93:4b:9b:a6:
                    e0:c3:65:04:b6:1a:fe:74:3d:59:c6:b4:52:4e:6d:
                    ce:83:63:13:a9:96:1f:b6:d3:bb:07:68:66:03:88:
                    c0:a1:66:05:b9:ac:0b:dc:bf:1f:3c:15:ff:34:6c:
                    2f:c6:29:fb:b1:70:7c:86:2e:67:98:fc:80:13:6a:
                    48:54:23:0c:59:26:e9:0e:d1:71:0f:e5:c8:71:e3:
                    f5:87:6e:3a:04:dc:88:81:55:83:1f:a3:f1:a6:6f:
                    16:a1:f5:db:b8:7c:58:86:fb:26:9a:72:d0:31:03:
                    24:4d:61:44:fe:a2:57:6b:95:a1:a5:31:85:75:ab:
                    09:62:c1:51:4f:a7:68:b8:b1:ed:22:7c:ae:ac:3c:
                    80:7c:11:35:67:cc:fc:bc:20:23:7a:c6:5d:dc:39:
                    58:81:d5:41:cb:b1:7a:f3:7a:a9:54:99:79:65:58:
                    02:50:20:7f:18:ca:b4:ce:5c:29:9c:8c:f9:6f:34:
                    09:27:55:28:49:2e:d5:70:34:b2:ae:a1:8d:a8:89:
                    76:8e:a1:a6:6f:01:2c:80:76:93:b2:d6:69:f7:39:
                    30:04:ac:b7:b5:d4:2e:4e:2e:0c:2f:ce:61:c1:33:
                    df:e0:b7:47:29:f5:a8:19:6d:af:58:00:df:ee:4a:
                    00:db:a1:8d:80:99:1f:83:f4:a5:7d:3c:23:53:29:
                    79:53:fa:87:c5:e1:6c:13:c4:66:0d:ad:c5:b5:46:
                    96:70:d7:10:5a:09:37:d1:20:e6:21:0a:44:1a:3f:
                    1c:a8:5f:94:ff:38:96:3b:41:4a:60:a7:b0:5e:8a:
                    ff:2d:ef:9d:18:ae:f8:15:4d:3f:92:8c:ab:5e:fd:
                    0c:da:84:cb:bb:52:ca:0c:e0:63:8d:66:de:94:85:
                    a2:37:6f:03:9f:4d:25:62:68:cb:df:cd:67:37:65:
                    53:89:f3
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Subject Key Identifier:
                5B:90:0C:5D:78:3F:CC:F4:9A:2C:7C:12:3B:FE:CB:C2:80:5E:65:85
            X509v3 Authority Key Identifier:
                keyid:5B:90:0C:5D:78:3F:CC:F4:9A:2C:7C:12:3B:FE:CB:C2:80:5E:65:85
                DirName:/C=US/ST=CA/L=SVL/O=gRPC/CN=test-server_ca
                serial:D4:2F:C9:BE:ED:80:B0:E4

            X509v3 Key Usage: critical
                Certificate Sign
    Signature Algorithm: sha256WithRSAEncryption
         a4:e1:39:f8:38:5a:8b:29:48:de:c6:83:fa:b7:6d:0a:9f:48:
         c4:f9:5d:31:42:ed:c8:f4:8f:77:16:0d:90:67:9a:a6:f4:1c:
         01:4f:9b:31:d3:fb:0d:e4:05:53:42:cb:1d:6b:89:ab:4e:85:
         a6:51:1f:b4:40:d2:f6:9c:65:04:f9:a4:3d:8c:19:4e:3c:ff:
         31:5f:c0:15:7f:a8:71:0d:41:6b:91:2c:3d:96:38:e1:14:dd:
         c8:e3:a0:83:41:72:57:cf:8a:6e:e6:97:fc:60:cf:40:02:3a:
         05:32:7f:ef:f6:86:44:bb:54:ba:b8:a7:d4:10:54:67:f6:4d:
         0b:7a:e4:36:40:f6:65:32:95:94:dd:b0:17:f8:01:f1:20:ec:
         85:ef:81:a6:29:ab:ab:58:98:d8:bb:7d:70:f3:47:f2:d3:f9:
         f3:dd:e1:e8:d0:01:1c:bd:53:bb:95:85:e7:32:1f:dc:d3:50:
         cc:e6:d8:5a:ff:a7:b1:e1:c9:d4:97:37:08:45:25:47:cf:5b:
         0e:d7:60:ed:25:c8:0f:bc:3f:15:12:00:a5:ce:7e:c4:f8:e1:
         63:6d:1c:07:91:9b:f1:a4:9a:7f:22:9c:17:5e:74:58:9a:ff:
         e3:3a:ee:4e:39:6f:fc:18:44:5d:75:43:3b:da:31:1d:f7:8e:
         6a:dd:65:6c:42:73:45:e3:83:ca:19:20:fe:d3:f4:c7:c6:15:
         42:c4:18:2a:41:dc:d6:38:62:39:d1:ae:27:7f:25:9f:de:76:
         13:d1:da:23:a2:12:27:a4:e2:a0:1a:c5:e2:27:7b:70:91:d9:
         da:82:33:e9:fe:15:5d:73:0a:ce:27:dd:2d:28:b3:8f:a4:92:
         79:d2:e5:a3:51:e7:06:69:aa:96:5f:1b:40:a8:17:25:84:f9:
         d1:39:d5:c1:9c:b9:bd:70:4f:5a:22:e5:7a:f1:83:f8:91:b5:
         e0:9f:77:7d:ca:cb:ae:bc:15:82:80:fe:d1:c7:76:a9:15:7c:
         17:e5:01:35:3a:88:eb:70:74:8c:ef:14:99:35:1f:59:94:c2:
         92:6b:1f:f6:8d:38:7c:d5:09:0e:4a:4a:97:1c:fa:45:dc:39:
         03:da:7d:31:c5:07:ac:cc:54:9e:8e:f9:59:bb:4a:ba:a2:48:
         43:4c:d9:bd:5f:04:ac:47:22:cc:4e:b7:0a:ca:bb:36:d1:85:
         4e:7e:78:02:02:ab:03:1d:ea:07:08:a5:3f:c9:49:1d:6b:f9:
         47:c7:55:ca:d7:6a:56:de:6b:39:93:d8:34:b9:08:53:9e:8a:
         78:f8:59:0b:ed:3d:83:9d:6d:8b:16:83:c6:3d:cc:63:78:2f:
         dc:c7:e8:6a:2f:33:5c:03


# server的证书server_cert.pem中可以看到Subject Alternative Name: DNS:*.test.example.com
- $ openssl x509 -in server_cert.pem -noout -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 1000 (0x3e8)
        Signature Algorithm: sha256WithRSAEncryption
        Issuer: C = US, ST = CA, L = SVL, O = gRPC, CN = test-server_ca
        Validity
            Not Before: Mar 18 21:44:58 2022 GMT
            Not After : Mar 15 21:44:58 2032 GMT
        Subject: C = US, ST = CA, L = SVL, O = gRPC, CN = test-server1 
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                RSA Public-Key: (4096 bit)
                Modulus:
                    00:be:46:05:6c:3e:a9:f5:f2:7a:57:a5:60:bf:d1:
                    0c:0f:c5:93:81:80:f0:39:5c:05:08:01:3a:30:5a:
                    5c:25:43:30:02:63:eb:7b:0d:fa:e6:ca:06:da:61:
                    59:ee:98:f5:25:8b:25:ad:a6:b7:c6:bf:65:34:19:
                    9b:64:79:14:f4:a9:f6:bc:1d:af:4e:14:42:09:8b:
                    d7:5c:21:0c:29:8e:fb:09:11:51:e4:d8:c2:ca:9c:
                    89:d1:07:47:03:b1:a1:cb:72:3b:e9:70:81:8c:3d:
                    f3:74:ff:7e:9f:37:aa:d3:52:e1:bc:3e:d6:42:70:
                    ac:bb:45:76:0a:24:29:df:54:18:8b:a0:b3:c0:53:
                    16:a1:3f:ec:2c:45:45:74:d8:b0:dc:bf:82:3d:29:
                    a7:8d:75:0e:d8:24:9e:5d:54:79:f3:dd:8d:ca:36:
                    42:89:fc:7f:70:e6:e6:20:df:03:25:0c:d6:82:fc:
                    f9:f9:d1:42:3f:eb:e8:fd:d7:0b:80:0e:88:1b:05:
                    dd:db:b1:88:1c:d4:a7:ce:d8:82:6c:03:39:a6:bc:
                    a5:a2:8e:c7:65:0d:2c:0f:de:3f:f6:8f:52:17:7b:
                    1e:4c:a0:a8:e5:fe:b5:cd:c7:f1:db:d8:38:da:3a:
                    96:1c:10:0f:52:6a:c8:6e:88:38:e5:9e:da:e7:f9:
                    61:72:8a:d9:3a:61:3c:f4:67:09:6e:10:20:40:f8:
                    3e:e5:d8:1e:32:2f:19:80:38:29:fd:16:99:bd:a1:
                    3e:4e:9b:c4:98:9d:36:82:e9:f6:5b:09:ca:a1:9f:
                    93:74:27:4f:93:d1:25:58:0e:d2:69:73:a3:f5:2c:
                    66:8e:fb:e6:3f:b2:ab:ce:f2:2d:ee:4e:2f:61:eb:
                    65:dd:4a:4d:f8:3b:48:a3:ad:30:9f:4d:ea:d9:b6:
                    1f:0d:a3:c3:2a:4a:cf:ec:76:8e:51:19:3a:05:c7:
                    99:86:f9:69:46:88:fb:8c:46:fb:d7:ea:73:cc:77:
                    49:ca:07:40:63:6e:8c:5d:ea:1c:ee:52:c6:75:e4:
                    b9:f3:54:49:56:d6:6e:fd:41:95:ef:29:7a:22:b2:
                    1b:a0:05:42:fb:5c:b3:b2:86:a4:4a:05:ce:d1:92:
                    95:a0:01:c1:2b:f1:3b:c2:77:68:93:c0:12:89:38:
                    38:8e:48:00:62:4d:cb:dd:97:68:13:0b:b7:2b:76:
                    ef:41:42:f0:9b:54:9c:bf:89:12:e6:d5:b7:cb:58:
                    e5:c0:9c:18:ae:c4:b5:11:db:e3:b9:8c:d0:af:07:
                    12:af:88:53:69:c0:ab:af:f6:f4:a6:b4:21:99:64:
                    6f:97:cf:a1:6c:d0:01:91:08:54:a8:0b:34:7b:78:
                    f7:a8:3d
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Basic Constraints: critical
                CA:FALSE
            X509v3 Subject Key Identifier:
                DC:77:E8:95:EC:08:8C:D5:58:23:B0:D4:44:ED:6B:19:4D:EF:D5:EB
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment, Key Agreement
            X509v3 Subject Alternative Name:
                DNS:*.test.example.com
    Signature Algorithm: sha256WithRSAEncryption
         01:04:66:59:fb:96:c4:bf:1c:8c:ac:34:f2:64:b5:cd:06:76:
         ed:e7:5f:50:5c:03:7c:17:87:de:58:4e:e8:1c:c5:24:9e:2c:
         bc:76:4b:29:17:9f:05:fe:70:99:dc:73:53:8a:6d:ba:0e:f4:
         30:39:99:9f:a1:b8:27:ac:18:83:8a:79:49:20:6b:34:19:2f:
         a1:40:77:0f:4f:40:a7:6a:2a:0d:d1:81:24:38:5e:12:bd:f7:
         47:cf:72:69:81:85:25:d8:db:88:cd:36:14:47:c4:6c:39:94:
         60:6b:0f:a5:f6:49:d5:b7:9b:ae:b2:63:39:fd:84:c5:cb:e3:
         2a:f0:78:63:bb:38:2e:1b:77:22:12:0e:b2:62:3f:e4:ad:ff:
         17:82:a0:87:17:3f:6b:34:ed:d1:3a:8e:5d:95:a2:b2:42:5a:
         7e:09:4b:12:9f:0e:13:cf:a2:16:4a:97:a7:4b:f6:7a:64:98:
         f7:29:cf:6c:d1:01:8a:68:e2:db:15:74:00:10:0c:14:6e:8b:
         8e:15:be:6d:5e:04:f5:28:ec:41:f6:50:e1:21:ff:72:6d:4d:
         e0:d4:a6:57:4c:f1:f0:da:28:4d:24:4d:ad:53:b7:1c:ec:d9:
         ed:09:9b:73:38:fc:53:0e:c6:a4:9e:b3:22:65:77:0e:78:17:
         8f:b8:9b:6e:1e:fe:24:14:81:c7:80:ad:83:05:54:54:e2:a6:
         0a:cb:7d:4b:91:bd:1c:71:ff:b0:fb:57:dd:a9:a9:3c:ae:0f:
         74:fb:f3:e2:9d:4e:aa:e5:c6:f0:93:3a:a4:87:84:51:34:e6:
         35:83:45:dc:a7:4b:a1:2f:ff:54:19:2e:59:d6:80:cc:21:40:
         09:e3:85:cd:4d:2a:4b:28:fa:76:a6:17:56:66:1c:7f:68:c2:
         24:d1:06:55:83:30:d7:07:30:34:5e:68:42:dd:42:61:54:be:
         eb:47:52:98:e2:76:dd:73:f8:6e:a5:cd:b3:f4:5c:1e:70:fb:
         4b:d6:e0:1b:13:d0:5a:9a:62:c1:5b:2b:97:ed:af:28:f5:fc:
         ad:15:72:da:d2:77:f9:1e:18:8d:2f:c6:b1:2b:03:22:c6:d8:
         1f:1b:74:af:3c:bb:6c:3b:0c:2b:97:8c:03:a9:f8:27:80:ec:
         75:26:93:dd:d3:b0:f1:00:86:bc:50:ea:54:79:d4:80:ba:2d:
         c2:b4:80:00:92:76:74:03:3c:6b:64:a1:15:f4:d0:34:ef:75:
         c3:a9:ce:69:12:81:a1:ab:fe:65:d3:aa:9d:a5:47:d2:2e:64:
         88:70:29:b8:52:de:50:ca:b0:06:7f:86:29:c1:c9:a0:70:2e:
         25:df:72:f6:6b:64:ec:32

# grpc使用tls教程
- https://dev.to/techschoolguru/how-to-secure-grpc-connection-with-ssl-tls-in-go-4ph