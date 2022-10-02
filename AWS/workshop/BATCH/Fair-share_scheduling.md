
[Fair-share scheduling on AWS Batch](https://catalog.us-east-1.prod.workshops.aws/workshops/c3d652f2-6de1-4014-9a1b-c1b3c8f08b8d/en-US)


#### スケジューリングポリシーの各パラメータ

各パラメータごとにワークショップが用意されており、動かしながら動作を理解できる作りになっている。

[Multi Workload](https://catalog.us-east-1.prod.workshops.aws/workshops/c3d652f2-6de1-4014-9a1b-c1b3c8f08b8d/en-US/fair/multi)

Blue と Green の shareIdentifier がある。ともに重みは 1。

この場合、Blue, Green とも 1:1 でリソースが割り当てられることが期待される。
なお、Blue を 0.25 にしたとすると、Blue 側によりリソースが割り当てられる。


[Prioritise Jobs](https://catalog.us-east-1.prod.workshops.aws/workshops/c3d652f2-6de1-4014-9a1b-c1b3c8f08b8d/en-US/fair/priority)

同じ shareIdentifier 内でジョブごとに優先度を変えることができる。
```--scheduling-priority-override``` で設定でき、より大きな値を設定するほど優先的に実行される。


[Reserve Capacity](https://catalog.us-east-1.prod.workshops.aws/workshops/c3d652f2-6de1-4014-9a1b-c1b3c8f08b8d/en-US/fair/capacity)

"computeReservation": 50　に設定した場合は Blue, Green にそれぞれ 25 % が予約され、残りの 50 % は誰もが使用することができる。

reserved ratio は computeReservation/100)^ActiveFairShares で計算される。


[Decay Factor](https://catalog.us-east-1.prod.workshops.aws/workshops/c3d652f2-6de1-4014-9a1b-c1b3c8f08b8d/en-US/fair/decay)

現在の情報だけでなく、shareDecaySeconds で設定した秒数分の過去の情報も考慮される。
例えば、Blue のジョブが大量に実行されている状況で Green のジョブを作成すると、しばらくは Green のジョブにリソースが優先的に割り当てられる。

