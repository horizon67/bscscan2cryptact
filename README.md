# bscscan2cryptact

Convert from [BscScan](https://bscscan.com/) Transaction Details to [cryptact CSV format](https://support.cryptact.com/hc/ja/articles/360002571312-%E3%82%AB%E3%82%B9%E3%82%BF%E3%83%A0%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%81%AE%E4%BD%9C%E6%88%90%E6%96%B9%E6%B3%95
).
(Only SELL is supported

## Installation

```
$ git clone https://github.com/horizon67/bscscan2cryptact.git
```

## Usage

```
$ cd bscscan2cryptact

# mac
$ ./main [TX]

# windows
$ ./main.exe [TX]
```

## example

```
./main 0x1ed1eab9e5c142dc8f31497831b12fc5e3a584a1383584646bf8dd009cdaXXXX
CreatedAt,Action,Source,Base,Volume,Price,Counter,Fee,FeeCcy
2021/4/23 00:54:58,SELL,PancakeSwap,ALPACA,40.63352568500912,0.7341372306024434,BUSD,0.38,US
```

