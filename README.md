# depoch

Convert epoch times from text on stdin into RFC3339 timestamps

## How To Use

If we take some log output like, you can see the epoch times after the `<info>` tag:

```
$ cat /var/log/syslog | tail -n 4
Mar 17 17:10:52 kohatu NetworkManager[997624]: <info>  [1742184652.2764] manager: (veth2b11ac5): new Veth device (/org/freedesktop/NetworkManager/Devices/59459)
Mar 17 17:10:52 kohatu NetworkManager[997624]: <info>  [1742184652.2853] device (vethf723e34): released from master device br-35c168b6c3a7
Mar 17 17:10:52 kohatu NetworkManager[997624]: <info>  [1742184652.7809] manager: (veth8ead586): new Veth device (/org/freedesktop/NetworkManager/Devices/59460)
Mar 17 17:10:52 kohatu NetworkManager[997624]: <info>  [1742184652.7892] device (veth311e632): released from master device br-db273219c5c3
```

Now if we pipe that through `depoch`:

```
$ cat /var/log/syslog | tail -n 4 | depoch
Mar 17 17:11:52 kohatu NetworkManager[997624]: <info>  [2025-03-17T17:11:52.9496] device (vethe72a0f7): carrier: link connected
Mar 17 17:11:52 kohatu NetworkManager[997624]: <info>  [2025-03-17T17:11:52.9499] device (br-db273219c5c3): carrier: link connected
Mar 17 17:11:57 kohatu NetworkManager[997624]: <info>  [2025-03-17T17:11:57.1027] manager: (veth0c4ed2e): new Veth device (/org/freedesktop/NetworkManager/Devices/59466)
Mar 17 17:11:57 kohatu NetworkManager[997624]: <info>  [2025-03-17T17:11:57.1122] device (vethe72a0f7): released from master device br-db273219c5c3
```

We can see it has converted the timestamps to RFC3339.
