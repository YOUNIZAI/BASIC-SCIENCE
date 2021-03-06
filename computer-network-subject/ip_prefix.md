 - IP 前缀意义
 ```$xslt
 一个IP地址与一个子网掩码按位与后就形成了用于路由的地址的网络/子网标识符（前缀），这是边界路由器需要的信息，
 以确定子网。Internet路由系统其余部分不需要识别子网掩码，因为站点之外的路由器做出路由决策只基于地址的网络号部分，
 并不需要网络/子网或主机部分。因此，子网掩码纯粹是站点内部的局部问题。可变长度子网掩码

```
```$xslt
 子网掩码的出现让网络划分的细粒度更高，提高了IP地址资源的利用率。将一个分配给站点的网络号进一步细分为多个可分配大小相同的子网，
 并根据网络管理员的合理要求使每个子网能支持相同数量的主机。这种方式在复杂的网络环境中，其IP地址资源还是会有不小的浪费。
 因此我们自然想到了放宽一个网络的子网掩码的限制，让其长度可变，即在同一站点的不同部分，将不同长度的子网掩码应用于相同网络号。
 虽然这样增加了地址配置管理的复杂性，但也提高了子网结构的灵活性，因为不同子网可以容纳不同数量的主机。
 这就是可变长度子网掩码（Variable Length Subnet Mask，VLSM），用于分割一个网络号，使每个子网支持不同数量的主机。
 VLSM可对子网进行层次化编址，使得多级子网成为可能，这种高级IP寻址技术允许网络管理员对已子网进行划分，以便最有效地利用现有的地址空间。
```

