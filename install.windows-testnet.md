# dump config
```
./build/galaxy --http  --http.addr '0.0.0.0' --http.corsdomain "*" --http.vhosts "*" --http.api="eth,debug,net,admin,web3,personal,txpool,axis,dag" --datadir=C:/works/axis-chain-galaxy/testnet --http.api "eth,debug,net,admin,web3,personal,txpool,axis,dag" --maxpeers 128 --maxpendpeers 128 --txpool.globalqueue 8192  dumpconfig > C:/works/axis-chain-galaxy/testnet/config.toml
```

# add validator
./build/galaxy --datadir C:/works/axis-chain-galaxy/testnet validator new

# make genesis
./build/galaxy inittest C:/works/axis-chain-galaxy/testnet/genesis.g

# create password file
1#	echo "71ec9927a0" > C:/works/axis-chain-galaxy/testnet/password.txt


# start validator
./build/galaxy --config C:/works/axis-chain-galaxy/testnet/config.toml --genesis C:/works/axis-chain-galaxy/testnet/genesis.g --datadir C:/works/axis-chain-galaxy/testnet  --validator.id 1 --validator.pubkey 0xc004cdbafa4f4d7d48f305a723b71b774fc45d54996acfe562f97867f6ae0cbaf21193cb553176a8833bef038284d73a9eecef996ddcfe5096ee08933e4ddffe6301 --validator.password C:/works/axis-chain-galaxy/testnet/password.txt








# rebuild and copy to binary repo
```
cd /root/axis-chain
git pull
make linux
cp ./build/galaxy ../axis-binary/galaxy-ubuntu-20.04-lts
cd /root/axis-binary
ls

```

# make a storage
```
rm -rf /root/galaxy
cd /root/galaxy
chmod 777 /root/galaxy
```

# make a validator
```

cd /root/axis-binary
/root/axis-binary/galaxy-ubuntu-20.04-lts --datadir /root/galaxy  validator new
```


# make genesis on ubuntu
```
/root/axis-binary/galaxy-ubuntu-20.04-lts init /root/axis-binary/genesis.g
rm -rf /root/galaxy/chaindata
rm -rf /root/galaxy/go-galaxy
```

# dump config
```
/root/axis-binary/galaxy-ubuntu-20.04-lts --http  --http.addr '0.0.0.0' --http.corsdomain "*" --http.vhosts "*" --http.api="eth,debug,net,admin,web3,personal,txpool,axis,dag" --datadir=/root/galaxy --http.api "eth,debug,net,admin,web3,personal,txpool,axis,dag" --maxpeers 128 --maxpendpeers 128 --txpool.globalqueue 8192  dumpconfig > /root/axis-binary/config.toml

```

# git push

```
cd /root/axis-binary
git add .
git commit -m "updated ubuntu 20.04 lts"
git push
```

# git clone on other rpc
```
git clone https://github.com/galaxy-digital/axis-binary.git

```

# create password file
```
echo "123456" > /root/galaxy/password.txt
```

# start validators 
```
/root/axis-binary/galaxy-ubuntu-20.04-lts --config /root/axis-binary/config.toml --genesis /root/axis-binary/genesis.g --datadir /root/galaxy  --validator.id 1 --validator.pubkey 0xc0047cf4039d716c107f389b2b7ae1bc3775cf6d4aad3b3ab3f663d8f34bef892fa6397e4dbd60c5205464db38d2578a8201813c281304018b102b7530fc5f4e16ee --validator.password /root/galaxy/password.txt

/root/axis-binary/galaxy-ubuntu-20.04-lts --config /root/axis-binary/config.toml --genesis /root/axis-binary/genesis.g --datadir /root/galaxy --validator.id 2 --validator.pubkey 0xc004d8597293ef427f5fd03a9ea33b4bd77de037eb43f46553a1f6f57bee71b1b913c3a3bac451bd02b552ce376a07c28b540723acd08d2284b98a90925425b8e88f --validator.password /root/galaxy/password.txt

/root/axis-binary/galaxy-ubuntu-20.04-lts --config /root/axis-binary/config.toml --genesis /root/axis-binary/genesis.g --datadir /root/galaxy --validator.id 3 --validator.pubkey 0xc004b0df54022cf14df39e3caba140a6c83ab17d98055dacd6048c41ecc5ae2ceff164787fa5e1585e2670bd43d9fec86c055e13bc944eded3f90e3bc844c9a4b18c --validator.password /root/galaxy/password.txt
```


## linux 
/root/galaxy/galaxy.ipc

```
/root/axis-binary/galaxy-ubuntu-20.04-lts --config /root/axis-binary/config.toml --genesis /root/axis-binary/genesis.g



/root/axis-binary/galaxy-ubuntu-20.04-lts --config /root/axis-binary/config.toml --genesis /root/axis-binary/genesis.g


/root/axis-binary/galaxy-ubuntu-20.04-lts --config /root/axis-binary/config.toml --genesis /root/axis-binary/genesis.g --validator.id 2 --validator.pubkey 0xc00484a101d22ac8b715683219bcea75c5b192821e6a9a5a6080eff983bfe0a4b8e5fbd1452300f9c9fae116987e030a80cd1fafd2e9bff2b2955af623dacf2fb6f2 --validator.password /home/axis/.galaxy/password.txt

/root/axis-binary/galaxy-ubuntu-20.04-lts --config /root/axis-binary/config.toml --genesis /root/axis-binary/genesis.g --validator.id 3 --validator.pubkey 0xc0047301efa2653128b50490d37bd7a58a1af88ba397b1f2a362487e1d3a02113d6450ce444ca39071867ce874172e8c465bf41bdf2f68ff5fc667eeae366df4377d --validator.password /home/axis/.galaxy/password.txt

```

# testnet
./build/galaxy inittest D:/works-axis/chaintest/testnet/genesis.g
## windows 
```
./build/galaxy --config D:/works-axis/chaintest/testnet/config.toml --genesis D:/works-axis/chaintest/testnet/testnet.g --datadir D:/works-axis/chaintest/testnet --validator.id 1 --validator.pubkey 0xc00439b9f3f5a56c6aa8c79e01094d496d4e5b0b2116f6e26790177fb7639ffdf473ed428b71eec45e9789e3210cd46e663b9852d2f58ce7070bf1c928ace37d904a --validator.password D:/works-axis/chaintest/testnet/password-1.txt

./build/galaxy --config D:/works-axis/chaintest/testnet/config.toml --genesis D:/works-axis/chaintest/testnet/testnet.g --datadir D:/works-axis/chaintest/testnet --validator.id 2 --validator.pubkey 0xc004be3ddfc6d48ad5d0ab793968f30e412cfaf0a1e1bdf3af63f542c4082191349ef8d2d13a1a1ce2b1512526f1ff53bbb8365d7f2e953b64fcc8cc93ce6ab60d9d --validator.password D:/works-axis/chaintest/testnet/password-2.txt

./build/galaxy --config D:/works-axis/chaintest/testnet/config.toml --genesis D:/works-axis/chaintest/testnet/testnet.g --datadir D:/works-axis/chaintest/testnet --validator.id 3 --validator.pubkey 0xc00454c530e6781b0c7bb378199d903651745c68239c150c793751d4fbb2bf923eb7e3fd155b5235140e8534fdafd726f0dd213ab5bc07d51bcc598aa67bb260901a --validator.password D:/works-axis/chaintest/testnet/password-3.txt

./build/galaxy --config D:/works-axis/chaintest/testnet/config.toml --genesis D:/works-axis/chaintest/testnet/testnet.g --datadir D:/works-axis/chaintest/testnet --validator.id 4 --validator.pubkey 0xc004847343604e986ba2b2fdd64c905503b85918fa206cac3df141e8bb61afee7c07e77196b4372d0b9ff9331b1eeff0a989ca22ab5c95cbfec4f77dc584a96f7278 --validator.password D:/works-axis/chaintest/testnet/password-4.txt

./build/galaxy --config D:/works-axis/chaintest/testnet/config.toml --genesis D:/works-axis/chaintest/testnet/testnet.g --datadir D:/works-axis/chaintest/testnet --validator.id 5 --validator.pubkey 0xc004303bbac1433b0d46feae674db33fd2b62794803f53f0335d8743ae7f6005608bee87c7f73a25fe296fd184536bd83a79e60e6c35d53ca208d2aa00f0b18d36d2 --validator.password D:/works-axis/chaintest/testnet/password-5.txt

./build/galaxy --config D:/works-axis/chaintest/testnet/config.toml --genesis D:/works-axis/chaintest/testnet/testnet.g --datadir D:/works-axis/chaintest/testnet --validator.id 6 --validator.pubkey 0xc0043ecdc02c855c322be643aee9b8f735e82bc664746b09304a9883e553cf64ba2e91c1f3ab39ae24ae27ed3d41c93d947b0e64cb78c06dd408532f99ef2d207895 --validator.password D:/works-axis/chaintest/testnet/password-6.txt

./build/galaxy --config D:/works-axis/chaintest/testnet/config.toml --genesis D:/works-axis/chaintest/testnet/testnet.g --datadir D:/works-axis/chaintest/testnet --validator.id 7 --validator.pubkey 0xc00486fb9204c56ce4fc2e2de29db5c7df9917ea22cdd39b21234dd35a191d3e9677e199142799001101d79b5b8c4a2966c072b3a9f7c06c55151cd2ad27a3d3cd8b --validator.password D:/works-axis/chaintest/testnet/password-7.txt

./build/galaxy --config D:/works-axis/chaintest/testnet/config.toml --genesis D:/works-axis/chaintest/testnet/testnet.g --datadir D:/works-axis/chaintest/testnet --validator.id 8 --validator.pubkey 0xc004be2adee5c7b3d15cdb8ae0a099a759117cc6d5bbe45018fe6e0b05d645ca43069051819a3173c7e87b8947d1d6d3ae85c9dbf725f6778ef417caf186bd8fcac9 --validator.password D:/works-axis/chaintest/testnet/password-8.txt

```

## linux 
```
/home/blockchain/go-axis/build/galaxy --config /home/blockchain/.galaxy/config-rpc.toml --genesis /home/blockchain/.galaxy/testnet.g

/home/axischain/go-axis/build/galaxy --config /home/axischain/.galaxy/config-1.toml --genesis /home/axischain/.galaxy/testnet.g --validator.id 1 --validator.pubkey 0xc00439b9f3f5a56c6aa8c79e01094d496d4e5b0b2116f6e26790177fb7639ffdf473ed428b71eec45e9789e3210cd46e663b9852d2f58ce7070bf1c928ace37d904a --validator.password /home/axischain/.galaxy/password.txt

/home/axischain/go-axis/build/galaxy --config /home/axischain/.galaxy/config-1.toml --genesis /home/axischain/.galaxy/testnet.g --validator.id 2 --validator.pubkey 0xc004be3ddfc6d48ad5d0ab793968f30e412cfaf0a1e1bdf3af63f542c4082191349ef8d2d13a1a1ce2b1512526f1ff53bbb8365d7f2e953b64fcc8cc93ce6ab60d9d --validator.password /home/axischain/.galaxy/password.txt

/home/axischain/go-axis/build/galaxy --config /home/axischain/.galaxy/config-1.toml --genesis /home/axischain/.galaxy/testnet.g --validator.id 3 --validator.pubkey 0xc00454c530e6781b0c7bb378199d903651745c68239c150c793751d4fbb2bf923eb7e3fd155b5235140e8534fdafd726f0dd213ab5bc07d51bcc598aa67bb260901a --validator.password /home/axischain/.galaxy/password.txt

/home/axischain/go-axis/build/galaxy --config /home/axischain/.galaxy/config-1.toml --genesis /home/axischain/.galaxy/testnet.g --validator.id 4 --validator.pubkey 0xc004847343604e986ba2b2fdd64c905503b85918fa206cac3df141e8bb61afee7c07e77196b4372d0b9ff9331b1eeff0a989ca22ab5c95cbfec4f77dc584a96f7278 --validator.password /home/axischain/.galaxy/password.txt

/home/axischain/go-axis/build/galaxy --config /home/axischain/.galaxy/config-1.toml --genesis /home/axischain/.galaxy/testnet.g --validator.id 5 --validator.pubkey 0xc004303bbac1433b0d46feae674db33fd2b62794803f53f0335d8743ae7f6005608bee87c7f73a25fe296fd184536bd83a79e60e6c35d53ca208d2aa00f0b18d36d2 --validator.password /home/axischain/.galaxy/password.txt

/home/axischain/go-axis/build/galaxy --config /home/axischain/.galaxy/config-1.toml --genesis /home/axischain/.galaxy/testnet.g --validator.id 6 --validator.pubkey 0xc0043ecdc02c855c322be643aee9b8f735e82bc664746b09304a9883e553cf64ba2e91c1f3ab39ae24ae27ed3d41c93d947b0e64cb78c06dd408532f99ef2d207895 --validator.password /home/axischain/.galaxy/password.txt

/home/agaxiscnode/go-axis/build/galaxy --config /home/agaxiscnode/.galaxy/config-2.toml --genesis /home/agaxiscnode/.galaxy/testnet.g --validator.id 7 --validator.pubkey 0xc00486fb9204c56ce4fc2e2de29db5c7df9917ea22cdd39b21234dd35a191d3e9677e199142799001101d79b5b8c4a2966c072b3a9f7c06c55151cd2ad27a3d3cd8b --validator.password /home/agaxiscnode/.galaxy/password.txt

/home/agaxiscnode/go-axis/build/galaxy --config /home/agaxiscnode/.galaxy/config-2.toml --genesis /home/agaxiscnode/.galaxy/testnet.g --validator.id 8 --validator.pubkey 0xc004be2adee5c7b3d15cdb8ae0a099a759117cc6d5bbe45018fe6e0b05d645ca43069051819a3173c7e87b8947d1d6d3ae85c9dbf725f6778ef417caf186bd8fcac9 --validator.password /home/agaxiscnode/.galaxy/password.txt

```









# Show info about validator #1
sfc.stakers(1)

sfc = web3.axis.contract(abi).at("0x1c1cb00000000000000000000000000000000000")
sfc.lastValidatorID()
personal.unlockAccount("0x44e54291a2a87D1651A2cb41b557677b3d6FA3b8", "123456789", 60)
tx = sfc.createValidator("0xc004c2c967b2880ff5fb6cfd1373e733d644701382c3ea86162b995baf0d42a44bd0c79bb58503bd83a4441ca87843cc2f2a0099562b99d8a50d3b14193b79e77b90", {from:"0x44e54291a2a87D1651A2cb41b557677b3d6FA3b8", value: web3.toWei("3175000.0", "axis")})
axis.getTransactionReceipt(tx)






Node 1 :

Public IP: 54.217.64.74
Validator Key : 0xc004ce49f7c6afdf7a457f31fc6e6216c8c0df3f46d0f7a80ccd80b35a4f91313184b345efd3724b4639736405a2c3bfb4b1aabea58a34fa2a58cdc3bbea9d80f234

Node 2 :

Public IP: 54.78.74.161
Validator Key : 0xc004653c606752353c3196193ca132466d137863c2966ac0479dbb989893e41bde818e7fcc6e1ec3d22f5262fb1c82814edf6dfb1368f34c6db400d7b44186449719

Node 3:

Public IP: 34.243.200.199
Validator Key : 0xc004a4008e352eabeff85db3f1eea164f6970d6d0d7846061b061c666037e482f68e0b52f4df3bfdc309658f6b52665791e1da10b890f8f7715b218ee28ae5c09ed7

Node 4 :

Public IP: 54.217.135.190
Validator Key : 0xc004939c39a101d5520339bff351a5cba585d85a72db013c187f8a329962c5ce0ffbc9745e680e9cbef2e448aec19809ad8e5fbdf318720d205e89897d4704849950

Node 5 :

Public IP: 63.33.48.74
Validator Key : 0xc00405fb835b662f599d4aeb1952a3e5ddd8002e6043fea0e0b1126adbd0ea8cebe5c6a5024fb8fc069c139b405ed261ea3f1a779c3a83b1041a37e1f9a8dfb163be

Node 6 :

Public IP: 54.170.123.82
Validator Key : 0xc00417b350522e3f8411e171ae3d4f7709ce0b10a02900e71581a7ac268c80472f97a6e6c87d5e2b1edbfa35357032d3627ccf3f90ee8f51ce49360fb1855e12f8ef

Node 7 :

Public IP: 63.34.170.18
Validator Key : 0xc004f704a31767b8f6d3606e166c35adf6ec725675196047b5012c0e05f849901d327bdaf44585a826ef7e986fbf060b8682341385b2db1bda7394edaca00752cd29

Node 8 :

Public IP: 34.244.40.128
Validator Key : 0xc004171eb754a53a173c500bf3c5b1e04a632924a04198b73f1881a19574c79e63e561b17f244dbbfcaec1e45a9310487d4c88bb78df187f17dfc1da98f619b184de



