# Notes

The first packet header had a length of `0x4e` (118 in decimal). After the following 118 bytes, the next packet header was the following:
```
4098 d05797ab 04004c00  
00004c00 0000 
```

The third and fourth 4-byte values of the above are the same value `4c00 0000` which means that the length and untruncated length are the same (what we were looking for)! The number of bytes for the raw packet data that follows is `0x4c` which is 76 in decimal.
